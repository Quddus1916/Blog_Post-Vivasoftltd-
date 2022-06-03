package routes

import (
	"blogpost.com/controllers"
	"blogpost.com/middlewares"
	"github.com/labstack/echo/v4"
)

func Userroutes(e *echo.Echo) {
	user := e.Group("/user")

	user.POST("/register", controllers.Register)
	user.POST("/login", controllers.LogIn)
	user.PATCH("/update/:id", controllers.UpdateUser, middlewares.Authenticate)
	user.GET("/logout/:id", controllers.LogOut, middlewares.Authenticate)
	user.POST("/upload-profile-pic/:id", controllers.UploadProfilePicture, middlewares.Authenticate)

}
