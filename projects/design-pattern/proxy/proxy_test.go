package proxy

import "testing"

func TestProxy(t *testing.T) {
	proxy := NewProxy()
	proxy.Do()
}
