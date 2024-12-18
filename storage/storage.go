package storage

import (
	"bank-api/models"
	"errors"
	"sync"
	"time"
)

var (
	accounts     = make(map[int]*models.Account)
	transactions = make(map[int][]models.Transaction)
	accountIDSeq = 1
	transIDSeq   = 1
	mu           sync.Mutex
)

func CreateAccount(owner string, initialBalance float64) *models.Account {
	mu.Lock()
	defer mu.Unlock()

	account := &models.Account{
		ID:      accountIDSeq,
		Owner:   owner,
		Balance: initialBalance,
	}
	accounts[accountIDSeq] = account
	accountIDSeq++
	return account
}

func GetAccount(id int) (*models.Account, error) {
	mu.Lock()
	defer mu.Unlock()

	account, exists := accounts[id]
	if !exists {
		return nil, errors.New("account not found")
	}
	return account, nil
}

func ListAccounts() []*models.Account {
	mu.Lock()
	defer mu.Unlock()

	var allAccounts []*models.Account
	for _, acc := range accounts {
		allAccounts = append(allAccounts, acc)
	}
	return allAccounts
}

func CreateTransaction(accountID int, txnType string, amount float64) (*models.Transaction, error) {
	mu.Lock()
	defer mu.Unlock()

	account, exists := accounts[accountID]
	if !exists {
		return nil, errors.New("account not found")
	}

	if txnType == "withdrawal" && account.Balance < amount {
		return nil, errors.New("insufficient funds")
	}

	if txnType == "deposit" {
		account.Balance += amount
	} else if txnType == "withdrawal" {
		account.Balance -= amount
	} else {
		return nil, errors.New("invalid transaction type")
	}

	transaction := models.Transaction{
		ID:        transIDSeq,
		AccountID: accountID,
		Type:      txnType,
		Amount:    amount,
		Timestamp: time.Now().Format(time.RFC3339),
	}
	transIDSeq++
	transactions[accountID] = append(transactions[accountID], transaction)

	return &transaction, nil
}

func GetTransactions(accountID int) []models.Transaction {
	mu.Lock()
	defer mu.Unlock()

	return transactions[accountID]
}

func TransferFunds(fromID, toID int, amount float64) error {
	mu.Lock()
	defer mu.Unlock()

	fromAccount, fromExists := accounts[fromID]
	toAccount, toExists := accounts[toID]

	if !fromExists || !toExists {
		return errors.New("account not found")
	}

	if fromAccount.Balance < amount {
		return errors.New("insufficient funds")
	}

	fromAccount.Balance -= amount
	toAccount.Balance += amount
	return nil
}
