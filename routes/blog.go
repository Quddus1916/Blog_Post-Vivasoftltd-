package routes

import (
	"blogpost.com/controllers"
	"github.com/labstack/echo/v4"
)

func Blogroutes(e *echo.Echo) {
	user := e.Group("/blog")

	user.POST("/post", controllers.Createblog)
	user.GET("/bycategory", controllers.Getblogbycategory)

}
