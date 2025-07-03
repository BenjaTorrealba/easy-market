<script>
  export let mensaje = "¿Estás seguro?";
  export let onConfirm = () => {};
  export let onCancel = () => {};
</script>

<div class="fixed inset-0 bg-black/40 z-50 flex items-center justify-center min-h-screen" on:click={onCancel}>
  <div class="bg-white rounded-lg shadow-xl p-8 w-full max-w-sm relative" on:click|stopPropagation>
    <h2 class="text-lg font-bold mb-4 text-gray-800">{mensaje}</h2>
    <div class="flex justify-end gap-4">
      <button class="px-4 py-2 rounded bg-gray-200 hover:bg-gray-300 text-gray-700" on:click={onCancel}>Cancelar</button>
      <button class="px-4 py-2 rounded bg-red-600 hover:bg-red-700 text-white" on:click={onConfirm}>Confirmar</button>
    </div>
  </div>
</div>