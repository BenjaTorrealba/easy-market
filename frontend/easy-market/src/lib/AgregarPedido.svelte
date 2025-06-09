<script>
  export let onAgregar;
  export let onClose;

  let distribuidor = "";
  let productosInput = "";
  let fechaEstimada = "";

  function agregar() {
    if (!distribuidor || !productosInput || !fechaEstimada) return;
    onAgregar({
      id: Math.floor(Math.random() * 10000),
      distribuidor,
      productos: productosInput.split(",").map(p => p.trim()).filter(Boolean),
      estado: "Pendiente",
      fecha: new Date().toISOString().slice(0, 10),
      fechaEstimada
    });
    distribuidor = "";
    productosInput = "";
    fechaEstimada = "";
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
    on:keydown={(e) => { if (e.key === 'Enter' || e.key === ' ') { onClose(); } }}
    tabindex="0"
    style="z-index:51;"
  ></button>
  <div
    class="bg-white rounded-lg shadow-xl p-8 w-full max-w-lg relative"
    role="dialog"
    aria-modal="true"
    tabindex="-1"
    on:click|stopPropagation
    on:keydown|stopPropagation
    style="z-index:52;"
  >
    <button class="absolute top-2 right-2 text-gray-400 hover:text-gray-700 text-xl" on:click={onClose} aria-label="Cerrar modal">&times;</button>
    <h2 class="text-xl font-bold text-blue-800 mb-4">Agregar nuevo pedido</h2>
    <div class="space-y-4">
      <input
        class="border rounded px-3 py-2 focus:outline-none focus:ring w-full"
        placeholder="Distribuidor"
        bind:value={distribuidor}
      />
      <input
        class="border rounded px-3 py-2 focus:outline-none focus:ring w-full"
        placeholder="Productos (separados por coma)"
        bind:value={productosInput}
      />
      <input
        type="date"
        class="border rounded px-3 py-2 focus:outline-none focus:ring w-full"
        placeholder="Fecha estimada de llegada"
        bind:value={fechaEstimada}
      />
      <button
        class="bg-blue-600 hover:bg-blue-700 text-white px-6 py-2 rounded shadow font-semibold transition w-full"
        on:click={agregar}
      >
        Agregar pedido
      </button>
    </div>
  </div>
</div>