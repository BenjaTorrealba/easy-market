package main

import (
	"backend/db"
	"log"
	"net/http"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var Conn *sql.DB

func main() {
	var err error
	db.Conn, err = sql.Open("sqlite3", "./inventario.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Conn.Close()

	if err := createTables(db.Conn); err != nil {
		log.Fatalf("Error creating tables: %s", err)
	}

	setupRoutes()

	log.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}