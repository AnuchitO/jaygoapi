package todo

import (
	"net/http"

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
	return c.JSON(http.StatusOK, items)
}

func CreateTodosHandler(c echo.Context) error {
	t := Todo{}
	if err := c.Bind(&t); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	id := len(todos)
	id++
	t.ID = id
	todos[t.ID] = &t

	return c.JSON(http.StatusCreated, "created todo.")
}

func GetTodoByIdHandler(c echo.Context) error {
	var id int
	err := echo.PathParamsBinder(c).Int("id", &id).BindError()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	t, ok := todos[id]
	if !ok {
		return c.JSON(http.StatusOK, map[int]string{})
	}
	return c.JSON(http.StatusOK, t)
}

func UpdateTodosHandler(c echo.Context) error {
	var id int
	err := echo.PathParamsBinder(c).Int("id", &id).BindError()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	t := todos[id]
	if err := c.Bind(t); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, t)
}

func DeleteTodosHandler(c echo.Context) error {
	var id int
	err := echo.PathParamsBinder(c).Int("id", &id).BindError()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	delete(todos, id)
	return c.JSON(http.StatusOK, "deleted todo.")
}
