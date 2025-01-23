package main

import (
	"voting-system/config"
	"voting-system/routes"

	"github.com/labstack/echo/v4"
)

func main() {

	db := config.InitDatabase()

	e := echo.New()

	routes.RegisterRoutes(e, db)

	e.Logger.Fatal(e.Start(":8080"))
}
