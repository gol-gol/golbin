package golbin

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

/*
If a command is present in available system path.
*/
func IsSystemCmd(cmd string) bool {
	_, err := os.Stat(cmd)
	if err == nil {
		return true
	}

	sysPath := os.Getenv("PATH")

	pathSeparator := ":"
	if runtime.GOOS == "windows" {
		pathSeparator = ";"
	}

	for _, e := range strings.Split(sysPath, pathSeparator) {
		cmdPath := fmt.Sprintf("%s%s%s", e, string(os.PathSeparator), cmd)
		_, err := os.Stat(cmdPath)
		if err == nil {
			return true
		}
	}
	return false
}
