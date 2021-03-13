package todo

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func UpdateTodosHandler(c echo.Context) error {
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

	if err := c.Bind(t); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	updateStmt, err := db.Prepare("UPDATE todos SET status=$2 WHERE id=$1;")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	if _, err := updateStmt.Exec(t.ID, t.Status); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, t)
}
