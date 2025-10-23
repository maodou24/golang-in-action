package generics

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestGenericFieldInStruct(t *testing.T) {
	userJson := `{"status":"ok","data":{"username":"test","age":20}}`

	var userData DataBody[UserData]
	_ = json.Unmarshal([]byte(userJson), &userData)
	fmt.Println(userData)

	orderJson := `{"status":"ok","data":{"username":"test","age":20}}`

	var orderData DataBody[OrderData]
	_ = json.Unmarshal([]byte(orderJson), &orderData)
	fmt.Println(orderData)
}

func TestGenericFunc(t *testing.T) {
	fmt.Println(Max[int](1, 2))
	fmt.Println(Max(1, 2))
}

func TestGenericFuncParamSlice(t *testing.T) {
	s := []int{1, 2, 3}
	fmt.Println(Sum(s))
}

func TestGenericFuncUnderlying(t *testing.T) {
	// AInt, Bint底层类型为int，可以使用
	aa, ab := AInt(1), AInt(2)
	fmt.Println(MaxIntUnderlying[AInt](aa, ab))

	ba, bb := BInt(3), BInt(2)
	fmt.Println(MaxIntUnderlying(ba, bb))

	var a, b int = 4, 5
	fmt.Println(MaxIntUnderlying(a, b))
	fmt.Println(MaxInt(a, b))
}
