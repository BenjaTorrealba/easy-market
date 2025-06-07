<!-- +layout.svelte -->
<script>
  import Navbar from '$lib/Navbar.svelte';
  import Sidebar from '$lib/Sidebar.svelte';
  import '../app.css';
</script>

<div class="flex flex-col min-h-screen">
  <!-- Navbar -->
  <Navbar />

  <!-- Main layout: Sidebar + Page Content -->
  <div class="flex flex-1">
    <!-- Sidebar -->
    <Sidebar />

    <!-- Page Content -->
    <main class="flex-1 p-6">
      <slot /> <!-- This is where route content gets rendered -->
    </main>
  </div>
</div>
