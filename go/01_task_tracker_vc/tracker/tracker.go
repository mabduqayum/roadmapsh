package tracker

import (
	"fmt"

	"github.com/mabduqayum/roadmapsh/01_task_tracker_vc/models"
	"github.com/mabduqayum/roadmapsh/01_task_tracker_vc/storage"
)

type TaskTracker struct {
	Storage *storage.Storage
}

func NewTaskTracker(file string) *TaskTracker {
	return &TaskTracker{
		Storage: storage.NewStorage(file),
	}
}

func (t *TaskTracker) AddTask(description string) (int, error) {
	tasks, err := t.Storage.Load()
	if err != nil {
		return 0, err
	}
	newTaskId := t.LastId() + 1
	newTask := models.NewTask(description, newTaskId)
	tasks.Tasks = append(tasks.Tasks, newTask)
	err = t.Storage.Save(tasks)
	return newTask.ID, err
}

func (t *TaskTracker) UpdateTask(id int, description string) error {
	tasks, err := t.Storage.Load()
	if err != nil {
		return err
	}
	for i, task := range tasks.Tasks {
		if task.ID == id {
			tasks.Tasks[i].UpdateDescription(description)
			return t.Storage.Save(tasks)
		}
	}
	return fmt.Errorf("task not found")
}

func (t *TaskTracker) DeleteTask(id int) error {
	tasks, err := t.Storage.Load()
	if err != nil {
		return err
	}
	for i, task := range tasks.Tasks {
		if task.ID == id {
			tasks.Tasks = append(tasks.Tasks[:i], tasks.Tasks[i+1:]...)
			return t.Storage.Save(tasks)
		}
	}
	return fmt.Errorf("task not found")
}

func (t *TaskTracker) MarkTaskStatus(id int, status string) error {
	tasks, err := t.Storage.Load()
	if err != nil {
		return err
	}
	for i, task := range tasks.Tasks {
		if task.ID == id {
			tasks.Tasks[i].UpdateStatus(status)
			return t.Storage.Save(tasks)
		}
	}
	return fmt.Errorf("task not found")
}

func (t *TaskTracker) ListTasks(status string) []models.Task {
	tasks, _ := t.Storage.Load()
	if status == "" {
		return tasks.Tasks
	}
	var filteredTasks []models.Task
	for _, task := range tasks.Tasks {
		if task.Status == status {
			filteredTasks = append(filteredTasks, task)
		}
	}
	return filteredTasks
}

func (t *TaskTracker) LastId() int {
	tasks, _ := t.Storage.Load()
	if len(tasks.Tasks) == 0 {
		return 0
	}
	return tasks.Tasks[len(tasks.Tasks)-1].ID
}
