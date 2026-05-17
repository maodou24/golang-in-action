package ratelimit

import (
	"sync/atomic"
	"time"
)

type TokenBucket struct {
	tokens atomic.Int64

	rate int64     // 令牌生成速率（个/s）
	prev time.Time // 上一次令牌发送时间
}

func NewTokenBucket(capacity int64, rate int64) *TokenBucket {
	bucket := &TokenBucket{
		rate: rate,
	}

	bucket.tokens.Store(capacity)

	go func(bucket *TokenBucket) {
		ticker := time.NewTicker(time.Second)

		for {
			select {
			case <-ticker.C:
				num := bucket.tokens.Load()
				if val := num + rate; val >= bucket.rate {
					bucket.tokens.Store(bucket.rate)
				} else {
					bucket.tokens.Store(val)
				}
			}
		}
	}(bucket)

	return bucket
}

func (t *TokenBucket) HasToken() bool {
	if t == nil {
		return false
	}

	for {
		tokenNum := t.tokens.Load()
		if tokenNum <= 0 {
			return false
		}

		if t.tokens.CompareAndSwap(tokenNum, tokenNum-1) {
			t.prev = time.Now()
			return true
		}
	}

	return false
}
