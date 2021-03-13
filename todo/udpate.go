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
	t := todos[id]
	if err := c.Bind(t); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, t)
}
