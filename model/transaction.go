package model

//Transaction represents operations that the user made.
type Transaction struct {
	ID string `json:"id"`
	TransactionType string `json:"transactionType"`
	Amount float32 `json:"amount"`
}

//NewTransaction creates new transaction
func NewTransaction(id string, transactionType string, amount float32) Transaction {
	return Transaction{ID: id, TransactionType: transactionType, Amount: amount}
}