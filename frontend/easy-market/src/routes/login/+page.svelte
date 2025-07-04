<script>
  import { api } from "$lib/api.js";
  let nombreUsuario = "";
  let password = "";
  let error = "";
  let mensaje = "";

  async function handleLogin() {
    error = "";
    mensaje = "";
    try {
      const res = await api.login(nombreUsuario, password);
      if (res.success) {
        mensaje = "¡Login exitoso!";
        localStorage.setItem("rol", res.rol); // Guarda el rol
        window.location.href = "/"; // Descomenta para redirigir
      } else {
        error = res.message || "Error en el login";
      }
    } catch (e) {
      error = "Usuario o contraseña incorrectos";
    }
  }
</script>

<div class="min-h-screen flex items-center justify-center bg-gray-50">
  <form
    on:submit|preventDefault={handleLogin}
    class="bg-white rounded-2xl shadow-lg p-8 w-full max-w-sm flex flex-col gap-6 border border-gray-100"
  >
    <h1 class="text-2xl font-bold text-gray-700 mb-2 text-center">
      Iniciar Sesión
    </h1>
    <div class="flex flex-col gap-2">
      <label class="text-gray-700 font-medium">Usuario</label>
      <input
        type="text"
        bind:value={nombreUsuario}
        required
        class="px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-green-200"
        placeholder="Ingresa tu usuario"
      />
    </div>
    <div class="flex flex-col gap-2">
      <label class="text-gray-700 font-medium">Contraseña</label>
      <input
        type="password"
        bind:value={password}
        required
        class="px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-green-200"
        placeholder="Ingresa tu contraseña"
      />
    </div>
    <button
      type="submit"
      class="bg-green-300 hover:bg-green-500 text-gray-700 font-semibold px-6 py-2 rounded-md text-base transition"
    >
      Ingresar
    </button>
    {#if error}
      <div
        class="bg-red-100 text-red-700 rounded px-3 py-2 text-center text-sm"
      >
        {error}
      </div>
    {/if}
    {#if mensaje}
      <div
        class="bg-green-100 text-green-700 rounded px-3 py-2 text-center text-sm"
      >
        {mensaje}
      </div>
    {/if}
  </form>
</div>
