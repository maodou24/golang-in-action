package chainid

import (
	"fmt"
	"strings"
)

type Work struct {
	ChainID ChainID
	Valid   bool
}

type TrieNode struct {
	IsEnd    bool
	ChainID  ChainID
	Children map[string]*TrieNode
}

func NewTrieNode(chainIDs []string) *TrieNode {
	root := &TrieNode{Children: make(map[string]*TrieNode)}

	for _, chainID := range chainIDs {
		parts := strings.Split(chainID, ".")
		node := root
		for _, part := range parts {
			if node.Children[part] == nil {
				node.Children[part] = &TrieNode{Children: make(map[string]*TrieNode)}
			}
			node = node.Children[part]
		}
		node.IsEnd = true
		node.ChainID = ChainID(chainID)
	}

	return root
}

func Traverse(root *TrieNode) {
	traverseRecursive(root)
}

func traverseRecursive(node *TrieNode) {
	if node == nil {
		return
	}
	if node.IsEnd {
		fmt.Println(node.ChainID)
	}

	// 递归遍历所有子节点
	for _, child := range node.Children {
		traverseRecursive(child)
	}
}

func BFSWithLevels(root *TrieNode) [][]ChainID {
	if root == nil {
		return nil
	}

	var result [][]ChainID
	queue := []*TrieNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		var currentLevel []ChainID

		for i := 0; i < levelSize; i++ {
			// 出队
			currentNode := queue[0]
			queue = queue[1:]

			// 处理当前节点
			if currentNode.IsEnd {
				currentLevel = append(currentLevel, currentNode.ChainID)
			}

			// 将所有子节点入队
			for _, child := range currentNode.Children {
				queue = append(queue, child)
			}
		}

		result = append(result, currentLevel)
	}

	return result
}
