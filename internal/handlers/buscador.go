package handlers

import (
	"net/http"
	"pec2/internal/db"
	"strconv"
	"strings"
)

type ResultadoBusqueda struct {
	Titulo string
	Tipo   string
	URL    string
}

func BuscadorHandler(w http.ResponseWriter, r *http.Request) {
	query := strings.ToLower(r.URL.Query().Get("q"))
	var resultados []ResultadoBusqueda

	if query != "" {
		for _, m := range db.Maquinas {
			if strings.Contains(strings.ToLower(m.Nombre), query) || strings.Contains(strings.ToLower(m.Descripcion), query) || strings.Contains(strings.ToLower(m.Marca), query) {
				resultados = append(resultados, ResultadoBusqueda{
					Titulo: m.Nombre + " (" + m.Marca + ")",
					Tipo:   "Maquinaria",
					URL:    "/maquinaria-detalle?id=" + strconv.Itoa(m.ID),
				})
			}
		}
		for _, s := range db.Servicios {
			if strings.Contains(strings.ToLower(s.Nombre), query) || strings.Contains(strings.ToLower(s.DescripcionBreve), query) || strings.Contains(strings.ToLower(s.DescripcionLarga), query) {
				resultados = append(resultados, ResultadoBusqueda{
					Titulo: s.Nombre,
					Tipo:   "Servicio",
					URL:    "/servicio-detalle?id=" + strconv.Itoa(s.ID),
				})
			}
		}
	}

	RenderTemplate(w, "buscar", map[string]interface{}{
		"Query":      r.URL.Query().Get("q"),
		"Resultados": resultados,
	})
}

