package storage

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sync"

	"github.com/mabduqayum/roadmapsh/03_expense_tracker/internal/transaction"
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

func (s *Storage) Load() ([]transaction.Transaction, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	data, err := os.ReadFile(s.file)
	if err != nil {
		if os.IsNotExist(err) {
			return []transaction.Transaction{}, nil
		}
		return nil, err
	}

	var transactions []transaction.Transaction
	err = json.Unmarshal(data, &transactions)
	return transactions, err
}

func (s *Storage) Save(transactions []transaction.Transaction) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	data, err := json.MarshalIndent(transactions, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(s.file, data, 0644)
}
