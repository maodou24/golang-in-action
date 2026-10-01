package sorting

// 快速排序
func quickSort(s []int) {
	quick(s, 0, len(s)-1)
}

func quick(s []int, left, right int) {
	if left >= right {
		return
	}

	p := partition(s, left, right)
	quick(s, left, p-1)
	quick(s, p+1, right)
}

func partition(s []int, left, right int) int {
	i, j := left, right
	for i < j {
		for i < j && s[j] >= s[left] {
			j--
		}

		for i < j && s[i] <= s[left] {
			i++
		}

		s[i], s[j] = s[j], s[i]
	}
	s[i], s[left] = s[left], s[i]

	return i
}
