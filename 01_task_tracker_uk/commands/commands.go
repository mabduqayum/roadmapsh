package commands

import (
	"fmt"
	"strings"

	"github.com/mabduqayum/roadmapsh/01_task_tracker_uk/tracker"
	"github.com/urfave/cli/v2"
)

func AddCommand(t *tracker.TaskTracker) *cli.Command {
	return &cli.Command{
		Name:  "add",
		Usage: "Add a new task",
		Action: func(c *cli.Context) error {
			description := strings.Join(c.Args().Slice(), " ")
			id, err := t.AddTask(description)
			if err != nil {
				return fmt.Errorf("error adding task: %w", err)
			}
			fmt.Printf("Task added successfully (ID: %d)\n", id)
			return nil
		},
	}
}

func UpdateCommand(t *tracker.TaskTracker) *cli.Command {
	return &cli.Command{
		Name:  "update",
		Usage: "Update a task",
		Flags: []cli.Flag{
			&cli.IntFlag{Name: "id", Required: true},
		},
		Action: func(c *cli.Context) error {
			id := c.Int("id")
			description := strings.Join(c.Args().Slice(), " ")
			err := t.UpdateTask(id, description)
			if err != nil {
				return fmt.Errorf("error updating task: %w", err)
			}
			fmt.Println("Task updated successfully")
			return nil
		},
	}
}

func DeleteCommand(t *tracker.TaskTracker) *cli.Command {
	return &cli.Command{
		Name:  "delete",
		Usage: "Delete a task",
		Flags: []cli.Flag{
			&cli.IntFlag{Name: "id", Required: true},
		},
		Action: func(c *cli.Context) error {
			id := c.Int("id")
			err := t.DeleteTask(id)
			if err != nil {
				return fmt.Errorf("error deleting task: %w", err)
			}
			fmt.Println("Task deleted successfully")
			return nil
		},
	}
}

func MarkInProgressCommand(t *tracker.TaskTracker) *cli.Command {
	return &cli.Command{
		Name:  "mark-in-progress",
		Usage: "Mark a task as in progress",
		Flags: []cli.Flag{
			&cli.IntFlag{Name: "id", Required: true},
		},
		Action: func(c *cli.Context) error {
			id := c.Int("id")
			err := t.MarkTaskStatus(id, "in-progress")
			if err != nil {
				return fmt.Errorf("error marking task as in progress: %w", err)
			}
			fmt.Println("Task marked as in progress successfully")
			return nil
		},
	}
}

func MarkDoneCommand(t *tracker.TaskTracker) *cli.Command {
	return &cli.Command{
		Name:  "mark-done",
		Usage: "Mark a task as done",
		Flags: []cli.Flag{
			&cli.IntFlag{Name: "id", Required: true},
		},
		Action: func(c *cli.Context) error {
			id := c.Int("id")
			err := t.MarkTaskStatus(id, "done")
			if err != nil {
				return fmt.Errorf("error marking task as done: %w", err)
			}
			fmt.Println("Task marked as done successfully")
			return nil
		},
	}
}

func ListCommand(t *tracker.TaskTracker) *cli.Command {
	return &cli.Command{
		Name:  "list",
		Usage: "List tasks",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "status"},
		},
		Action: func(c *cli.Context) error {
			status := c.String("status")
			tasks := t.ListTasks(status)
			if len(tasks) == 0 {
				fmt.Println("No tasks found")
				return nil
			}
			for _, task := range tasks {
				task.Print()
			}
			return nil
		},
	}
}
