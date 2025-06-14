package list

import "github.com/maodou24/algorithm-go/internel/list"

func intersection(list1 *list.Node, list2 *list.Node) *list.Node {
	var s1, s2 []*list.Node
	for node := list1; node != nil; node = node.Next {
		s1 = append(s1, node)
	}

	for node := list2; node != nil; node = node.Next {
		s2 = append(s2, node)
	}

	i1, i2 := len(s1)-1, len(s2)-1
	for i1 >= 0 && i2 >= 0 && s1[i1] == s2[i2] {
		i1--
		i2--
	}

	if i1+1 == len(s1) {
		return nil
	}

	return s1[i1+1]
}
