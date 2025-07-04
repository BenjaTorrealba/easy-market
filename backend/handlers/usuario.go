package handlers

import (
    "backend/db"
    "encoding/json"
    "net/http"
)

type LoginRequest struct {
    NombreUsuario string `json:"nombreUsuario"`
    Password      string `json:"password"`
}

type LoginResponse struct {
    Success bool   `json:"success"`
    Message string `json:"message"`
    Rol     string `json:"rol,omitempty"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
        http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
        return
    }

    var req LoginRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Datos inválidos", http.StatusBadRequest)
        return
    }

    var id int
    var rol string
    var dbPassword string
    err := db.Conn.QueryRow("SELECT ID, Password, Rol FROM Usuario WHERE NombreUsuario = ?", req.NombreUsuario).Scan(&id, &dbPassword, &rol)
    if err != nil {
        http.Error(w, "Usuario o contraseña incorrectos", http.StatusUnauthorized)
        return
    }

    if req.Password != dbPassword {
        http.Error(w, "Usuario o contraseña incorrectos", http.StatusUnauthorized)
        return
    }

    resp := LoginResponse{
        Success: true,
        Message: "Login exitoso",
        Rol:     rol,
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(resp)
}