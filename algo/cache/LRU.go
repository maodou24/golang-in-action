package cache

import (
	"container/list"
)

type KV struct {
	key   string
	value any
}

// keys保存缓存列表，最近访问的在队尾，最久没有访问的在队首
type LRU struct {
	cap  int
	keys map[string]*list.Element
	root *list.List
}

func NewLRU(capacity int) *LRU {
	return &LRU{
		cap:  capacity,
		keys: make(map[string]*list.Element, capacity),
		root: list.New(),
	}
}

func (l *LRU) Get(key string) any {
	ele, ok := l.keys[key]
	if !ok {
		return nil
	}

	l.root.MoveToBack(ele)
	return ele.Value.(KV).value
}

// key已存在：
// 		1. 更新key对应的value值
//      2. 将key移到队尾
// key不存在：
//		1. 判断size是否超过cap，超过时需要删除最久的
// 		2. 将key添加到队尾
func (l *LRU) Put(key string, v any) {
	ele, ok := l.keys[key]
	if ok {
		ele.Value = KV{key: key, value: v}
		l.root.MoveToBack(ele)
		return
	}

	if l.root.Len() >= l.cap {
		kv := l.root.Remove(l.root.Front())
		delete(l.keys, kv.(KV).key)
	}

	l.keys[key] = l.root.PushBack(KV{key: key, value: v})
}
