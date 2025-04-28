package basic

import (
	"testing"
)

func TestConvertToTypeT(t *testing.T) {
	m := map[string]any{
		"name":      "andy",
		"nilString": nil,
	}

	if "andy" != ConvertToTypeT[string](m["name"]) {
		t.Fail()
	}

	if "" != ConvertToTypeT[string](m["nilString"]) {
		t.Fail()
	}
}
