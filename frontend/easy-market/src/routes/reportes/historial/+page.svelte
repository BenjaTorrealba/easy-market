<script>
  import { onMount } from 'svelte';
  import { api } from '$lib/api.js';
  import { getRangoSemana } from '$lib/reporte.js'; // <--- Importa aquí

  let historialVentas = [];
  let filtroInicio = '';
  let filtroFin = '';
  let historialFiltrado = [];

  async function filtrarHistorial() {
    historialFiltrado = await api.getHistorialVentas(filtroInicio, filtroFin);
  }

  onMount(async () => {
    const { fecha_inicio, fecha_fin } = getRangoSemana(); // <--- Usa aquí
    historialVentas = await api.getHistorialVentas(fecha_inicio, fecha_fin);
    filtroInicio = fecha_inicio;
    filtroFin = fecha_fin;
    historialFiltrado = historialVentas;
  });
</script>

<h2 class="text-xl font-bold mb-4 text-blue-800">Historial de ventas</h2>
<div class="flex gap-4 mb-4">
  <div>
    <label class="block text-sm font-semibold mb-1">Desde:</label>
    <input type="date" bind:value={filtroInicio} class="border rounded px-2 py-1" />
  </div>
  <div>
    <label class="block text-sm font-semibold mb-1">Hasta:</label>
    <input type="date" bind:value={filtroFin} class="border rounded px-2 py-1" />
  </div>
  <button class="bg-blue-600 text-white px-4 py-2 rounded mt-5" on:click={filtrarHistorial}>
    Filtrar
  </button>
</div>
<div class="overflow-x-auto">
  <table class="min-w-full rounded-lg shadow border border-blue-200 bg-white">
    <thead>
      <tr>
        <th class="px-6 py-3 text-left bg-blue-100 text-blue-800 font-semibold">ID</th>
        <th class="px-6 py-3 text-left bg-blue-100 text-blue-800 font-semibold">Fecha</th>
        <th class="px-6 py-3 text-left bg-blue-100 text-blue-800 font-semibold">Total</th>
      </tr>
    </thead>
    <tbody>
      {#each historialFiltrado as venta}
        <tr class="hover:bg-blue-100/40 transition">
          <td class="border-t px-6 py-3">{venta.id}</td>
          <td class="border-t px-6 py-3">{venta.fecha.slice(0, 16).replace('T', ' ')}</td>
          <td class="border-t px-6 py-3">${venta.total.toLocaleString()}</td>
        </tr>
      {/each}
    </tbody>
  </table>
</div>