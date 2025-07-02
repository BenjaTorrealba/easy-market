package main

import (
	"backend/db"
	"backend/handlers"
    "backend/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func setupRoutes() {
	http.HandleFunc("/", corsMiddleware(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Backend API is running!")
	}))
    http.HandleFunc("/api/productos", corsMiddleware(productosHandler))
	http.HandleFunc("/api/categorias", corsMiddleware(categoriasHandler))
	http.HandleFunc("/api/ventas", corsMiddleware(ventasHandler))
	http.HandleFunc("/api/boletas", corsMiddleware(boletasHandler))
	http.HandleFunc("/api/productos/bajo-stock", corsMiddleware(getProductosBajoStock))
	http.HandleFunc("/api/pedidos", corsMiddleware(pedidosHandler))
	http.HandleFunc("/api/pedidos/estado", corsMiddleware(updateEstadoPedido))

}

func categoriasHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		handlers.GetCategorias(w, r)
	case "POST":
		handlers.CreateCategoria(w, r)
	default:
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}

func productosHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		handlers.GetProductos(w, r)
	case "POST":
		handlers.CreateProducto(w, r)
	case "PUT":
		handlers.UpdateProducto(w, r)
	case "DELETE":
		handlers.DeleteProducto(w, r)
	default:
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}

func getProductosBajoStock(w http.ResponseWriter, r *http.Request) {
    const umbral = 5 // Puedes ajustar este valor
    rows, err := db.Conn.Query("SELECT ID, Nombre, CategoriaID, Precio, Stock, COALESCE(CodigoBarras, '') FROM Producto WHERE Stock < ?", umbral)
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
func ventasHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		handlers.CreateVenta(w, r)
	default:
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}

func boletasHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		handlers.GetBoletas(w, r)
	default:
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}
func pedidosHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "GET":
        handlers.GetPedidos(w, r)
    case "POST":
        handlers.CreatePedido(w, r)
    default:
        http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
    }
}
func updateEstadoPedido(w http.ResponseWriter, r *http.Request) {
    if r.Method != "PUT" {
        http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
        return
    }
    
    idStr := r.URL.Query().Get("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "ID inválido", http.StatusBadRequest)
        return
    }
    
    var req struct {
        Estado string `json:"estado"`
    }
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    
    _, err = db.Conn.Exec("UPDATE Pedido SET Estado = ? WHERE ID = ?", req.Estado, id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"message": "Estado actualizado"})
}