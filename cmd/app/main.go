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
	http.HandleFunc("/resenas", handlers.GuardarResenaHandler)
	http.HandleFunc("/buscar", handlers.BuscadorHandler)
	http.HandleFunc("/calculadora", handlers.CalculadoraHandler)
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/logout", handlers.LogoutHandler)
	http.HandleFunc("/reservas", handlers.ReservasHandler)
	http.HandleFunc("/reservas/reservar", handlers.ProcesarReservaHandler)
	http.HandleFunc("/reservas/cancelar", handlers.ProcesarCancelacionHandler)
	http.HandleFunc("/tienda/tramitar", handlers.TramitarPedidoHandler)
	http.HandleFunc("/tienda/confirmar", handlers.ConfirmarPedidoHandler)

	log.Println("Servidor iniciado en http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Error iniciando servidor: ", err)
	}
}
