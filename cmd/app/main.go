package main

import (
	"log"
	"net/http"
	"path/filepath"
	"pec2/internal/db"
	"pec2/internal/handlers"
	"strings"
)

func main() {
	// Initialize database
	db.InitDB()

	// Intentar encontrar la carpeta web
	staticDir := "web/static"
	if _, err := http.Dir(staticDir).Open("."); err != nil {
		staticDir = "../../web/static"
	}

	fs := http.FileServer(http.Dir(staticDir))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir(filepath.Join(staticDir, "css")))))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir(filepath.Join(staticDir, "img")))))
	http.Handle("/scss/", http.StripPrefix("/scss/", http.FileServer(http.Dir(filepath.Join(staticDir, "scss")))))

	http.HandleFunc("/registro", handlers.RegistroHandler)
	http.HandleFunc("/resenas", handlers.GuardarResenaHandler)
	http.HandleFunc("/buscar", handlers.BuscadorHandler)
	http.HandleFunc("/calculadora", handlers.CalculadoraHandler)
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/logout", handlers.LogoutHandler)
	http.HandleFunc("/reservas", handlers.ReservasHandler)
	http.HandleFunc("/reservas/reservar", handlers.ProcesarReservaHandler)
	http.HandleFunc("/reservas/cancelar", handlers.ProcesarCancelacionHandler)
	http.HandleFunc("/tienda/tramitar", handlers.TramitarPedidoHandler)

	// API REST para Artículos (CRUD)
	http.HandleFunc("/api/articulos", handlers.APIArticulosHandler)
	http.HandleFunc("/api/articulos/", handlers.APIArticulosHandler)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/css/") || strings.HasPrefix(r.URL.Path, "/img/") || strings.HasPrefix(r.URL.Path, "/scss/") {
			fs.ServeHTTP(w, r)
			return
		}
		handlers.PageHandler(w, r)
	})

	log.Println("Servidor iniciado en http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error iniciando servidor: ", err)
	}
}
