package handlers

import (
	"log"
	"net/http"
	"path/filepath"
	"pec2/internal/db"
	"pec2/internal/models"
	"pec2/internal/services"
	"strconv"
	"strings"
)

// GetWebPath construye la ruta absoluta a un fichero dentro del directorio web/,
// detectando automáticamente si el binario se ejecuta desde la raíz del proyecto
// o desde un subdirectorio como cmd/app/.
// Recibe segmentos de ruta relativos a web/ y devuelve la ruta completa.
func GetWebPath(relative ...string) string {
	webBase := "web"
	// Si no existe ./web, probar ../../web (ejecución desde cmd/app/)
	if _, err := http.Dir(webBase).Open("."); err != nil {
		webBase = "../../web"
	}
	parts := append([]string{webBase}, relative...)
	return filepath.Join(parts...)
}

// IndexHandler gestiona la página de inicio (/). Carga las reseñas de la base de
// datos y, si hay sesión activa, incluye los datos del usuario autenticado.
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// El mux de Go enruta a "/" cualquier ruta no registrada; filtramos las desconocidas.
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	log.Printf("Usuario accedió a la página: index")

	resenas := db.ObtenerResenas()
	var usuario *models.Socio
	cookie, err := r.Cookie("session_id")
	if err == nil {
		usuario = db.ObtenerSocioPorCorreo(cookie.Value)
	}

	RenderTemplate(w, "index", map[string]interface{}{
		"Resenas": resenas,
		"Usuario": usuario,
	})
}

// MaquinariaHandler gestiona el listado de maquinaria (/maquinaria).
func MaquinariaHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Usuario accedió a la página: maquinaria")
	RenderTemplate(w, "maquinaria", db.Maquinas)
}

// MaquinariaDetalleHandler gestiona el detalle de una máquina (/maquinaria-detalle?id=N).
// Busca la máquina por ID en la colección en memoria y la pasa a la plantilla.
func MaquinariaDetalleHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Usuario accedió a la página: maquinaria-detalle")
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)
	var data interface{}
	for _, m := range db.Maquinas {
		if m.ID == id {
			data = m
			break
		}
	}
	RenderTemplate(w, "maquinaria-detalle", data)
}

// ServiciosHandler gestiona el listado de servicios (/servicios).
func ServiciosHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Usuario accedió a la página: servicios")
	RenderTemplate(w, "servicios", db.Servicios)
}

// ServicioDetalleHandler gestiona el detalle de un servicio (/servicio-detalle?id=N).
// Busca el servicio por ID en la colección en memoria y lo pasa a la plantilla.
func ServicioDetalleHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Usuario accedió a la página: servicio-detalle")
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)
	var data interface{}
	for _, s := range db.Servicios {
		if s.ID == id {
			data = s
			break
		}
	}
	RenderTemplate(w, "servicio-detalle", data)
}

// EquipoHandler gestiona la página del equipo (/equipo).
func EquipoHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Usuario accedió a la página: equipo")
	RenderTemplate(w, "equipo", db.EquipoCompleto)
}

// ReglasHandler gestiona la página de reglas del centro (/reglas).
func ReglasHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Usuario accedió a la página: reglas")
	RenderTemplate(w, "reglas", nil)
}

// ApuntateHandler gestiona la página de registro de nuevos socios (/apuntate).
func ApuntateHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Usuario accedió a la página: apuntate")
	RenderTemplate(w, "apuntate", nil)
}

// TiendaHandler gestiona la página de la tienda (/tienda).
func TiendaHandler(w http.ResponseWriter, r *http.Request) {
	// Evitar que rutas desconocidas bajo /tienda/ lleguen aquí
	if r.URL.Path != "/tienda" && r.URL.Path != "/tienda/" {
		http.NotFound(w, r)
		return
	}
	log.Printf("Usuario accedió a la página: tienda")
	RenderTemplate(w, "tienda", map[string]interface{}{
		"Articulos": services.CatalogoTienda(),
	})
}

// PageHandler es un handler comodín que cubre páginas simples sin lógica de datos.
// Redirige a NotFound si la URL contiene una extensión de fichero desconocida.
// Se mantiene únicamente como fallback para páginas no registradas explícitamente.
func PageHandler(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Path[1:]
	if page == "" || page == "index.html" {
		page = "index"
	}
	page = strings.TrimSuffix(page, ".html")

	if strings.Contains(page, ".") {
		http.NotFound(w, r)
		return
	}

	log.Printf("Usuario accedió a la página (fallback): %s", page)
	RenderTemplate(w, page, nil)
}
