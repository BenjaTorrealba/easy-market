<script>
  import Alerta from '$lib/Alerta.svelte';
  import AgregarPedido from '$lib/AgregarPedido.svelte';
  import { api } from '$lib/api.js';
  import { onMount } from 'svelte';

  let productosBajoStock = [];

  let pedidos = [];

  let mostrarAlerta = false;
  let mensajeAlerta = "";
  let mostrarModal = false; // 
  onMount(async () => {
    try {
      productosBajoStock = await api.getProductosBajoStock();
      pedidos = await api.getPedidos();      
    } catch (e) {
      productosBajoStock = [];
      pedidos = [];
    }
  });



async function agregarPedido(pedido) {
  try {
    const nuevoPedido = await api.createPedido(pedido);
    pedidos = [nuevoPedido, ...pedidos]; // <-- Esto agrega el nuevo pedido arriba
    mensajeAlerta = "¡Pedido agregado exitosamente!";
    mostrarAlerta = true;
    setTimeout(() => mostrarAlerta = false, 2000);
  } catch (e) {
    mensajeAlerta = "Error al agregar pedido";
    mostrarAlerta = true;
    setTimeout(() => mostrarAlerta = false, 2000);
  }
}

  async function confirmarEntrega(id) {
    try {
      await api.updateEstadoPedido(id, "Entregado");
      pedidos = pedidos.map(p =>
        p.id === id ? { ...p, estado: "Entregado" } : p
      );
      
      mensajeAlerta = "¡Entrega confirmada exitosamente!";
      mostrarAlerta = true;
      setTimeout(() => mostrarAlerta = false, 2000);
    } catch (e) {
      mensajeAlerta = "Error al confirmar entrega";
      mostrarAlerta = true;
      setTimeout(() => mostrarAlerta = false, 2000);
    }
  }
</script>

{#if mostrarAlerta}
  <Alerta mensaje={mensajeAlerta} tipo="success" />
{/if}

{#if mostrarModal}
  <AgregarPedido
    onAgregar={agregarPedido}
    onClose={() => mostrarModal = false}
  />
{/if}

<div class="p-6 space-y-10">
  <!-- Productos con bajo stock (tabla) -->
  <div>
    <h2 class="text-2xl font-bold mb-4 text-yellow-700 flex items-center gap-2">
      <svg class="w-6 h-6 text-yellow-500" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M12 9v2m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" /></svg>
      Productos con poco stock
    </h2>
    <div class="overflow-x-auto">
      <table class="min-w-full rounded-lg shadow border border-yellow-200 bg-yellow-50">
        <thead>
          <tr>
            <th class="px-6 py-3 text-left bg-yellow-100 text-yellow-800 font-semibold">Producto</th>
            <th class="px-6 py-3 text-left bg-yellow-100 text-yellow-800 font-semibold">Stock</th>
          </tr>
        </thead>
        <tbody>
          {#each productosBajoStock as producto}
            <tr class="hover:bg-yellow-200/40 transition">
              <td class="border-t px-6 py-3 font-semibold">{producto.nombre}</td>
              <td class="border-t px-6 py-3">{producto.stock}</td>
            </tr>
          {/each}
          {#if productosBajoStock.length === 0}
            <tr>
              <td colspan="2" class="text-gray-400 px-6 py-3">No hay productos con bajo stock.</td>
            </tr>
          {/if}
        </tbody>
      </table>
    </div>
  </div>

  <!-- Botón para abrir el modal de agregar pedido -->
  <div class="flex justify-end">
    <button
      class="bg-blue-600 hover:bg-blue-700 text-white px-6 py-2 rounded shadow font-semibold transition"
      on:click={() => mostrarModal = true}>
      Agregar pedido
    </button>
  </div>

  <!-- Pedidos realizados (tabla) -->
  <div>
    <h2 class="text-2xl font-bold text-blue-800 flex items-center gap-2 mb-4">
      <svg class="w-6 h-6 text-blue-500" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M3 7v4a1 1 0 001 1h3m10-5h3a1 1 0 011 1v4a1 1 0 01-1 1h-3m-10 4h10m-10 4h10" /></svg>
      Pedidos realizados
    </h2>
    <div class="overflow-x-auto">
      <table class="min-w-full rounded-lg shadow border border-blue-200 bg-white">
        <thead>
          <tr>
            <th class="px-6 py-3 text-left bg-blue-100 text-blue-800 font-semibold">Fecha pedido</th>
            <th class="px-6 py-3 text-left bg-blue-100 text-blue-800 font-semibold">Distribuidor</th>
            <th class="px-6 py-3 text-left bg-blue-100 text-blue-800 font-semibold">Productos</th>
            <th class="px-6 py-3 text-left bg-blue-100 text-blue-800 font-semibold">Fecha estimada</th>
            <th class="px-6 py-3 text-left bg-blue-100 text-blue-800 font-semibold">Estado</th>
            <th class="px-6 py-3 text-left bg-blue-100 text-blue-800 font-semibold"></th>
          </tr>
        </thead>
        <tbody>
          {#each pedidos as pedido}
            <tr class="hover:bg-blue-100/40 transition">
              <td class="border-t px-6 py-3">{pedido.fecha}</td>
              <td class="border-t px-6 py-3">{pedido.distribuidor}</td>
              <td class="border-t px-6 py-3">
  {#if Array.isArray(pedido.productos) && pedido.productos.length > 0}
    {pedido.productos.map(p => `${p.nombre} (x${p.cantidad})`).join(", ")}
  {:else}
    <span class="text-gray-400">Sin productos</span>
  {/if}
</td>
              <td class="border-t px-6 py-3">{pedido.fechaEstimada}</td>
              <td class="border-t px-6 py-3">
                {#if pedido.estado === "Entregado"}
                  <span class="bg-green-100 text-green-800 px-3 py-1 rounded-full text-xs font-semibold">Entregado</span>
                {:else}
                  <span class="bg-yellow-100 text-yellow-800 px-3 py-1 rounded-full text-xs font-semibold">Pendiente</span>
                {/if}
              </td>
              <td class="border-t px-6 py-3">
                {#if pedido.estado === "Pendiente"}
                  <button
                    class="bg-green-600 hover:bg-green-700 text-white px-3 py-1 rounded text-xs"
                    on:click={() => confirmarEntrega(pedido.id)}>
                    Confirmar entrega
                  </button>
                {/if}
              </td>
            </tr>
          {/each}
          {#if pedidos.length === 0}
            <tr>
              <td colspan="6" class="text-gray-400 px-6 py-3">No hay pedidos registrados.</td>
            </tr>
          {/if}
        </tbody>
      </table>
    </div>
  </div>
</div>