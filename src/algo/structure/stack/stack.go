package stack

type Stack interface {
	Push(int)
	Pop() int
	Show() []int
}
