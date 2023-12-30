package routes

import (
	"bez/bez_server/internal/services"
	"bez/bez_server/internal/types"
	"bez/bez_server/internal/utils"
	"bez/bez_server/templates"
	"encoding/json"
	"strconv"

	"github.com/labstack/echo/v4"
)

func usersInit() {
	authRouters.GET("/users/getOne/:id", getUser)
	authRouters.GET("/users", getUsers)
	authRouters.POST("/users/create", createUser)
	authRouters.POST("/users/update", updateUser)
	authRouters.DELETE("/users/delete/:id", deleteUser)
}

func getUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(400, err.Error())
	}

	user, err := services.GetUser(id)
	if err != nil {
		return c.JSON(400, err.Error())
	} else {
		return c.JSON(200, user)
	}
}

func getUsers(c echo.Context) error {
	sort := c.QueryParam("sort")
	dir := c.QueryParam("dir")
	users, err := services.GetUsers(sort, dir)
	if err != nil {
		errorComponent := templates.Error(err.Error())
		errorComponent.Render(c.Request().Context(), c.Response().Writer)
	}

	usersComponent := templates.Users(users)
	usersComponent.Render(c.Request().Context(), c.Response().Writer)
	return nil
}

func createUser(c echo.Context) error {
	var user types.User
	json_map := make(map[string]interface{})
	if err := json.NewDecoder(c.Request().Body).Decode(&json_map); err != nil {
		return c.JSON(400, err.Error())
	}
	pass, err := utils.HashAndSalt([]byte(json_map["Password"].(string)))
	if err != nil {
		return c.JSON(400, err.Error())
	}

	user.FirstName = json_map["FirstName"].(string)
	user.LastName = json_map["LastName"].(string)
	user.Email = json_map["Email"].(string)
	user.Password = pass
	userId, err := services.CreateUser(user)

	if err != nil {
		return c.JSON(400, err.Error())
	} else {
		return c.JSON(200, userId)
	}
}

func updateUser(c echo.Context) error {
	return c.JSON(200, "update user")
}

func deleteUser(c echo.Context) error {
	return c.JSON(200, "delete user")
}
