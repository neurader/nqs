//go:build windows

package process

import (
	"os/exec"
	"strconv"
)

func Kill(pid int) error {
	cmd := exec.Command("taskkill", "/PID", strconv.Itoa(pid), "/F", "/T")
	return cmd.Run()
}
