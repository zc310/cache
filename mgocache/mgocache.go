package mgocache

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/zc310/cache"
	"time"
)

type Cache struct {
	col *mgo.Collection
}

func (p *Cache) Get(key []byte) ([]byte, bool) {
	var q struct {
		Key   []byte `bson:"_id"`
		Value []byte `bson:"v"`
	}

	err := p.col.FindId(key).One(&q)
	if err != nil {
		return nil, false
	}
	var r cache.Value
	if err = cache.Unmarshal(q.Value, &r); err != nil {
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
func (p *Cache) Set(key []byte, value []byte) error {
	return p.SetTimeout(key, value, 0)
}

func (p *Cache) SetTimeout(key []byte, value []byte, timeout time.Duration) error {
	var r cache.Value
	tm1 := time.Now()
	r.Value = value
	r.Size = int64(len(value))
	if timeout > 0 {
		tm1 = tm1.Add(timeout)
		r.Timeout = tm1.Unix()
	} else {
		tm1 = tm1.AddDate(99, 0, 0)
	}

	b, err := cache.Marshal(&r)
	if err != nil {
		return err
	}
	_, err = p.col.UpsertId(key, bson.M{"v": b, "t": tm1})
	return err
}

func (p *Cache) Delete(key []byte) error {
	return p.col.RemoveId(key)
}
func (p *Cache) Clean() error {
	return nil
}
func (p *Cache) Empty() error {
	_, err := p.col.RemoveAll(nil)
	return err
}
func New(col *mgo.Collection) *Cache {
	_ = col.EnsureIndex(mgo.Index{Key: []string{"t"}, ExpireAfter: time.Hour})
	return &Cache{col}
}
