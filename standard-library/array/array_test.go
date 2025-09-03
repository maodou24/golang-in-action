package array

import (
	"fmt"
	"testing"
)

func TestModifyArrVal(t *testing.T) {
	arr := [4]int{1, 2, 3, 4}
	// arr[0] = 55
	ModifyArrVal(arr)

	fmt.Println(arr)
	if arr != [4]int{1, 2, 3, 4} {
		t.Error("ModifyArrVal modify arr")
	}
}
