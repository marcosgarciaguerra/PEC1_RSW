package handlers

import (
	"html/template"
	"log"
	"net/http"
	"sync"
)

// templateCache almacena las plantillas ya compiladas, indexadas por nombre de página.
// Se rellena una sola vez al arrancar el servidor mediante InitTemplates.
var templateCache = make(map[string]*template.Template)

// templateCacheMu protege el mapa de caché en escenarios de lectura concurrente.
var templateCacheMu sync.RWMutex

// pageTitles define el título legible que se muestra en el <title> de cada página.
// Si una página no está en el mapa, se usa su nombre de ruta como fallback.
var pageTitles = map[string]string{
	"index":            "Inicio",
	"tienda":           "Tienda",
	"maquinaria":       "Maquinaria",
	"maquinaria-detalle": "Detalle Maquinaria",
	"servicios":        "Servicios",
	"servicio-detalle": "Detalle Servicio",
	"equipo":           "Equipo",
	"reservas":         "Reservas",
	"calculadora":      "Calculadora FFMI",
	"apuntate":         "Apúntate",
	"login":            "Login",
	"buscar":           "Buscar",
	"reglas":           "Reglas",
	"tramitar-pedido":  "Tramitar Pedido",
}

// InitTemplates parsea todas las plantillas del proyecto una sola vez al arrancar
// el servidor y las guarda en templateCache. Debe llamarse desde main() antes de
// registrar los handlers. Recibe la lista de nombres de página a precargar.
func InitTemplates(pages []string) {
	layoutFile := GetWebPath("templates", "layout.html")
	navbarFile := GetWebPath("templates", "navbar.html")
	footerFile := GetWebPath("templates", "footer.html")

	templateCacheMu.Lock()
	defer templateCacheMu.Unlock()

	for _, page := range pages {
		tmplFile := GetWebPath("templates", page+".html")
		tmpl, err := template.ParseFiles(layoutFile, tmplFile, navbarFile, footerFile)
		if err != nil {
			log.Fatalf("Error precargando plantilla '%s': %v", page, err)
		}
		templateCache[page] = tmpl
	}
	log.Printf("Plantillas precargadas: %d", len(templateCache))
}

// RenderTemplate ejecuta la plantilla correspondiente a la página indicada,
// pasando data al contexto de la plantilla. Utiliza el caché de plantillas
// inicializado por InitTemplates para evitar accesos a disco en cada petición.
// Si la plantilla no está en caché, devuelve un error HTTP 500.
func RenderTemplate(w http.ResponseWriter, page string, data interface{}) {
	templateCacheMu.RLock()
	tmpl, ok := templateCache[page]
	templateCacheMu.RUnlock()

	if !ok {
		log.Printf("Plantilla no encontrada en caché: %s", page)
		http.Error(w, "Página no encontrada", http.StatusNotFound)
		return
	}

	title, exists := pageTitles[page]
	if !exists {
		title = page
	}

	err := tmpl.Execute(w, map[string]interface{}{
		"Title": title,
		"Data":  data,
	})
	if err != nil {
		log.Printf("Error ejecutando plantilla %s: %v", page, err)
	}
}
