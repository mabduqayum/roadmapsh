package main

import (
	"fmt"
	"log"
	"os"
	"task_tracker_uk/commands"
	"task_tracker_uk/config"
	"task_tracker_uk/tracker"

	"github.com/urfave/cli/v2"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	t := tracker.NewTaskTracker(cfg)

	app := &cli.App{
		Name:  "task-cli",
		Usage: "A simple CLI task tracker",
		Commands: []*cli.Command{
			commands.AddCommand(t),
			commands.UpdateCommand(t),
			commands.DeleteCommand(t),
			commands.MarkInProgressCommand(t),
			commands.MarkDoneCommand(t),
			commands.ListCommand(t),
		},
	}

	err = app.Run(os.Args)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
