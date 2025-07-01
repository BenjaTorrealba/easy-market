<script>
  import { api } from "$lib/api.js";
  import { onMount } from "svelte";

  export let onAgregar;
  export let onClose;
  export let categorias = [];

  let productosDisponibles = [];
  let distribuidor = "";
  let fechaEstimada = "";
  let seleccionados = [];
  let fechaPedido = new Date().toISOString().slice(0, 10);

  // Cargar productos al montar el componente
  onMount(async () => {
    productosDisponibles = await api.getProductos();
  });

  function agregarProducto(producto) {
    if (!seleccionados.find((p) => p.id === producto.id)) {
      seleccionados = [...seleccionados, { ...producto, cantidad: 1 }];
    }
  }

  function quitarProducto(producto) {
    seleccionados = seleccionados.filter((p) => p.id !== producto.id);
  }

  function cambiarCantidad(producto, cantidad) {
    seleccionados = seleccionados.map((p) =>
      p.id === producto.id
        ? { ...p, cantidad: Math.max(1, Number(cantidad) || 1) }
        : p,
    );
  }

  function obtenerNombreCategoria(id) {
    const cat = categorias.find((c) => c.id === id);
    return cat ? cat.nombre : id;
  }

    function agregar() {
    if (!distribuidor || seleccionados.length === 0 || !fechaEstimada) return;

    // Transformar productos al formato correcto
    const productosParaEnvio = seleccionados.map(p => ({
      producto_id: p.id,
      cantidad: p.cantidad
    }));

    onAgregar({
      distribuidor,
      fecha: fechaPedido,
      fechaEstimada,
      estado: "Pendiente",
      productos: productosParaEnvio
    });

    distribuidor = "";
    fechaEstimada = "";
    seleccionados = [];
    onClose();
  }
</script>

<div
  class="fixed inset-0 bg-black/40 z-50 flex items-center justify-center"
  role="dialog"
  aria-modal="true"
>
  <button
    type="button"
    class="fixed inset-0 w-full h-full bg-transparent cursor-default"
    aria-label="Cerrar modal"
    on:click={onClose}
    tabindex="0"
    style="z-index:51;"
  ></button>
  <div
    class="bg-white rounded-lg shadow-xl p-8 w-full max-w-3xl relative flex flex-col gap-6"
    style="z-index:52;">
    <button
      class="absolute top-2 right-2 text-gray-400 hover:text-gray-700 text-xl"
      on:click={onClose}
      aria-label="Cerrar modal">&times;</button>
    <h2 class="text-xl font-bold text-blue-800 mb-2">Agregar nuevo pedido</h2>
    <div class="flex gap-6">
      <!-- Lista de productos disponibles -->
      <div class="w-1/2 border rounded p-4">
        <h3 class="font-semibold mb-2">Productos disponibles</h3>
        <ul class="space-y-2 max-h-64 overflow-y-auto">
          {#each productosDisponibles as producto (producto.id)}
            {#if !seleccionados.find((p) => p.id === producto.id)}
              <li class="flex justify-between items-center border-b pb-1">
                <div>
                  <div class="font-semibold">{producto.nombre}</div>
                  <div class="text-s text-gray-500">
                    Precio: ${producto.precio?.toLocaleString() ?? "-"}
                    | Stock: {producto.stock ?? "-"}
                  </div>
                </div>
                <button
                  class="bg-green-500 text-white px-2 py-1 rounded text-xs"
                  on:click={() => agregarProducto(producto)}>
                  Agregar
                </button>
              </li>
            {/if}
          {/each}
        </ul>
      </div>
      <!-- Lista de productos seleccionados -->
      <div class="w-1/2 border rounded p-4">
        <h3 class="font-semibold mb-2">Productos a pedir</h3>
        <ul class="space-y-2 max-h-64 overflow-y-auto">
          {#each seleccionados as producto (producto.id)}
            <li class="flex justify-between items-center border-b pb-1">
              <div>
                <div class="font-semibold">{producto.nombre}</div>
                <div class="text-xs text-gray-500">
                  Precio: ${producto.precio?.toLocaleString() ?? "-"}
                  | Stock: {producto.stock ?? "-"}
                </div>
              </div>
              <div class="flex items-center gap-2">
                <label class="text-xs text-gray-600">Cantidad:</label>
                <input
                  type="number"
                  min="1"
                  class="border rounded px-1 py-0.5 w-10 text-xs"
                  bind:value={producto.cantidad}
                  on:input={(e) => cambiarCantidad(producto, e.target.value)}
                  style="font-size: 0.8rem;"
                />
                <button
                  class="bg-red-500 text-white px-2 py-1 rounded text-xs ml-2"
                  on:click={() => quitarProducto(producto)}
                >
                  Quitar
                </button>
              </div>
            </li>
          {/each}
        </ul>
      </div>
    </div>
    <!-- Campos de distribuidor y fechas -->
    <div class="flex flex-col gap-3 mt-4">
      <input
        class="border rounded px-3 py-2 focus:outline-none focus:ring w-full"
        placeholder="Distribuidor"
        bind:value={distribuidor}
      />
      <div class="flex gap-4">
        <div class="flex-1">
          <label class="block text-sm text-gray-600 mb-1">Fecha de pedido</label>
          <input
            type="date"
            class="border rounded px-3 py-2 focus:outline-none focus:ring w-full"
            bind:value={fechaPedido}
            disabled
          />
        </div>
        <div class="flex-1">
          <label class="block text-sm text-gray-600 mb-1"
            >Fecha estimada de llegada</label>
          <input
            type="date"
            class="border rounded px-3 py-2 focus:outline-none focus:ring w-full"
            bind:value={fechaEstimada}
          />
        </div>
      </div>
      <button
        class="bg-blue-600 hover:bg-blue-700 text-white px-6 py-2 rounded shadow font-semibold transition w-full mt-2"
        on:click={agregar}>Agregar pedido</button>
    </div>
  </div>
</div>
