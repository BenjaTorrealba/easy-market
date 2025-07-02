package handlers

import (
	"encoding/json"
	"net/http"
	"backend/db"
	"backend/models"
)

func GetBoletas(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Conn.Query("SELECT ID, Fecha, Total FROM Boleta ORDER BY Fecha DESC")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var boletas []models.Boleta
	for rows.Next() {
		var boleta models.Boleta
		err := rows.Scan(&boleta.ID, &boleta.Fecha, &boleta.Total)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		boletas = append(boletas, boleta)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(boletas)
}