package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	
	// Importa tu paquete de database
	"easy-market-api/internal/database"
)

func main() {
	log.Println("=== INICIANDO EASY MARKET API ===")
	
	// Cargar variables de entorno
	log.Println("Cargando variables de entorno...")
	if err := godotenv.Load(); err != nil {
		log.Println("No se encontró archivo .env, usando variables del sistema")
	} else {
		log.Println("Archivo .env cargado correctamente")
	}

	// Obtener configuración desde variables de entorno
	dbPath := getEnv("DB_PATH", "./data/easy_market.db")
	port := getEnv("PORT", "8080")
	host := getEnv("HOST", "localhost")
	
	log.Printf("Configuración cargada - DB: %s, Host: %s, Port: %s", dbPath, host, port)

	// Inicializar la base de datos
	log.Println("Inicializando la base de datos...")
	if err := database.Initialize(dbPath); err != nil {
		log.Fatalf("Error inicializando la base de datos: %v", err)
	}
	defer database.Close()

	log.Println("Base de datos inicializada correctamente")

	// Configurar el router
	router := mux.NewRouter()
	
	// Rutas de ejemplo
	router.HandleFunc("/", homeHandler).Methods("GET")
	router.HandleFunc("/health", healthHandler).Methods("GET")
	
	// Rutas API
	api := router.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/categorias", getCategoriasHandler).Methods("GET")

	// Configurar el servidor
	addr := host + ":" + port
	log.Printf("Servidor configurado para http://%s", addr)
	log.Printf("Rutas disponibles:")
	log.Printf("  GET  /")
	log.Printf("  GET  /health")
	log.Printf("  GET  /api/v1/categorias")
	
	log.Printf("🚀 Servidor iniciando en http://%s", addr)
	log.Println("Presiona Ctrl+C para detener el servidor")

	// Manejar señales para shutdown graceful
	go func() {
		if err := http.ListenAndServe(addr, router); err != nil {
			log.Fatal("Error iniciando el servidor:", err)
		}
	}()

	// Esperar señal de interrupción
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	log.Println("Cerrando servidor...")
}

// getEnv obtiene una variable de entorno con valor por defecto
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// Handlers
func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "Easy Market API", "status": "running"}`))
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"status": "healthy"}`))
}

func getCategoriasHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"categorias": [], "message": "Endpoint de categorías - por implementar"}`))
}
