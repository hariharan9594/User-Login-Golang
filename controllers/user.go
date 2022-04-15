package controllers

import (
	"UserAuth/models"
	"UserAuth/storage"
	"net/http"

	"fmt"

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
	err := curs.VerifyUser(user)
	if err != nil {
		fmt.Println("Incorrect Login details.")
		return c.String(http.StatusCreated, "username or password doesn't match")

	}
	return c.String(http.StatusCreated, "Logged in...")
}
