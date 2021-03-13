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
	delete(todos, id)
	return c.JSON(http.StatusOK, "deleted todo.")
}
