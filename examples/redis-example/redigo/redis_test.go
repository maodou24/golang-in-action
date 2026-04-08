package redigo

import (
	"fmt"
	"testing"
	"time"

	"github.com/gomodule/redigo/redis"
)

var pool redis.Pool

func init() {
	pool = redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379",
				redis.DialConnectTimeout(10*time.Second),
				redis.DialReadTimeout(10*time.Second),
				redis.DialWriteTimeout(10*time.Second),
			)
		},
	}
}

func GetConn() redis.Conn {
	return pool.Get()
}

func TestRedisString(t *testing.T) {
	conn := GetConn()
	defer conn.Close()

	r, err := conn.Do("SET", "key", "value")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(r)

	r, err = redis.String(conn.Do("GET", "key"))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(r)
}
