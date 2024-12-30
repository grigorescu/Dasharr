package main

import (
	"backend/handlers"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/initdb", handlers.InitDB)
	e.GET("/update", handlers.Update)
	e.GET("/stats", handlers.GetStats)
	e.Logger.Fatal(e.Start(":1323"))
}
