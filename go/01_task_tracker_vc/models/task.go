package models

import (
	"fmt"
	"time"
)

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type TaskList struct {
	Tasks []Task `json:"tasks"`
}

func NewTask(description string, id int) Task {
	return Task{
		ID:          id,
		Description: description,
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func (t *Task) UpdateDescription(description string) {
	t.Description = description
	t.UpdatedAt = time.Now()
}

func (t *Task) UpdateStatus(status string) {
	t.Status = status
	t.UpdatedAt = time.Now()
}

func (t *Task) Print() {
	fmt.Printf("[%d] %s (Status: %s)\n", t.ID, t.Description, t.Status)
}
