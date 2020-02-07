package cliparams

import "os"

// Get returns the n th parameter enter by the user
func Get(pos int) string {
	if pos+1 > len(os.Args) {
		return ""
	}

	return os.Args[pos]
}
