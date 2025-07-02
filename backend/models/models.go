package models

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

type ProductoPedido struct {
    ID       int    `json:"id"`
    Nombre   string `json:"nombre"`
    Cantidad int    `json:"cantidad"`
}

type Pedido struct {
    ID            int              `json:"id"`
    Distribuidor  string           `json:"distribuidor"`
    Fecha         string           `json:"fecha"`
    FechaEstimada string           `json:"fechaEstimada"`
    Estado        string           `json:"estado"`
    Productos     []ProductoPedido `json:"productos,omitempty"`
}

type PedidoRequest struct {
    Distribuidor  string `json:"distribuidor"`
    Fecha         string `json:"fecha"`
    FechaEstimada string `json:"fechaEstimada"`
    Estado        string `json:"estado"`
    Productos     []struct {
        ProductoID int `json:"producto_id"`
        Cantidad   int `json:"cantidad"`
    } `json:"productos"`
}