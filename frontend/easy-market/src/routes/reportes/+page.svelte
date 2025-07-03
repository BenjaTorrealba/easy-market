<script>
  import { onMount } from "svelte";
  import { api } from "$lib/api.js";

  let ventasSemana = [];

  onMount(async () => {
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

<h2 class="text-xl font-bold mb-4 text-blue-800">Ventas por día</h2>
<div class="space-y-2">
  {#each ventasSemana as venta}
    <div>
      <div class="flex justify-between text-sm mb-1">
        <span class="font-semibold">{venta.dia}</span>
        <span>{venta.cantidad} productos / ${venta.total.toLocaleString()}</span
        >
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
