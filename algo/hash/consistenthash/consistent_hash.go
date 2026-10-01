package consistenthash

import (
	"fmt"
	"sort"
)

type Hash func([]byte) uint32

type Map struct {
	hash     Hash
	replicas int
	keys     []int
	hashMap  map[int]string
}

func New(replicas int, hash Hash) *Map {
	m := &Map{
		replicas: replicas,
		hash:     hash,
		hashMap:  make(map[int]string, 0),
	}

	return m
}

func (m *Map) Add(keys ...string) {
	for _, key := range keys {
		for i := 0; i < m.replicas; i++ {
			nKey := fmt.Sprintf("%v_%v", key, i)
			hash := int(m.hash([]byte(nKey)))

			m.keys = append(m.keys, hash)
			m.hashMap[hash] = key
		}
	}

	sort.Ints(m.keys)
}


func (m *Map) Get(key string) string {
	if len(key) == 0 {
		return ""
	}

	hash := int(m.hash([]byte(key)))

	idx := sort.Search(len(m.keys), func(i int) bool {
		return m.keys[i] >= int(hash)
	})

	if idx == len(m.keys) {
		idx = 0
	}

	return m.hashMap[m.keys[idx]]
}