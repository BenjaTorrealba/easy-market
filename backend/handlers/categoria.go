package handlers

import (
    "encoding/json"
    "net/http"
    "backend/db"
    "backend/models"
)

func GetCategorias(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Conn.Query("SELECT ID, Nombre FROM Categoria")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var categorias []models.Categoria
	for rows.Next() {
		var cat models.Categoria
		err := rows.Scan(&cat.ID, &cat.Nombre)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		categorias = append(categorias, cat)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categorias)
}

func CreateCategoria(w http.ResponseWriter, r *http.Request) {
	var cat models.Categoria
	if err := json.NewDecoder(r.Body).Decode(&cat); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := db.Conn.Exec("INSERT INTO Categoria (Nombre) VALUES (?)", cat.Nombre)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, _ := result.LastInsertId()
	cat.ID = int(id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cat)
}