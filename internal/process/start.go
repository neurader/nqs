package process

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"syscall"

	"github.com/neurader/nqs/internal/state"
)

func buildCommand(cmd []string) *exec.Cmd {
	if runtime.GOOS == "windows" {
		return exec.Command("cmd", "/C", strings.Join(cmd, " "))
	}
	return exec.Command(cmd[0], cmd[1:]...)
}

func Start(name string, command []string) {
	data := state.Load()

	if _, exists := data[name]; exists {
		fmt.Println("❌ Name already exists:", name)
		return
	}

	cmd := buildCommand(command)

	logFile, err := os.Create(name + ".log")
	if err != nil {
		fmt.Println("Error creating log file:", err)
		return
	}

	cmd.Stdout = logFile
	cmd.Stderr = logFile

	err = cmd.Start()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	data[name] = state.Process{
		Name:    name,
		Pid:     cmd.Process.Pid,
		Command: strings.Join(command, " "),
	}

	state.Save(data)

	fmt.Println("✅ Started:", name, "PID:", cmd.Process.Pid)
}
