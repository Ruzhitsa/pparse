package shell

import (
	"log"
	"os"
	"os/exec"
)

// ExecuteCommand is a generic function to execute command in the shell
func ExecuteCommand(path, commandName, param string) string {
	if path != "" {
		// Change to project path
		err := os.Chdir(path)
		if err != nil {
			log.Println(err)
			return ""
		}
	}

	// Execute git log
	cmd := exec.Command(commandName, param)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}

	return string(out)
}
