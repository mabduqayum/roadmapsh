package commands

import (
	"github.com/mabduqayum/roadmapsh/03_expense_tracker/internal/app"
	"github.com/urfave/cli/v2"
)

func UpdateCommand(expenseApp *app.App) *cli.Command {
	return &cli.Command{
		Name:  "update",
		Usage: "Update an existing transaction",
		Flags: []cli.Flag{
			&cli.IntFlag{Name: "id", Required: true},
			&cli.StringFlag{Name: "description", Aliases: []string{"d"}, Required: true},
			&cli.Float64Flag{Name: "amount", Aliases: []string{"a"}, Required: true},
		},
		Action: func(c *cli.Context) error {
			return expenseApp.UpdateTransaction(c.Int("id"), c.String("description"), c.Float64("amount"))
		},
	}
}
