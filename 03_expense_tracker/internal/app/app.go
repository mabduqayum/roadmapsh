package app

import (
	"errors"
	"fmt"
	"time"

	"github.com/mabduqayum/roadmapsh/03_expense_tracker/internal/expense"
	"github.com/mabduqayum/roadmapsh/03_expense_tracker/internal/storage"
)

type App struct {
	storage *storage.Storage
}

func NewApp(storageFile string) *App {
	return &App{
		storage: storage.NewStorage(storageFile),
	}
}

func (a *App) AddExpense(description string, amount float64) error {
	expenses, err := a.storage.Load()
	if err != nil {
		return err
	}

	newID := 1
	if len(expenses) > 0 {
		newID = expenses[len(expenses)-1].ID + 1
	}

	newExpense := expense.Expense{
		ID:          newID,
		Date:        time.Now(),
		Description: description,
		Amount:      amount,
	}

	expenses = append(expenses, newExpense)
	err = a.storage.Save(expenses)
	if err != nil {
		return err
	}

	fmt.Printf("Expense added successfully (ID: %d)\n", newID)
	return nil
}

func (a *App) UpdateExpense(id int, description string, amount float64) error {
	expenses, err := a.storage.Load()
	if err != nil {
		return err
	}

	for i, e := range expenses {
		if e.ID == id {
			expenses[i].Description = description
			expenses[i].Amount = amount
			err = a.storage.Save(expenses)
			if err != nil {
				return err
			}
			fmt.Printf("Expense updated successfully (ID: %d)\n", id)
			return nil
		}
	}

	return errors.New("expense not found")
}

func (a *App) DeleteExpense(id int) error {
	expenses, err := a.storage.Load()
	if err != nil {
		return err
	}

	for i, e := range expenses {
		if e.ID == id {
			expenses = append(expenses[:i], expenses[i+1:]...)
			err = a.storage.Save(expenses)
			if err != nil {
				return err
			}
			fmt.Printf("Expense deleted successfully (ID: %d)\n", id)
			return nil
		}
	}

	return errors.New("expense not found")
}

func (a *App) ListExpenses() error {
	expenses, err := a.storage.Load()
	if err != nil {
		return err
	}

	fmt.Println("ID  Date       Description  Amount")
	for _, e := range expenses {
		fmt.Printf("%d   %s  %-12s $%.2f\n", e.ID, e.Date.Format("2006-01-02"), e.Description, e.Amount)
	}

	return nil
}

func (a *App) Summary() error {
	expenses, err := a.storage.Load()
	if err != nil {
		return err
	}

	total := 0.0
	for _, e := range expenses {
		total += e.Amount
	}

	fmt.Printf("Total expenses: $%.2f\n", total)
	return nil
}

func (a *App) MonthlySummary(month int) error {
	expenses, err := a.storage.Load()
	if err != nil {
		return err
	}

	total := 0.0
	currentYear := time.Now().Year()
	for _, e := range expenses {
		if e.Date.Month() == time.Month(month) && e.Date.Year() == currentYear {
			total += e.Amount
		}
	}

	fmt.Printf("Total expenses for %s: $%.2f\n", time.Month(month), total)
	return nil
}
