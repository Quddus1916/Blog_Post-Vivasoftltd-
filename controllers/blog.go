package controllers

import (
	"blogpost.com/models"
	"blogpost.com/repositories"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"time"
)

func CreatePost(c echo.Context) error {
	//bind data
	var blog = new(models.Post)
	if err := c.Bind(blog); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	//set time
	blog.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	blog.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	// send res
	res := repositories.CreatePost(blog)
	return c.JSON(http.StatusCreated, res)
}

func CreateCategory(c echo.Context) error {
	//bind data
	var category = new(models.Category)
	if err := c.Bind(category); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	//set time
	category.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	category.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	//send res
	res := repositories.CreateCategory(category)

	return c.JSON(http.StatusCreated, res)
}

func GetPostByCategory(c echo.Context) error {
	//bind param
	category_name := c.Param("category_name")
	//query
	categories := repositories.GetByCategory(category_name)
	// checking
	if len(categories) == 0 {
		return c.JSON(http.StatusFound, "failed to fetch data or no data available in this category")
	}
	//res
	return c.JSON(http.StatusOK, categories)
}

func GetAll(c echo.Context) error {
	res := repositories.GetAll()
	return c.JSON(http.StatusOK, res)
}

func UpdatePost(c echo.Context) error {
	var newblog = new(models.Post)
	if err := c.Bind(newblog); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	id := c.Param("id")
	newblog.Id, _ = strconv.Atoi(id)
	newblog.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	res := repositories.UpdatePost(newblog)
	return c.JSON(http.StatusOK, res)
}

func DeletePost(c echo.Context) error {
	var newblog = new(models.Post)
	id := c.Param("id")
	newblog.Id, _ = strconv.Atoi(id)
	res := repositories.DeletePost(newblog)
	return c.JSON(http.StatusOK, res)
}

func PostComment(c echo.Context) error {
	var new_comment = new(models.Comment)
	if err := c.Bind(new_comment); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	new_comment.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	res := repositories.PostComment(new_comment)
	return c.JSON(http.StatusCreated, res)
}

func UpdateComment(c echo.Context) error {
	var new_comment = new(models.Comment)
	err := c.Bind(new_comment)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	id := c.Param("user_id")
	comment_id, _ := strconv.Atoi(c.Param("comment_id"))
	new_comment.Id = comment_id
	saved_comment := repositories.GetCommentDetails(comment_id)
	user_id, _ := strconv.Atoi(id)
	user := repositories.GetUserDetails(user_id)
	if user.Role == "user" {
		if user.Id != saved_comment.User_id {
			return c.JSON(http.StatusUnauthorized, user)
		}
	}

	res := repositories.UpdateComment(new_comment)

	return c.JSON(http.StatusOK, res)
}
