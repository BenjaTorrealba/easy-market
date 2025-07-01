package main

import (
	"fmt"
	"net/http"
	"encoding/json"
)

func setupRoutes() {
	http.HandleFunc("/", corsMiddleware(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Backend API is running!")
	}))

	http.HandleFunc("/api/categorias", corsMiddleware(categoriasHandler))
	http.HandleFunc("/api/productos", corsMiddleware(productosHandler))
	http.HandleFunc("/api/ventas", corsMiddleware(ventasHandler))
	http.HandleFunc("/api/boletas", corsMiddleware(boletasHandler))
	http.HandleFunc("/api/productos/bajo-stock", corsMiddleware(getProductosBajoStock))
}

func categoriasHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getCategorias(w, r)
	case "POST":
		createCategoria(w, r)
	default:
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}

func productosHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getProductos(w, r)
	case "POST":
		createProducto(w, r)
	case "PUT":
		updateProducto(w, r)
	case "DELETE":
		deleteProducto(w, r)
	default:
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}

func getProductosBajoStock(w http.ResponseWriter, r *http.Request) {
    const umbral = 5 // Puedes ajustar este valor
    rows, err := db.Query("SELECT ID, Nombre, CategoriaID, Precio, Stock, COALESCE(CodigoBarras, '') FROM Producto WHERE Stock < ?", umbral)
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
func ventasHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		createVenta(w, r)
	default:
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}

func boletasHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getBoletas(w, r)
	default:
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}
