<script>
  import { onMount } from "svelte";
  let rol = null;

  onMount(() => {
    rol = localStorage.getItem("rol");
  });

  function logout() {
    localStorage.removeItem("rol");
    rol = null;
    window.location.href = "/login";
  }
</script>

<nav
  class="bg-white dark:bg-gray-800 shadow-md h-16 px-6 flex items-center justify-between"
>
  <div class="text-2xl font-bold text-emerald-600 dark:text-white">
    <a href="/">Easy Market</a>
  </div>

  <div class="flex items-center space-x-4">
    {#if rol}
      <span class="text-gray-700 dark:text-gray-200 text-sm">Rol: {rol}</span>
      <button
        on:click={logout}
        class="bg-red-400 hover:bg-red-600 text-white px-4 py-2 rounded-md text-sm transition-shadow shadow-sm hover:shadow-md"
      >
        Cerrar sesión
      </button>
    {:else}
      <a
        href="/login"
        class="bg-emerald-500 hover:bg-emerald-600 text-white px-4 py-2 rounded-md text-sm transition-shadow shadow-sm hover:shadow-md"
      >
        Iniciar sesión
      </a>
    {/if}
  </div>
</nav>
