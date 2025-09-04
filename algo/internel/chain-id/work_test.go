package chainid

import (
	"fmt"
	"testing"
)

func TestChainID(t *testing.T) {
	chainIDs := []string{
		"1",
		"1.1",
		"1.2", "1.2.2",
		"1.4.2",
		"2",
	}
	root := NewTrieNode(chainIDs)

	Traverse(root)

	result := BFSWithLevels(root)
	for i := range result {
		fmt.Println(result[i])
	}
}
