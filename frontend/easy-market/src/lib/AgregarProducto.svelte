<script>
  export let onAgregar;
  export let onClose;

  let nombre = "";
  let precio = "";
  let stock = "";
  let codigoBarras = "";

  function agregar() {
    if (!nombre || !precio || !stock || !codigoBarras) return;
    onAgregar({
      id: Math.floor(Math.random() * 10000),
      nombre,
      precio: Number(precio),
      stock: Number(stock),
      codigoBarras
    });
    nombre = "";
    precio = "";
    stock = "";
    codigoBarras = "";
    onClose();
  }
</script>

<div class="fixed inset-0 bg-black/40 z-50 flex items-center justify-center min-h-screen" on:click={onClose}>
  <div class="bg-white rounded-lg shadow-xl p-8 w-full max-w-lg relative" on:click|stopPropagation>
    <button class="absolute top-2 right-2 text-gray-400 hover:text-gray-700 text-xl" on:click={onClose}>&times;</button>
    <h2 class="text-xl font-bold text-indigo-800 mb-4">Agregar nuevo producto</h2>
    <div class="space-y-4">
      <input
        class="border rounded px-3 py-2 focus:outline-none focus:ring w-full"
        placeholder="Nombre"
        bind:value={nombre}
      />
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
        class="bg-indigo-600 hover:bg-indigo-700 text-white px-6 py-2 rounded shadow font-semibold transition w-full"
        on:click={agregar}
      >
        Agregar producto
      </button>
    </div>
  </div>
</div>