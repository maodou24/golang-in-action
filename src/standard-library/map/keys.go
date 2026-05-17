package _map

import "fmt"

func ChannelAsMapKey() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	m := make(map[chan int]int)
	m[ch1] = 1
	m[ch2] = 2

	fmt.Println(m[ch1])
	fmt.Println(m[ch2])
}

func MapAsMapKey() {
	type mkey struct {
		m1 map[int]int
	}

	// m := make(map[mkey]int) // compile error: invalid map key type mkey
}
