package domain

import (
	"github.com/agileengine/financialAccountServer/model"
	"github.com/agileengine/financialAccountServer/util"
)

//Transaction manages all user transaction
type Transaction struct{
	transactions map[string][]*model.Transaction
}

//NewTransaction creates an instance of Transaction
func NewTransaction() *Transaction {
	return &Transaction{ transactions: make(map[string][]*model.Transaction)}
}

func (t *Transaction) GetTransactions(userID string) ([]*model.Transaction,  error) {

	if val, ok := t.transactions[userID]; ok {
		return val, nil
	}

	return nil, nil

}

func (t *Transaction) InitializeUserAccount(userID string){
	t.transactions[userID] = make([]*model.Transaction, 0)
}

func (t *Transaction) AddTransaction(userID string, transaction *model.Transaction) {
	transaction.ID, _ = util.NewUUID()
	if val, ok := t.transactions[userID]; ok {
		t.transactions[userID] = append(val, transaction)
	}
}
