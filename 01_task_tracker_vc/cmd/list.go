package cmd

import (
	"fmt"

	"github.com/mabduqayum/roadmapsh/01_task_tracker_vc/tracker"
	"github.com/spf13/cobra"
)

func NewListCmd(taskTracker *tracker.TaskTracker) *cobra.Command {
	return &cobra.Command{
		Use:   "list [status]",
		Short: "List all tasks or tasks with specific status",
		Args:  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			var status string
			if len(args) > 0 {
				status = args[0]
			}
			tasks := taskTracker.ListTasks(status)
			if len(tasks) == 0 {
				fmt.Println("No tasks found")
				return
			}
			for _, task := range tasks {
				task.Print()
			}
		},
	}
}
