package main

import (
	"database/sql"
	"example/crud/handlers"

	"github.com/labstack/echo"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	e := echo.New()
	db := initDB("storage.db")
	migrate(db)

	e.GET("/computers", handlers.GetComputers(db))
	e.POST("/computers", handlers.PutComputer(db))
	e.PUT("/computers", handlers.EditComputer(db))
	e.DELETE("/computers/:id", handlers.DeleteComputer(db))

	e.Logger.Fatal(e.Start(":8000"))
}

func initDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)

	if err != nil {
		panic(err)
	}

	if db == nil {
		panic("db nil")
	}

	return db
}

func migrate(db *sql.DB) {
	sql := `
    CREATE TABLE IF NOT EXISTS computers(
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name VARCHAR NOT NULL,
		price INTEGER
    );
    `

	_, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
}
