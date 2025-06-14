package array

func Max(arr []int) int {
	max := arr[0]
	for _, v := range arr { // range is faster then fori
		if v > max {
			max = v
		}
	}
	return max
}

func Sum(arr []int) int {
	left, right := 0, len(arr)-1

	var sum int
	for left < right {
		sum += arr[left] + arr[right]
		left++
		right--
	}

	return sum
}

func Average(arr []int) int {
	n := len(arr)

	var sum int
	for i := range arr {
		sum += arr[i]
	}
	return sum / n
}

func Reverse(arr []int) []int {
	n := len(arr)

	for i := 0; i < n/2; i++ {
		arr[i], arr[n-i-1] = arr[n-i-1], arr[i]
	}

	return arr
}

func Copy(arr []int) []int {
	newArr := make([]int, len(arr))

	for i := range arr {
		newArr[i] = arr[i]
	}

	return newArr
}