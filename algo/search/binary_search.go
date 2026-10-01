package search

import "math"

func BinarySearch(s []int, ele int) int {
	low := 0
	high := len(s) - 1
	for low <= high {
		mid := low + (high-low)/2
		midValue := s[mid]
		if ele == midValue {
			return ele
		}
		if ele < midValue {
			high = mid - 1
		}
		if ele > midValue {
			low = mid + 1
		}
	}
	return math.MaxInt32 // not find
}

func RecurBinarySearch(s []int, ele int) int {
	return recurBs(s, 0, len(s)-1, ele)
}

// 递归
func recurBs(s []int, low, high, ele int) int {
	mid := low + (high-low)/2
	if low > high {
		return math.MaxInt32 // not find
	}

	if ele < s[mid] {
		return recurBs(s, low, mid-1, ele)
	} else if ele > s[mid] {
		return recurBs(s, mid+1, high, ele)
	} else {
		return ele
	}
}
