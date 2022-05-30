package controllers

import (
	"blogpost.com/models"
	"blogpost.com/repositories"
	"blogpost.com/types"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func Create_blog(c echo.Context) error {
	var blog = new(models.Blog)
	if err := c.Bind(blog); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	blog.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	blog.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	res := repositories.Create_Blog(blog)

	return c.JSON(http.StatusOK, res)
}

func Get_blog_by_category(c echo.Context) error {

	var category = new(types.Category)

	if err := c.Bind(category); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	fmt.Println(category.Category)

	blogs := repositories.Get_By_Category(category.Category)
	if len(blogs) == 0 {

		return c.JSON(http.StatusFound, "failed to fetch data or no data available in this category")
	}

	return c.JSON(http.StatusOK, blogs)
}

func Get_all(c echo.Context) error {

	res := repositories.Get_all()

	return c.JSON(http.StatusOK, res)
}

func Update_blog(c echo.Context) error {
	var newblog = new(models.Blog)
	if err := c.Bind(newblog); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	newblog.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	res := repositories.Update_blog(newblog)
	return c.JSON(http.StatusOK, res)
}

func Delete_blog(c echo.Context) error {
	var newblog = new(models.Blog)
	if err := c.Bind(newblog); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	//newblog.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	res := repositories.Delete_blog(newblog)
	return c.JSON(http.StatusOK, res)
}
