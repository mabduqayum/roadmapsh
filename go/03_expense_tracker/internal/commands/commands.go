package commands

import (
	"github.com/mabduqayum/roadmapsh/03_expense_tracker/internal/app"
)

func GetCommands(expenseApp *app.App) []*cli.Command {
	return []*cli.Command{
		AddCommand(expenseApp),
		UpdateCommand(expenseApp),
		DeleteCommand(expenseApp),
		ListCommand(expenseApp),
		SummaryCommand(expenseApp),
	}
}
