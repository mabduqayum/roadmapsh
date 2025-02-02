package commands

import (
	"github.com/mabduqayum/roadmapsh/03_expense_tracker/internal/app"
	"github.com/urfave/cli/v2"
)

func SummaryCommand(expenseApp *app.App) *cli.Command {
	return &cli.Command{
		Name:  "summary",
		Usage: "Show transaction summary",
		Flags: []cli.Flag{
			&cli.IntFlag{Name: "month", Aliases: []string{"m"}},
		},
		Action: func(c *cli.Context) error {
			if c.IsSet("month") {
				return expenseApp.MonthlySummary(c.Int("month"))
			}
			return expenseApp.Summary()
		},
	}
}
