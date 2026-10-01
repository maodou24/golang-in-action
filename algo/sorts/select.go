package sorting

func selectSort(s []int) {
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[j] < s[i] {
				exch(s, j, i)
			}
		}
	}
}
