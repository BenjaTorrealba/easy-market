<script>
  import { onMount } from "svelte";
  import { api } from "$lib/api.js";
  import { getRangoSemana } from "$lib/reporte.js";

  let historialVentas = [];
  let filtroInicio = "";
  let filtroFin = "";
  let historialFiltrado = [];

  function formatearFechaHora(fechaIso) {
    const fecha = new Date(fechaIso);
    const dia = String(fecha.getDate()).padStart(2, "0");
    const mes = String(fecha.getMonth() + 1).padStart(2, "0");
    const anio = fecha.getFullYear();
    const hora = String(fecha.getHours()).padStart(2, "0");
    const min = String(fecha.getMinutes()).padStart(2, "0");
    return `${dia}/${mes}/${anio} ${hora}:${min}`;
  }

  async function filtrarHistorial() {
    historialFiltrado = await api.getHistorialVentas(filtroInicio, filtroFin);
  }

  onMount(async () => {
    const { fecha_inicio, fecha_fin } = getRangoSemana();
    historialVentas = await api.getHistorialVentas(fecha_inicio, fecha_fin);
    filtroInicio = fecha_inicio;
    filtroFin = fecha_fin;
    historialFiltrado = historialVentas;
  });
</script>

<h2 class="text-xl font-bold mb-4 text-green-700">Historial de ventas</h2>
<div class="flex gap-4 mb-4">
  <div>
    <label class="block text-sm font-semibold mb-1">Desde:</label>
    <input
      type="date"
      bind:value={filtroInicio}
      class="border border-gray-400 rounded px-2 py-1 focus:outline-none focus:ring-2 focus:ring-green-200"
    />
  </div>
  <div>
    <label class="block text-sm font-semibold mb-1">Hasta:</label>
    <input
      type="date"
      bind:value={filtroFin}
      class="border border-gray-400 rounded px-2 py-1 focus:outline-none focus:ring-2 focus:ring-green-200"
    />
  </div>
  <button
    class="bg-green-300 hover:bg-green-500 text-gray-700 px-4 py-2 rounded mt-5 transition"
    on:click={filtrarHistorial}
  >
    Filtrar
  </button>
</div>
<div class="overflow-x-auto">
  <table class="min-w-full rounded-lg shadow border border-green-200 bg-white">
    <thead>
      <tr>
        <th
          class="px-6 py-3 text-left bg-green-100 text-green-700 font-semibold"
          >ID</th
        >
        <th
          class="px-6 py-3 text-left bg-green-100 text-green-700 font-semibold"
          >Fecha</th
        >
        <th
          class="px-6 py-3 text-left bg-green-100 text-green-700 font-semibold"
          >Total</th
        >
      </tr>
    </thead>
    <tbody>
      {#each historialFiltrado as venta, i}
        <tr
          class={`border-t transition ${i % 2 === 0 ? "bg-white" : "bg-gray-100"} hover:bg-green-100/40`}
        >
          <td class="border-t px-6 py-3">{venta.id}</td>
          <td class="border-t px-6 py-3">{formatearFechaHora(venta.fecha)}</td>
          <td class="border-t px-6 py-3">${venta.total.toLocaleString()}</td>
        </tr>
      {/each}
    </tbody>
  </table>
</div>
