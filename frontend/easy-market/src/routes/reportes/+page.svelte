<script>
  // Mock data
  let ventasSemana = [
    { dia: "Lunes", cantidad: 12, total: 12000 },
    { dia: "Martes", cantidad: 8, total: 8000 },
    { dia: "Miércoles", cantidad: 15, total: 15000 },
    { dia: "Jueves", cantidad: 10, total: 10000 },
    { dia: "Viernes", cantidad: 20, total: 20000 },
    { dia: "Sábado", cantidad: 18, total: 18000 },
    { dia: "Domingo", cantidad: 7, total: 7000 }
  ];

  let productosMasVendidos = [
    { nombre: "Refresco", cantidad: 30 },
    { nombre: "Galletas", cantidad: 25 },
    { nombre: "Pan de caja", cantidad: 18 }
  ];

  let productosBajoStock = [
    { nombre: "Pan de caja", stock: 2 },
    { nombre: "Leche", stock: 1 }
  ];

  // Cálculos
  $: totalVendidos = ventasSemana.reduce((sum, v) => sum + v.cantidad, 0);
  $: totalVentas = ventasSemana.reduce((sum, v) => sum + v.total, 0);
</script>

<div class="p-8 space-y-10">
  <h1 class="text-3xl font-bold mb-6 text-blue-900">Dashboard de Reportes</h1>

  <!-- Métricas principales -->
  <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
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

  <!-- Productos con bajo stock -->
  <div>
    <h2 class="text-xl font-bold mb-4 text-yellow-800">Productos con bajo stock</h2>
    <ul class="list-disc pl-6 space-y-1">
      {#each productosBajoStock as producto}
        <li class="text-yellow-700">{producto.nombre} <span class="font-semibold">(Stock: {producto.stock})</span></li>
      {/each}
      {#if productosBajoStock.length === 0}
        <li class="text-gray-400">No hay productos con bajo stock.</li>
      {/if}
    </ul>
  </div>
</div>