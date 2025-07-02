package handlers

import (
    "encoding/json"
    "net/http"
	"backend/db"
	"backend/models"
    "strconv"
	"time"
)

func CreateVenta(w http.ResponseWriter, r *http.Request) {
	var venta models.VentaRequest

	if err := json.NewDecoder(r.Body).Decode(&venta); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tx, err := db.Conn.Begin()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer tx.Rollback()

	result, err := tx.Exec("INSERT INTO Boleta (Fecha, Total) VALUES (?, ?)", time.Now(), venta.Total)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	boletaID, _ := result.LastInsertId()


	for _, prod := range venta.Productos {

		var stockActual int
		err = tx.QueryRow("SELECT Stock FROM Producto WHERE ID = ?", prod.ID).Scan(&stockActual)
		if err != nil {
			http.Error(w, "Producto no encontrado", http.StatusBadRequest)
			return
		}

		if stockActual < prod.Cantidad {
			http.Error(w, "Stock insuficiente para el producto ID "+strconv.Itoa(prod.ID), http.StatusBadRequest)
			return
		}


		_, err = tx.Exec("INSERT INTO DetalleVenta (ID_Producto, Cant_Producto, ID_Boleta) VALUES (?, ?, ?)",
			prod.ID, prod.Cantidad, boletaID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}


		_, err = tx.Exec("UPDATE Producto SET Stock = Stock - ? WHERE ID = ?", prod.Cantidad, prod.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}


	if err = tx.Commit(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := models.VentaResponse{
		Message:  "Venta creada exitosamente",
		BoletaID: boletaID,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
