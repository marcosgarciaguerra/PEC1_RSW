package handlers

import (
	"log"
	"net/http"
	"path/filepath"
	"pec2/internal/auth"
	"pec2/internal/db"
	"pec2/internal/models"
	"strconv"
	"strings"
)

func GetWebPath(relative ...string) string {
	webBase := "web"
	// Si no existe ./web, probar ../../web (si corremos desde cmd/servidor)
	if _, err := http.Dir(webBase).Open("."); err != nil {
		webBase = "../../web"
	}
	parts := append([]string{webBase}, relative...)
	return filepath.Join(parts...)
}

func PageHandler(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Path[1:]
	if page == "" || page == "index.html" {
		page = "index"
	}

	// Quitar .html si el usuario o un enlace navega directamente con la extensión
	page = strings.TrimSuffix(page, ".html")

	if strings.Contains(page, ".") {
		http.NotFound(w, r)
		return
	}

	log.Printf("Usuario accedió a la página: %s", page)

	var data interface{}

	switch page {
	case "index":
		resenas := db.ObtenerResenas()
		var usuario *models.Socio
	cookie, err := r.Cookie("session_id")
	if err == nil {
		if email, ok := auth.Get(cookie.Value); ok {
			usuario = db.ObtenerSocioPorCorreo(email)
		}
	}
		data = map[string]interface{}{
			"Resenas": resenas,
			"Usuario": usuario,
		}
	case "maquinaria":
		data = db.Maquinas
	case "maquinaria-detalle":
		idStr := r.URL.Query().Get("id")
		id, _ := strconv.Atoi(idStr)
		for _, m := range db.Maquinas {
			if m.ID == id {
				data = m
				break
			}
		}
	case "servicios":
		data = db.Servicios
	case "equipo":
		data = db.EquipoCompleto
	case "servicio-detalle":
		idStr := r.URL.Query().Get("id")
		id, _ := strconv.Atoi(idStr)
		for _, s := range db.Servicios {
			if s.ID == id {
				data = s
				break
			}
		}
	case "tienda/tramitar":
		cookie, err := r.Cookie("session_id")
		if err == nil {
			if email, ok := auth.Get(cookie.Value); ok {
				usuario := db.ObtenerSocioPorCorreo(email)
				if usuario != nil {
					data = map[string]interface{}{"Usuario": usuario}
					page = "tramitar-pedido"
				} else {
					http.Redirect(w, r, "/login", http.StatusSeeOther)
					return
				}
			} else {
				http.Redirect(w, r, "/login?return=/tienda/tramitar", http.StatusSeeOther)
				return
			}
		} else {
			http.Redirect(w, r, "/login?return=/tienda/tramitar", http.StatusSeeOther)
			return
		}
	}

	RenderTemplate(w, page, data)
}

