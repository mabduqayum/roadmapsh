package commands

import (
	"github.com/mabduqayum/roadmapsh/03_expense_tracker/internal/app"
	"github.com/mabduqayum/roadmapsh/03_expense_tracker/internal/transaction"
	"github.com/urfave/cli/v2"
)

func AddCommand(expenseApp *app.App) *cli.Command {
	return &cli.Command{
		Name:  "add",
		Usage: "Add a new transaction",
		Subcommands: []*cli.Command{
			{
				Name:  "expense",
				Usage: "Add a new expense",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "description", Aliases: []string{"d"}, Required: true},
					&cli.Float64Flag{Name: "amount", Aliases: []string{"a"}, Required: true},
				},
				Action: func(c *cli.Context) error {
					return expenseApp.AddTransaction(c.String("description"), c.Float64("amount"), transaction.TypeExpense)
				},
			},
			{
				Name:  "top-up",
				Usage: "Add a new top-up",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "description", Aliases: []string{"d"}, Required: true},
					&cli.Float64Flag{Name: "amount", Aliases: []string{"a"}, Required: true},
				},
				Action: func(c *cli.Context) error {
					return expenseApp.AddTransaction(c.String("description"), c.Float64("amount"), transaction.TypeTopUp)
				},
			},
			{
				Name:  "transfer",
				Usage: "Add a new transfer",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "description", Aliases: []string{"d"}, Required: true},
					&cli.Float64Flag{Name: "amount", Aliases: []string{"a"}, Required: true},
				},
				Action: func(c *cli.Context) error {
					return expenseApp.AddTransaction(c.String("description"), c.Float64("amount"), transaction.TypeTransfer)
				},
			},
		},
	}
}
