package routes

import (
	"github.com/agileengine/financialAccountServer/app"
	"github.com/labstack/echo"
	"net/http"
)

//GetTransactionHandler Gets user's transactions 
func GetTransactionHandler(c echo.Context) error {
	userID := c.Param("userID")
	result, _ := app.TransactionManager.GetTransactions(userID)
	return c.JSON(http.StatusOK, interface{}(result))
}