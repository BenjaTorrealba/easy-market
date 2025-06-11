// lib/api.js - Cliente API para tu backend Go
const API_BASE = 'http://localhost:8080/api';

class ApiClient {
  async request(endpoint, options = {}) {
    const url = `${API_BASE}${endpoint}`;
    const config = {
      headers: {
        'Content-Type': 'application/json',
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
      console.error('API Error:', error);
      throw error;
    }
  }

  // Categorías
  async getCategorias() {
    return this.request('/categorias');
  }

  async createCategoria(nombre) {
    return this.request('/categorias', {
      method: 'POST',
      body: JSON.stringify({ nombre }),
    });
  }

  // Productos
  async getProductos() {
    return this.request('/productos');
  }

  async createProducto(producto) {
    return this.request('/productos', {
      method: 'POST',
      body: JSON.stringify(producto),
    });
  }

  async updateProducto(id, producto) {
    return this.request(`/productos?id=${id}`, {
      method: 'PUT',
      body: JSON.stringify(producto),
    });
  }

  // Ventas
  async createVenta(productos, total) {
    return this.request('/ventas', {
      method: 'POST',
      body: JSON.stringify({ productos, total }),
    });
  }
}

export const api = new ApiClient();

// Ejemplo de componente Svelte para mostrar productos
export const productosStore = {
  // Store para manejar el estado de productos
  subscribe: null, // Se implementaría con writable store de Svelte
  
  async load() {
    try {
      const productos = await api.getProductos();
      return productos;
    } catch (error) {
      console.error('Error cargando productos:', error);
      return [];
    }
  },

  async create(producto) {
    try {
      const nuevoProducto = await api.createProducto(producto);
      return nuevoProducto;
    } catch (error) {
      console.error('Error creando producto:', error);
      throw error;
    }
  }
};