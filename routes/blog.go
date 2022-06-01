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
	blog.GET("/by-category/:category_name", controllers.Get_blog_by_category)
	blog.GET("/all-blogs", controllers.Get_all)
	blog.PATCH("/update/:id", controllers.Update_blog)
	blog.DELETE("/delete/:id", controllers.Delete_blog)
	blog.POST("/comment", controllers.Post_comment)
	blog.POST("/create-category", controllers.Create_category)

}
