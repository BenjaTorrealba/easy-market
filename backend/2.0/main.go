package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

const schema = `
CREATE TABLE IF NOT EXISTS Categoria (
    ID INTEGER PRIMARY KEY AUTOINCREMENT,
    Nombre TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS Producto (
    ID INTEGER PRIMARY KEY AUTOINCREMENT,
    Nombre TEXT NOT NULL,
    CategoriaID INTEGER NOT NULL,
    Precio REAL NOT NULL,
    Stock INTEGER NOT NULL,
    CodigoBarras TEXT,
    FOREIGN KEY (CategoriaID) REFERENCES Categoria(ID)
);

CREATE TABLE IF NOT EXISTS Boleta (
    ID INTEGER PRIMARY KEY AUTOINCREMENT,
    Fecha DATETIME NOT NULL,
    Total REAL NOT NULL
);

CREATE TABLE IF NOT EXISTS DetalleVenta (
    ID INTEGER PRIMARY KEY AUTOINCREMENT,
    ID_Producto INTEGER NOT NULL,
    Cant_Producto INTEGER NOT NULL,
    ID_Boleta INTEGER NOT NULL,
    FOREIGN KEY (ID_Producto) REFERENCES Producto(ID),
    FOREIGN KEY (ID_Boleta) REFERENCES Boleta(ID)
);
`

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("sqlite3", "./inventario.db")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(schema)
	if err != nil {
		log.Fatalf("Error creating tables: %s", err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Backend with SQLite is running!")
	})

	log.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
