# 切片

切片的底层数据是数组，切片相对于数组来说非常灵活，可以支持动态扩缩容

切片初始化
```go
s1 := []int{1, 2, 3}
s2 := make([]int, 2)  // 指定初始长度
s3 := make([]int, 0, 2) // 指定初始Cap
```

下标灵活访问切片
```go
s := []int{1, 2, 3, 4}
s1 := s[:2]  // 取切片s的第二位以及之前的元素
s2 := s[2:]  // 取切片s的第三位以及之后的元素
```

索引操作`[left:right]`，取`left`索引与`right`索引之间的元素，包括`left`所在索引元素，不包括`right`所在索引元素

## 切片的数据结构
`
切片的底层是数组的引用，底层数组可以被多个切片同时`指向，对其中的一个切片进行操作时会影响到其他的切片

```go
// runtime/slice.go

type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}
```

## 切片扩容原理

`runtime/slice.go`中`growslice`

`go1.18`之前，容量小于1024时，容量翻倍；容量大于等于1024时，容量扩容1.25倍
`go1.18`之后，容量小于256时，容量翻倍；容量大于等于256时，容量扩容为`newCap += (newcap + 3*threshold) >> 2`
