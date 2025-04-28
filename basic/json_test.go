package basic

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserUnmarshalJSON(t *testing.T) {
	s := `{"name":"maodou", "password":"abc"}`

	var u User
	err := json.Unmarshal([]byte(s), &u)
	assert.NoError(t, err)

	fmt.Println(u)
}

func TestUserMarshalJSON(t *testing.T) {
	u := User{Name: "maodou", Password: "abc"}

	d, err := json.Marshal(u)
	assert.NoError(t, err)

	fmt.Println(string(d))
}
