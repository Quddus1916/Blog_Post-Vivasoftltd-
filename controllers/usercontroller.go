package controllers

import (
	"blogpost.com/models"
	"blogpost.com/token"
	"blogpost.com/types"
	"blogpost.com/utils"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Register(c echo.Context) error {
	var user = new(models.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	user.Password = utils.Encrypt(user.Password)

	res := models.CreateUser(user)

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
	userdata = models.Getuserbyemail(user.Email)
	Ok := utils.Verifypassword(user.Password, userdata.Password)
	if Ok {
		fmt.Println("verified user")
	}
	//handover tokens
	token, refreshtoken, err := token.Generatetokens(userdata.Email, userdata.Name, userdata.Id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	userdata.Token = token
	userdata.Refreshtoken = refreshtoken
	//update user
	response := models.Updateusertoken(userdata)
	if !response {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, userdata.Token)
}
