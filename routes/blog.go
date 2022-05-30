package routes

import (
	"blogpost.com/controllers"
	"blogpost.com/middlewares"
	"github.com/labstack/echo/v4"
)

func Blogroutes(e *echo.Echo) {
	user := e.Group("/blog")
	user.Use(middlewares.Authenticate)

	user.POST("/post", controllers.Create_blog)
	user.GET("/by_category", controllers.Get_blog_by_category)
	user.GET("/all_blogs", controllers.Get_all)
	user.POST("/update", controllers.Update_blog)

}
