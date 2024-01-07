package middlewares

import (
	"bez/bez_server/internal/types"
	"slices"
	"strings"

	"github.com/labstack/echo/v4"
)

func Perm(perm string) func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			redirectPath := "/"

			user := c.Get("user").(types.User)

			if user.Id == 0 {
				return c.Redirect(302, redirectPath)
			}

			roles := strings.Split(user.Roles, ";")

			if !slices.Contains(roles, perm) {
				return c.Redirect(302, redirectPath)
			}

			return next(c)
		}
	}
}
