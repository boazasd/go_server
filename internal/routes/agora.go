package routes

import (
	"bez/bez_server/internal/services"
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

	pageSize, err := strconv.ParseUint(pageSizeStr, 10, 16)
	if err != nil {
		pageSize = 10
	}

	pageNumber, err := strconv.ParseUint(pageNumberStr, 10, 16)
	if err != nil {
		pageNumber = 0
	}

	agoraData, err := services.GetAgoraData(sort, dir, uint(pageSize), uint(pageNumber))
	if err != nil {
		errorComponent := templates.Error(err.Error())
		errorComponent.Render(c.Request().Context(), c.Response().Writer)
	}

	usersComponent := templates.AgoraData(agoraData)
	usersComponent.Render(c.Request().Context(), c.Response().Writer)
	return nil
}
