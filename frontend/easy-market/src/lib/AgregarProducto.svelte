<script>
  import { onMount, onDestroy } from "svelte";
  import { api } from "$lib/api.js";
  import { iniciarScanner, detenerScanner } from "$lib/codigoBarras.js";
  export let onAgregar;
  export let onClose;

  let nombre = "";
  let precio = "";
  let stock = "";
  let codigoBarras = "";
  let categoriaID = "";
  let categorias = [];

  // Nueva función para limpiar todos los campos
  function limpiarCampos() {
    nombre = "";
    precio = "";
    stock = "";
    codigoBarras = "";
    categoriaID = categorias.length > 0 ? categorias[0].id : "";
  }

  // Nueva función para cerrar y limpiar
  function closeAndReset() {
    detenerScanner();
    limpiarCampos();
    onClose();
  }

  onMount(async () => {
    categorias = await api.getCategorias();
    limpiarCampos();

    // Inicia el escáner SOLO después de que el div esté en el DOM
    iniciarScanner((codigo) => {
      codigoBarras = codigo;
      detenerScanner();
    });
  });

  onDestroy(() => {
    detenerScanner();
    limpiarCampos();
  });

  function agregar() {
    console.log("Intentando agregar producto...");
    console.log({ nombre, precio, stock, categoriaID, codigoBarras });

    if (!nombre || !precio || !stock || !categoriaID) {
      console.log("Faltan campos obligatorios");
      return;
    }
    const producto = {
      nombre,
      precio: Number(precio),
      stock: Number(stock),
      codigo_barras: codigoBarras,
      categoria_id: Number(categoriaID),
    };
    console.log("Producto a agregar:", producto);

    onAgregar(producto);
    detenerScanner();
    limpiarCampos();
    onClose();
  }
  function volverAEscanear() {
    codigoBarras = "";
    detenerScanner();
    iniciarScanner((codigo) => {
      codigoBarras = codigo;
      detenerScanner();
    });
  }
</script>

<div
  class="fixed inset-0 bg-black/40 z-50 flex items-center justify-center min-h-screen"
  on:click={closeAndReset}
>
  <div
    class="bg-white rounded-lg shadow-xl p-8 w-full max-w-lg relative"
    on:click|stopPropagation
  >
    <button
      class="absolute top-2 right-2 text-gray-400 hover:text-gray-700 text-xl"
      on:click={closeAndReset}>&times;</button
    >
    <h2 class="text-xl font-bold text-indigo-800 mb-4">
      Agregar nuevo producto
    </h2>
    <div class="space-y-4">
      <input
        class="border rounded px-3 py-2 focus:outline-none focus:ring w-full"
        placeholder="Nombre"
        bind:value={nombre}
      />
      <select
        class="border rounded px-3 py-2 focus:outline-none focus:ring w-full"
        bind:value={categoriaID}
      >
        {#each categorias as cat}
          <option value={cat.id}>{cat.nombre}</option>
        {/each}
      </select>
      <input
        type="number"
        min="0"
        class="border rounded px-3 py-2 focus:outline-none focus:ring w-full"
        placeholder="Precio"
        bind:value={precio}
      />
      <input
        type="number"
        min="0"
        class="border rounded px-3 py-2 focus:outline-none focus:ring w-full"
        placeholder="Stock"
        bind:value={stock}
      />
      <input
        class="border rounded px-3 py-2 focus:outline-none focus:ring w-full"
        placeholder="Código de barras"
        bind:value={codigoBarras}
      />
      <button
        type="button"
        class="bg-gray-200 hover:bg-gray-300 text-gray-800 px-4 py-2 rounded mb-2 w-full"
        on:click={volverAEscanear}
      >
        Volver a escanear código de barras
      </button>
      <!-- Contenedor para el escáner -->
      <div id="scanner-container"></div>
      <button
        class="bg-indigo-600 hover:bg-indigo-700 text-white px-6 py-2 rounded shadow font-semibold transition w-full"
        on:click={agregar}
      >
        Agregar producto
      </button>
    </div>
  </div>
</div>

<style>
  #scanner-container {
    position: relative;
    width: 100%;
    max-width: 320px;
    height: 160px;
    margin: 0 auto 1rem auto;
    display: block;
    border-radius: 8px;
    overflow: hidden;
  }
  #scanner-container video,
  #scanner-container canvas {
    position: static !important;
    pointer-events: none !important;
    width: 100% !important;
    height: 100% !important;
    z-index: 1 !important;
    display: block;
  }
</style>
