package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mabduqayum/roadmapsh/01_task_tracker_vc/tracker"
	"github.com/spf13/cobra"
)

func NewUpdateCmd(taskTracker *tracker.TaskTracker) *cobra.Command {
	return &cobra.Command{
		Use:   "update <id> <new description>",
		Short: "Update a task",
		Args:  cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			id, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Println("Invalid task ID")
				return
			}
			description := strings.Join(args[1:], " ")
			err = taskTracker.UpdateTask(id, description)
			if err != nil {
				fmt.Println("Error updating task:", err)
				return
			}
			fmt.Println("Task updated successfully")
		},
	}
}
