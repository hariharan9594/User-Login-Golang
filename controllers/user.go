package controllers

import (
	"UserAuth/models"
	"UserAuth/storage"
	"net/http"

	"github.com/labstack/echo/v4"
)

func SaveUser(c echo.Context) error {
	user := new(models.User)

	if err := c.Bind(user); err != nil {
		return err
	}
	curs := storage.GetCursor()
	_, err := curs.CreateUser(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, user)
}

func LoginUser(c echo.Context) error {
	user := new(models.User)

	if err := c.Bind(user); err != nil {
		return err
	}
	curs := storage.GetCursor()
	login := curs.VerifyUser(user)
	if login["message"] == "Successfully Logged In" {

		return c.JSON(http.StatusCreated, login)

	}
	return c.JSON(http.StatusCreated, login)
}
