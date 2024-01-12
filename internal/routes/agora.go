package routes

import (
	"bez/bez_server/internal/services"
	"bez/bez_server/internal/types"
	"bez/bez_server/templates"
	"strconv"

	"github.com/labstack/echo/v4"
)

func agoraInit() {
	authRouters.GET("/agoradata", getAgoraData)
}

func getAgoraData(c echo.Context) error {
	sort := c.QueryParam("sort")
	dir := c.QueryParam("dir")
	pageSizeStr := c.QueryParam("pageSize")
	pageNumberStr := c.QueryParam("pageNumber")
	user := c.Get("user").(types.User)

	pageSize, err := strconv.ParseUint(pageSizeStr, 10, 16)
	if err != nil {
		pageSize = 200
	}

	pageNumber, err := strconv.ParseUint(pageNumberStr, 10, 16)
	if err != nil {
		pageNumber = 0
	}

	agoraData, err := services.GetAgoraData(sort, dir, uint(pageSize), uint(pageNumber))
	if err != nil {
		return Render(c, templates.Error(err.Error()))
	}

	return Render(c, templates.AgoraData(user, agoraData))
}
