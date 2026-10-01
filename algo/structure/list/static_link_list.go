package list

const maxSize = 10

type staticLinkList struct {
	arr  [maxSize]staticLinkNode
	size int
}

type staticLinkNode struct {
	val     int
	nextCur int
}

// 数组第一个元素 arr[0].nextCur存放下一个备用结点的下标
// 数组最后一个元素 arr[maxSize-1].nextCur存放链表第一个元素的下标（相当于头结点）
func NewStaticLinkList() *staticLinkList {
	list := &staticLinkList{}
	for i := range list.arr {
		list.arr[i].nextCur = i + 1
	}

	list.arr[maxSize-1].nextCur = 0
	return list
}

// first element index is 1
// 插入一个结点的步骤：
//     1. arr[0].nextCur获取待插入的结点 arr[n]，更新 arr[0].nextCur的值为 arr[n].nextCur
//     2. arr[maxSize-1].nextCur找到头结点，通过头结点找到目标位置（n）的前一个节点（k = n-1）。
//     3. 目标结点 arr[n].nextCur设为 arr[k].nextCur，并将 arr[k].nextCur设为n
func (l *staticLinkList) Insert(n int, v int) {
	if n < 1 || n > l.size+1 {
		return
	}
	next := l.next()
	l.arr[next].val = v
	k := maxSize - 1
	for i := 1; i <= n-1; i++ {
		k = l.arr[k].nextCur
	}
	l.arr[next].nextCur = l.arr[k].nextCur
	l.arr[k].nextCur = next

	l.size++
}

// first element index is 1
// 删除一个结点的步骤：
//     1. arr[maxSize-1].nextCur找到头结点，通过头结点找到待删除位置（n）的前一个节点（k = n-1）。
//     2. 目标结点 arr[k].nextCur设为 arr[n].nextCur
//     3. 删除后，该位置变为优先插入的空闲结点；arr[n].nextCur设为 arr[0].nextCur，并将 arr[0].nextCur设为n
func (l *staticLinkList) Remove(n int) int {
	if n < 1 || n > l.size+1 {
		return 0
	}
	k := maxSize - 1
	for i := 1; i <= n-1; i++ {
		k = l.arr[k].nextCur
	}
	j := l.arr[k].nextCur
	l.arr[k].nextCur = l.arr[j].nextCur
	return l.freeNode(j)
}

func (l *staticLinkList) next() int {
	nextCur := l.arr[0].nextCur
	l.arr[0].nextCur = l.arr[nextCur].nextCur
	return nextCur
}

func (l *staticLinkList) freeNode(cur int) int {
	result := l.arr[cur].val
	l.arr[cur].nextCur = l.arr[0].nextCur
	l.arr[0].nextCur = cur

	l.size--
	return result
}

func (l *staticLinkList) show() []int {
	k := maxSize - 1
	var result []int

	node := l.arr[l.arr[k].nextCur]
	for node.nextCur != l.arr[0].nextCur {
		result = append(result, node.val)
		node = l.arr[node.nextCur]
	}
	return result
}
