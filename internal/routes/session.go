package routes

import (
	"bez/bez_server/internal/services"
	"bez/bez_server/templates"

	"github.com/labstack/echo/v4"
)

func sessionInit() {
	router.POST("/sessionLogin", sessionLogin)
	router.POST("/sessionLogout", sessionLogout)
}

func sessionLogin(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	id, err := services.Login(email, password)
	if err != nil {
		errorComponent := templates.Error(err.Error())
		errorComponent.Render(c.Request().Context(), c.Response().Writer)
		return nil
	}

	cookie, err := services.CreateOrRefreshSession(id)
	if err != nil {
		errorComponent := templates.Error(err.Error())
		errorComponent.Render(c.Request().Context(), c.Response().Writer)
		return nil
	}

	c.SetCookie(cookie)
	c.Response().Header().Set("HX-Redirect", "/")
	return c.String(200, "Success")
}

func sessionLogout(c echo.Context) error {
	deletionCookie := services.CreateSessionCookie("")

	currentCookie, err := c.Cookie("session")
	if err != nil {
		print(err.Error())
	} else {
		services.DeleteSession(currentCookie.Value)
	}

	c.SetCookie(deletionCookie)
	c.Response().Header().Set("HX-Redirect", "/login")
	return c.String(200, "Success")
}
