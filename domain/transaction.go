package domain

import (
	"github.com/agileengine/financialAccountServer/model"
	"github.com/agileengine/financialAccountServer/util"
	"github.com/labstack/echo"
	"net/http"
	"sync"
)

//Transaction manages all user transaction
type Transaction struct{
	transactions map[string][]*model.Transaction
	mux sync.Mutex
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

func (t *Transaction) AddTransaction(userID string, transaction *model.Transaction) (*model.Transaction, error) {
	transaction.ID, _ = util.NewUUID()

	if transaction.Amount < 0 {
		return nil, echo.NewHTTPError(http.StatusBadGateway, "Bad Request")
	}

	if val, ok := t.transactions[userID]; ok {
		t.mux.Lock()
		// Lock so only one goroutine at a time can write the user transactions
		t.transactions[userID] = append(val, transaction)
		t.mux.Unlock()

		return transaction, nil
	}

	return nil, nil
}
