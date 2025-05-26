<script>
    import Icon from '@iconify/svelte';

    let productos = [
      {
        id: 1,
        nombre: 'Pan marraqueta',
        categoria: 'Panadería',
        stock: 35,
        precio: 500,
        activo: true
      },
      {
        id: 2,
        nombre: 'Coca Cola 1.5L',
        categoria: 'Bebidas',
        stock: 12,
        precio: 1500,
        activo: true
      }
    ];
  

  </script>
  
  <h2>Lista de Productos</h2>
  
  <table>
    <thead>
      <tr>
        <th>ID</th>
        <th>Nombre</th>
        <th>Categoría</th>
        <th>Stock</th>
        <th>Precio ($)</th>
        <th>Activo</th>
      </tr>
    </thead>
    <tbody>
      {#each productos as producto, index}
        <tr>
          <td>{producto.id}</td>
          <td>{producto.nombre}</td>
          <td>{producto.categoria}</td>
          <td>
            <input type="number" bind:value={producto.stock} min="0" />
            <button>
                <Icon icon="mdi:refresh" width="18" />
              </button>
          </td>
          <td>
            <input type="number" bind:value={producto.precio} min="0" />
            <button>
                <Icon icon="mdi:refresh" width="18" />
              </button>
        </td>
          
          <td>
            <input type="checkbox" bind:checked={producto.activo} />
          </td>

        </tr>
      {/each}
    </tbody>
  </table>
  
  <style>
    table {
      width: 100%;
      border-collapse: collapse;
      background: white;
      box-shadow: 0 2px 8px rgba(0,0,0,0.05);
      border-radius: 10px;
      overflow: hidden;
    }
  
    th, td {
      padding: 0.8rem;
      text-align: left;
      border-bottom: 1px solid #e5e7eb;
    }
  
    th {
      background-color: #f3f4f6;
      color: #111827;
    }
  
    input[type="number"] {
      width: 80px;
      padding: 0.3rem;
      border: 1px solid #d1d5db;
      border-radius: 4px;
    }
  
    input[type="checkbox"] {
      transform: scale(1.2);
    }
  
    button {
      background-color: #38bdf8;
      border: none;
      padding: 0.4rem 1rem;
      color: white;
      border-radius: 6px;
      cursor: pointer;
      transition: 0.2s;
    }
  
    button:hover {
      background-color: #0ea5e9;
    }
  </style>
  