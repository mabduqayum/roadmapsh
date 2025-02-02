package storage

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/mabduqayum/roadmapsh/01_task_tracker_vc/models"
)

//type TaskStorage interface {
//	LoadTasks() models.TaskList
//	SaveTasks(tasks models.TaskList) error
//}

type Storage struct {
	file string
}

func NewStorage(file string) *Storage {
	dir := filepath.Dir(file)
	if err := os.MkdirAll(dir, 0755); err != nil {
		panic(err)
	}
	return &Storage{file: file}
}

func (s *Storage) Load() (*models.TaskList, error) {
	data, err := os.ReadFile(s.file)
	if err != nil {
		if os.IsNotExist(err) {
			return &models.TaskList{}, nil
		}
		return nil, err
	}

	var taskList *models.TaskList
	err = json.Unmarshal(data, &taskList)
	return taskList, err
}

func (s *Storage) Save(taskList *models.TaskList) error {
	data, err := json.MarshalIndent(taskList, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(s.file, data, 0644)
}
