package array_string

func removeDuplicates(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}

	var k int
	for i := 1; i < n; i++ {
		if nums[i] != nums[k] {
			k++
			nums[k] = nums[i]
		}
	}
	return k + 1
}

func removeDuplicates2(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return n
	}

	k := 1
	for i := 2; i < n; i++ {
		if nums[i] != nums[k] || nums[k-1] != nums[i] {
			k++
			nums[k] = nums[i]
		}
	}
	return k + 1
}
