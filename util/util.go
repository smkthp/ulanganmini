package util

import (
	"encoding/json"
	"os"
	"os/exec"
	"runtime"

	"github.com/smkthp/ulanganmini/system"
)

func ClearTerminal() {
	osname := runtime.GOOS

	switch osname {
	case "linux", "darwin":
		cmd := exec.Command("clear")
        cmd.Stdout = os.Stdout
        cmd.Run()
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
        cmd.Stdout = os.Stdout
        cmd.Run()
	}
}

func unmarshalTasks(p []byte) ([]system.Task, error) {
	var tasks []system.Task

	if err := json.Unmarshal(p, &tasks); err != nil {
		return tasks, err
	}

	return tasks, nil
}
