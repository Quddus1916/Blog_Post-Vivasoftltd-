package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Createblog(c echo.Context) error {
	return c.JSON(http.StatusOK, "create blog works")
}

func Getblogbycategory(c echo.Context) error {
	return c.JSON(http.StatusOK, "get blog by category")
}
