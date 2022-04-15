package routers

import (
	"UserAuth/controllers"

	"github.com/labstack/echo/v4"
)

func InitRoutes() *echo.Echo {
	e := echo.New()

	e.POST("/User", controllers.SaveUser)
	e.POST("/login", controllers.LoginUser)
	return e
}
