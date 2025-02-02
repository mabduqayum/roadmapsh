package main

import (
	"fmt"
	"os"

	"github.com/mabduqayum/roadmapsh/01_task_tracker/commands"
	"github.com/mabduqayum/roadmapsh/01_task_tracker/tracker"
)

func main() {
	if len(os.Args) < 2 || os.Args[1] == "--help" {
		commands.PrintHelp()
		return
	}

	command := os.Args[1]
	args := os.Args[2:]
	t := tracker.NewTaskTracker()

	if !commands.HandleCommand(t, command, args) {
		fmt.Printf("'%s' is not a Task Tracker command.\n", command)
	}
}
