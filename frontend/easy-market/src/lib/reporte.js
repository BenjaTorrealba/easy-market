import { api } from "./api.js";

export function getRangoSemana() {
  const hoy = new Date();
  const diaSemana = hoy.getDay();
  const diffLunes = diaSemana === 0 ? 6 : diaSemana - 1;
  const lunes = new Date(hoy);
  lunes.setDate(hoy.getDate() - diffLunes);
  const domingo = new Date(lunes);
  domingo.setDate(lunes.getDate() + 6);

  const fecha_inicio = lunes.toISOString().slice(0, 10);
  const fecha_fin = new Date(domingo.getTime() + 24 * 60 * 60 * 1000)
    .toISOString()
    .slice(0, 10); // exclusivo

  return { fecha_inicio, fecha_fin };
}

/**
 * Devuelve un arreglo con los días de la semana y ventas agrupadas por día.
 * @param {Array} historialVentas - Array de ventas con campo fecha y total.
 * @param {String} fecha_inicio - Fecha de inicio (YYYY-MM-DD)
 */

export function agruparVentasPorDia(historialVentas, fecha_inicio) {
  const dias = [
    "Lunes",
    "Martes",
    "Miércoles",
    "Jueves",
    "Viernes",
    "Sábado",
    "Domingo",
  ];
  let ventasPorDia = {};
  const inicio = new Date(fecha_inicio);
  for (let i = 0; i < 7; i++) {
    const d = new Date(inicio);
    d.setDate(inicio.getDate() + i);
    const key = d.toISOString().slice(0, 10);
    ventasPorDia[key] = { dia: dias[i], cantidad: 0, total: 0 };
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

export async function getVentasSemana() {
  try {
    const { fecha_inicio, fecha_fin } = getRangoSemana();
    const historialVentas = await api.getHistorialVentas(
      fecha_inicio,
      fecha_fin
    );
    return agruparVentasPorDia(historialVentas, fecha_inicio);
  } catch (e) {
    console.error("Error en getVentasSemana:", e);
    return [];
  }
}
