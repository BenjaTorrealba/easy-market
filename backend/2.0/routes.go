package main

import (
	"fmt"
	"net/http"
)

func setupRoutes() {
	// Ruta raíz
	http.HandleFunc("/", corsMiddleware(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Backend API is running!")
	}))

	// Rutas para Categorías
	http.HandleFunc("/api/categorias", corsMiddleware(categoriasHandler))

	// Rutas para Productos
	http.HandleFunc("/api/productos", corsMiddleware(productosHandler))

	// Rutas para Ventas
	http.HandleFunc("/api/ventas", corsMiddleware(ventasHandler))

	// Rutas para Boletas
	http.HandleFunc("/api/boletas", corsMiddleware(boletasHandler))
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