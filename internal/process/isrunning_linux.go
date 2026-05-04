//go:build linux

package process

import "syscall"

func IsRunning(pid int) bool {
	return syscall.Kill(pid, 0) == nil
}
