import { api } from './api.js';

/**
 * Devuelve el rango de fechas para la semana actual (últimos 7 días, incluyendo hoy).
 */
export function getRangoSemana() {
  const hoy = new Date();
  const hace7 = new Date();
  hace7.setDate(hoy.getDate() - 6);

  // Sumar un día a hoy para incluir todo el día actual
  const manana = new Date(hoy);
  manana.setDate(hoy.getDate() + 1);

  const fecha_inicio = hace7.toISOString().slice(0, 10);
  const fecha_fin = manana.toISOString().slice(0, 10); // Exclusivo

  return { fecha_inicio, fecha_fin };
}

/**
 * Devuelve un arreglo con los días de la semana y ventas agrupadas por día.
 * @param {Array} historialVentas - Array de ventas con campo fecha y total.
 * @param {String} fecha_inicio - Fecha de inicio (YYYY-MM-DD)
 */
export function agruparVentasPorDia(historialVentas, fecha_inicio) {
  const dias = [
    "Domingo",
    "Lunes",
    "Martes",
    "Miércoles",
    "Jueves",
    "Viernes",
    "Sábado",
  ];
  let ventasPorDia = {};
  const inicio = new Date(fecha_inicio);
  for (let i = 0; i < 7; i++) {
    const d = new Date(inicio);
    d.setDate(inicio.getDate() + i);
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
  return Object.values(ventasPorDia);
}

/**
 * Obtiene las ventas de la semana agrupadas por día, listo para usar en dashboard o gráficos.
 */
export async function getVentasSemana() {
  const { fecha_inicio, fecha_fin } = getRangoSemana();
  const historialVentas = await api.getHistorialVentas(fecha_inicio, fecha_fin);
  return agruparVentasPorDia(historialVentas, fecha_inicio);
}

/**
 * Otros métodos de reporte pueden ir aquí...
 */