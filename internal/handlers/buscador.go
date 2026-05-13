package handlers

import (
	"net/http"
	"strings"
)

func BuscadorHandler(w http.ResponseWriter, r *http.Request) {
	query := strings.TrimSpace(r.URL.Query().Get("q"))
	resultados := BuscarContenido(query)

	RenderTemplate(w, "buscar", map[string]interface{}{
		"Query":      query,
		"Resultados": resultados,
	})
}

