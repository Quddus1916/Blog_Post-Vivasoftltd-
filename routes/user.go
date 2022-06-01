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
	user.PATCH("/update/:id", controllers.Update_user, middlewares.Authenticate)
	user.GET("/logout/:id", controllers.Log_out, middlewares.Authenticate)
	user.POST("/upload-profile-pic/:id", controllers.Upload_profile_picture, middlewares.Authenticate)

}
