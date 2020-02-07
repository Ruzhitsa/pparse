package history

import "github.com/ruze/pparse/history/singlecommit"

// History contains all commits from git log
type History struct {
	List []singlecommit.SingleHistory
}
