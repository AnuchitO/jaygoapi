package main

import (
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

var todos = map[int]*Todo{
	1: &Todo{ID: 1, Title: "pay phone bills", Status: "active"},
}

func getTodosHandler(c echo.Context) error {
	items := []*Todo{}
	for _, item := range todos {
		items = append(items, item)
	}
	return c.JSON(http.StatusOK, items)
}

func createTodosHandler(c echo.Context) error {
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

func getTodoByIdHandler(c echo.Context) error {
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

func updateTodosHandler(c echo.Context) error {
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

func deleteTodosHandler(c echo.Context) error {
	var id int
	err := echo.PathParamsBinder(c).Int("id", &id).BindError()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	delete(todos, id)
	return c.JSON(http.StatusOK, "deleted todo.")
}

func authMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Println("start middleware : check authentication")
		token := c.Request().Header.Get("Authorization")
		if token != "ABC" {
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Unauthorized"})
		}
		return next(c)
	}
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(authMiddleware)

	e.GET("/todos", getTodosHandler)
	e.GET("/todos/:id", getTodoByIdHandler)
	e.POST("/todos", createTodosHandler)
	e.PUT("todos/:id", updateTodosHandler)
	e.DELETE("todos/:id", deleteTodosHandler)

	port := os.Getenv("PORT")
	log.Println("port:", port)
	e.Start(":" + port)
}
