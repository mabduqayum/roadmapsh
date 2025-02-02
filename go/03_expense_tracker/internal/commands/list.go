package commands

import (
	"github.com/mabduqayum/roadmapsh/03_expense_tracker/internal/app"
	"github.com/urfave/cli/v2"
)

func ListCommand(expenseApp *app.App) *cli.Command {
	return &cli.Command{
		Name:  "list",
		Usage: "List all transactions",
		Action: func(c *cli.Context) error {
			return expenseApp.ListTransactions()
		},
	}
}
