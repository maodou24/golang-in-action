package go1_24

import (
	"strings"
	"testing"
)

func TestRootReadFile(t *testing.T) {
	err := ReadFile("./testdata/a.txt")
	if err != nil {
		t.Fatal(err)
	}
}

func TestRootReadFileOutDir(t *testing.T) {
	err := ReadFile("/root/testdata/a.txt") // path escapes
	if err == nil || !strings.Contains(err.Error(), "path escapes from parent") {
		t.Fatal(err)
	}
}
