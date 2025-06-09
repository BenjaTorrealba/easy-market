package database

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

// DB es la instancia global de la base de datos
var DB *sql.DB

// Initialize inicializa la conexión a la base de datos
func Initialize(dbPath string) error {
	// Crear el directorio si no existe
	if err := createDBDirectory(dbPath); err != nil {
		return fmt.Errorf("error creando directorio de la base de datos: %v", err)
	}

	// Abrir la conexión a SQLite
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return fmt.Errorf("error abriendo la base de datos: %v", err)
	}

	// Verificar la conexión
	if err := db.Ping(); err != nil {
		return fmt.Errorf("error conectando a la base de datos: %v", err)
	}

	// Configurar la conexión
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)

	DB = db

	// Ejecutar migraciones
	if err := RunMigrations(db); err != nil {
		return fmt.Errorf("error ejecutando migraciones: %v", err)
	}

	return nil
}

// Close cierra la conexión a la base de datos
func Close() error {
	if DB != nil {
		return DB.Close()
	}
	return nil
}

// createDBDirectory crea el directorio para la base de datos si no existe
func createDBDirectory(dbPath string) error {
	dir := filepath.Dir(dbPath)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return os.MkdirAll(dir, 0755)
	}
	return nil
}
