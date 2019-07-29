package main

import (
	"github.com/agileengine/financialAccountServer/app"
	"github.com/agileengine/financialAccountServer/routes"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/echo"
	"net/http"
)

func main() {
	app.InitApplication()
	initializeCORS()

	e := echo.New()

	e.Use(middleware.CORSWithConfig(initializeCORS()))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/users/:userID/transactions", routes.GetTransactionHandler)
	e.POST("/users/:userID/transactions", routes.CreateTransactionHandler)
	e.Logger.Fatal(e.Start(":1323"))
}

func initializeCORS() middleware.CORSConfig {
	config := middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}

	return config
}