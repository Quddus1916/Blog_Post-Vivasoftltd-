package main

import (
	"blogpost.com/config"
	"blogpost.com/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	routes.Userroutes(e)

	routes.Blogroutes(e)

	config := config.Getconfig()

	e.Start(":" + config.Port)

}
