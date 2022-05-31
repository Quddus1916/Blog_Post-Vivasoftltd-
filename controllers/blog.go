package controllers

import (
	"blogpost.com/models"
	"blogpost.com/repositories"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
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
	return c.JSON(http.StatusCreated, res)
}

func Create_category(c echo.Context) error {
	var category = new(models.Category)
	if err := c.Bind(category); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	category.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	category.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	res := repositories.Create_category(category)

	return c.JSON(http.StatusCreated, res)
}

func Get_blog_by_category(c echo.Context) error {
	category_name := c.Param("category_name")
	categories := repositories.Get_By_Category(category_name)
	if len(categories) == 0 {
		return c.JSON(http.StatusFound, "failed to fetch data or no data available in this category")
	}
	return c.JSON(http.StatusOK, categories)
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
	id := c.Param("id")
	newblog.Post_id, _ = strconv.Atoi(id)
	newblog.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	res := repositories.Update_blog(newblog)
	return c.JSON(http.StatusOK, res)
}

func Delete_blog(c echo.Context) error {
	var newblog = new(models.Blog)
	id := c.Param("id")
	newblog.Post_id, _ = strconv.Atoi(id)
	res := repositories.Delete_blog(newblog)
	return c.JSON(http.StatusOK, res)
}

func Post_comment(c echo.Context) error {
	var new_comment = new(models.Comment)
	if err := c.Bind(new_comment); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	new_comment.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	res := repositories.Post_comment(new_comment)
	return c.JSON(http.StatusCreated, res)
}
