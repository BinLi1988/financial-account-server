package model

import "time"

//Transaction represents operations that the user made.
type Transaction struct {
	ID string `json:"id"`
	TransactionType string `json:"transactionType"`
	Amount float32 `json:"amount"`
	EffectiveDate time.Time `json:"effectiveDate"`
}

//NewTransaction creates new transaction
func NewTransaction(id string, transactionType string, amount float32) Transaction {
	return Transaction{ID: id, TransactionType: transactionType, Amount: amount, EffectiveDate: time.Now()}
}