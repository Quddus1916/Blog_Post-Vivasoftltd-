package routes

import (
	"blogpost.com/controllers"
	"blogpost.com/middlewares"
	"github.com/labstack/echo/v4"
)

func Userroutes(e *echo.Echo) {
	user := e.Group("/user")

	user.POST("/register", controllers.Register)
	user.POST("/login", controllers.Login)
	user.POST("/update", controllers.Update_user, middlewares.Authenticate)

}
