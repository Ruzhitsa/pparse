package path

import "os"

// GetGoPath returns the name of then path which contains the go project
func GetGoPath() string {
	home, _ := os.UserHomeDir()

	return home + "/go/src/"
}

// CheckIfPathExists returns true if the given path exists
func CheckIfPathExists(path string) bool {
	_, e := os.Stat(path)

	return e == nil
}
