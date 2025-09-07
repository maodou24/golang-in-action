package chainid

import (
	"context"
	"testing"
	"time"
)

func TestChainID(t *testing.T) {
	tasks := []*Device{
		{ID: "1"},
		{ID: "1.1"},
		{ID: "1.2"}, {ID: "1.2.2"}, {ID: "1.2.3"},
		{ID: "1.4.2"},
		{ID: "2"},
	}
	root := NewTrieNode(tasks)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	wf := NewBatch(ctx)

	var dfs func(node *TrieNode) []TaskID
	dfs = func(node *TrieNode) []TaskID {
		if node == nil {
			return []TaskID{}
		}

		var s []TaskID
		for _, child := range node.Children {
			s = append(s, dfs(child)...)
		}

		if node.IsEnd {
			node.Task.Dependencies = s
			node.Task.State = TaskStatePending
			node.Task.Execute = func() error {
				return nil
			}

			wf.Add(node.Task)
			s = append(s, node.Task.ID)
		}
		return s
	}

	dfs(root)

	if err := wf.Start(); err != nil {
		t.Fatal(err)
	}
}
