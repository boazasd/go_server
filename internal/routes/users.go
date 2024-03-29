package routes

import (
	"bez/bez_server/internal/services"
	"bez/bez_server/internal/types"
	"bez/bez_server/internal/utils"
	"bez/bez_server/templates"
	"log"
	"strconv"

	"github.com/labstack/echo/v4"
)

func usersInit() {
	authRouters.GET("/users/getOne/:id", getUser, P("super"))
	authRouters.GET("/users", getUsers, P("super"))
	authRouters.GET("/users/create", createUser)
	authRouters.POST("/users/createSubmit", createUserSubmit)
	authRouters.POST("/users/update", updateUser)
	authRouters.DELETE("/users/delete/:id", deleteUser)
}

func getUser(c echo.Context) error {

	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)

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
		return Render(c, errorComponent)
	}

	return Render(c, templates.Users(c.Get("user").(types.User), users))
}

func createUser(c echo.Context) error {
	cmp := templates.CreateUser(c.Get("user").(types.User))
	Render(c, cmp)
	return nil
}
func createAgent(c echo.Context) error {
	cmp := templates.CreateAgent(c.Get("user").(types.User))
	Render(c, cmp)
	return nil
}

func createUserSubmit(c echo.Context) error {
	var user types.User
	// json_map := make(map[string]interface{})

	// if err := json.NewDecoder(c.Request().Body).Decode(&json_map); err != nil {
	// 	return c.HTML(400, err.Error())
	// }
	pass := c.FormValue("password")
	if pass != "" {
		repeatPass := c.FormValue("repeatPassword")
		if pass != repeatPass {
			return c.HTML(200, "passwords do not match")
		}

	} else {
		p, err := utils.RandomString(10, "")
		pass = p
		if err != nil {
			return c.HTML(200, err.Error())
		}
	}

	hashed, err := utils.HashAndSalt([]byte(pass))

	if err != nil {
		return c.HTML(200, err.Error())
	}

	user.FirstName = c.FormValue("firstName")
	user.LastName = c.FormValue("lastName")
	user.Email = c.FormValue("email")
	user.Roles = c.FormValue("roles")
	user.Password = hashed
	log.Println(user.Email)
	_, err = services.CreateUser(user)

	if err != nil {
		return c.HTML(200, err.Error())
	}

	if err != nil {
		return c.HTML(200, err.Error())
	} else {
		c.Response().Header().Set("HX-Redirect", "/users")
		return c.HTML(200, pass)
	}
}

func updateUser(c echo.Context) error {
	return c.String(200, "update user")
}

func deleteUser(c echo.Context) error {
	return c.String(200, "delete user")
}
