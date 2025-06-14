package sorting

func exch(s []int, i, j int) {
	temp := s[i]
	s[i] = s[j]
	s[j] = temp
}
