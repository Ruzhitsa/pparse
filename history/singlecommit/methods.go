package singlecommit

import (
	"strings"
)

// NewCommit builds and returns a new Commit
func NewCommit() SingleHistory {
	return SingleHistory{}
}

// SetCommitHash reads the hash in the received line
func (s *SingleHistory) SetCommitHash(line string) {
	pos := strings.Index(line, " ")
	if pos > -1 {
		line = strings.Trim(line[pos+1:], " ")

		pos = strings.Index(line, " ")
		if pos == -1 {
			s.Commit = line
		} else {
			s.Commit = strings.Trim(line[:pos], " ")
		}
	}
}

// SetAuthor reads the author in the received line
func (s *SingleHistory) SetAuthor(line string) {
	pos := strings.Index(line, " ")
	if pos > -1 {
		s.Author = strings.Trim(line[pos+1:], " ")
	}
}

// SetDate reads the date in the received line
func (s *SingleHistory) SetDate(line string) {
	pos := strings.Index(line, " ")
	if pos > -1 {
		s.Date = strings.Trim(line[pos+1:], " ")
		s.SetInternalDate()
	}
}

// SetComment reads the comment in the received line
func (s *SingleHistory) SetComment(line string) {
	pos := strings.Index(line, " ")
	if pos > -1 {
		comment := strings.Trim(line[pos+1:], " ")

		if comment == "" {
			return
		}

		if s.Comments == "" {
			s.Comments = comment
			return
		}

		s.Comments = s.Comments + "\n" + comment
	}
}

// SetInternalDate transfroms the date into time.Time
func (s *SingleHistory) SetInternalDate() {
	// s.Date -> Fri Feb 7 08:08:46 2020 +0100
	// s.InternalDate = ...
}
