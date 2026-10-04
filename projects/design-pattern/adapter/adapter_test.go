package adapter

import (
	"testing"
)

func TestBuilder(t *testing.T) {
	adaptee := NewAdaptee()
	adapter := NewAdapter(adaptee)
	if adapter.Play() != "PlayMP3" {
		t.Error("Expected PlayMP3")
	}
}
