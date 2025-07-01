package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

// Handlers para Categorías
func getCategorias(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT ID, Nombre FROM Categoria")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var categorias []Categoria
	for rows.Next() {
		var cat Categoria
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

func createCategoria(w http.ResponseWriter, r *http.Request) {
	var cat Categoria
	if err := json.NewDecoder(r.Body).Decode(&cat); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := db.Exec("INSERT INTO Categoria (Nombre) VALUES (?)", cat.Nombre)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, _ := result.LastInsertId()
	cat.ID = int(id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cat)
}

// Handlers para Productos
func getProductos(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT ID, Nombre, CategoriaID, Precio, Stock, COALESCE(CodigoBarras, '') FROM Producto")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var productos []Producto
	for rows.Next() {
		var prod Producto
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

func createProducto(w http.ResponseWriter, r *http.Request) {
	var prod Producto
	if err := json.NewDecoder(r.Body).Decode(&prod); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := db.Exec("INSERT INTO Producto (Nombre, CategoriaID, Precio, Stock, CodigoBarras) VALUES (?, ?, ?, ?, ?)",
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

func updateProducto(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var prod Producto
	if err := json.NewDecoder(r.Body).Decode(&prod); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = db.Exec("UPDATE Producto SET Nombre=?, CategoriaID=?, Precio=?, Stock=?, CodigoBarras=? WHERE ID=?",
		prod.Nombre, prod.CategoriaID, prod.Precio, prod.Stock, prod.CodigoBarras, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	prod.ID = id
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(prod)
}

func deleteProducto(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	_, err = db.Exec("DELETE FROM Producto WHERE ID=?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Producto eliminado"})
}

func createVenta(w http.ResponseWriter, r *http.Request) {
	var venta VentaRequest

	if err := json.NewDecoder(r.Body).Decode(&venta); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tx, err := db.Begin()
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

	response := VentaResponse{
		Message:  "Venta creada exitosamente",
		BoletaID: boletaID,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func getBoletas(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT ID, Fecha, Total FROM Boleta ORDER BY Fecha DESC")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var boletas []Boleta
	for rows.Next() {
		var boleta Boleta
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