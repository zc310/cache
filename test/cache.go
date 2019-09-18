package test

import (
	"github.com/stretchr/testify/assert"
	"github.com/zc310/cache"
	"testing"
	"time"
)

func Cache(t *testing.T, c cache.Cache) {
	var err error
	key := []byte("key1")
	value := []byte("1234567890")

	err = c.SetTimeout(key, value, 0)
	assert.Equal(t, err, nil)

	b, ok := c.Get(key)
	assert.Equal(t, ok, true)
	assert.Equal(t, b, value)

	key = []byte("key3")
	err = c.SetTimeout(key, value, time.Minute)
	assert.Equal(t, err, nil)

	b, ok = c.Get(key)
	assert.Equal(t, ok, true)
	assert.Equal(t, b, value)

	k := []byte("abc")
	v := []byte("0123456789")
	err = c.SetTimeout(k, v, time.Hour*24)
	assert.Equal(t, err, nil)

	b, ok = c.Get(k)
	assert.Equal(t, ok, true)
	assert.Equal(t, b, v)

	b, ok = c.GetRange(k, 6, 9)
	assert.Equal(t, ok, true)
	assert.Equal(t, b, []byte("678"))

	err = c.SetTimeout(k, []byte("abc"), time.Hour*24)
	assert.Equal(t, err, nil)

	b, ok = c.Get(k)
	assert.Equal(t, ok, true)
	assert.Equal(t, b, []byte("abc"))

	err = c.Delete(k)
	assert.Equal(t, err, nil)

	b, ok = c.Get(k)
	assert.Equal(t, ok, false)
}
