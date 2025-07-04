<script>
  export let mensaje = "";
  export let tipo = "success";
</script>

<div class={`fixed left-1/2 top-1/3 transform -translate-x-1/2 -translate-y-1/2 z-50 px-8 py-4 rounded-xl shadow-2xl text-xl font-bold max-w-lg w-full text-center
  ${tipo === "success" ? "bg-green-100 text-green-700 border-2 border-green-400" : ""}
  ${tipo === "error" ? "bg-red-100 text-red-700 border-2 border-red-400" : ""}
  ${tipo === "info" ? "bg-blue-100 text-blue-700 border-2 border-blue-400" : ""}
`}>
  {mensaje}
</div>