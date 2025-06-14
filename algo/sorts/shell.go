package sorting

func shell(s []int) {
	l := len(s)
	h := 1
	for h < l/3 {
		h = 3*h + 1
	}

	for h >= 1 {
		for i := h; i < l; i++ {
			for j := i; j >= h; j = j - h {
				if s[j] < s[j-h] {
					exch(s, j, j-h)
				}
			}
		}
		h = h / 3
	}
}
