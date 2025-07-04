<script>
  import { onMount } from 'svelte';
  import { api } from '$lib/api.js';

  let productosMasVendidos = [];

  onMount(async () => {
    productosMasVendidos = await api.getProductosMasVendidos();
  });
</script>

<h2 class="text-xl font-bold mb-4 text-green-700">Productos más vendidos</h2>
<div class="overflow-x-auto">
  <table class="min-w-full rounded-lg shadow border border-green-200 bg-white">
    <thead>
      <tr>
        <th class="px-6 py-3 text-left bg-green-100 text-green-700 font-semibold">Producto</th>
        <th class="px-6 py-3 text-left bg-green-100 text-green-700 font-semibold">Cantidad vendida</th>
      </tr>
    </thead>
    <tbody>
      {#each productosMasVendidos as producto}
        <tr class="hover:bg-green-100/40 transition">
          <td class="border-t px-6 py-3 font-semibold">{producto.nombre}</td>
          <td class="border-t px-6 py-3">{producto.cantidad}</td>
        </tr>
      {/each}
    </tbody>
  </table>
</div>