package standard_library

func ConvertToTypeT[T string](a any) T {
	v, ok := a.(T)
	if ok {
		return v
	}

	var t T
	return t
}

func Max[T int | int16](a, b T) bool {
	return a > b
}
