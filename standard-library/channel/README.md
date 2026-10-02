# channel 通道

**创建channel**

```go
ch1 := make(chan int) // 无缓存通道
ch2 := make(chan int, 10) // 有缓存通道
```

**channel发送接收数据**
```go
ch := make(chan int, 2)

ch <- 1 // 发送数据到channel中
a := <- ch // 从channel中读取数据
```

对channel操作的总结

| 操作      | nil  | closed              | normal                                             |
|---------|------|---------------------|----------------------------------------------------|
| close   | panic | panic               | 正常关闭                                               |
| 读（<-ch） | 阻塞   | 读到对应类型的零值（ok为false） | 阻塞或者正常读取数据。缓存channel中没有数据或者无缓存channel发送者没准备好时读取会阻塞 |
| 写（ch<-） | 阻塞 | panic               | 阻塞或者正常写入数据。缓存channel中数据满或者无缓存channel接收者没准备好时读取会阻塞  |

## channel的数据结构

```go
type hchan struct {
    qcount   uint // channel中元素数量
    dataqsiz uint // chan底层循环数组的长度
    buf      unsafe.Pointer // 指向底层循环数组的指针，只有缓存的channel才有
    elemsize uint16         // 元素大小
    closed   uint32         // chan是否关闭
    elemtype *_type // 元素类型
    sendx    uint   // 已发送元素在循环数组中的索引
    recvx    uint   // 已接收元素在循环数组中的索引
    recvq    waitq  // 等待接收的goroutine队列
    sendq    waitq // 等待发送的goroutine队列
    lock mutex     // 锁
}
```

## 练习题

### 读取nil channel会导致阻塞，设计一个方法避免阻塞

可以使用`select`关键字读取
