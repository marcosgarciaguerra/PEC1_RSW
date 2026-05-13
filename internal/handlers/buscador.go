package handlers

import (
	"net/http"
	"pec2/internal/services"
	"strings"
)

func BuscadorHandler(w http.ResponseWriter, r *http.Request) {
	query := strings.TrimSpace(r.URL.Query().Get("q"))
	resultados := services.BuscarContenido(query)

	RenderTemplate(w, "buscar", map[string]interface{}{
		"Query":      query,
		"Resultados": resultados,
	})
}

