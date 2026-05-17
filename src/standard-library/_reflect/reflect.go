package _reflect

import (
	"fmt"
	"reflect"
)

// Law1 反射可以将interface类型变量转换成反射对象
func Law1() {
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	fmt.Println("Type:", v.Type())   // float
	fmt.Println("Value:", v.Float()) // 3.4
}

// Law2 反射可以将反射对象还原成interface对象
func Law2() {
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	y := v.Interface().(float64) // convert reflect.Value back to float64
	fmt.Println(y)               // 3.4
}

// Law3 反射对象可修改，value值必须是可设置的
func Law3() {
	var x float64 = 3.4
	v := reflect.ValueOf(&x).Elem()
	if v.CanSet() {
		v.SetFloat(2.4)
		fmt.Println(x) // 2.4
	}
}
