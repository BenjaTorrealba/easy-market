package handlers

import (
    "encoding/json"
    "net/http"
    "backend/db"
    "backend/models"

)

func GetPedidos(w http.ResponseWriter, r *http.Request) {
    rows, err := db.Conn.Query("SELECT ID, Distribuidor, Fecha, FechaEstimada, Estado FROM Pedido ORDER BY Fecha DESC")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var pedidos []models.Pedido
    for rows.Next() {
        var p models.Pedido
        err := rows.Scan(&p.ID, &p.Distribuidor, &p.Fecha, &p.FechaEstimada, &p.Estado)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        // Cargar productos asociados (id, nombre, cantidad)
        prodRows, err := db.Conn.Query(`
            SELECT Producto.ID, Producto.Nombre, PedidoProducto.Cantidad
            FROM PedidoProducto
            JOIN Producto ON Producto.ID = PedidoProducto.ProductoID
            WHERE PedidoProducto.PedidoID = ?`, p.ID)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        var productos []models.ProductoPedido
        for prodRows.Next() {
            var prod models.ProductoPedido
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

func CreatePedido(w http.ResponseWriter, r *http.Request) {
    var req models.PedidoRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    tx, err := db.Conn.Begin()
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
    var p models.Pedido
    err = db.Conn.QueryRow("SELECT ID, Distribuidor, Fecha, FechaEstimada, Estado FROM Pedido WHERE ID = ?", pedidoID).
        Scan(&p.ID, &p.Distribuidor, &p.Fecha, &p.FechaEstimada, &p.Estado)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Cargar productos asociados (id, nombre, cantidad)
    prodRows, err := db.Conn.Query(`
        SELECT Producto.ID, Producto.Nombre, PedidoProducto.Cantidad
        FROM PedidoProducto
        JOIN Producto ON Producto.ID = PedidoProducto.ProductoID
        WHERE PedidoProducto.PedidoID = ?`, p.ID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    var productos []models.ProductoPedido
    for prodRows.Next() {
        var prod models.ProductoPedido
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

func DeletePedido(w http.ResponseWriter, r *http.Request) {
    idStr := r.URL.Query().Get("id")
    if idStr == "" {
        http.Error(w, "ID no proporcionado", http.StatusBadRequest)
        return
    }

    _, err := db.Conn.Exec("DELETE FROM Pedido WHERE ID = ?", idStr)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}