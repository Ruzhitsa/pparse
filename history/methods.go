package history

import (
	"fmt"
	"os"
	"strings"

	"github.com/ruze/pparse/history/singlecommit"
	"github.com/ruze/pparse/path"
	"github.com/ruze/pparse/shell"
)

// NewHistory builds and returns a new History
func NewHistory() *History {
	return &History{}
}

// Search is looking for ticketNr in the commits of the project
func (h History) Search(projectName, ticketNr string) {
	projectPath := path.GetGoPath() + projectName

	// Check if path exists
	if !path.CheckIfPathExists(projectPath) {
		fmt.Printf("project path '%s' not found!\n", projectPath)
		os.Exit(-1)
	}

	// Read commits data
	h.Read(projectPath)

	// Search for ticketNr
	for _, history := range h.List {
		if strings.Contains(history.Comments, ticketNr) {
			fmt.Printf("found ticket in commit %s (pushed on %s by %s)\n", history.Commit, history.Date, history.Author)
		}
	}
}

// Read the git logs of the project
func (h *History) Read(projectPath string) {
	// Read Git Log
	result := shell.ExecuteCommand(projectPath, "git", "log")

	// Parse result
	h.parse(result)
}

func (h *History) parse(res string) {
	list := strings.Split(res, "\n")
	l := len(list)

	startIndex := -1

	for i, line := range list {
		// Handle commit
		if strings.Index(line, "commit") == 0 {
			if startIndex > -1 {
				h.parseCommit(list[startIndex : i-1])
			}
			startIndex = i
		}

		// Solve last commit
		if i == l-1 && startIndex > -1 {
			h.parseCommit(list[startIndex:i])
		}
	}
}

func (h *History) parseCommit(commit []string) {
	// Create a new single history
	s := singlecommit.NewCommit()

	// Fill single history with data
	for _, line := range commit {
		// Commit line
		if strings.Index(line, "commit") == 0 {
			s.SetCommitHash(line)
			continue
		}

		// Author line
		if strings.Index(line, "Author") == 0 {
			s.SetAuthor(line)
			continue
		}

		// Date line
		if strings.Index(line, "Date") == 0 {
			s.SetDate(line)
			continue
		}

		// Comment line
		if strings.Index(line, " ") == 0 {
			s.SetComment(line)
		}
	}

	// Add the single history in the list
	h.List = append(h.List, s)
}
