<script>
  import AgregarProducto from "$lib/AgregarProducto.svelte";
  import Alerta from "$lib/Alerta.svelte";
  import { api } from "$lib/api.js";
  import { onMount } from "svelte";

  let productos = [];
  let categorias = [];
  let mostrarModal = false;
  let mostrarAlerta = false;
  let mensajeAlerta = "";

  onMount(async () => {
    categorias = await api.getCategorias();
    productos = await api.getProductos();
  });

  async function agregarProducto(producto) {
    try {
      const nuevo = await api.createProducto(producto);
      productos = [nuevo, ...productos];
      mensajeAlerta = "¡Producto agregado exitosamente!";
      mostrarAlerta = true;
      setTimeout(() => (mostrarAlerta = false), 2000);
    } catch (e) {
      mensajeAlerta = "Error al agregar producto";
      mostrarAlerta = true;
      setTimeout(() => (mostrarAlerta = false), 2000);
    }
  }

  // Editar precio
  async function editarPrecio(producto) {
    const nuevoPrecio = prompt("Nuevo precio:", producto.precio);
    if (nuevoPrecio && !isNaN(nuevoPrecio)) {
      const actualizado = { ...producto, precio: Number(nuevoPrecio) };
      await api.updateProducto(producto.id, actualizado);
      productos = productos.map(p => p.id === producto.id ? actualizado : p);
    }
  }

  // Editar stock
  async function editarStock(producto) {
    const nuevoStock = prompt("Nuevo stock:", producto.stock);
    if (nuevoStock && !isNaN(nuevoStock)) {
      const actualizado = { ...producto, stock: Number(nuevoStock) };
      await api.updateProducto(producto.id, actualizado);
      productos = productos.map(p => p.id === producto.id ? actualizado : p);
    }
  }

  // Marcar poco stock
  function marcarPocoStock(producto) {
    mensajeAlerta = `¡Quedan pocos de "${producto.nombre}"!`;
    mostrarAlerta = true;
    setTimeout(() => (mostrarAlerta = false), 2000);
  }

function obtenerNombreCategoria(id) {
  const cat = categorias.find(c => c.id === id);
  return cat ? cat.nombre : "-";
}
</script>

{#if mostrarAlerta}
  <Alerta mensaje={mensajeAlerta} tipo="success" />
{/if}

{#if mostrarModal}
  <AgregarProducto
    onAgregar={agregarProducto}
    onClose={() => (mostrarModal = false)}
  />
{/if}

<div class="p-8 max-w-6xl mx-auto">
  <h1 class="text-3xl font-bold mb-6 text-gray-800">Lista de Productos</h1>
  <div class="mb-4 flex justify-end">
    <button
      class="bg-indigo-600 hover:bg-indigo-800 text-white px-6 py-2 rounded-md text-sm transition"
      on:click={() => (mostrarModal = true)}
    >
      Agregar Producto
    </button>
  </div>
  <div class="overflow-x-auto">
    <table class="w-full bg-white rounded-2xl shadow-lg overflow-hidden">
      <thead class="bg-gray-100 text-gray-700 text-sm uppercase">
        <tr>
          <th class="py-4 px-6 text-left">ID</th>
          <th class="py-4 px-6 text-left">Nombre</th>
          <th class="py-4 px-6 text-left">Categoría</th>
          <th class="py-4 px-6 text-left">Precio</th>
          <th class="py-4 px-6 text-left">Stock</th>
          <th class="py-4 px-6 text-left">Código de barras</th>
          <th class="py-4 px-6 text-left">Acciones</th>
        </tr>
      </thead>
<tbody class="text-gray-800 text-base">
  {#each productos as producto}
    <tr class="border-t border-gray-200 hover:bg-gray-50 transition">
      <td class="py-4 px-6">{producto.id}</td>
      <td class="py-4 px-6">{producto.nombre}</td>
      <td class="py-4 px-6">{obtenerNombreCategoria(producto.categoria_id)}</td>
      <td class="py-4 px-6">${producto.precio.toLocaleString()}</td>
      <td class="py-4 px-6">{producto.stock}</td>
      <td class="py-4 px-6">{producto.codigo_barras}</td>
      <td class="py-4 px-6 flex gap-2 items-center">
  <button
    class="flex items-center gap-1 bg-blue-100 hover:bg-blue-200 text-blue-700 px-3 py-1 rounded transition"
    title="Editar precio"
    on:click={() => editarPrecio(producto)}
  >
    <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15.232 5.232l3.536 3.536M9 13l6-6m2 2l-6 6m-2 2h6a2 2 0 002-2v-6a2 2 0 00-2-2h-6a2 2 0 00-2 2v6a2 2 0 002 2z" /></svg>
    Precio
  </button>
  <button
    class="flex items-center gap-1 bg-yellow-100 hover:bg-yellow-200 text-yellow-700 px-3 py-1 rounded transition"
    title="Editar stock"
    on:click={() => editarStock(producto)}
  >
    <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6 0a9 9 0 11-18 0 9 9 0 0118 0z" /></svg>
    Stock
  </button>
  <button
    class="flex items-center gap-1 bg-red-100 hover:bg-red-200 text-red-700 px-3 py-1 rounded transition"
    title="Marcar poco stock"
    on:click={() => marcarPocoStock(producto)}
  >
    <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" /></svg>
    ¡Poco!
  </button>
</td>
    </tr>
  {/each}
</tbody>
    </table>
  </div>
</div>