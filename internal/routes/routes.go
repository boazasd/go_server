package routes

import (
	"bez/bez_server/templates"

	"github.com/labstack/echo/v4"
)

var router = echo.New()

func Init() {
	router.GET("/v", func(c echo.Context) error {
		return c.String(200, "v0.0.2")
	})

	router.GET("/favicon.ico", func(c echo.Context) error {
		return c.String(200, "")
	})

	router.GET("/", func(c echo.Context) error {
		templates.Login().Render(c.Request().Context(), c.Response().Writer)
		return nil
	})

	router.RouteNotFound("/*", func(c echo.Context) error {
		templates.NotFound().Render(c.Request().Context(), c.Response().Writer)
		return nil
	})

	usersInit()
	router.Logger.Fatal(router.Start(":8080"))
}
