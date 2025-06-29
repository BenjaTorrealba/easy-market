<script>
  import Navbar from '$lib/Navbar.svelte';
  import Sidebar from '$lib/Sidebar.svelte';
  import '../app.css';

</script>

<div class="flex flex-col min-h-screen">
  <Navbar />

  <div class="flex flex-1">
    <Sidebar />
    <main class="flex-1 p-6">
      <slot />
    </main>
  </div>
</div>