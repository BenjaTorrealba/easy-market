<script>
    const sections = [
        {
            title: "Ver Productos",
            description: "Explora nuestro catálogo de productos disponibles.",
            link: "/productos",
            icon: "🛒"
        },
        {
            title: "Realizar Venta",
            description: "Registra una nueva venta de manera rápida y sencilla.",
            link: "/ventas",
            icon: "💵"
        },
        {
            title: "Ver Reportes",
            description: "Consulta reportes detallados de ventas y productos.",
            link: "/reportes",
            icon: "📊"
        },
        {
            title: "Ver Pedidos",
            description: "Revisa y gestiona los pedidos realizados.",
            link: "/pedidos",
            icon: "📦"
        }
    ];
</script>

<div class="min-h-screen bg-gray-50 flex flex-col items-center justify-center py-12 px-4">
    <h1 class="text-4xl font-bold mb-8 text-gray-800">Easy Market</h1>
    <div class="grid grid-cols-1 sm:grid-cols-2 gap-8 w-full max-w-3xl">
        {#each sections as section}
            <a href={section.link} class="bg-white rounded-xl shadow-md p-8 flex flex-col items-center hover:shadow-xl transition-shadow group">
                <div class="text-5xl mb-4 group-hover:scale-110 transition-transform">{section.icon}</div>
                <h2 class="text-2xl font-semibold mb-2 text-gray-700">{section.title}</h2>
                <p class="text-gray-500 text-center">{section.description}</p>
            </a>
        {/each}
    </div>
</div>