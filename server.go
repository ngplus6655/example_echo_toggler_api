package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"hello/routes"
	"hello/middlewares"
)

func main() {
	e := echo.New()
	
    // Middleware 
	e.Use(middleware.Logger())
	e.Use(middlewares.DatabaseService())

	//Routes
	routes.Init(e)

	e.Logger.Fatal(e.Start(":8080"))
}