package process

import (
	"fmt"

	"github.com/neurader/nqs/internal/state"
)

func Delete(name string) {
	data := state.Load()

	p, exists := data[name]
	if !exists {
		fmt.Println("❌ Not found:", name)
		return
	}

	err := Kill(p.Pid)
	if err != nil {
		fmt.Println("Error killing process:", err)
	}

	delete(data, name)
	state.Save(data)

	fmt.Println("🗑️ Deleted:", name)
}
