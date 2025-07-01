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

func getPedidos(w http.ResponseWriter, r *http.Request) {
    rows, err := db.Query("SELECT ID, Distribuidor, Fecha, FechaEstimada, Estado FROM Pedido ORDER BY Fecha DESC")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var pedidos []Pedido
    for rows.Next() {
        var p Pedido
        err := rows.Scan(&p.ID, &p.Distribuidor, &p.Fecha, &p.FechaEstimada, &p.Estado)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        // Cargar productos asociados (id, nombre, cantidad)
        prodRows, err := db.Query(`
            SELECT Producto.ID, Producto.Nombre, PedidoProducto.Cantidad
            FROM PedidoProducto
            JOIN Producto ON Producto.ID = PedidoProducto.ProductoID
            WHERE PedidoProducto.PedidoID = ?`, p.ID)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        var productos []ProductoPedido
        for prodRows.Next() {
            var prod ProductoPedido
            if err := prodRows.Scan(&prod.ID, &prod.Nombre, &prod.Cantidad); err != nil {
                prodRows.Close()
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }
            productos = append(productos, prod)
        }
        prodRows.Close()
        p.Productos = productos

        pedidos = append(pedidos, p)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(pedidos)
}

func createPedido(w http.ResponseWriter, r *http.Request) {
    var req PedidoRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    tx, err := db.Begin()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer tx.Rollback()

    result, err := tx.Exec(
        "INSERT INTO Pedido (Distribuidor, Fecha, FechaEstimada, Estado) VALUES (?, ?, ?, ?)",
        req.Distribuidor, req.Fecha, req.FechaEstimada, req.Estado)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    pedidoID, err := result.LastInsertId()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    for _, prod := range req.Productos {
        _, err := tx.Exec(
            "INSERT INTO PedidoProducto (PedidoID, ProductoID, Cantidad) VALUES (?, ?, ?)",
            pedidoID, prod.ProductoID, prod.Cantidad)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
    }

    if err := tx.Commit(); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }


    // Consulta el pedido recién creado para devolverlo completo
    var p Pedido
    err = db.QueryRow("SELECT ID, Distribuidor, Fecha, FechaEstimada, Estado FROM Pedido WHERE ID = ?", pedidoID).
        Scan(&p.ID, &p.Distribuidor, &p.Fecha, &p.FechaEstimada, &p.Estado)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Cargar productos asociados (id, nombre, cantidad)
    prodRows, err := db.Query(`
        SELECT Producto.ID, Producto.Nombre, PedidoProducto.Cantidad
        FROM PedidoProducto
        JOIN Producto ON Producto.ID = PedidoProducto.ProductoID
        WHERE PedidoProducto.PedidoID = ?`, p.ID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    var productos []ProductoPedido
    for prodRows.Next() {
        var prod ProductoPedido
        if err := prodRows.Scan(&prod.ID, &prod.Nombre, &prod.Cantidad); err != nil {
            prodRows.Close()
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        productos = append(productos, prod)
    }
    prodRows.Close()
    p.Productos = productos

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(p)
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

    // Si el estado es "Entregado", sumar stock
    if req.Estado == "Entregado" {
        // Obtener productos y cantidades del pedido
        rows, err := db.Query("SELECT ProductoID, Cantidad FROM PedidoProducto WHERE PedidoID = ?", id)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        defer rows.Close()
        for rows.Next() {
            var productoID, cantidad int
            if err := rows.Scan(&productoID, &cantidad); err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }
            // Sumar stock
            _, err = db.Exec("UPDATE Producto SET Stock = Stock + ? WHERE ID = ?", cantidad, productoID)
            if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }
        }
    }

    // Actualizar estado del pedido
    _, err = db.Exec("UPDATE Pedido SET Estado = ? WHERE ID = ?", req.Estado, id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"message": "Estado actualizado"})
}