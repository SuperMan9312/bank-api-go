package models

type Account struct {
	ID      int     `json:"id"`
	Owner   string  `json:"owner"`
	Balance float64 `json:"balance"`
}

type CreateAccountInput struct {
	Owner         string  `json:"owner"`
	InitialBalance float64 `json:"initial_balance"`
}
