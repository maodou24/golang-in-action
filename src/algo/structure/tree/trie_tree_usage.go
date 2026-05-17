package tree

var root = &TrieTree[rune]{
	Children: make(map[rune]*TrieTree[rune]),
}

type Dictionary TrieTree[rune]

func (d *Dictionary) Insert(word string) {
	var parent = root
	for _, ch := range word {
		child, ok := parent.Children[ch]
		if !ok {
			child = &TrieTree[rune]{
				Val:      ch,
				Children: make(map[rune]*TrieTree[rune]),
			}
			parent.Children[ch] = child
		}
		parent = child
	}
	parent.IsEnd = true
}

func (d *Dictionary) Contains(word string) bool {
	var parent = root
	for _, ch := range word {
		child, ok := parent.Children[ch]
		if !ok {
			return false
		}
		parent = child
	}

	return parent != nil && parent.IsEnd
}
