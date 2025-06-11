package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("sqlite3", "./inventario.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Crear tablas
	if err := createTables(db); err != nil {
		log.Fatalf("Error creating tables: %s", err)
	}

	// Configurar rutas
	setupRoutes()

	log.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}