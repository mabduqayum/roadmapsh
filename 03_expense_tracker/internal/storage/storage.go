package storage

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sync"

	"github.com/mabduqayum/roadmapsh/03_expense_tracker/internal/expense"
)

type Storage struct {
	file  string
	mutex sync.Mutex
}

func NewStorage(file string) *Storage {
	dir := filepath.Dir(file)
	if err := os.MkdirAll(dir, 0755); err != nil {
		panic(err)
	}
	return &Storage{file: file}
}

func (s *Storage) Load() ([]expense.Expense, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	data, err := os.ReadFile(s.file)
	if err != nil {
		if os.IsNotExist(err) {
			return []expense.Expense{}, nil
		}
		return nil, err
	}

	var expenses []expense.Expense
	err = json.Unmarshal(data, &expenses)
	return expenses, err
}

func (s *Storage) Save(expenses []expense.Expense) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	data, err := json.MarshalIndent(expenses, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(s.file, data, 0644)
}
