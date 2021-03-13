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

	t, ok := todos[id]
	if !ok {
		return c.JSON(http.StatusOK, map[int]string{})
	}
	return c.JSON(http.StatusOK, t)
}
