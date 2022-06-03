package routes

import (
	"blogpost.com/controllers"
	"blogpost.com/middlewares"
	"github.com/labstack/echo/v4"
)

func Blogroutes(e *echo.Echo) {
	blog := e.Group("/post")
	blog.Use(middlewares.Authenticate)

	blog.POST("/create", controllers.CreatePost)
	blog.GET("/category/:category_name", controllers.GetPostByCategory)
	blog.GET("/all", controllers.GetAll)
	blog.PATCH("/update/:id", controllers.UpdatePost)
	blog.DELETE("/delete/:id", controllers.DeletePost)
	blog.POST("/comment", controllers.PostComment)
	blog.POST("/category/create", controllers.CreateCategory)
	blog.PATCH("/comment/:user_id/:comment_id", controllers.UpdateComment)

}
