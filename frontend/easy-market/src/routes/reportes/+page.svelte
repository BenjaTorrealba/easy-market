<script>
  import { onMount } from 'svelte';
  import { api } from '$lib/api.js';

  let ventasTotales = { dia: 0, semana: 0, mes: 0 };
  let historialVentas = [];
  let productosBajoStock = [];
  let productosMasVendidos = [];
  let ventasSemana = [];
  let filtroInicio = '';
  let filtroFin = '';
  let historialFiltrado = [];

  let activeTab = 'productos';

  $: totalVendidos = ventasSemana.reduce((sum, v) => sum + v.cantidad, 0);
  $: totalVentas = ventasSemana.reduce((sum, v) => sum + v.total, 0);

  async function filtrarHistorial() {
    historialFiltrado = await api.getHistorialVentas(filtroInicio, filtroFin);
  }

  onMount(async () => {
    ventasTotales = await api.getVentasTotales();
    productosBajoStock = await api.getProductosBajoStock();

    // Historial de ventas de los últimos 7 días
    const hoy = new Date();
    const hace7 = new Date();
    hace7.setDate(hoy.getDate() - 6);
    const fecha_inicio = hace7.toISOString().slice(0, 10);
    const fecha_fin = hoy.toISOString().slice(0, 10);

    historialVentas = await api.getHistorialVentas(fecha_inicio, fecha_fin);
    filtroInicio = fecha_inicio;
    filtroFin = fecha_fin;
    historialFiltrado = historialVentas;

    // Procesar ventasSemana para la gráfica
    const dias = ['Domingo', 'Lunes', 'Martes', 'Miércoles', 'Jueves', 'Viernes', 'Sabado'];
    let ventasPorDia = {};
    for (let i = 0; i < 7; i++) {
      const d = new Date(hace7);
      d.setDate(hace7.getDate() + i);
      const key = d.toISOString().slice(0, 10);
      ventasPorDia[key] = { dia: dias[d.getDay()], cantidad: 0, total: 0 };
    }
    for (const venta of historialVentas) {
      const fecha = venta.fecha.slice(0, 10);
      if (ventasPorDia[fecha]) {
        ventasPorDia[fecha].cantidad += 1;
        ventasPorDia[fecha].total += venta.total;
      }
    }
    ventasSemana = Object.values(ventasPorDia);

    // Si tienes endpoint de productos más vendidos, descomenta:
    // productosMasVendidos = await api.getProductosMasVendidos();
  });
</script>

<div class="p-8 space-y-10">
  <h1 class="text-3xl font-bold mb-6 text-blue-900">Dashboard de Reportes</h1>

  <!-- Métricas principales -->
  <div class="grid grid-cols-1 md:grid-cols-3 gap-6 mb-8">
    <div class="bg-blue-50 border-l-4 border-blue-400 rounded-lg p-6 shadow flex flex-col items-center">
      <span class="text-5xl font-bold text-blue-700">{totalVendidos}</span>
      <span class="text-lg text-blue-900 mt-2">Productos vendidos esta semana</span>
    </div>
    <div class="bg-green-50 border-l-4 border-green-400 rounded-lg p-6 shadow flex flex-col items-center">
      <span class="text-5xl font-bold text-green-700">${totalVentas.toLocaleString()}</span>
      <span class="text-lg text-green-900 mt-2">Total ventas esta semana</span>
    </div>
    <div class="bg-yellow-50 border-l-4 border-yellow-400 rounded-lg p-6 shadow flex flex-col items-center">
      <span class="text-5xl font-bold text-yellow-700">{productosBajoStock.length}</span>
      <span class="text-lg text-yellow-900 mt-2">Productos con bajo stock</span>
    </div>
  </div>

  <!-- Tabs -->
<div class="mb-6 flex w-full gap-2 border-b">
  <button
    class="flex-1 px-4 py-2 font-semibold border-b-2 transition-colors"
    class:border-blue-600={activeTab === 'productos'}
    class:text-blue-700={activeTab === 'productos'}
    on:click={() => activeTab = 'productos'}
  >Productos más vendidos</button>
  <button
    class="flex-1 px-4 py-2 font-semibold border-b-2 transition-colors"
    class:border-blue-600={activeTab === 'ventas'}
    class:text-blue-700={activeTab === 'ventas'}
    on:click={() => activeTab = 'ventas'}
  >Ventas por día</button>
  <button
    class="flex-1 px-4 py-2 font-semibold border-b-2 transition-colors"
    class:border-blue-600={activeTab === 'historial'}
    class:text-blue-700={activeTab === 'historial'}
    on:click={() => activeTab = 'historial'}
  >Historial de ventas</button>
</div>

  {#if activeTab === 'productos'}
    <!-- Productos más vendidos -->
    <div>
      <h2 class="text-xl font-bold mb-4 text-blue-800">Productos más vendidos</h2>
      <div class="overflow-x-auto">
        <table class="min-w-full rounded-lg shadow border border-blue-200 bg-white">
          <thead>
            <tr>
              <th class="px-6 py-3 text-left bg-blue-100 text-blue-800 font-semibold">Producto</th>
              <th class="px-6 py-3 text-left bg-blue-100 text-blue-800 font-semibold">Cantidad vendida</th>
            </tr>
          </thead>
          <tbody>
            {#each productosMasVendidos as producto}
              <tr class="hover:bg-blue-100/40 transition">
                <td class="border-t px-6 py-3 font-semibold">{producto.nombre}</td>
                <td class="border-t px-6 py-3">{producto.cantidad}</td>
              </tr>
            {/each}
          </tbody>
        </table>
      </div>
    </div>
  {/if}

  {#if activeTab === 'ventas'}
    <!-- Ventas por día (gráfica simple de barras) -->
    <div>
      <h2 class="text-xl font-bold mb-4 text-blue-800">Ventas por día</h2>
      <div class="space-y-2">
        {#each ventasSemana as venta}
          <div>
            <div class="flex justify-between text-sm mb-1">
              <span class="font-semibold">{venta.dia}</span>
              <span>{venta.cantidad} productos / ${venta.total.toLocaleString()}</span>
            </div>
            <div class="bg-blue-100 rounded h-4 relative">
              <div
                class="bg-blue-500 h-4 rounded"
                style="width: {Math.max(venta.cantidad * 5, 10)}px; max-width: 100%;"
              ></div>
            </div>
          </div>
        {/each}
      </div>
    </div>
  {/if}

  {#if activeTab === 'historial'}
    <!-- Historial de ventas -->
    <div>
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
    </div>
  {/if}
</div>