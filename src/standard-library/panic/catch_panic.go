package _panic

import (
	"os"
	"runtime/debug"
)

func SetGlobalCatchPanic() {
	f, _ := os.Create("panic.txt")
	_ = debug.SetCrashOutput(f, debug.CrashOptions{})
}
