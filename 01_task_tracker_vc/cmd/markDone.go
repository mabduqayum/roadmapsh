package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var markDoneCmd = &cobra.Command{
	Use:   "mark-done <id>",
	Short: "Mark a task as done",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid task ID")
			return
		}
		err = taskTracker.MarkTaskStatus(id, "done")
		if err != nil {
			fmt.Println("Error marking task:", err)
			return
		}
		fmt.Println("Task marked as done successfully")
	},
}

func init() {
	rootCmd.AddCommand(markDoneCmd)
}
