package todo

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var _db *sql.DB

func Conn() *sql.DB {
	if _db != nil {
		return _db
	}

	var err error
	_db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Connect to database error", err)
	}
	return _db
}
