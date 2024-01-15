package routes

import (
	"bez/bez_server/internal/models"
	"bez/bez_server/internal/services"
	"bez/bez_server/internal/types"
	"bez/bez_server/templates"
	"strconv"

	"github.com/labstack/echo/v4"
)

func agoraInit() {
	authRouters.GET("/agoradata", getAgoraData)
	authRouters.GET("/getAgoraAgents", getAgoraAgents)
	authRouters.GET("/createAgent", createAgent)
	authRouters.POST("/createAgentSubmit", createAgentSubmit)
	authRouters.POST("/deleteAgoraAgent/:id", deleteAgoraAgent)
}

func getAgoraData(c echo.Context) error {
	sort := c.QueryParam("sort")
	dir := c.QueryParam("dir")
	pageSizeStr := c.QueryParam("pageSize")
	pageNumberStr := c.QueryParam("pageNumber")
	user := c.Get("user").(types.User)

	pageSize, err := strconv.ParseUint(pageSizeStr, 10, 16)
	if err != nil {
		pageSize = 50
	}

	pageNumber, err := strconv.ParseUint(pageNumberStr, 10, 16)
	if err != nil {
		pageNumber = 0
	}

	if dir == "" {
		dir = "desc"
	}

	agoraData, err := services.GetAgoraData(sort, dir, uint(pageSize), uint(pageNumber))
	if err != nil {
		return Render(c, templates.Error(err.Error()))
	}

	return Render(c, templates.AgoraData(user, agoraData))
}

// agora agents

func deleteAgoraAgent(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)

	if err != nil {
		return c.JSON(400, err.Error())
	}

	am := models.IAgoraAgents{}
	err = am.Delete(id)

	if err != nil {
		return Render(c, templates.Error(err.Error()))
	}

	c.Response().Header().Set("HX-Redirect", "/")
	return nil
}

func getAgoraAgents(c echo.Context) error {
	id := c.Get("userId").(int64)

	agents, err := services.GetAgoraAgents(id)

	if err != nil {
		Render(c, templates.Error(err.Error()))
		return nil
	}

	html := ""
	for _, agent := range agents {
		html += agent.SearchTxt + "<br/>"
	}

	return c.HTML(200, html)
}

func createAgentSubmit(c echo.Context) error {
	id := c.Get("userId").(int64)

	um := models.IUser{}
	user, err := um.GetById(id)

	if err != nil {
		return Render(c, templates.Error(err.Error()))
	}

	agent := types.AgoraAgent{}
	agent.SearchTxt = c.FormValue("searchTxt")
	agent.Category = c.FormValue("category")
	agent.Condition = c.FormValue("condition")
	agent.Area = c.FormValue("area")
	agent.WithImage = c.FormValue("withImage") == "on"
	agent.UserId = id
	agent.UserEmail = user.Email

	_, err = services.AddAgoraAgent(agent)

	if err != nil {
		return Render(c, templates.Error(err.Error()))
	}

	if err != nil {
		return Render(c, templates.Error(err.Error()))
	}

	c.Response().Header().Set("HX-Redirect", "/")
	return nil
}
