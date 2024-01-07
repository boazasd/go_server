package middlewares

import (
	"bez/bez_server/internal/services"
	"log"

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
			println(err.Error())
			return c.Redirect(302, redirectPath)
		}

		userId, err := services.CheckSession(cookie.Value)

		if err != nil {
			println(err.Error())
		}

		if userId == -1 {
			return c.Redirect(302, redirectPath)
		}
		c.Set("userId", userId)

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

		userId, err := services.CheckSession(cookie.Value)

		if err != nil {
			log.Println(err.Error())
		}

		if userId != -1 {
			return c.Redirect(302, redirectPath)
		}

		return next(c)
	}
}

func AddUserData(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := c.Get("userId").(int64)

		user, err := services.GetUser(userId)

		if err != nil {
			println(err.Error())
		}

		c.Set("user", user)

		return next(c)
	}
}
