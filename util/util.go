package util

import (
	"os"
	"os/exec"
	"runtime"
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
