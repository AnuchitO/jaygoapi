package todo

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func Conn() *sql.DB {
	var err error
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Connect to database error", err)
	}
	return db
}

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

	rowId := 1
	row := stmt.QueryRow(rowId)
	var title, status string

	err = row.Scan(&title, &status)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	t := Todo{
		ID:     id,
		Title:  title,
		Status: status,
	}
	return c.JSON(http.StatusOK, t)
}
