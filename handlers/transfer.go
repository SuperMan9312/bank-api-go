package handlers

import (
	"bank-api/models"
	"bank-api/storage"
	"encoding/json"
	"net/http"
)

func TransferFunds(w http.ResponseWriter, r *http.Request) {
	var input models.TransferInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	err := storage.TransferFunds(input.FromAccountID, input.ToAccountID, input.Amount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "success"})
}
