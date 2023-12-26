package middlewares

import (
	"bez/bez_server/internal/services"

	"github.com/labstack/echo/v4"
)

func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		redirectPath := "/login"
		cookie, err := c.Cookie("session")

		// if c.Request().URL.Path != "/" {
		// 	redirectPath = redirectPath + "?redirect=" + c.Request().URL.Path
		// }

		if err != nil {
			return c.Redirect(302, redirectPath)
		}

		logged, err := services.CheckSession(cookie.Value)

		if !logged {
			return c.Redirect(302, redirectPath)
		}

		return next(c)
	}
}

func NoAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		redirectPath := "/"
		cookie, err := c.Cookie("session")

		if err != nil {
			return next(c)
		}

		logged, err := services.CheckSession(cookie.Value)

		if !logged {
			return next(c)
		}

		return c.Redirect(302, redirectPath)
	}
}
