package encoding

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// use impl UnmarshalJSON to skip password filed
func TestJSONImpUnmarshalJSON(t *testing.T) {
	s := `{"name":"maodou", "password":"abc"}`

	var u User
	err := json.Unmarshal([]byte(s), &u)
	assert.NoError(t, err)

	fmt.Printf("%+v", u)
}

func TestJSONMarshal(t *testing.T) {
	u := User{Name: "maodou", Password: "abc"}

	d, err := json.Marshal(u)
	assert.NoError(t, err)

	fmt.Println(string(d))
}
