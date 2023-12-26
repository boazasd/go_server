package routes

import (
	"bez/bez_server/internal/middlewares"
	"bez/bez_server/templates"

	"github.com/labstack/echo/v4"
)

var router = echo.New()
var authRouters = router.Group("")
var noAuthRouters = router.Group("")

func Init() {

	authRouters.Use(middlewares.Auth)
	noAuthRouters.Use(middlewares.NoAuth)

	router.GET("/v", func(c echo.Context) error {
		return c.String(200, "v0.0.2")
	})

	router.GET("/favicon.ico", func(c echo.Context) error {
		return c.String(200, "")
	})

	authRouters.GET("/", func(c echo.Context) error {
		templates.Home().Render(c.Request().Context(), c.Response().Writer)
		return nil
	})

	noAuthRouters.GET("/login", func(c echo.Context) error {
		templates.Login().Render(c.Request().Context(), c.Response().Writer)
		return nil
	})

	router.RouteNotFound("/*", func(c echo.Context) error {
		templates.NotFound().Render(c.Request().Context(), c.Response().Writer)
		return nil
	})

	usersInit()
	sessionInit()
	router.Logger.Fatal(router.Start(":8080"))
}
