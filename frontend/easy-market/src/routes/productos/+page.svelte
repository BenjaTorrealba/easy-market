<script>
  let productos = [
    { id: 1, nombre: "Laptop", precio: 750000, stock: 5 },
    { id: 2, nombre: "Mouse", precio: 12000, stock: 50 },
    { id: 3, nombre: "Teclado Mecánico", precio: 30000, stock: 20 },
  ];
</script>

<div class="p-8 max-w-6xl mx-auto"><script>
  let productos = [
    { id: 1, nombre: "Laptop", precio: 750000, stock: 5 },
    { id: 2, nombre: "Mouse", precio: 12000, stock: 50 },
    { id: 3, nombre: "Teclado Mecánico", precio: 30000, stock: 20 },
  ];
</script>

<div class="p-8 max-w-6xl mx-auto">
  <h1 class="text-3xl font-bold mb-6 text-gray-800 dark:text-black">
    Lista de Productos
  </h1>
  <div class="mb-4">
    <button
      class="bg-indigo-600 hover:bg-indigo-800 text-white px-6 py-2 rounded-md text-sm transition"
    >
      Agregar Producto
    </button>
  </div>
  <div class="overflow-x-auto">
    <table
      class="w-full bg-white dark:bg-gray-800 rounded-2xl shadow-lg overflow-hidden"
    >
      <thead
        class="bg-gray-100 dark:bg-gray-700 text-gray-700 dark:text-gray-200 text-sm uppercase"
      >
        <tr>
          <th class="py-4 px-6 text-left">ID</th>
          <th class="py-4 px-6 text-left">Nombre</th>
          <th class="py-4 px-6 text-left">Precio</th>
          <th class="py-4 px-6 text-left">Stock</th>
        </tr>
      </thead>
      <tbody class="text-gray-800 dark:text-gray-100 text-base">
        {#each productos as producto}
          <tr
            class="border-t border-gray-200 dark:border-gray-700 hover:bg-gray-50 dark:hover:bg-gray-600 transition"
          >
            <td class="py-4 px-6">{producto.id}</td>
            <td class="py-4 px-6">{producto.nombre}</td>
            <td class="py-4 px-6"
              >${producto.precio.toLocaleString()}
              <a
                href="/login"
                class="bg-green-600 hover:bg-green-800 text-white px-4 py-2 rounded-md text-sm transition"
              >
                Editar
              </a></td>

            <td class="py-4 px-6">{producto.stock}
              <a href="/login"
                class="bg-blue-600 hover:bg-blue-800 text-white px-4 py-2 rounded-md text-sm transition">
                Notificar </a>
              </td>
          </tr>
        {/each}
      </tbody>
    </table>
  </div>
</div>

  <h1 class="text-3xl font-bold mb-6 text-gray-800 dark:text-black">
    Lista de Productos
  </h1>
  <div class="mb-4">
    <button
      class="bg-indigo-600 hover:bg-indigo-800 text-white px-6 py-2 rounded-md text-sm transition"
    >
      Agregar Producto
    </button>
  </div>
  <div class="overflow-x-auto">
    <table
      class="w-full bg-white dark:bg-gray-800 rounded-2xl shadow-lg overflow-hidden"
    >
      <thead
        class="bg-gray-100 dark:bg-gray-700 text-gray-700 dark:text-gray-200 text-sm uppercase"
      >
        <tr>
          <th class="py-4 px-6 text-left">ID</th>
          <th class="py-4 px-6 text-left">Nombre</th>
          <th class="py-4 px-6 text-left">Precio</th>
          <th class="py-4 px-6 text-left">Stock</th>
        </tr>
      </thead>
      <tbody class="text-gray-800 dark:text-gray-100 text-base">
        {#each productos as producto}
          <tr
            class="border-t border-gray-200 dark:border-gray-700 hover:bg-gray-50 dark:hover:bg-gray-600 transition"
          >
            <td class="py-4 px-6">{producto.id}</td>
            <td class="py-4 px-6">{producto.nombre}</td>
            <td class="py-4 px-6"
              >${producto.precio.toLocaleString()}
              <a
                href="/login"
                class="bg-green-600 hover:bg-green-800 text-white px-4 py-2 rounded-md text-sm transition"
              >
                Editar
              </a></td>

            <td class="py-4 px-6">{producto.stock}
              <a href="/login"
                class="bg-blue-600 hover:bg-blue-800 text-white px-4 py-2 rounded-md text-sm transition">
                Notificar </a>
              </td>
          </tr>
        {/each}
      </tbody>
    </table>
  </div>
</div>
