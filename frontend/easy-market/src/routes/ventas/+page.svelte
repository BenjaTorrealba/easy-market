<script>
  let productosDisponibles = [
    { id: 1, nombre: "Mouse", precio: 15000 },
    { id: 2, nombre: "Teclado", precio: 25000 },
    { id: 3, nombre: "Joystick", precio: 30000 },
  ];

  /** @type {{ id: number, nombre: string, precio: number, cantidad: number }[]} */
  let carrito = [];
  let idInput = "";

  function agregarProductoPorId() {
    const id = parseInt(idInput);
    const producto = productosDisponibles.find(p => p.id === id);

    if (producto) {
      // Ver si ya está en el carrito
      const itemExistente = carrito.find(p => p.id === producto.id);
      if (itemExistente) {
        itemExistente.cantidad++;
      } else {
        carrito = [...carrito, { ...producto, cantidad: 1 }];
      }
    } else {
      alert("Producto no encontrado");
    }

    idInput = "";
  }
</script>

<input
  type="number"
  bind:value={idInput}
  placeholder="Ingresa ID del producto"
  on:keydown={(e) => e.key === 'Enter' && agregarProductoPorId()}
/>
<button on:click={agregarProductoPorId}>Agregar</button>

<table>
  <thead>
    <tr>
      <th>ID</th>
      <th>Nombre</th>
      <th>Precio</th>
      <th>Cantidad</th>
    </tr>
  </thead>
  <tbody>
    {#each carrito as producto}
      <tr>
        <td>{producto.id}</td>
        <td>{producto.nombre}</td>
        <td>${producto.precio}</td>
        <td>{producto.cantidad}</td>
      </tr>
    {/each}
    {#if carrito.length === 0}
      <tr>
        <td colspan="4">Carrito vacío</td>
      </tr>
    {/if}
  </tbody>
</table>
<h4>Total: ${carrito.reduce((total, p) => total + p.precio * p.cantidad, 0)}</h4>
