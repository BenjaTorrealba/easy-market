package handlers

import (
    "encoding/json"
    "net/http"
	"backend/db"
	"backend/models"
    "strconv"
)


func GetProductos(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Conn.Query("SELECT ID, Nombre, CategoriaID, Precio, Stock, COALESCE(CodigoBarras, '') FROM Producto")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var productos []models.Producto
	for rows.Next() {
		var prod models.Producto
		err := rows.Scan(&prod.ID, &prod.Nombre, &prod.CategoriaID, &prod.Precio, &prod.Stock, &prod.CodigoBarras)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		productos = append(productos, prod)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(productos)
}

func CreateProducto(w http.ResponseWriter, r *http.Request) {
	var prod models.Producto
	if err := json.NewDecoder(r.Body).Decode(&prod); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := db.Conn.Exec("INSERT INTO Producto (Nombre, CategoriaID, Precio, Stock, CodigoBarras) VALUES (?, ?, ?, ?, ?)",
		prod.Nombre, prod.CategoriaID, prod.Precio, prod.Stock, prod.CodigoBarras)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, _ := result.LastInsertId()
	prod.ID = int(id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(prod)
}

func UpdateProducto(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var prod models.Producto
	if err := json.NewDecoder(r.Body).Decode(&prod); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = db.Conn.Exec("UPDATE Producto SET Nombre=?, CategoriaID=?, Precio=?, Stock=?, CodigoBarras=? WHERE ID=?",
		prod.Nombre, prod.CategoriaID, prod.Precio, prod.Stock, prod.CodigoBarras, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	prod.ID = id
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(prod)
}

func DeleteProducto(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}
	_, err = db.Conn.Exec("DELETE FROM Producto WHERE ID=?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Producto eliminado"})
}
