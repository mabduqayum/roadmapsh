package cmd

import (
	"fmt"
	"strconv"

	"github.com/mabduqayum/roadmapsh/01_task_tracker_vc/tracker"
	"github.com/spf13/cobra"
)

func NewDeleteCmd(taskTracker *tracker.TaskTracker) *cobra.Command {
	return &cobra.Command{
		Use:   "delete <id>",
		Short: "Delete a task",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			id, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Println("Invalid task ID")
				return
			}
			err = taskTracker.DeleteTask(id)
			if err != nil {
				fmt.Println("Error deleting task:", err)
				return
			}
			fmt.Println("Task deleted successfully")
		},
	}
}
