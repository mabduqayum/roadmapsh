package storage

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/mabduqayum/roadmapsh/01_task_tracker_uk/models"
)

type TaskStorage interface {
	LoadTasks() models.TaskList
	SaveTasks(tasks models.TaskList) error
}

type FileStorage struct {
	trackerDir string
	tasksFile  string
}

func NewFileStorage(trackerDir string, taskFile string) *FileStorage {
	os.MkdirAll(trackerDir, os.ModePerm)
	return &FileStorage{
		trackerDir: trackerDir,
		tasksFile:  taskFile,
	}
}

func (fs *FileStorage) LoadTasks() models.TaskList {
	var taskList models.TaskList
	data, err := os.ReadFile(filepath.Join(fs.trackerDir, fs.tasksFile))
	if err != nil {
		if os.IsNotExist(err) {
			return taskList
		}
		panic(err)
	}
	if err := json.Unmarshal(data, &taskList); err != nil {
		panic(err)
	}
	return taskList
}

func (fs *FileStorage) SaveTasks(taskList models.TaskList) error {
	data, err := json.MarshalIndent(taskList, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(fs.trackerDir, fs.tasksFile), data, 0644)
}
