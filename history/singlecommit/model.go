package singlecommit

import "time"

// SingleHistory contains a commit data from git log
type SingleHistory struct {
	Commit       string
	Author       string
	Date         string
	InternalDate time.Time
	Comments     string
	Tag          string
}
