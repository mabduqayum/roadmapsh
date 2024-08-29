package main

import (
	"fmt"
	"os"

	"github.com/mabduqayum/roadmapsh/03_expense_tracker/config"
	"github.com/mabduqayum/roadmapsh/03_expense_tracker/internal/app"
	"github.com/mabduqayum/roadmapsh/03_expense_tracker/internal/commands"
	"github.com/urfave/cli/v2"
)

func main() {
	cfg := config.Load()
	expenseApp := app.NewApp(cfg.Storage.File)

	cliApp := &cli.App{
		Name:     "expense-tracker",
		Usage:    "A simple expense tracker CLI application",
		Commands: commands.GetCommands(expenseApp),
	}

	if err := cliApp.Run(os.Args); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
