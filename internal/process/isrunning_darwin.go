//go:build darwin

package process

import "syscall"

func IsRunning(pid int) bool {
	return syscall.Kill(pid, 0) == nil
}
