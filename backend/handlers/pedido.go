package handlers

import (
    "encoding/json"
    "net/http"
    "backend/db"
    "backend/models"
    "strconv"
    "log"
    "strings"
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

func UpdateEstadoPedido(w http.ResponseWriter, r *http.Request) {
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
    log.Println("Estado recibido:", req.Estado)

    // Si el estado es "Entregado", sumar cantidades al stock
    if strings.TrimSpace(strings.ToLower(req.Estado)) == "entregado" {
    rows, err := db.Conn.Query(`
        SELECT ProductoID, Cantidad
        FROM PedidoProducto
        WHERE PedidoID = ?`, id)
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
        _, err = db.Conn.Exec(`
            UPDATE Producto
            SET Stock = Stock + ?
            WHERE ID = ?`, cantidad, productoID)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
    }
}
    
    _, err = db.Conn.Exec("UPDATE Pedido SET Estado = ? WHERE ID = ?", req.Estado, id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"message": "Estado actualizado"})
}