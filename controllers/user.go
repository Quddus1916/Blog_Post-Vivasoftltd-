package controllers

import (
	"blogpost.com/models"
	"blogpost.com/repositories"
	"blogpost.com/token"
	"blogpost.com/types"
	"blogpost.com/utils"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func Register(c echo.Context) error {
	var user = new(models.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	user.Password = utils.Encrypt(user.Password)

	res := repositories.Create_user(user)

	return c.JSON(http.StatusOK, res)
}

func Login(c echo.Context) error {
	//given details
	var user = new(types.UserLogIn)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	//fetch details
	var userdata = new(models.User)
	userdata = repositories.Get_by_email(user.Email)

	if ok := utils.Verify_password(user.Password, userdata.Password); ok {
		fmt.Println("verified user")
	}
	//handover tokens to db
	token, refreshtoken, err := token.Generate_tokens(userdata.Email, userdata.Name, userdata.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	userdata.Token = token
	userdata.Refresh_token = refreshtoken
	//update user
	response := repositories.Set_token(userdata)
	if !response {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, userdata.Token)
}

func Update_user(c echo.Context) error {
	var new_user = new(models.User)
	if err := c.Bind(new_user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	new_user.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	res := repositories.Update_user(new_user)
	return c.JSON(http.StatusOK, res)
}

func Log_out(c echo.Context) error {
	var user = new(models.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	user.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	res := repositories.Log_out(user)
	return c.JSON(http.StatusOK, res)
}
