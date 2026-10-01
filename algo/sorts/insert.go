package sorting

func insert(s []int) {
	for i := 1; i < len(s); i++ {
		for j := i; j > 0 && s[j] < s[j-1]; j-- {
			exch(s, j, j-1)
		}
	}
}
