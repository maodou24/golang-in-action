package cache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLRUOverCapacity(t *testing.T) {
	lru := NewLRU(2)

	lru.Put("1", 1)
	lru.Put("2", 2)
	lru.Put("3", 3)

	assert.Nil(t, lru.Get("1"))
	assert.Equal(t, 2, lru.Get("2"))
	assert.Equal(t, 3, lru.Get("3"))
}


func TestLRUPutSame(t *testing.T) {
	lru := NewLRU(2)

	lru.Put("url", "baidu")
	lru.Put("url", "google")

	assert.Equal(t, "google", lru.Get("url"))
}

func TestLRUGet(t *testing.T) {
	lru := NewLRU(2)

	lru.Put("1", 1)
	lru.Put("2", 2)
	lru.Put("2", 3)
	lru.Put("4", 1)

	assert.Nil(t, lru.Get("1"))
	assert.Equal(t, 3, lru.Get("2"))
}