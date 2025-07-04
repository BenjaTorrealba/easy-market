<script>
  import Icon from "@iconify/svelte";
  import { onMount } from "svelte";

  let autorizado = false;
  let rol = null;

  const sections = [
    {
      title: "Ver Productos",
      description: "Explora nuestro catálogo de productos disponibles.",
      link: "/productos",
      icon: "mdi:cart-outline",
    },
    {
      title: "Realizar Venta",
      description: "Registra una nueva venta de manera rápida y sencilla.",
      link: "/ventas",
      icon: "mdi:cash-register",
    },
    {
      title: "Ver Pedidos",
      description: "Revisa y gestiona los pedidos realizados.",
      link: "/pedidos",
      icon: "mdi:package-variant",
    },
    {
      title: "Ver Reportes",
      description: "Consulta reportes detallados de ventas y productos.",
      link: "/reportes",
      icon: "mdi:chart-bar",
    },
  ];

  let visibleSections = sections;

  onMount(() => {
    rol = localStorage.getItem("rol");
    if (!rol) {
      window.location.href = "/login";
    } else {
      autorizado = true;
      if (rol === "vendedor") {
        visibleSections = sections.filter((s) => s.title !== "Ver Reportes");
      }
    }
  });
</script>

{#if autorizado}
  <div class="min-h-screen bg-white flex flex-col items-center pt-12 px-4">
    <h1 class="text-5xl font-extrabold mb-2 text-gray-700 drop-shadow">
      Easy Market
    </h1>
    <p class="mb-8 text-lg text-gray-600">Bienvenido, ¿qué deseas hacer hoy?</p>
    <div
      class={`grid gap-8 w-full max-w-3xl ${visibleSections.length === 3 ? "grid-cols-1 sm:grid-cols-3" : "grid-cols-1 sm:grid-cols-2"}`}
    >
      {#each visibleSections as section}
        <a
          href={section.link}
          class="bg-gradient-to-t from-green-50 to-white rounded-2xl shadow-lg p-8 flex flex-col items-center border border-gray-200 group transition-all hover:shadow-xl hover:border-green-400 hover:scale-105"
        >
          <Icon
            icon={section.icon}
            class="text-6xl mb-4 text-green-700 group-hover:scale-125 transition-transform"
          />
          <h2 class="text-2xl font-bold mb-1 text-green-700">
            {section.title}
          </h2>
          <p class="text-gray-500 text-center">{section.description}</p>
        </a>
      {/each}
    </div>
  </div>
{/if}
