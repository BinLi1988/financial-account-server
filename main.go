package main

import (
	"github.com/agileengine/financialAccountServer/app"
	"github.com/agileengine/financialAccountServer/routes"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/echo"
)

func main() {
	app.InitApplication()
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/users/:userID/transactions", routes.GetTransactionHandler)

	e.Logger.Fatal(e.Start(":1323"))
}