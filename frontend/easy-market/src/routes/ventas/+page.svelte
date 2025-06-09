<script>
  // Mock data de productos
  let productos = [
    { id: 1, nombre: "Producto 1", precio: 10000 },
    { id: 2, nombre: "Producto 2", precio: 15000 },
    { id: 3, nombre: "Producto 3", precio: 5000 }
  ];

  let filtro = "";
  let carrito = [];

  // Productos filtrados por búsqueda
  $: productosFiltrados = productos.filter(p =>
    p.nombre.toLowerCase().includes(filtro.toLowerCase())
  );

  // Agregar producto al carrito
  function agregarAlCarrito(producto) {
    const idx = carrito.findIndex(item => item.id === producto.id);
    if (idx !== -1) {
      carrito[idx].cantidad += 1;
      carrito = [...carrito];
    } else {
      carrito = [...carrito, { ...producto, cantidad: 1 }];
    }
  }

  // Eliminar producto del carrito
  function eliminarDelCarrito(id) {
    carrito = carrito.filter(item => item.id !== id);
  }

  // Calcular total
  $: total = carrito.reduce((sum, item) => sum + item.precio * item.cantidad, 0);

  // Confirmar venta (mock)
  function confirmarVenta() {
    alert("¡Venta confirmada!");
    carrito = [];
  }
</script>

<div class="p-6">
  <h1 class="text-2xl font-bold mb-4">Realizar Venta</h1>

  <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
    <!-- 🛍️ Columna de productos -->
    <div class="bg-white shadow-md rounded-xl p-4">
      <h2 class="text-xl font-semibold mb-3">Productos</h2>
      <input
        type="text"
        placeholder="Buscar producto..."
        class="w-full mb-4 px-3 py-2 border rounded-md focus:outline-none focus:ring"
        bind:value={filtro}
      />

      <div class="space-y-2">
        {#each productosFiltrados as producto}
          <div class="flex items-center justify-between border rounded-md px-3 py-2">
            <div>
              <p class="font-medium">{producto.nombre}</p>
              <p class="text-sm text-gray-500">${producto.precio.toLocaleString()}</p>
            </div>
            <button
              class="bg-blue-600 hover:bg-blue-700 text-white px-3 py-1 rounded-md"
              on:click={() => agregarAlCarrito(producto)}
            >
              +
            </button>
          </div>
        {/each}
        {#if productosFiltrados.length === 0}
          <p class="text-gray-400 text-sm">No hay productos.</p>
        {/if}
      </div>
    </div>

    <!-- 🧺 Columna del carrito -->
    <div class="bg-white shadow-md rounded-xl p-4">
      <h2 class="text-xl font-semibold mb-3">Carrito</h2>

      <div class="space-y-2">
        {#each carrito as item}
          <div class="flex items-center justify-between border px-3 py-2 rounded-md">
            <div>
              <p class="font-medium">{item.nombre} x{item.cantidad}</p>
              <p class="text-sm text-gray-500">
                Subtotal: ${ (item.precio * item.cantidad).toLocaleString() }
              </p>
            </div>
            <button
              class="text-red-600 hover:text-red-800 text-sm"
              on:click={() => eliminarDelCarrito(item.id)}
            >
              Eliminar
            </button>
          </div>
        {/each}
        {#if carrito.length === 0}
          <p class="text-gray-400 text-sm">El carrito está vacío.</p>
        {/if}

        <!-- Total y botón de venta -->
        <div class="mt-4 border-t pt-4">
          <p class="text-lg font-semibold">Total: ${total.toLocaleString()}</p>
          <button
            class="mt-2 w-full bg-green-600 hover:bg-green-700 text-white py-2 rounded-md"
            on:click={confirmarVenta}
            disabled={carrito.length === 0}
          >
            Confirmar Venta
          </button>
        </div>
      </div>
    </div>
  </div>
</div>