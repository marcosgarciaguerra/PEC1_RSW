package main

import (
	"log"
	"net/http"
	"path/filepath"
	"pec2/internal/db"
	"pec2/internal/handlers"
)

func main() {
	// Inicializar la base de datos
	db.InitDB()

	// Precargar todas las plantillas una sola vez al arrancar el servidor.
	// De esta forma, RenderTemplate no accede a disco en cada petición.
	handlers.InitTemplates([]string{
		"index",
		"tienda",
		"maquinaria",
		"maquinaria-detalle",
		"servicios",
		"servicio-detalle",
		"equipo",
		"reglas",
		"apuntate",
		"login",
		"buscar",
		"reservas",
		"calculadora",
		"tramitar-pedido",
		"admin",
		"admin-clases",
		"admin-articulos",
		"admin-maquinarias",
	})

	// Servir ficheros estáticos (CSS, imágenes, SCSS).
	// El mux de Go siempre hace match con la ruta más larga registrada,
	// por lo que estas rutas tienen prioridad sobre la ruta comodín "/".
	staticDir := "web/static"
	if _, err := http.Dir(staticDir).Open("."); err != nil {
		staticDir = "../../web/static"
	}
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir(filepath.Join(staticDir, "css")))))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir(filepath.Join(staticDir, "img")))))
	http.Handle("/scss/", http.StripPrefix("/scss/", http.FileServer(http.Dir(filepath.Join(staticDir, "scss")))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir(filepath.Join(staticDir, "js")))))

	// Handlers de vistas — cada ruta tiene su propio controlador dedicado.
	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/tienda", handlers.TiendaHandler)
	http.HandleFunc("/maquinaria", handlers.MaquinariaHandler)
	http.HandleFunc("/maquinaria-detalle", handlers.MaquinariaDetalleHandler)
	http.HandleFunc("/servicios", handlers.ServiciosHandler)
	http.HandleFunc("/servicio-detalle", handlers.ServicioDetalleHandler)
	http.HandleFunc("/equipo", handlers.EquipoHandler)
	http.HandleFunc("/reglas", handlers.ReglasHandler)
	http.HandleFunc("/apuntate", handlers.ApuntateHandler)

	// Handlers de acciones / formularios
	http.HandleFunc("/registro", handlers.RegistroHandler)
	http.HandleFunc("/resenas", handlers.RequireSocioAuth(handlers.GuardarResenaHandler))
	http.HandleFunc("/buscar", handlers.BuscadorHandler)
	http.HandleFunc("/calculadora", handlers.CalculadoraHandler)
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/logout", handlers.LogoutHandler)
	http.HandleFunc("/reservas", handlers.RequireSocioAuth(handlers.ReservasHandler))
	http.HandleFunc("/reservas/reservar", handlers.RequireSocioAuth(handlers.ProcesarReservaHandler))
	http.HandleFunc("/reservas/cancelar", handlers.RequireSocioAuth(handlers.ProcesarCancelacionHandler))
	http.HandleFunc("/tienda/tramitar", handlers.RequireUsuarioAuth(handlers.TramitarPedidoHandler))
	http.HandleFunc("/tienda/confirmar", handlers.RequireUsuarioAuth(handlers.ConfirmarPedidoHandler))

	// API REST — enrutamiento Go 1.22+ (método + patrón), CORS y auth de socio
	registerAPIRoutes()

	// Admin panel
	http.HandleFunc("/admin", handlers.RequireSocioAuth(handlers.AdminHandler))
	http.HandleFunc("/admin/clases", handlers.AdminClasesHandler)
	http.HandleFunc("/admin/articulos", handlers.AdminArticulosHandler)
	http.HandleFunc("/admin/maquinarias", handlers.AdminMaquinariasHandler)

	log.Println("Servidor iniciado en http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Error iniciando servidor: ", err)
	}
}

// registerAPIRoutes registra rutas REST con verbos HTTP explícitos y preflight OPTIONS.
func registerAPIRoutes() {
	noop := func(http.ResponseWriter, *http.Request) {}

	// Artículos
	http.HandleFunc("OPTIONS /api/articulos", handlers.WithCORS(noop))
	http.HandleFunc("OPTIONS /api/articulos/{id}", handlers.WithCORS(noop))
	http.HandleFunc("GET /api/articulos", handlers.APIRoute(handlers.ListArticulos))
	http.HandleFunc("POST /api/articulos", handlers.APIRoute(handlers.CreateArticulo))
	http.HandleFunc("GET /api/articulos/{id}", handlers.APIRoute(handlers.GetArticulo))
	http.HandleFunc("PUT /api/articulos/{id}", handlers.APIRoute(handlers.UpdateArticulo))
	http.HandleFunc("DELETE /api/articulos/{id}", handlers.APIRoute(handlers.DeleteArticulo))

	// Clases
	http.HandleFunc("OPTIONS /api/clases", handlers.WithCORS(noop))
	http.HandleFunc("OPTIONS /api/clases/{id}", handlers.WithCORS(noop))
	http.HandleFunc("GET /api/clases", handlers.APIRoute(handlers.ListClases))
	http.HandleFunc("POST /api/clases", handlers.APIRoute(handlers.CreateClase))
	http.HandleFunc("GET /api/clases/{id}", handlers.APIRoute(handlers.GetClase))
	http.HandleFunc("PUT /api/clases/{id}", handlers.APIRoute(handlers.UpdateClase))
	http.HandleFunc("DELETE /api/clases/{id}", handlers.APIRoute(handlers.DeleteClase))

	// Maquinarias
	http.HandleFunc("OPTIONS /api/maquinarias", handlers.WithCORS(noop))
	http.HandleFunc("OPTIONS /api/maquinarias/{id}", handlers.WithCORS(noop))
	http.HandleFunc("GET /api/maquinarias", handlers.APIRoute(handlers.ListMaquinarias))
	http.HandleFunc("POST /api/maquinarias", handlers.APIRoute(handlers.CreateMaquinaria))
	http.HandleFunc("GET /api/maquinarias/{id}", handlers.APIRoute(handlers.GetMaquinaria))
	http.HandleFunc("PUT /api/maquinarias/{id}", handlers.APIRoute(handlers.UpdateMaquinaria))
	http.HandleFunc("DELETE /api/maquinarias/{id}", handlers.APIRoute(handlers.DeleteMaquinaria))
}
