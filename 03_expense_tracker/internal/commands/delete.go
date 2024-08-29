package commands

import (
	"github.com/mabduqayum/roadmapsh/03_expense_tracker/internal/app"
	"github.com/urfave/cli/v2"
)

func DeleteCommand(expenseApp *app.App) *cli.Command {
	return &cli.Command{
		Name:  "delete",
		Usage: "Delete a transaction",
		Flags: []cli.Flag{
			&cli.IntFlag{Name: "id", Required: true},
		},
		Action: func(c *cli.Context) error {
			return expenseApp.DeleteTransaction(c.Int("id"))
		},
	}
}
