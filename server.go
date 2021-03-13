package main

import (
	"log"
	"net/http"
	"os"

	"github.com/anuchito/jaygoapi/todo"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

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
	e.GET("/todos", todo.GetTodosHandler)

	port := os.Getenv("PORT")
	log.Println("port:", port)
	e.Start(":" + port) // listen and serve on 127.0.0.0:8080
}
