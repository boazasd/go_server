package routes

import (
	"bez/bez_server/internal/middlewares"
	"bez/bez_server/internal/services"
	"bez/bez_server/templates"
	"log"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

var P = middlewares.Perm
var router = echo.New()
var authRouters = router.Group("")
var noAuthRouters = router.Group("")

func Render(c echo.Context, component templ.Component) error {
	component.Render(c.Request().Context(), c.Response().Writer)
	return nil
}

func Init() {

	authRouters.Use(middlewares.Auth)
	authRouters.Use(middlewares.AddUserData)
	noAuthRouters.Use(middlewares.NoAuth)

	router.GET("/v", func(c echo.Context) error {
		return c.String(200, "v0.0.2")
	})

	router.GET("/favicon", func(c echo.Context) error {
		return c.File("assets/favicon_io/favicon.ico")
	})

	router.GET("/style.css", func(c echo.Context) error {
		return c.File("assets/style.css")
	})

	authRouters.GET("/", func(c echo.Context) error {
		userId := c.Get("userId").(int64)
		user, err := services.GetUser(userId)
		if err != nil {
			Render(c, templates.Error(err.Error()))
			return nil
		}

		agents, err := services.GetAgoraAgents(userId)

		if err != nil {
			// Render(c, templates.Error(err.Error()))
			// return nil
			log.Println(err.Error())
		}

		agentsStr := []string{}
		for _, agent := range agents {
			agentsStr = append(agentsStr, agent.SearchTxt)
		}

		Render(c, templates.Home(user, agentsStr))
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
	agoraInit()
	router.Logger.Fatal(router.Start(":8081"))
}
