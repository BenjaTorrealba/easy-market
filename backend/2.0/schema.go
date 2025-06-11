package main

import "database/sql"

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

func createTables(database *sql.DB) error {
	_, err := database.Exec(schema)
	return err
}