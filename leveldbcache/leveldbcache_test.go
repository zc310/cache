package leveldbcache

import (
	"github.com/stretchr/testify/assert"
	"github.com/zc310/cache/test"
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
