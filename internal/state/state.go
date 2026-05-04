package state

import (
	"encoding/json"
	"os"
)

type Process struct {
	Name    string `json:"name"`
	Pid     int    `json:"pid"`
	Command string `json:"command"`
}

var file = "state.json"

func Load() map[string]Process {
	data := make(map[string]Process)

	bytes, err := os.ReadFile(file)
	if err != nil {
		return data
	}

	_ = json.Unmarshal(bytes, &data)
	return data
}

func Save(data map[string]Process) {
	bytes, _ := json.MarshalIndent(data, "", "  ")
	_ = os.WriteFile(file, bytes, 0644)
}
