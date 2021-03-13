package todo

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetTodoByIdHandler(c echo.Context) error {
	var id int
	err := echo.PathParamsBinder(c).Int("id", &id).BindError()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	db := Conn()
	stmt, err := db.Prepare("SELECT title, status FROM todos where id=$1")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	row := stmt.QueryRow(id)

	t := Todo{}

	err = row.Scan(&t.ID, &t.Title, &t.Status)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, t)
}
