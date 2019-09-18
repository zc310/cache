package mgocache

import (
	"github.com/globalsign/mgo"
	"github.com/zc310/cache/test"
	"testing"
)

func TestCache_Get(t *testing.T) {
	session, err := mgo.Dial("mongodb://127.0.0.1:27017/lot")
	if err != nil {
		t.Errorf("%v", err.Error())
	}
	defer session.Close()
	c := New(session.DB("cache").C("test_cache"))
	test.Cache(t, c)
}
