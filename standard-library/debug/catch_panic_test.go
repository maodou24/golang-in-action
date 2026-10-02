package _panic

import (
	"fmt"
	"os"
	"runtime/debug"
	"testing"
)

func TestDebugSetCrashOutput(t *testing.T) {
	f, _ := os.Create("panic.log")
	_ = debug.SetCrashOutput(f, debug.CrashOptions{})

	defer func() {
		err := recover()
		fmt.Printf("recover goroutine panic: %v\n", err)
	}()
	panic("test panic")
}

// If a panic occurs and it not captured, the crash log will be written to the file specified by SetCrashOutput.
func TestDebugSetCrashOutputCrash(t *testing.T) {
	f, _ := os.Create("panic.log")
	_ = debug.SetCrashOutput(f, debug.CrashOptions{})

	panic("test panic")
}
