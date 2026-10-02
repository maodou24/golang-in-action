package net

import (
	"testing"
)

func TestNetListenMulticastUDP(t *testing.T) {
	multicastAddr := "224.0.0.1:12345"

	go func() {
		err := JoinMulticastGroup("node-test", multicastAddr)
		if err != nil {
			t.Error(err)
		}
	}()

	err := SendMulticastGroupWaitUnicastResponse(multicastAddr)
	if err != nil {
		t.Error(err)
	}
}
