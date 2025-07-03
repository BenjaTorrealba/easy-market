<script>
  import { goto } from "$app/navigation";
  import { page } from "$app/stores";
  import { onMount } from "svelte";
  import { api } from "$lib/api.js";

  $: path = $page.url.pathname;

  // Métricas principales
  let ventasTotales = { dia: 0, semana: 0, mes: 0 };
  let productosBajoStock = [];
  let ventasSemana = [];

  $: totalVendidos = ventasSemana.reduce((sum, v) => sum + v.cantidad, 0);
  $: totalVentas = ventasSemana.reduce((sum, v) => sum + v.total, 0);

  onMount(async () => {
    ventasTotales = await api.getVentasTotales();
    productosBajoStock = await api.getProductosBajoStock();

    // Para productos vendidos esta semana
    const hoy = new Date();
    const hace7 = new Date();
    hace7.setDate(hoy.getDate() - 6);
    const fecha_inicio = hace7.toISOString().slice(0, 10);
    const fecha_fin = hoy.toISOString().slice(0, 10);

    const historialVentas = await api.getHistorialVentas(
      fecha_inicio,
      fecha_fin
    );

    const dias = [
      "Domingo",
      "Lunes",
      "Martes",
      "Miércoles",
      "Jueves",
      "Viernes",
      "Sabado",
    ];
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
  });
</script>

<!-- DASHBOARD HEADER -->
<div class="p-4 pb-0">
  <h1 class="text-3xl font-bold mb-6 text-blue-900">Dashboard de Reportes</h1>
</div>

<!-- MÉTRICAS PRINCIPALES -->
<div class="grid grid-cols-1 md:grid-cols-3 gap-6 mb-8">
  <div
    class="bg-blue-50 border-l-4 border-blue-400 rounded-lg p-6 shadow flex flex-col items-center"
  >
    <span class="text-5xl font-bold text-blue-700">{totalVendidos}</span>
    <span class="text-lg text-blue-900 mt-2"
      >Productos vendidos esta semana</span
    >
  </div>
  <div
    class="bg-green-50 border-l-4 border-green-400 rounded-lg p-6 shadow flex flex-col items-center"
  >
    <span class="text-5xl font-bold text-green-700"
      >${totalVentas.toLocaleString()}</span
    >
    <span class="text-lg text-green-900 mt-2">Total ventas esta semana</span>
  </div>
  <div
    class="bg-yellow-50 border-l-4 border-yellow-400 rounded-lg p-6 shadow flex flex-col items-center"
  >
    <span class="text-5xl font-bold text-yellow-700"
      >{productosBajoStock.length}</span
    >
    <span class="text-lg text-yellow-900 mt-2">Productos con bajo stock</span>
  </div>
</div>

<!-- BARRA DE TABS -->
<div class="mb-6 flex w-full gap-2 border-b">
  <a
    class="flex-1 px-4 py-2 font-semibold border-b-2 text-center transition-colors cursor-pointer"
    class:border-blue-600={path.endsWith("/productos")}
    class:text-blue-700={path.endsWith("/productos")}
    on:click={() => goto("/reportes/productos")}>Productos más vendidos</a
  >
  <a
    class="flex-1 px-4 py-2 font-semibold border-b-2 text-center transition-colors cursor-pointer"
    class:border-blue-600={path === "/reportes" || path.endsWith("/ventas")}
    class:text-blue-700={path === "/reportes" || path.endsWith("/ventas")}
    on:click={() => goto("/reportes")}>Ventas por día</a
  >
  <a
    class="flex-1 px-4 py-2 font-semibold border-b-2 text-center transition-colors cursor-pointer"
    class:border-blue-600={path.endsWith("/historial")}
    class:text-blue-700={path.endsWith("/historial")}
    on:click={() => goto("/reportes/historial")}>Historial de ventas</a
  >
</div>

<slot />
