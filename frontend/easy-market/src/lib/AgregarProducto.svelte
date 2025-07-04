<script>
  import { onMount, onDestroy, tick } from "svelte";
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
  let mostrarScanner = false;

  function limpiarCampos() {
    nombre = "";
    precio = "";
    stock = "";
    codigoBarras = "";
    categoriaID = categorias.length > 0 ? categorias[0].id : "";
  }

  function closeAndReset() {
    detenerScanner();
    mostrarScanner = false;
    limpiarCampos();
    onClose();
  }

  onMount(async () => {
    categorias = await api.getCategorias();
    limpiarCampos();
  });

  onDestroy(() => {
    detenerScanner();
    mostrarScanner = false;
    limpiarCampos();
  });

  function agregar() {
    if (!nombre || !precio || !stock || !categoriaID) return;
    const producto = {
      nombre,
      precio: Number(precio),
      stock: Number(stock),
      codigo_barras: codigoBarras,
      categoria_id: Number(categoriaID),
    };
    onAgregar(producto);
    detenerScanner();
    mostrarScanner = false;
    limpiarCampos();
    onClose();
  }

  async function iniciarEscaner() {
    mostrarScanner = true;
    await tick(); // Espera a que el DOM actualice y el div esté presente
    iniciarScanner((codigo) => {
      codigoBarras = codigo;
      detenerScanner();
      mostrarScanner = false;
    });
  }

  function detenerEscaner() {
    detenerScanner();
    mostrarScanner = false;
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
    <h2 class="text-xl font-bold text-gray-700 mb-4">
      Agregar nuevo producto
    </h2>
    <div class="space-y-4">
      <input
        class="border border-gray-400 rounded px-3 py-2 focus:outline-none focus:ring-2 focus:ring-green-200 w-full"
        placeholder="Nombre"
        bind:value={nombre}
      />
      <select
        class="border border-gray-400 rounded px-3 py-2 focus:outline-none focus:ring-2 focus:ring-green-200 w-full"
        bind:value={categoriaID}
      >
        {#each categorias as cat}
          <option value={cat.id}>{cat.nombre}</option>
        {/each}
      </select>
      <input
        type="number"
        min="0"
        class="border border-gray-400 rounded px-3 py-2 focus:outline-none focus:ring-2 focus:ring-green-200 w-full"
        placeholder="Precio"
        bind:value={precio}
      />
      <input
        type="number"
        min="0"
        class="border border-gray-400 rounded px-3 py-2 focus:outline-none focus:ring-2 focus:ring-green-200 w-full"
        placeholder="Stock"
        bind:value={stock}
      />
      <input
        class="border border-gray-400 rounded px-3 py-2 focus:outline-none focus:ring-2 focus:ring-green-200 w-full"
        placeholder="Código de barras"
        bind:value={codigoBarras}
      />
      <div class="flex gap-2 mb-2">
        <button
          type="button"
          class="bg-green-300 hover:bg-green-500 text-gray-700 px-4 py-2 rounded w-full transition"
          on:click={iniciarEscaner}
          disabled={mostrarScanner}
        >
          Iniciar escáner
        </button>
        <button
          type="button"
          class="bg-red-100 hover:bg-red-200 text-red-700 px-4 py-2 rounded w-full transition"
          on:click={detenerEscaner}
          disabled={!mostrarScanner}
        >
          Detener escáner
        </button>
      </div>
      {#if mostrarScanner}
        <div id="scanner-container" class="scanner-video"></div>
      {/if}
      <button
        class="bg-green-300 hover:bg-green-500 text-gray-700 px-6 py-2 rounded shadow font-semibold transition w-full"
        on:click={agregar}
      >
        Agregar producto
      </button>
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