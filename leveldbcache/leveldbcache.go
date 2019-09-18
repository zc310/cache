package leveldbcache

import (
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/zc310/cache"
	"time"
)

type Cache struct {
	db *leveldb.DB
}

func (p *Cache) Get(key []byte) ([]byte, bool) {
	b, err := p.db.Get(key, nil)
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

	return p.db.Put(key, b, nil)
}
func (p *Cache) Delete(key []byte) error {
	return p.db.Delete(key, nil)
}
func (p *Cache) Empty() error {
	return nil
}
func (p *Cache) Clean() error {
	var cv cache.Value
	var err error
	now := time.Now().Unix()
	iter := p.db.NewIterator(nil, nil)
	for iter.Next() {
		key := iter.Key()
		err = cache.Unmarshal(iter.Value(), &cv)
		if err != nil {
			err = p.Delete(key)
			continue
		}
		if cv.Timeout == 0 {
			continue
		}
		if now < cv.Timeout {
			err = p.Delete(key)
		}
	}
	iter.Release()
	return err
}

func New(path string) (*Cache, error) {
	var err error
	db, err := leveldb.OpenFile(path, nil)

	if err != nil {
		return nil, err
	}
	return NewWithDB(db), nil
}

func NewWithDB(db *leveldb.DB) *Cache {
	c := &Cache{db}
	return c
}
