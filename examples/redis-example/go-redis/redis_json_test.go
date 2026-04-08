package goredis

import (
	"fmt"
	"testing"

	"github.com/redis/go-redis/v9"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestReidsJson(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	defer rdb.Close()

	u := User{
		Name: "maodou",
		Age:  24,
	}
	err := rdb.JSONSet(ctx, "user", "$", u).Err()
	if err != nil {
		t.Fatal(err)
	}

	val, err := rdb.JSONGet(ctx, "user", "$").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("json", val)
}
