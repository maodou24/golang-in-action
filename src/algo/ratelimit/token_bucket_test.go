package ratelimit

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestTokenBucket(t *testing.T) {
	bucket := NewTokenBucket(2, 2)

	goroutineNumb := 5

	var tokenSuccessNum int
	var tokenFailNum int

	for i := 0; i < goroutineNumb; i++ {
		go func(bucket *TokenBucket) {
			for {
				if bucket.HasToken() {
					tokenSuccessNum++
				} else {
					tokenFailNum++
				}

				n := rand.Intn(4)
				time.Sleep(time.Second * time.Duration(n))
			}
		}(bucket)
	}

	time.Sleep(time.Second*10 + time.Millisecond*500)

	fmt.Println("tokenSuccessNum: ", tokenSuccessNum)
	fmt.Println("tokenFailNum: ", tokenFailNum)
	fmt.Println("tokenLeft: ", bucket.tokens.Load())
}
