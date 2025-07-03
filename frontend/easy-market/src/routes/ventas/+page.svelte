<script>
  import { onMount, onDestroy, tick } from "svelte";
  import { iniciarScanner, detenerScanner } from "$lib/codigoBarras.js";
  import { api } from "$lib/api.js";
  import Alerta from "$lib/Alerta.svelte";
  let productos = [];
  let filtro = "";
  let carrito = [];
  let mensaje = "";
  let tipoAlerta = "success"; // success, error, info
  let codigoBarras = "";
  let mostrarScanner = false;

  onMount(async () => {
    try {
      productos = await api.getProductos();
    } catch (e) {
      mensaje = "Error al cargar productos";
    }
  });
  onDestroy(() => {
    detenerScanner();
    mostrarScanner = false;
  });

  $: productosFiltrados = productos.filter((p) =>
    p.nombre.toLowerCase().includes(filtro.toLowerCase())
  );

  function agregarAlCarrito(producto) {
    const idx = carrito.findIndex((item) => item.id === producto.id);
    if (idx !== -1) {
      carrito[idx].cantidad += 1;
      carrito = [...carrito];
    } else {
      carrito = [...carrito, { ...producto, cantidad: 1 }];
    }
  }

  function eliminarDelCarrito(id) {
    carrito = carrito.filter((item) => item.id !== id);
  }

  function actualizarCantidad(id, nuevaCantidad) {
    carrito = carrito.map((item) =>
      item.id === id
        ? {
            ...item,
            cantidad: Math.max(1, Math.min(nuevaCantidad, item.stock)),
          }
        : item
    );
  }

  $: total = carrito.reduce(
    (sum, item) => sum + item.precio * item.cantidad,
    0
  );

  async function confirmarVenta() {
    try {
      await api.createVenta(
        carrito.map(({ id, cantidad }) => ({ id, cantidad })),
        total
      );
      mensaje = "¡Venta confirmada!";
      carrito = [];
      productos = await api.getProductos();
    } catch (e) {
      mensaje = "Error al realizar la venta";
    }
    setTimeout(() => (mensaje = ""), 2000);
  }

async function agregarPorCodigo(codigo) {
  codigoBarras = codigo;
  const producto = productos.find((p) => p.codigo_barras === codigo);
  if (producto) {
    agregarAlCarrito(producto);
    mensaje = `Producto "${producto.nombre}" agregado al carrito`;
    tipoAlerta = "success";
  } else {
    mensaje = "Producto no encontrado";
    tipoAlerta = "error";
  }
  setTimeout(() => (mensaje = ""), 2000);
}

  async function iniciarEscaner() {
    mostrarScanner = true;
    await tick();
    iniciarScanner((codigo) => {
      codigoBarras = codigo;
      agregarPorCodigo(codigo);
      detenerScanner();
      mostrarScanner = false;
    });
  }

  function detenerEscaner() {
    detenerScanner();
    mostrarScanner = false;
  }
</script>

<div class="p-6">
  <h1 class="text-2xl font-bold mb-4">Realizar Venta</h1>

  <div class="mb-6">
    <div
      class="flex items-center gap-2 bg-blue-50 border border-blue-200 text-blue-800 px-4 py-3 rounded-md shadow-sm"
    >
      <svg
        xmlns="http://www.w3.org/2000/svg"
        class="h-5 w-5"
        fill="none"
        viewBox="0 0 24 24"
        stroke="currentColor"
      >
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          stroke-width="2"
          d="M3 7h2M3 17h2M7 3v2m10-2v2m4 4h2m-2 10h2M7 21v-2m10 2v-2M12 8v4l3 3"
        />
      </svg>
      <span><strong>Escanee el producto</strong> para que se agregue al carrito</span>
    </div>
    <div class="flex gap-2 mt-4 mb-2">
      <button
        class="bg-green-500 hover:bg-green-600 text-white px-4 py-2 rounded"
        on:click={iniciarEscaner}
        disabled={mostrarScanner}
        type="button"
      >
        Iniciar escáner
      </button>
      <button
        class="bg-red-500 hover:bg-red-600 text-white px-4 py-2 rounded"
        on:click={detenerEscaner}
        disabled={!mostrarScanner}
        type="button"
      >
        Detener escáner
      </button>
    </div>
    {#if mostrarScanner}
      <div
        id="scanner-container"
        class="scanner-video mb-4"
      ></div>
    {/if}
    {#if codigoBarras}
      <div class="text-sm text-gray-600 mb-2">
        Último código escaneado: <span class="font-mono">{codigoBarras}</span>
      </div>
    {/if}
  </div>

  {#if mensaje}
    <Alerta {mensaje} tipo={tipoAlerta}  />
  {/if}


  <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
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
          <div
            class="flex items-center justify-between border rounded-md px-3 py-2"
          >
            <div>
              <p class="font-medium">{producto.nombre}</p>
              <p class="text-sm text-gray-500">
                Precio: ${producto.precio.toLocaleString()} | Stock: {producto.stock}
              </p>
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

    <div class="bg-white shadow-md rounded-xl p-4">
      <h2 class="text-xl font-semibold mb-3">Carrito</h2>

      <div class="space-y-2">
        {#each carrito as item}
          <div
            class="flex items-center justify-between border px-3 py-2 rounded-md"
          >
            <div>
              <p class="font-medium">
                {item.nombre}
                <input
                  type="number"
                  min="1"
                  max={item.stock}
                  class="w-16 ml-2 border rounded px-1 py-0.5 text-center"
                  bind:value={item.cantidad}
                  on:input={(e) => actualizarCantidad(item.id, +e.target.value)}
                />
                <span class="text-xs text-gray-400 ml-1"
                  >/ {item.stock} disp.</span
                >
              </p>
              <p class="text-sm text-gray-500">
                Subtotal: ${(item.precio * item.cantidad).toLocaleString()}
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


<style>
  .scanner-video {
    position: relative;
    width: 100%;
    max-width: 320px;
    height: 200px;
    margin: 0 auto 1rem auto;
    border-radius: 8px;
    overflow: hidden;
    background: #000;
  }
  #scanner-container video,
  #scanner-container canvas {
    position: absolute !important;
    top: 0; left: 0;
    width: 100% !important;
    height: 100% !important;
    object-fit: cover !important;
    z-index: 1 !important;
    display: block;
    background: #000;
  }
</style>