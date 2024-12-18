package models

type Transaction struct {
	ID        int     `json:"id"`
	AccountID int     `json:"account_id"`
	Type      string  `json:"type"`
	Amount    float64 `json:"amount"`
	Timestamp string  `json:"timestamp"`
}

type TransactionInput struct {
	Type   string  `json:"type"`
	Amount float64 `json:"amount"`
}
