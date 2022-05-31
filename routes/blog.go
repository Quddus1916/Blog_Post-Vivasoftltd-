package routes

import (
	"blogpost.com/controllers"
	"blogpost.com/middlewares"
	"github.com/labstack/echo/v4"
)

func Blogroutes(e *echo.Echo) {
	blog := e.Group("/blog")
	blog.Use(middlewares.Authenticate)

	blog.POST("/post", controllers.Create_blog)
	blog.GET("/by_category", controllers.Get_blog_by_category)
	blog.GET("/all_blogs", controllers.Get_all)
	blog.POST("/update", controllers.Update_blog)
	blog.DELETE("/delete", controllers.Delete_blog)
	blog.POST("/comment", controllers.Post_comment)

}
