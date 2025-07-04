<script>
  export let producto;
  export let categorias = [];
  export let onGuardar;
  export let onClose;

  let nombre = "";
  let precio = "";
  let stock = "";
  let codigoBarras = "";
  let categoriaID = "";

  $: if (
    producto &&
    nombre === "" &&
    precio === "" &&
    stock === "" &&
    codigoBarras === "" &&
    categoriaID === ""
  ) {
    nombre = producto.nombre ?? "";
    precio = producto.precio ?? "";
    stock = producto.stock ?? "";
    codigoBarras = producto.codigo_barras ?? "";
    categoriaID = producto.categoria_id ?? categorias[0]?.id ?? "";
  }

  function guardar() {
    if (!nombre || precio === "" || stock === "" || !categoriaID) return;
    onGuardar({
      ...producto,
      nombre,
      precio: Number(precio),
      stock: Number(stock),
      codigo_barras: codigoBarras,
      categoria_id: Number(categoriaID),
    });
    onClose();

    nombre = "";
    precio = "";
    stock = "";
    codigoBarras = "";
    categoriaID = "";
  }
</script>

<div
  class="fixed inset-0 bg-black/40 z-50 flex items-center justify-center min-h-screen"
  on:click={onClose}
>
  <div
    class="bg-white rounded-lg shadow-xl p-8 w-full max-w-md relative"
    on:click|stopPropagation
  >
    <button
      class="absolute top-2 right-2 text-gray-400 hover:text-gray-700 text-xl"
      on:click={onClose}>&times;</button
    >
    <h2 class="text-xl font-bold text-green-700 mb-4">Editar producto</h2>
    <div class="space-y-4">
      <div>
        <label class="block text-gray-700 mb-1">Nombre</label>
        <input
          class="border border-gray-400 rounded px-3 py-2 w-full bg-gray-100"
          bind:value={nombre}
          disabled
        />
      </div>
      <div>
        <label class="block text-gray-700 mb-1">Categoría</label>
        <select
          class="border border-gray-400 rounded px-3 py-2 w-full focus:outline-none focus:ring-2 focus:ring-green-200"
          bind:value={categoriaID}
        >
          {#each categorias as cat}
            <option value={cat.id}>{cat.nombre}</option>
          {/each}
        </select>
      </div>
      <div>
        <label class="block text-gray-700 mb-1">Precio</label>
        <input
          type="number"
          min="0"
          class="border border-gray-400 rounded px-3 py-2 w-full focus:outline-none focus:ring-2 focus:ring-green-200"
          bind:value={precio}
        />
      </div>
      <div>
        <label class="block text-gray-700 mb-1">Stock</label>
        <input
          type="number"
          min="0"
          class="border border-gray-400 rounded px-3 py-2 w-full focus:outline-none focus:ring-2 focus:ring-green-200"
          bind:value={stock}
        />
      </div>
      <div>
        <label class="block text-gray-700 mb-1">Código de barras</label>
        <input
          class="border border-gray-400 rounded px-3 py-2 w-full focus:outline-none focus:ring-2 focus:ring-green-200"
          bind:value={codigoBarras}
        />
      </div>
      <button
        class="bg-green-300 hover:bg-green-500 text-gray-700 px-6 py-2 rounded shadow font-semibold transition w-full"
        on:click={guardar}
      >
        Guardar cambios
      </button>
    </div>
  </div>
</div>
