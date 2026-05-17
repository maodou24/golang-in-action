package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUsageDictionary(t *testing.T) {
	dict := &Dictionary{}

	dict.Insert("hello")
	dict.Insert("world")

	assert.True(t, dict.Contains("hello"))
	assert.True(t, dict.Contains("world"))
	assert.False(t, dict.Contains("golang"))
}
