<script>
  import Icon from "@iconify/svelte";
  import { onMount } from "svelte";
  let rol = null;
  let autorizado = false;

  onMount(() => {
    rol = localStorage.getItem("rol");
    if (rol) {
      autorizado = true;
    }
  });
</script>

{#if autorizado}
  <div class="flex min-h-screen">
    <aside
      class="w-64 bg-white border-r border-gray-100 p-4 shadow-sm flex flex-col"
    >
      <nav class="flex flex-col space-y-2">
        <a
          href="/"
          class="flex items-center gap-3 text-gray-700 hover:text-emerald-600 text-lg px-3 py-2 rounded transition-colors duration-200 hover:bg-emerald-50 font-medium"
        >
          <Icon icon="mdi:home-outline" class="text-2xl" />
          Inicio
        </a>
        <a
          href="/productos"
          class="flex items-center gap-3 text-gray-700 hover:text-emerald-600 text-lg px-3 py-2 rounded transition-colors duration-200 hover:bg-emerald-50 font-medium"
        >
          <Icon icon="mdi:cart-outline" class="text-2xl" />
          Productos
        </a>
        <a
          href="/ventas"
          class="flex items-center gap-3 text-gray-700 hover:text-emerald-600 text-lg px-3 py-2 rounded transition-colors duration-200 hover:bg-emerald-50 font-medium"
        >
          <Icon icon="mdi:cash-register" class="text-2xl" />
          Realizar Venta
        </a>
        <a
          href="/pedidos"
          class="flex items-center gap-3 text-gray-700 hover:text-emerald-600 text-lg px-3 py-2 rounded transition-colors duration-200 hover:bg-emerald-50 font-medium"
        >
          <Icon icon="mdi:package-variant" class="text-2xl" />
          Pedidos
        </a>
        {#if rol === "admin"}
          <a
            href="/reportes"
            class="flex items-center gap-3 text-gray-700 hover:text-emerald-600 text-lg px-3 py-2 rounded transition-colors duration-200 hover:bg-emerald-50 font-medium"
          >
            <Icon icon="mdi:chart-bar" class="text-2xl" />
            Reportes
          </a>
        {/if}
      </nav>
    </aside>
  </div>
{/if}
