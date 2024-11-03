# Map

在计算机科学里，被称为相关数组、map、符号表或者字典，是由一组 <key, value> 对组成的抽象数据结构，并且同一个key 只会出现一次。

Go 语言采用的是哈希查找表，并且使用链表解决哈希冲突。

map初始化
```go
m := make(map[int]int, 0)  // 使用map声明初始化
m1 := map[string]string{"key": "value"} // 直接声明
var m2 = map[string]string{"key": "value"}
```

获取map中数据
```go
m := make(map[string]string, 1)
m["key"] = "value"

v1 := m["key"]
v2, ok := m["key"]
// 省略...
```

## Map的数据结构

```go
type hmap struct {
	count     int // map的大小，调用len(map)返回该值
	flags     uint8
	B         uint8  // bucket存放2^B元素
	noverflow uint16 // overflow 的 bucket 近似数
	hash0     uint32 // hash seed

	buckets    unsafe.Pointer // array of 2^B Buckets. may be nil if count==0.
	oldbuckets unsafe.Pointer // 扩容场景老的bucket
	nevacuate  uintptr        // 指示扩容进度，小于此地址的buckets迁移完成

	extra *mapextra // optional fields
}
```

buckets数组的长度为`2^B`，bucket里面存储了key和value，buckets是一个指针，指向结构体`bmap`的数组
```go
type bmap struct {
	tophash [bucketCnt]uint8
}
```

`bmap`就是通常说的桶，桶里面最多装8个key，value，这些key之所以会落入同一个桶，是因为经过哈希计算后，哈希结果是一样的。在桶内，又会根据hash值的高8位来决定key落入桶内的哪个位置

编译期间，动态创建出新的结构`runtime/hashmap.go`
```go
type bmap struct {
    topbits  [8]uint8
    keys     [8]keytype
    values   [8]valuetype
    pad      uintptr
    overflow uintptr
}
```
当同一个hash超过8个以后，会在`overflow`扩展一个桶

## 扩容流程

负载因子达到当前界限，将会动态扩容当前大小的两倍作为新的容量

扩容条件：
- 没有正在进行扩容
- 触发最大LoadFactor装载因子
- 过多的溢出桶

扩容规则
- 负载因子超过当前界限，将会动态扩容当前大小的两倍作为新容量
- 溢出桶场景，扩容规则是`sameSizeGrow`，不改变大小的扩容动作
```go
	bigger := uint8(1)
	if !overLoadFactor(h.count+1, h.B) {
		bigger = 0
		h.flags |= sameSizeGrow
	}
	oldbuckets := h.buckets
	newbuckets, nextOverflow := makeBucketArray(t, h.B+bigger, nil)
```

新申请的扩容空间后，不会立马进行迁移，而是采用增量库容的方式，当方位到具体的bucket时，才会逐渐的进行迁移（oldbucket迁移到bucket）

为什么是增量扩容？

全量扩容的话，当hmap的容量比较大的时候，扩容会花费大量的时间和内存，导致卡顿延迟。

## Map并发安全吗

在查找、复制、遍历、删除的过程中都会检查写标识(`hashWriting`)，一旦发现写标识置位，则直接panic

```go
// runtime/map.go
    // ...
	if h.flags&hashWriting != 0 {
		fatal("concurrent map read and map write")
	}
    // ...
```

赋值，删除都会在操作之前将写标识(`hashWriting`)置位，然后才会执行后续的操作，执行完成后写标识复位
```go
func mapclear(t *maptype, h *hmap) {
	// 省略之前的代码
	// 写标识置位
	h.flags ^= hashWriting
	
	// 省略中间的删除操作

	if h.flags&hashWriting == 0 {
		fatal("concurrent map writes")
	}
        // 写标识复位·
	h.flags &^= hashWriting
}
```