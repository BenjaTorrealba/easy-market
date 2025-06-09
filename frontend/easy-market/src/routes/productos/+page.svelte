<script>
  import AgregarProducto from '$lib/AgregarProducto.svelte';
  import Alerta from '$lib/Alerta.svelte';

  let productos = [
    { id: 1, nombre: "Laptop", precio: 750000, stock: 5, codigoBarras: "1234567890" },
    { id: 2, nombre: "Mouse", precio: 12000, stock: 50, codigoBarras: "2345678901" },
    { id: 3, nombre: "Teclado Mecánico", precio: 30000, stock: 20, codigoBarras: "3456789012" },
  ];

  let mostrarModal = false;
  let mostrarAlerta = false;
  let mensajeAlerta = "";

  function agregarProducto(producto) {
    productos = [producto, ...productos];
    mensajeAlerta = "¡Producto agregado exitosamente!";
    mostrarAlerta = true;
    setTimeout(() => mostrarAlerta = false, 2000);
  }
</script>

{#if mostrarAlerta}
  <Alerta mensaje={mensajeAlerta} tipo="success" />
{/if}

{#if mostrarModal}
  <AgregarProducto
    onAgregar={agregarProducto}
    onClose={() => mostrarModal = false}
  />
{/if}

<div class="p-8 max-w-6xl mx-auto">
  <h1 class="text-3xl font-bold mb-6 text-gray-800">
    Lista de Productos
  </h1>
  <div class="mb-4 flex justify-end">
    <button
      class="bg-indigo-600 hover:bg-indigo-800 text-white px-6 py-2 rounded-md text-sm transition"
      on:click={() => mostrarModal = true}
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
          <th class="py-4 px-6 text-left">Precio</th>
          <th class="py-4 px-6 text-left">Stock</th>
          <th class="py-4 px-6 text-left">Código de barras</th>
        </tr>
      </thead>
      <tbody class="text-gray-800 text-base">
        {#each productos as producto}
          <tr class="border-t border-gray-200 hover:bg-gray-50 transition">
            <td class="py-4 px-6">{producto.id}</td>
            <td class="py-4 px-6">{producto.nombre}</td>
            <td class="py-4 px-6">${producto.precio.toLocaleString()}</td>
            <td class="py-4 px-6">{producto.stock}</td>
            <td class="py-4 px-6">{producto.codigoBarras}</td>
          </tr>
        {/each}
      </tbody>
    </table>
  </div>
</div>