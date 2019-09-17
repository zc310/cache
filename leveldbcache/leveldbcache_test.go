package leveldbcache

import (
	"cache/test"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestCache_Get(t *testing.T) {
	dir := "/tmp/test_cache"
	c, err := New(dir)
	assert.Equal(t, err, nil)
	test.Cache(t, c)
	assert.Equal(t, os.RemoveAll(dir), nil)
}
