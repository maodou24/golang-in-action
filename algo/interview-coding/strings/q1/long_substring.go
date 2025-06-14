package q1

func longSubstring(s string, k int) int {
	var ans int

	n := len(s)
	for i := 0; i < n; i++ {
		m := make(map[byte]bool)

		count := 0
		for j := i; j < n; j++ {
			if len(m) == k {
				if !m[s[j]] {
					break
				}
			}
			count++
			m[s[j]] = true
		}
		ans = max(ans, count)
	}
	return ans
}

func longSubstring2(s string, k int) int {
	var ans int

	left, right := 0, 0
	n := len(s)
	for left < n {
		right = left
		m := make(map[byte]bool)
		for len(m) <= k && right < n {
			m[s[right]] = true
			right++
		}

		if right == n {
			if len(m) <= k {
				ans = max(ans, right-left)
			}
			if len(m)-1 == k {
				ans = max(ans, right-left-1)
			}
			break
		}

		ans = max(ans, right-left-1)
		left++
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
