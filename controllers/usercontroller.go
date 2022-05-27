package controllers

import (
	"blogpost.com/models"
	"blogpost.com/token"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Register(c echo.Context) error {
	var user = new(models.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	token, refreshtoken, err := token.Generatetokens(user.Email, user.Name, user.Id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	user.Token = token
	user.Refreshtoken = refreshtoken

	res := models.CreateUser(user)

	return c.JSON(http.StatusOK, res)
}

func Login(c echo.Context) error {
	return c.JSON(http.StatusOK, "login works")
}
