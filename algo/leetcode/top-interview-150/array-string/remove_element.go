package array_string

func removeElement(nums []int, val int) int {
	var idx int
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[idx] = nums[i]
			idx++
		}
	}
	return idx
}
