package main

import (
	"log"
	"net/http"
	"os"

	"github.com/anuchito/jaygoapi/todo"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

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

	e.GET("/todos", todo.GetTodosHandler)
	e.GET("/todos/:id", todo.GetTodoByIdHandler)
	e.POST("/todos", todo.CreateTodosHandler)
	e.PUT("/todos/:id", todo.UpdateTodosHandler)
	e.DELETE("todos/:id", todo.DeleteTodosHandler)

	port := os.Getenv("PORT")
	log.Println("port:", port)
	e.Start(":" + port)
}
