package todo

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetTodosHandler(c echo.Context) error {
	items := []*Todo{}
	stmt, err := Conn().Prepare("SELECT id, title, status FROM todos")
	if err != nil {
		log.Fatal("can't prepare query all todos statment", err)
	}

	rows, err := stmt.Query()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	for rows.Next() {
		t := &Todo{}
		err := rows.Scan(&t.ID, &t.Title, &t.Status)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		items = append(items, t)
	}
	return c.JSON(http.StatusOK, items)
}
