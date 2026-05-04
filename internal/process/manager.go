package process

import (
	"fmt"

	"github.com/neurader/nqs/internal/state"
)

func List() {
	data := state.Load()

	fmt.Println("📦 NQS Dashboard")
	fmt.Println()

	if len(data) == 0 {
		fmt.Println("No processes running")
		return
	}

	for name, p := range data {
		status := "🟢 running"
		if !IsRunning(p.Pid) {
			status = "🔴 stopped"
		}
		fmt.Printf("[%s] %s (PID: %d)\n", status, name, p.Pid)
	}
}
