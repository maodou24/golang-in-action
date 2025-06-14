package list

func IntSliceToListNode(input []int) *Node {
	if len(input) == 0 {
		return nil
	}
	result := &Node{}
	head := result
	for i := range input {
		head.Val = input[i]
		if i < len(input)-1 {
			head.Next = &Node{}
			head = head.Next
		}
	}
	return result
}
