import { localStorageStore } from '$lib/localStorage';

type Venta = {
  id: number;
  cantidad: number;  
};
type ProductosNotificar = {
  id: number;
};

export const tareas = localStorageStore<Venta[]>('tareas', []);