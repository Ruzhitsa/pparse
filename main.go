package main

import (
	"fmt"
	"os"

	"github.com/ruze/pparse/cliparams"
	"github.com/ruze/pparse/history"
)

func main() {
	projectName := cliparams.Get(1)
	ticketNr := cliparams.Get(2)

	if projectName == "" {
		fmt.Println("project name is missing")
		os.Exit(-1)
	}

	if ticketNr == "" {
		fmt.Println("ticket nr is missing")
		os.Exit(-1)
	}

	hist := history.NewHistory()
	hist.Search(projectName, ticketNr)

	fmt.Println("done")
}
