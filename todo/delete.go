package todo

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func DeleteTodosHandler(c echo.Context) error {
	var id int
	err := echo.PathParamsBinder(c).Int("id", &id).BindError()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	db := Conn()
	stmt, err := db.Prepare("DELETE FROM todos WHERE id = $1")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	if _, err := stmt.Exec(id); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "deleted todo.")
}
