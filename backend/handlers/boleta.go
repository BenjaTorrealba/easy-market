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

type VentasTotales struct {
    Dia    float64 `json:"dia"`
    Semana float64 `json:"semana"`
    Mes    float64 `json:"mes"`
}

func GetVentasTotales(w http.ResponseWriter, r *http.Request) {
    var totales VentasTotales

    // Total día
    row := db.Conn.QueryRow("SELECT IFNULL(SUM(Total),0) FROM Boleta WHERE DATE(Fecha) = DATE('now')")
    row.Scan(&totales.Dia)

    // Total semana (últimos 7 días)
    row = db.Conn.QueryRow("SELECT IFNULL(SUM(Total),0) FROM Boleta WHERE Fecha >= DATE('now', '-6 days')")
    row.Scan(&totales.Semana)

    // Total mes (mes actual)
    row = db.Conn.QueryRow("SELECT IFNULL(SUM(Total),0) FROM Boleta WHERE strftime('%Y-%m', Fecha) = strftime('%Y-%m', 'now')")
    row.Scan(&totales.Mes)

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(totales)
}

type GananciaNeta struct {
    Ganancia float64 `json:"ganancia"`
}

func GetGananciaNeta(w http.ResponseWriter, r *http.Request) {
    row := db.Conn.QueryRow(`
        SELECT IFNULL(SUM((p.Precio - p.Costo) * dv.Cant_Producto),0)
        FROM DetalleVenta dv
        JOIN Producto p ON dv.ID_Producto = p.ID
        JOIN Boleta b ON dv.ID_Boleta = b.ID
        WHERE strftime('%Y-%m', b.Fecha) = strftime('%Y-%m', 'now')
    `)
    var ganancia float64
    row.Scan(&ganancia)

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(GananciaNeta{Ganancia: ganancia})
}



func GetHistorialVentas(w http.ResponseWriter, r *http.Request) {
    fechaInicio := r.URL.Query().Get("fecha_inicio")
    fechaFin := r.URL.Query().Get("fecha_fin")

    query := "SELECT ID, Fecha, Total FROM Boleta WHERE 1=1"
    var args []interface{}

    if fechaInicio != "" {
        query += " AND Fecha >= ?"
        args = append(args, fechaInicio)
    }
    if fechaFin != "" {
        query += " AND Fecha <= ?"
        args = append(args, fechaFin)
    }
    query += " ORDER BY Fecha DESC"

    rows, err := db.Conn.Query(query, args...)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var boletas []models.Boleta
    for rows.Next() {
        var boleta models.Boleta
        if err := rows.Scan(&boleta.ID, &boleta.Fecha, &boleta.Total); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        boletas = append(boletas, boleta)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(boletas)
}