package distkvcache

import (
	"cache"
	"crypto/md5"
	"encoding/hex"
	"github.com/peterbourgon/diskv"

	"github.com/zc310/utils"

	"time"
)

type Cache struct {
	db *diskv.Diskv
}

func getKey(k []byte) string {
	h := md5.New()
	h.Write(k)
	return hex.EncodeToString(h.Sum(nil))
}

func (p *Cache) Get(key []byte) ([]byte, bool) {
	k := getKey(key)
	if ok := p.db.Has(k); !ok {
		return nil, false
	}
	b, err := p.db.Read(k)
	if err != nil {
		return nil, false
	}
	var r cache.Value
	if err = cache.Unmarshal(b, &r); err != nil {
		return nil, false
	}
	return r.Value, r.Valid()
}
func (p *Cache) GetRange(key []byte, low, high int64) (b []byte, ok bool) {
	if b, ok = p.Get(key); !ok {
		return
	}
	return b[low:high], ok
}
func (p *Cache) Set(key, value []byte) error {
	return p.SetTimeout(key, value, 0)
}
func (p *Cache) SetTimeout(key []byte, value []byte, timeout time.Duration) error {
	var r cache.Value
	r.Value = value
	r.Size = int64(len(value))
	if timeout > 0 {
		r.Timeout = time.Now().Add(timeout).Unix()
	}

	b, err := cache.Marshal(&r)
	if err != nil {
		return err
	}

	return p.db.Write(getKey(key), b)
}
func (p *Cache) Delete(key []byte) error {
	return p.db.Erase(getKey(key))
}
func (p *Cache) Empty() error {
	return nil
}
func (p *Cache) Clean() error {
	return nil
}

func New(path string) (*Cache, error) {
	transformBlockSize := 2
	db := diskv.New(diskv.Options{
		BasePath: path,
		Transform: func(s string) []string {
			var (
				sliceSize = utils.MinInt(6, len(s)) / transformBlockSize
				pathSlice = make([]string, sliceSize)
			)
			for i := 0; i < sliceSize; i++ {
				from, to := i*transformBlockSize, (i*transformBlockSize)+transformBlockSize
				pathSlice[i] = s[from:to]
			}
			return pathSlice
		},
		CacheSizeMax: 1024 * 1024 * 3,
	})

	return NewWithDB(db), nil
}

func NewWithDB(db *diskv.Diskv) *Cache {
	c := &Cache{db}
	return c
}
