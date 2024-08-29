package commands

import (
	"github.com/mabduqayum/roadmapsh/03_expense_tracker/internal/app"
	"github.com/urfave/cli/v2"
)

func GetCommands(expenseApp *app.App) []*cli.Command {
	return []*cli.Command{
		{
			Name:  "add",
			Usage: "Add a new expense",
			Flags: []cli.Flag{
				&cli.StringFlag{Name: "description", Aliases: []string{"d"}, Required: true},
				&cli.Float64Flag{Name: "amount", Aliases: []string{"a"}, Required: true},
			},
			Action: func(c *cli.Context) error {
				return expenseApp.AddExpense(c.String("description"), c.Float64("amount"))
			},
		},
		{
			Name:  "update",
			Usage: "Update an existing expense",
			Flags: []cli.Flag{
				&cli.IntFlag{Name: "id", Required: true},
				&cli.StringFlag{Name: "description", Aliases: []string{"d"}, Required: true},
				&cli.Float64Flag{Name: "amount", Aliases: []string{"a"}, Required: true},
			},
			Action: func(c *cli.Context) error {
				return expenseApp.UpdateExpense(c.Int("id"), c.String("description"), c.Float64("amount"))
			},
		},
		{
			Name:  "delete",
			Usage: "Delete an expense",
			Flags: []cli.Flag{
				&cli.IntFlag{Name: "id", Required: true},
			},
			Action: func(c *cli.Context) error {
				return expenseApp.DeleteExpense(c.Int("id"))
			},
		},
		{
			Name:  "list",
			Usage: "List all expenses",
			Action: func(c *cli.Context) error {
				return expenseApp.ListExpenses()
			},
		},
		{
			Name:  "summary",
			Usage: "Show expense summary",
			Flags: []cli.Flag{
				&cli.IntFlag{Name: "month", Aliases: []string{"m"}},
			},
			Action: func(c *cli.Context) error {
				if c.IsSet("month") {
					return expenseApp.MonthlySummary(c.Int("month"))
				}
				return expenseApp.Summary()
			},
		},
	}
}
