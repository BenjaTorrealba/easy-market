package main

type Categoria struct {
	ID     int    `json:"id"`
	Nombre string `json:"nombre"`
}

type Producto struct {
	ID           int     `json:"id"`
	Nombre       string  `json:"nombre"`
	CategoriaID  int     `json:"categoria_id"`
	Precio       float64 `json:"precio"`
	Stock        int     `json:"stock"`
	CodigoBarras string  `json:"codigo_barras"`
}

type Boleta struct {
	ID    int     `json:"id"`
	Fecha string  `json:"fecha"`
	Total float64 `json:"total"`
}

type DetalleVenta struct {
	ID           int `json:"id"`
	IDProducto   int `json:"id_producto"`
	CantProducto int `json:"cant_producto"`
	IDBoleta     int `json:"id_boleta"`
}

type VentaRequest struct {
	Productos []struct {
		ID       int `json:"id"`
		Cantidad int `json:"cantidad"`
	} `json:"productos"`
	Total float64 `json:"total"`
}

type VentaResponse struct {
	Message  string `json:"message"`
	BoletaID int64  `json:"boleta_id"`
}