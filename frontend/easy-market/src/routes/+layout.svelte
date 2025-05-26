<script>
    import { page } from "$app/stores";
</script>

<div class="layout">
    <aside class="sidebar">
        <div class="logo">🛒 Easy-Market</div>
        <nav>
            <a
                href="/productos"
                class={$page.url.pathname === "/productos" ? "active" : ""}
                >Productos</a
            >
            <a
                href="/ventas"
                class={$page.url.pathname === "/ventas" ? "active" : ""}
                >Realizar Venta</a
            >
            <a
                href="/informe"
                class={$page.url.pathname === "/informe" ? "active" : ""}
                >Informe Ventas</a
            >
            <a
                href="/pedidos"
                class={$page.url.pathname === "/pedidos" ? "active" : ""}
                >Pedidos</a
            >
        </nav>
    </aside>

    <main class="content">
        <header>
            <h1>
                {$page.url.pathname === "/"
                    ? "Dashboard"
                    : $page.url.pathname.slice(1).replace("/", "")}
            </h1>
        </header>
        <section class="page">
            <slot />
        </section>
    </main>
</div>

<style>
    .layout {
        display: flex;
        min-height: 100vh;
        font-family: "Segoe UI", Tahoma, Geneva, Verdana, sans-serif;
        background-color: #f8f9fa;
    }
    .sidebar {
        width: 220px;
        background-color: #1e293b;
        color: #fff;
        padding: 2rem 1.5rem;
        display: flex;
        flex-direction: column;
        box-shadow: 2px 0 5px rgba(0, 0, 0, 0.1);
    }

    .logo {
        font-size: 1.5rem;
        font-weight: bold;
        margin-bottom: 2rem;
        color: #38bdf8; 
    }

    nav {
        display: flex;
        flex-direction: column;
        gap: 0.5rem;
    }

    nav a {
        color: #cbd5e1;
        text-decoration: none;
        padding: 0.5rem 1rem;
        border-radius: 8px;
        transition:
            background 0.2s ease,
            color 0.2s ease;
    }

    nav a:hover {
        background-color: #334155;
        color: #fff;
    }

    nav a.active {
        background-color: #38bdf8;
        color: #1e293b;
        font-weight: bold;
    }

    .content {
        flex: 1;
        padding: 2rem;
    }

    header h1 {
        font-size: 2rem;
        margin-bottom: 1.5rem;
        color: #1e293b;
    }

    .page {
        background: white;
        border-radius: 12px;
        padding: 2rem;
        box-shadow: 0 5px 20px rgba(0, 0, 0, 0.05);
    }
</style>
