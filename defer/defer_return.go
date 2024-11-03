package _defer

// input = 1
// output = 4
func DeferFunc1(i int) (t int) {
	t = i

	defer func() {
		t = t + 3
	}()

	return t
}

// input = 2
// output = 4
func DeferFunc2(i int) (t int) {
	t = i // 第1步修改变量t的值为i(2)

	defer func() {
		t = t + 3 // 第3步修改变量t的值为t+3
	}()

	return 1 // 第2步修改变量t的值为1
}

// input = 1
// output = 1
func DeferFunc3(i int) int {
	t := i

	defer func() {
		t = t + 3
	}()

	return t
}

// input = 1
// output = 1
func DeferFunc4(i int) int {
	t := i

	defer func(t int) {
		t = t + 3
	}(t)

	return t
}

// input = 1
// output = 4
func DeferFunc5(i *int) *int {
	defer func() {
		*i = *i + 3
	}()

	return i
}

// input = 1
// output = 1
func DeferFunc6(i int) int {
	defer func() {
		i = i + 3
	}()

	return i
}
