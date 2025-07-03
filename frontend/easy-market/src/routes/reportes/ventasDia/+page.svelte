<script>
  import { onMount } from 'svelte';
  import Chart from 'chart.js/auto';
  import ChartDataLabels from 'chartjs-plugin-datalabels';
  import { getVentasSemana } from '$lib/reporte.js';

  let ventasSemana = [];
  let chart;
  let canvasEl;

  onMount(async () => {
    ventasSemana = await getVentasSemana();

    // Solo existe window en el cliente
    window.ChartDataLabels = ChartDataLabels;

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
      plugins: [ChartDataLabels]
    });
  });
</script>
<div class="relative" style="height: 400px;">
  <canvas bind:this={canvasEl}></canvas>
</div>