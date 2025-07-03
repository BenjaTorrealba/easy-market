<script>
  import { onMount } from 'svelte';
  import { api } from '$lib/api.js';
  import Chart from 'chart.js/auto';
  import ChartDataLabels from 'chartjs-plugin-datalabels';
  window.ChartDataLabels = ChartDataLabels;

  let ventasSemana = [];
  let chart;
  let canvasEl;

  onMount(async () => {
    const hoy = new Date();
    const hace7 = new Date();
    hace7.setDate(hoy.getDate() - 6);
    const fecha_inicio = hace7.toISOString().slice(0, 10);
    const fecha_fin = hoy.toISOString().slice(0, 10);

    const historialVentas = await api.getHistorialVentas(fecha_inicio, fecha_fin);

    const dias = ['Domingo', 'Lunes', 'Martes', 'Miércoles', 'Jueves', 'Viernes', 'Sábado'];
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

    // Dibuja el gráfico
    if (chart) chart.destroy();
    chart = new Chart(canvasEl, {
      type: 'bar',
      data: {
        labels: ventasSemana.map(v => v.dia),
        datasets: [{
          label: 'Total vendido ($)',
          data: ventasSemana.map(v => v.total),
          backgroundColor: 'rgba(37, 99, 235, 0.85)',
          borderColor: 'rgba(30, 64, 175, 1)',
          borderWidth: 2,
          borderRadius: 8,
          hoverBackgroundColor: 'rgba(59, 130, 246, 1)',
        }]
      },
      options: {
        responsive: true,
        maintainAspectRatio: false,
        plugins: {
          legend: { display: false },
          tooltip: {
            callbacks: {
              label: ctx => `$${ctx.parsed.y.toLocaleString()}`
            }
          },
            datalabels: {
            anchor: 'end',
            align: 'start',
            color: '#1e40af',
            font: { weight: 'bold', size: 16 },
            formatter: value => value === 0 ? '' : `$${value.toLocaleString()}`
          },
          title: {
            display: false
          }
        },
        layout: {
          padding: 24
        },
        scales: {
          x: {
            title: { display: true, text: 'Día de la semana', font: { size: 18, weight: 'bold' } },
            ticks: { font: { size: 16 } }
          },
          y: {
            title: { display: true, text: 'Total vendido ($)', font: { size: 18, weight: 'bold' } },
            beginAtZero: true,
            ticks: { font: { size: 16 } }
          }
        }
      },
      plugins: [window.ChartDataLabels]
    });
  });
</script>

<h2 class="text-2xl font-bold mb-6 text-blue-800 text-center">Ventas por día</h2>
<div class="flex justify-center">
  <div class="w-full max-w-4xl bg-white rounded shadow p-6">
    <canvas bind:this={canvasEl} style="height:400px; min-height:400px; max-height:500px;"></canvas>
  </div>
</div>