package controllers

import (
	"blogpost.com/models"
	"blogpost.com/repositories"
	"blogpost.com/token"
	"blogpost.com/types"
	"blogpost.com/utils"
	//"bytes"
	//	"image"
	//"bufio"
	//"encoding/base64"
	"fmt"
	"github.com/labstack/echo/v4"
	"io"
	//"io/ioutil"
	"net/http"
	"os"
	"path"
	"strconv"
	"time"
)

func Register(c echo.Context) error {
	var user = new(models.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	user.Password = utils.Encrypt(user.Password)

	res := repositories.Create_user(user)

	return c.JSON(http.StatusCreated, res)
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
	id := c.Param("id")
	new_user.Id, _ = strconv.Atoi(id)
	new_user.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	res := repositories.Update_user(new_user)
	return c.JSON(http.StatusOK, res)
}

func Log_out(c echo.Context) error {
	var user = new(models.User)
	param_id := c.Param("id")
	if param_id == "" {
		return c.JSON(http.StatusBadRequest, user.Id)
	}
	user.Id, _ = strconv.Atoi(param_id)
	user.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	err := repositories.Log_out(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, "logout successful")
}

func Upload_profile_picture(c echo.Context) error {
	var user = new(models.User)
	id := c.Param("id")
	user.Id, _ = strconv.Atoi(id)

	form, err := c.MultipartForm()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	files := form.File["image"]

	for _, file := range files {

		src, err := file.Open()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		defer src.Close()

		uploaded_filename := file.Filename
		uploaded_filepath := path.Join("./images", uploaded_filename)

		dst, err := os.Create(uploaded_filepath)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		defer dst.Close()
		if _, err = io.Copy(dst, src); err != nil {
			return err
		}
		user.Image_path = uploaded_filepath

		res := repositories.Upload_pro_pic(user)
		if res != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

	}

	return c.JSON(http.StatusOK, "upload successful")
}
