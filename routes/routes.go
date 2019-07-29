package routes

import (
	"github.com/agileengine/financialAccountServer/app"
	"github.com/agileengine/financialAccountServer/model"
	"github.com/labstack/echo"
	"net/http"
)

//GetTransactionHandler Gets user's transactions 
func GetTransactionHandler(c echo.Context) error {
	userID := c.Param("userID")
	result, _ := app.TransactionManager.GetTransactions(userID)
	return c.JSON(http.StatusOK, interface{}(result))
}

func CreateTransactionHandler(c echo.Context) error {
	userID := c.Param("userID")
	transaction := new(model.Transaction)
	if err := c.Bind(transaction); err != nil {
		return err
	}

	result, _ := app.TransactionManager.AddTransaction(userID, transaction)

	return c.JSON(http.StatusCreated, result)
}