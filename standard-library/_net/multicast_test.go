package _net

import (
	"fmt"
	"testing"
)

func TestMulticast(t *testing.T) {
	multicastAddr := "224.0.0.1:12345"

	go func() {
		err := JoinMulticastGroup(fmt.Sprintf("node-test"), multicastAddr)
		if err != nil {
			t.Error(err)
		}
	}()

	err := SendMulticastGroupWaitUnicastResponse(multicastAddr)
	if err != nil {
		t.Error(err)
	}
}
