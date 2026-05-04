//go:build linux

package process

import "syscall"

func Kill(pid int) error {
	return syscall.Kill(-pid, syscall.SIGKILL)
}
