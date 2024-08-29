package expense

import (
	"time"
)

type TransactionType string

const (
	TypeExpense  TransactionType = "expense"
	TypeTopUp    TransactionType = "top-up"
	TypeTransfer TransactionType = "transfer"
)

type Transaction struct {
	ID          int             `json:"id"`
	Date        time.Time       `json:"date"`
	Description string          `json:"description"`
	Amount      float64         `json:"amount"`
	Type        TransactionType `json:"type"`
}
