package basic

import (
	"fmt"
	"reflect"
)

func createSlice(a any) any {
	t := reflect.TypeOf(a)

	s := reflect.MakeSlice(t, 0, 0)

	reflect.Append(s, reflect.ValueOf(a))
	fmt.Println(t)
	return nil
}
