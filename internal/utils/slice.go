package utils

func FilterFunc[T any](s []T, predicate func(idx int, t T) bool) []T {
	result := make([]T, 0, len(s))
	for i, e := range s {
		if predicate(i, e) {
			result = append(result, e)
		}
	}
	return result
}

// Unique returns a new slice containing only the unique elements of the provided slice.
// The result has length at most len(s).
// The result is never nil.
func Unique[T comparable](s []T) []T {
	result := make([]T, 0, len(s))
	seen := make(map[T]struct{})

	for _, e := range s {
		if _, ok := seen[e]; !ok {
			seen[e] = struct{}{}
			result = append(result, e)
		}
	}
	return result
}

// UniqueFunc returns a new slice containing only the unique elements of the provided slice,
// as determined by the keyFunc function.
// The result has length at most len(s).
// The result is never nil.
func UniqueFunc[T any, E comparable](s []T, keyFunc func(T) E) []T {
	seen := make(map[E]struct{})
	result := make([]T, 0, len(s))

	for _, e := range s {
		key := keyFunc(e)
		if _, ok := seen[key]; !ok {
			seen[key] = struct{}{}
			result = append(result, e)
		}
	}
	return result
}
