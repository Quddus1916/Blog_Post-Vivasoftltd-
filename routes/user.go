package routes

import (
	"blogpost.com/controllers"
	"github.com/labstack/echo/v4"
)

func Userroutes(e *echo.Echo) {
	user := e.Group("/user")

	user.POST("/register", controllers.Register)
	user.POST("/login", controllers.Login)

}
