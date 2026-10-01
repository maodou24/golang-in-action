package tree

type TrieTree[T comparable] struct {
	Val      T
	IsEnd    bool
	Children map[T]*TrieTree[T]
}
