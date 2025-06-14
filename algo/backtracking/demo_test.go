package backtracking

import (
	"fmt"
	"testing"
)

func combinationSum(candidates []int, target int) [][]int {
	var res [][]int

	var dfs func(n int, temp, candidates []int)

	dfs = func(n int, temp, candidates []int) {
		if n > target {
			return
		}
		if n == target {
			res = append(res, temp)
			return
		}

		for _, candidate := range candidates {
			temp = append(temp, candidate)
			dfs(n+candidate, temp, candidates)

			temp = temp[:len(temp)-1]
		}

		fmt.Println()
	}

	dfs(0, []int{}, candidates)

	return res
}

func TestBackTracking(t *testing.T) {
	combinationSum([]int{2, 3, 6, 7}, 7)
}
