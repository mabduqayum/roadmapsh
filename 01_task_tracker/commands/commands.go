package commands

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mabduqayum/roadmapsh/01_task_tracker/tracker"
)

var commandDescriptions = map[string]struct {
	args string
	desc string
}{
	"add":              {"<description>", "Add a new task"},
	"update":           {"<id> <description>", "Update a task"},
	"delete":           {"<id>", "Delete a task"},
	"mark-in-progress": {"<id>", "Mark a task as in progress"},
	"mark-done":        {"<id>", "Mark a task as done"},
	"list":             {"[status]", "List all tasks or tasks with specific status"},
}

func HandleCommand(t *tracker.TaskTracker, command string, args []string) bool {
	switch command {
	case "add":
		handleAdd(t, args)
	case "update":
		handleUpdate(t, args)
	case "delete":
		handleDelete(t, args)
	case "mark-in-progress":
		handleMarkStatus(t, args, "in-progress")
	case "mark-done":
		handleMarkStatus(t, args, "done")
	case "list":
		handleList(t, args)
	default:
		return false
	}
	return true
}

func handleAdd(t *tracker.TaskTracker, args []string) {
	if len(args) == 0 {
		fmt.Println("Task description was not provided.")
		return
	}
	description := strings.Join(args, " ")
	id, err := t.AddTask(description)
	if err != nil {
		fmt.Println("Error adding task:", err)
		return
	}
	fmt.Printf("Task added successfully (ID: %d)\n", id)
}

func handleUpdate(t *tracker.TaskTracker, args []string) {
	if len(args) < 2 {
		fmt.Println("Invalid arguments. Usage: update <id> <new description>")
		return
	}
	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Invalid task ID")
		return
	}
	description := strings.Join(args[1:], " ")
	err = t.UpdateTask(id, description)
	if err != nil {
		fmt.Println("Error updating task:", err)
		return
	}
	fmt.Println("Task updated successfully")
}

func handleDelete(t *tracker.TaskTracker, args []string) {
	if len(args) != 1 {
		fmt.Println("Invalid arguments. Usage: delete <id>")
		return
	}
	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Invalid task ID")
		return
	}
	err = t.DeleteTask(id)
	if err != nil {
		fmt.Println("Error deleting task:", err)
		return
	}
	fmt.Println("Task deleted successfully")
}

func handleMarkStatus(t *tracker.TaskTracker, args []string, status string) {
	if len(args) != 1 {
		fmt.Printf("Invalid arguments. Usage: mark-%s <id>\n", status)
		return
	}
	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Invalid task ID")
		return
	}
	err = t.MarkTaskStatus(id, status)
	if err != nil {
		fmt.Println("Error marking task:", err)
		return
	}
	fmt.Printf("Task marked as %s successfully\n", status)
}

func handleList(t *tracker.TaskTracker, args []string) {
	var status string
	if len(args) > 0 {
		status = args[0]
	}
	tasks := t.ListTasks(status)
	if len(tasks) == 0 {
		fmt.Println("No tasks found")
		return
	}
	for _, task := range tasks {
		task.Print()
	}
}

func PrintHelp() {
	fmt.Println("These are Task Tracker commands:")

	// Find the maximum lengths for each column
	maxCmdLen, maxArgsLen := 0, 0
	for cmd, info := range commandDescriptions {
		if len(cmd) > maxCmdLen {
			maxCmdLen = len(cmd)
		}
		if len(info.args) > maxArgsLen {
			maxArgsLen = len(info.args)
		}
	}

	// Print each command with aligned columns
	for cmd, info := range commandDescriptions {
		fmt.Printf("  %-*s  %-*s  %s\n", maxCmdLen, cmd, maxArgsLen, info.args, info.desc)
	}
}
