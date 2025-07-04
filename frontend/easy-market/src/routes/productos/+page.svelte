<script>
  import AgregarProducto from "$lib/AgregarProducto.svelte";
  import EditarProducto from "$lib/EditarProducto.svelte";
  import Confirmar from "$lib/Confirmar.svelte";
  import Alerta from "$lib/Alerta.svelte";
  import { api } from "$lib/api.js";
  import { onMount } from "svelte";
  import Icon from "@iconify/svelte";
  import { requireAuth } from "$lib/guard.js";

  let productos = [];
  let categorias = [];
  let mostrarModal = false;
  let mostrarAlerta = false;
  let mensajeAlerta = "";
  let productoEditar = null;
  let mostrarConfirmar = false;
  let productoAEliminar = null;
  let filtro = "";
  let autorizado = false;
  let rol = null;

  $: productosFiltrados = productos.filter((p) =>
    p.nombre.toLowerCase().includes(filtro.toLowerCase())
  );

  function pedirConfirmacionEliminar(producto) {
    productoAEliminar = producto;
    mostrarConfirmar = true;
  }

  async function confirmarEliminar() {
    await eliminarProducto(productoAEliminar);
    mostrarConfirmar = false;
    productoAEliminar = null;
  }

  function cancelarEliminar() {
    mostrarConfirmar = false;
    productoAEliminar = null;
  }

  onMount(async () => {
    requireAuth();
    rol = localStorage.getItem("rol");
    if (!rol) {
      return;
    }
    autorizado = true;
    const Quagga = (await import("quagga")).default;
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

  async function guardarEdicionProducto(actualizado) {
    await api.updateProducto(actualizado.id, actualizado);
    productos = productos.map((p) =>
      p.id === actualizado.id ? actualizado : p
    );
    mensajeAlerta = "¡Producto actualizado!";
    mostrarAlerta = true;
    setTimeout(() => (mostrarAlerta = false), 2000);
  }
  function abrirEditarProducto(producto) {
    productoEditar = producto;
  }

  async function eliminarProducto(producto) {
    try {
      await api.request(`/productos?id=${producto.id}`, { method: "DELETE" });
      productos = productos.filter((p) => p.id !== producto.id);
      mensajeAlerta = "Producto eliminado";
      mostrarAlerta = true;
      setTimeout(() => (mostrarAlerta = false), 2000);
    } catch (e) {
      mensajeAlerta = "Error al eliminar producto";
      mostrarAlerta = true;
      setTimeout(() => (mostrarAlerta = false), 2000);
    }
  }

  function obtenerNombreCategoria(id) {
    const cat = categorias.find((c) => c.id === id);
    return cat ? cat.nombre : "-";
  }
</script>

{#if autorizado}
  {#if mostrarConfirmar}
    <Confirmar
      mensaje={`¿Seguro que deseas eliminar "${productoAEliminar?.nombre}"?`}
      onConfirm={confirmarEliminar}
      onCancel={cancelarEliminar}
    />
  {/if}

  {#if mostrarAlerta}
    <Alerta mensaje={mensajeAlerta} tipo="success" />
  {/if}

  {#if mostrarModal}
    <AgregarProducto
      onAgregar={agregarProducto}
      onClose={() => (mostrarModal = false)}
    />
  {/if}

  {#if productoEditar}
    <EditarProducto
      producto={productoEditar}
      {categorias}
      onGuardar={guardarEdicionProducto}
      onClose={() => (productoEditar = null)}
    />
  {/if}

  <div class="p-8 max-w-6xl mx-auto">
    <h1 class="text-3xl font-bold mb-6 text-gray-700">Lista de Productos</h1>
    <div
      class="mb-4 flex flex-col md:flex-row md:justify-between md:items-center gap-2"
    >
      <input
        type="text"
        placeholder="Buscar producto..."
        class="w-full md:w-1/3 px-3 py-2 border border-gray-400 rounded-md focus:outline-none focus:ring-2 focus:ring-green-200"
        bind:value={filtro}
      />
      {#if rol === "admin"}
        <button
          class="bg-green-300 hover:bg-green-500 text-gray-700 px-6 py-2 rounded-md text-sm transition"
          on:click={() => (mostrarModal = true)}
        >
          Agregar Producto
        </button>
      {/if}
    </div>
    <div class="overflow-x-auto">
      <table class="w-full bg-white rounded-2xl shadow-lg overflow-hidden">
        <thead class="bg-gray-100 text-gray-800 text-sm uppercase">
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
          {#each productosFiltrados as producto, i}
            <tr
              class={`border-t border-gray-100 transition ${i % 2 === 0 ? "bg-white" : "bg-gray-100"} hover:bg-green-50`}
            >
              <td class="py-4 px-6">{producto.id}</td>
              <td class="py-4 px-6">{producto.nombre}</td>
              <td class="py-4 px-6"
                >{obtenerNombreCategoria(producto.categoria_id)}</td
              >
              <td class="py-4 px-6">${producto.precio.toLocaleString()}</td>
              <td class="py-4 px-6">{producto.stock}</td>
              <td class="py-4 px-6">{producto.codigo_barras}</td>
              <td class="py-4 px-6 flex gap-2 items-center">
                <button
                  class="flex items-center gap-1 bg-green-100 hover:bg-green-200 text-green-700 px-3 py-1 rounded transition"
                  title="Editar producto"
                  on:click={() => abrirEditarProducto(producto)}
                >
                  <Icon icon="mdi:pencil" class="h-4 w-4" />
                  Editar
                </button>
                {#if rol === "admin"}
                  <button
                    class="flex items-center gap-1 bg-red-100 hover:bg-red-200 text-red-700 px-3 py-1 rounded transition"
                    title="Eliminar producto"
                    on:click={() => pedirConfirmacionEliminar(producto)}
                  >
                    <Icon icon="mdi:delete" class="h-4 w-4" />
                    Eliminar
                  </button>
                {/if}
              </td>
            </tr>
          {/each}
          {#if productosFiltrados.length === 0}
            <tr>
              <td colspan="7" class="text-gray-400 text-center py-6"
                >No hay productos.</td
              >
            </tr>
          {/if}
        </tbody>
      </table>
    </div>
  </div>
{/if}
