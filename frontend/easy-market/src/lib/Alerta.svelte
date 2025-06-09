<script>
  export let mensaje = "";
  export let tipo = "success"; // success, error, info
</script>

<div class={`fixed top-6 left-1/2 transform -translate-x-1/2 z-50 px-6 py-3 rounded shadow-lg
  ${tipo === "success" ? "bg-green-100 text-green-800 border border-green-300" : ""}
  ${tipo === "error" ? "bg-red-100 text-red-800 border border-red-300" : ""}
  ${tipo === "info" ? "bg-blue-100 text-blue-800 border border-blue-300" : ""}
`}>
  {mensaje}
</div>