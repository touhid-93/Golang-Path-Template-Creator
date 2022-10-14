package pathhelper

import (
	"os"
)

// Represents the exact path to the executable file
// Exact path to the executable not the path but the file.
func GetExecutablePath() string {
	exe, err := os.Executable()

	if err != nil {
		panic(err)
	}

	return exe
}
