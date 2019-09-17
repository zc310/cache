package distkvcache

import (
	"cache/test"
	"github.com/stretchr/testify/assert"
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
