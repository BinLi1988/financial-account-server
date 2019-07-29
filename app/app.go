package app

import (
	"github.com/agileengine/financialAccountServer/domain"
	"github.com/agileengine/financialAccountServer/model"
)

var TransactionManager *domain.Transaction

//InitApplication initialie Application
func InitApplication(){
	TransactionManager = domain.NewTransaction()

	//Initialize user test
	userID := "1234"
	TransactionManager.InitializeUserAccount(userID) // user test
	transaction1 := model.NewTransaction("", "credit", 20)
	transaction2 := model.NewTransaction("", "debit", 10)
	TransactionManager.AddTransaction(userID, &transaction1)
	TransactionManager.AddTransaction(userID, &transaction2)
}
