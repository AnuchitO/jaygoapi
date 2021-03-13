package todo

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateTodosHandler(c echo.Context) error {
	t := Todo{}
	if err := c.Bind(&t); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	row := Conn().QueryRow("INSERT INTO todos (title, status) values ($1, $2)  RETURNING id", t.Title, t.Status)
	err := row.Scan(&t.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, t)
}
