package generics

import "cmp"

type DataBody[T any] struct {
	Status string `json:"status"`
	Data   T      `json:"data"`
}

type UserData struct {
	Username string `json:"username"`
	Age      int    `json:"age"`
}

type OrderData struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func Max[T cmp.Ordered](a, b T) T {
	return max(a, b)
}

func Sum[S []E, E int | int32](s S) E {
	var sum E
	for _, v := range s {
		sum += v
	}
	return sum
}

type AInt int

type BInt int

func MaxIntUnderlying[T ~int](a, b T) T {
	return max(a, b)
}

func MaxInt[T int](a, b T) T {
	return max(a, b)
}
