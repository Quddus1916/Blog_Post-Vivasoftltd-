package controllers

import (
	"blogpost.com/models"
	"blogpost.com/types"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func Createblog(c echo.Context) error {
	var blog = new(models.Blog)
	if err := c.Bind(blog); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	blog.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	blog.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	res := models.CreateBlog(blog)

	return c.JSON(http.StatusOK, res)
}

func Getblogbycategory(c echo.Context) error {

	var category = new(types.Category)

	if err := c.Bind(category); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	fmt.Println(category.Category)

	blogs := models.Getblogbycategory(category.Category)
	if len(blogs) == 0 {

		return c.String(http.StatusNoContent, "failed to fetch data or no data available in this category")
	}

	return c.JSON(http.StatusOK, blogs)
}
