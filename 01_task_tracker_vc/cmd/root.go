package cmd

import (
	"github.com/mabduqayum/roadmapsh/01_task_tracker_vc/tracker"
	"github.com/spf13/cobra"
)

func NewRootCmd(file string) *cobra.Command {
	taskTracker := tracker.NewTaskTracker(file)

	cmd := &cobra.Command{
		Use:   "task-cli",
		Short: "A simple CLI task tracker",
		Long:  `Task-cli is a simple command line interface to track and manage your tasks.`,
	}

	cmd.AddCommand(NewAddCmd(taskTracker))
	cmd.AddCommand(NewListCmd(taskTracker))
	cmd.AddCommand(NewDeleteCmd(taskTracker))
	cmd.AddCommand(NewUpdateCmd(taskTracker))
	cmd.AddCommand(NewMarkDoneCmd(taskTracker))
	cmd.AddCommand(NewMarkInProgressCmd(taskTracker))

	return cmd
}
