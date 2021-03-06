package main

import (
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/middleware"
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

func getTodosHandler(c echo.Context) error {
	items := []*Todo{}
	for _, item := range todos {
		items = append(items, item)
	}
	return c.JSON(http.StatusOK, items)
}

func helloHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"message": "hello",
	})
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/hello", helloHandler)
	e.GET("/todos", getTodosHandler)

	port := os.Getenv("PORT")
	log.Println("port:", port)
	e.Start(":" + port) // listen and serve on 127.0.0.0:8080
}
