package handlers

import (
	"html/template"
	"log"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, page string, data interface{}) {
	layoutFile := GetWebPath("templates", "layout.html")
	tmplFile := GetWebPath("templates", page+".html")
	navbarFile := GetWebPath("templates", "navbar.html")
	footerFile := GetWebPath("templates", "footer.html")

	tmpl, err := template.ParseFiles(layoutFile, tmplFile, navbarFile, footerFile)
	if err != nil {
		log.Printf("Error cargando plantillas para %s: %v", page, err)
		http.Error(w, "Error interno del servidor", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, map[string]interface{}{
		"Title": page,
		"Data":  data,
	})
	if err != nil {
		log.Printf("Error ejecutando plantilla %s: %v", page, err)
	}
}
