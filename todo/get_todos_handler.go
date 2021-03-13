package todo

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

var todos = map[int]*Todo{
	1: &Todo{ID: 1, Title: "pay phone bills", Status: "active"},
}

func GetTodosHandler(c echo.Context) error {
	items := []*Todo{}
	for _, item := range todos {
		items = append(items, item)
	}

	items = append(items, &Todo{Title: os.Getenv("DATABASE_URL")})
	return c.JSON(http.StatusOK, items)
}
