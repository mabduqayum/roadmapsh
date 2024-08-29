package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var markInProgressCmd = &cobra.Command{
	Use:   "mark-in-progress <id>",
	Short: "Mark a task as in progress",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid task ID")
			return
		}
		err = taskTracker.MarkTaskStatus(id, "in-progress")
		if err != nil {
			fmt.Println("Error marking task:", err)
			return
		}
		fmt.Println("Task marked as in-progress successfully")
	},
}

func init() {
	rootCmd.AddCommand(markInProgressCmd)
}