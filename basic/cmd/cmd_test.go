package cmd

import (
	"fmt"
	"os/exec"
	"testing"
)

// 执行交互式命令，代码代替手动输入
func TestExecCommand(t *testing.T) {
	cmd := exec.Command("bash", "recursion.sh")

	stdinPipe, err := cmd.StdinPipe()
	if err != nil {
		return
	}

	if err := cmd.Start(); err != nil {
		return
	}

	go func() {
		defer stdinPipe.Close()
		_, _ = stdinPipe.Write([]byte("10"))
	}()

	_ = cmd.Wait()

	output, err := cmd.CombinedOutput()
	if err != nil {
		return
	}
	fmt.Println(string(output))
}
