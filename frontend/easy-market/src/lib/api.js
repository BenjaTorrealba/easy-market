const API_BASE = "http://localhost:8080/api";

class ApiClient {
  async request(endpoint, options = {}) {
    const url = `${API_BASE}${endpoint}`;
    const config = {
      headers: {
        "Content-Type": "application/json",
        ...options.headers,
      },
      ...options,
    };

    try {
      const response = await fetch(url, config);

      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }

      return await response.json();
    } catch (error) {
      console.error("API Error:", error);
      throw error;
    }
  }
  async updateEstadoPedido(id, estado) {
    return this.request(`/pedidos/estado?id=${id}`, {
      method: "PUT",
      body: JSON.stringify({ estado }),
    });
  }
  async getCategorias() {
    return this.request("/categorias");
  }

  async createCategoria(nombre) {
    return this.request("/categorias", {
      method: "POST",
      body: JSON.stringify({ nombre }),
    });
  }

  async getPedidos() {
    return this.request("/pedidos");
  }

  async createPedido(pedido) {
    return this.request("/pedidos", {
      method: "POST",
      body: JSON.stringify(pedido),
    });
  }
  async deletePedido(id) {
    return this.request(`/pedidos?id=${id}`, {
      method: "DELETE",
      
    });
  }
  async getProductos() {
    return this.request("/productos");
  }

  async createProducto(producto) {
    return this.request("/productos", {
      method: "POST",
      body: JSON.stringify(producto),
    });
  }

  async updateProducto(id, producto) {
    return this.request(`/productos?id=${id}`, {
      method: "PUT",
      body: JSON.stringify(producto),
    });
  }

  async createVenta(productos, total) {
    return this.request("/ventas", {
      method: "POST",
      body: JSON.stringify({ productos, total }),
    });
  }
  async getProductosBajoStock() {
    return this.request("/productos/bajo-stock");
  }
  async getVentasTotales() {
    return this.request("/reportes/ventas-totales");
  }

  async getHistorialVentas(fecha_inicio = "", fecha_fin = "") {
    let url = "/reportes/historial-ventas";
    const params = [];
    if (fecha_inicio) params.push(`fecha_inicio=${fecha_inicio}`);
    if (fecha_fin) params.push(`fecha_fin=${fecha_fin}`);
    if (params.length) url += "?" + params.join("&");
    return this.request(url);
  }
  async getProductosMasVendidos() {
    return this.request("/reportes/productos-mas-vendidos");
  }
}

export const api = new ApiClient();

export const productosStore = {
  subscribe: null,

  async load() {
    try {
      const productos = await api.getProductos();
      return productos;
    } catch (error) {
      console.error("Error cargando productos:", error);
      return [];
    }
  },

  async create(producto) {
    try {
      const nuevoProducto = await api.createProducto(producto);
      return nuevoProducto;
    } catch (error) {
      console.error("Error creando producto:", error);
      throw error;
    }
  },
};
