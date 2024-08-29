package app

import (
	"errors"
	"fmt"
	"time"

	"github.com/fatih/color"
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

func (a *App) AddTransaction(description string, amount float64, transactionType expense.TransactionType) error {
	transactions, err := a.storage.Load()
	if err != nil {
		return err
	}

	newID := 1
	if len(transactions) > 0 {
		newID = transactions[len(transactions)-1].ID + 1
	}

	newTransaction := expense.Transaction{
		ID:          newID,
		Date:        time.Now(),
		Description: description,
		Amount:      amount,
		Type:        transactionType,
	}

	transactions = append(transactions, newTransaction)
	err = a.storage.Save(transactions)
	if err != nil {
		return err
	}

	fmt.Printf("%s added successfully (ID: %d)\n", transactionType, newID)
	return nil
}

func (a *App) UpdateTransaction(id int, description string, amount float64) error {
	transactions, err := a.storage.Load()
	if err != nil {
		return err
	}

	for i, t := range transactions {
		if t.ID == id {
			transactions[i].Description = description
			transactions[i].Amount = amount
			err = a.storage.Save(transactions)
			if err != nil {
				return err
			}
			fmt.Printf("Transaction updated successfully (ID: %d)\n", id)
			return nil
		}
	}

	return errors.New("transaction not found")
}

func (a *App) DeleteTransaction(id int) error {
	transactions, err := a.storage.Load()
	if err != nil {
		return err
	}

	for i, t := range transactions {
		if t.ID == id {
			transactions = append(transactions[:i], transactions[i+1:]...)
			err = a.storage.Save(transactions)
			if err != nil {
				return err
			}
			fmt.Printf("Transaction deleted successfully (ID: %d)\n", id)
			return nil
		}
	}

	return errors.New("transaction not found")
}

func (a *App) ListTransactions() error {
	transactions, err := a.storage.Load()
	if err != nil {
		return err
	}

	fmt.Println("ID  Date       Type          Description  Amount")
	for _, t := range transactions {
		amount := fmt.Sprintf("$%.2f", t.Amount)
		switch t.Type {
		case expense.TypeExpense:
			color.Red("%d   %s  %-8s  %-12s %s\n", t.ID, t.Date.Format("2006-01-02"), t.Type, t.Description, amount)
		case expense.TypeTopUp:
			color.Green("%d   %s  %-8s  %-12s %s\n", t.ID, t.Date.Format("2006-01-02"), t.Type, t.Description, amount)
		case expense.TypeTransfer:
			color.Yellow("%d   %s  %-8s  %-12s %s\n", t.ID, t.Date.Format("2006-01-02"), t.Type, t.Description, amount)
		}
	}

	return nil
}

func (a *App) Summary() error {
	transactions, err := a.storage.Load()
	if err != nil {
		return err
	}

	var totalExpense, totalTopUp, totalTransfer float64
	for _, t := range transactions {
		switch t.Type {
		case expense.TypeExpense:
			totalExpense += t.Amount
		case expense.TypeTopUp:
			totalTopUp += t.Amount
		case expense.TypeTransfer:
			totalTransfer += t.Amount
		}
	}

	balance := totalTopUp - totalExpense - totalTransfer

	color.Red("Total expenses: $%.2f\n", totalExpense)
	color.Green("Total top-ups: $%.2f\n", totalTopUp)
	color.Yellow("Total transfers: $%.2f\n", totalTransfer)
	fmt.Printf("Current balance: $%.2f\n", balance)
	return nil
}

func (a *App) MonthlySummary(month int) error {
	transactions, err := a.storage.Load()
	if err != nil {
		return err
	}

	var totalExpense, totalTopUp, totalTransfer float64
	currentYear := time.Now().Year()
	for _, t := range transactions {
		if t.Date.Month() == time.Month(month) && t.Date.Year() == currentYear {
			switch t.Type {
			case expense.TypeExpense:
				totalExpense += t.Amount
			case expense.TypeTopUp:
				totalTopUp += t.Amount
			case expense.TypeTransfer:
				totalTransfer += t.Amount
			}
		}
	}

	balance := totalTopUp - totalExpense - totalTransfer

	fmt.Printf("Summary for %s:\n", time.Month(month))
	color.Red("Total expenses: $%.2f\n", totalExpense)
	color.Green("Total top-ups: $%.2f\n", totalTopUp)
	color.Yellow("Total transfers: $%.2f\n", totalTransfer)
	fmt.Printf("Balance: $%.2f\n", balance)
	return nil
}
