package go1_24

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

func TestJsonTagOmitZero(t *testing.T) {
	type T struct {
		Name string `json:"name"`
		Age  int    `json:"age,omitzero"`   // 值是其类型的零值时, 则会被忽略
		Data any    `json:"data,omitempty"` // 类型为nil时，则会被忽略
	}
	value := T{
		Name: "John Doe",
		Age:  42,
	}

	data, _ := json.Marshal(value)
	fmt.Println(string(data)) // {"name":"John Doe","age":42}

	value.Age = 0
	data, _ = json.Marshal(value)
	fmt.Println(string(data)) // {"name":"John Doe"}
}

func TestStringsLines(t *testing.T) {
	sequence := strings.Lines("aaaa\nbbbb\n")
	for s := range sequence {
		fmt.Println(s)
	}
	// output
	// aaaa
	//
	// bbbb
	//
}

func TestStringsSplitSeq(t *testing.T) {
	sequence := strings.SplitSeq("aaaa\nbbbb\n", "\n")
	for s := range sequence {
		fmt.Println(s)
	}
	// output
	// aaaa
	// bbbb
}

func TestStringsSplitAfterSeq(t *testing.T) {
	sequence := strings.SplitAfterSeq("aaaa\nbbbb\ncccc\ndddd", "\n")
	for s := range sequence {
		fmt.Println(s)
	}
	// output
	// aaaa
	//
	// bbbb
	//
}

func TestStringsFieldsSeq(t *testing.T) {
	sequence := strings.FieldsSeq("aaaa\nbbbb\n")
	for s := range sequence {
		fmt.Println(s)
	}
	// output
	// aaaa
	// bbbb
}

func TestStringsFieldsFuncSeq(t *testing.T) {
	sequence := strings.FieldsFuncSeq("aaaa\nbbbb\n", func(r rune) bool {
		return r == '\n'
	})
	for s := range sequence {
		fmt.Println(s)
	}
	// output
	// aaaa
	// bbbb
}
