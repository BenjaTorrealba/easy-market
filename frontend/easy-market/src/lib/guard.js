export function requireAuth(redirect = "/login") {
  if (typeof window !== "undefined") {
    const rol = localStorage.getItem("rol");
    if (!rol) {
      window.location.href = redirect;
    }
  }
}
