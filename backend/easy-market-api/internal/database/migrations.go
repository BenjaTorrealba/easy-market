
package database

import (
	"database/sql"
	"fmt"
	"log"
)

// Migration representa una migración de base de datos
type Migration struct {
	ID   int
	Name string
	SQL  string
}

func GetMigrations() []Migration {
	return []Migration{
		{
			ID:   1,
			Name: "create_categorias_table",
			SQL: `
				CREATE TABLE IF NOT EXISTS categorias (
					id INTEGER PRIMARY KEY AUTOINCREMENT,
					nombre VARCHAR(50) NOT NULL
				);
			`,
		},
	}
}

// RunMigrations ejecuta todas las migraciones pendientes
func RunMigrations(db *sql.DB) error {
	// Crear tabla de migraciones si no existe
	if err := createMigrationsTable(db); err != nil {
		return fmt.Errorf("error creando tabla de migraciones: %v", err)
	}

	migrations := GetMigrations()

	for _, migration := range migrations {
		// Verificar si la migración ya fue ejecutada
		if migrationExists(db, migration.ID) {
			log.Printf("Migración %d (%s) ya ejecutada, saltando...", migration.ID, migration.Name)
			continue
		}

		// Ejecutar la migración
		log.Printf("Ejecutando migración %d: %s", migration.ID, migration.Name)
		if err := executeMigration(db, migration); err != nil {
			return fmt.Errorf("error ejecutando migración %d: %v", migration.ID, err)
		}

		// Registrar la migración como ejecutada
		if err := recordMigration(db, migration.ID, migration.Name); err != nil {
			return fmt.Errorf("error registrando migración %d: %v", migration.ID, err)
		}
	}

	return nil
}

// createMigrationsTable crea la tabla para trackear migraciones
func createMigrationsTable(db *sql.DB) error {
	query := `
		CREATE TABLE IF NOT EXISTS migrations (
			id INTEGER PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			executed_at DATETIME DEFAULT CURRENT_TIMESTAMP
		);
	`
	_, err := db.Exec(query)
	return err
}

// migrationExists verifica si una migración ya fue ejecutada
func migrationExists(db *sql.DB, migrationID int) bool {
	var count int
	query := "SELECT COUNT(*) FROM migrations WHERE id = ?"
	err := db.QueryRow(query, migrationID).Scan(&count)
	if err != nil {
		return false
	}
	return count > 0
}

// executeMigration ejecuta una migración específica
func executeMigration(db *sql.DB, migration Migration) error {
	_, err := db.Exec(migration.SQL)
	return err
}

// recordMigration registra una migración como ejecutada
func recordMigration(db *sql.DB, id int, name string) error {
	query := "INSERT INTO migrations (id, name) VALUES (?, ?)"
	_, err := db.Exec(query, id, name)
	return err
}
