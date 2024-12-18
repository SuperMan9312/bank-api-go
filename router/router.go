package router

import (
	"bank-api/handlers"
	"github.com/gorilla/mux"
)

func InitializeRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/accounts", handlers.CreateAccount).Methods("POST")
	r.HandleFunc("/accounts/{id}", handlers.GetAccount).Methods("GET")
	r.HandleFunc("/accounts", handlers.ListAccounts).Methods("GET")

	r.HandleFunc("/accounts/{id}/transactions", handlers.CreateTransaction).Methods("POST")
	r.HandleFunc("/accounts/{id}/transactions", handlers.GetTransactions).Methods("GET")

	r.HandleFunc("/transfer", handlers.TransferFunds).Methods("POST")

	return r
}
