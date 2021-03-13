package todo

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetTodosHandler(c echo.Context) error {
	items := []*Todo{}
	for _, item := range todos {
		items = append(items, item)
	}
	return c.JSON(http.StatusOK, items)
}
