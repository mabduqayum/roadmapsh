package cmd

import (
	"fmt"
	"strings"

	"github.com/mabduqayum/roadmapsh/01_task_tracker_vc/tracker"
	"github.com/spf13/cobra"
)

func NewAddCmd(taskTracker *tracker.TaskTracker) *cobra.Command {
	return &cobra.Command{
		Use:   "add <description>",
		Short: "Add a new task",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			description := strings.Join(args, " ")
			id, err := taskTracker.AddTask(description)
			if err != nil {
				fmt.Println("Error adding task:", err)
				return
			}
			fmt.Printf("Task added successfully (ID: %d)\n", id)
		},
	}
}
