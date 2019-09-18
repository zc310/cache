package distkvcache

import (
	"github.com/stretchr/testify/assert"
	"github.com/zc310/cache/test"
	"os"
	"path/filepath"
	"testing"
)

func TestCache_Get(t *testing.T) {
	dir := filepath.Join(os.TempDir(), "diskv.cache")
	c, err := New(dir)
	assert.Equal(t, err, nil)
	test.Cache(t, c)
	assert.Equal(t, os.RemoveAll(dir), nil)
}
