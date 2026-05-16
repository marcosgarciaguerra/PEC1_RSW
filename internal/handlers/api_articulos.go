package handlers

import (
	"encoding/json"
	"net/http"
	"pec2/internal/db"
	"pec2/internal/models"
	"pec2/internal/validation"
	"strconv"
	"strings"
)

// HandleAPIArticulos maneja el CRUD REST de artículos de la tienda.
func HandleAPIArticulos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	path := strings.TrimPrefix(r.URL.Path, "/api/articulos")
	path = strings.TrimPrefix(path, "/")

	if path == "" {
		switch r.Method {
		case http.MethodGet:
			articulos := db.ObtenerArticulosTienda()
			json.NewEncoder(w).Encode(articulos)
		case http.MethodPost:
			var nuevo models.Articulo
			if err := decodeJSONBody(w, r, &nuevo); err != nil {
				return
			}
			if err := validation.ValidarArticulo(nuevo); err != nil {
				writeJSONError(w, http.StatusUnprocessableEntity, err.Error())
				return
			}
			id, err := db.CrearArticulo(nuevo)
			if err != nil {
				writeJSONError(w, http.StatusInternalServerError, "Error creando artículo")
				return
			}
			nuevo.ID = id
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(nuevo)
		default:
			writeJSONError(w, http.StatusMethodNotAllowed, "Método no permitido")
		}
		return
	}

	id, err := strconv.Atoi(path)
	if err != nil {
		writeJSONError(w, http.StatusBadRequest, "ID inválido")
		return
	}

	switch r.Method {
	case http.MethodGet:
		articulo, ok := db.ObtenerArticuloPorID(id)
		if !ok {
			writeJSONError(w, http.StatusNotFound, "Artículo no encontrado")
			return
		}
		json.NewEncoder(w).Encode(articulo)
	case http.MethodPut:
		var actualizado models.Articulo
		if err := decodeJSONBody(w, r, &actualizado); err != nil {
			return
		}
		actualizado.ID = id
		if err := validation.ValidarArticulo(actualizado); err != nil {
			writeJSONError(w, http.StatusUnprocessableEntity, err.Error())
			return
		}
		ok, err := db.ActualizarArticulo(actualizado)
		if err != nil {
			writeJSONError(w, http.StatusInternalServerError, "Error actualizando artículo")
			return
		}
		if !ok {
			writeJSONError(w, http.StatusNotFound, "Artículo no encontrado")
			return
		}
		json.NewEncoder(w).Encode(actualizado)
	case http.MethodDelete:
		ok, err := db.EliminarArticulo(id)
		if err != nil {
			writeJSONError(w, http.StatusInternalServerError, "Error eliminando artículo")
			return
		}
		if !ok {
			writeJSONError(w, http.StatusNotFound, "Artículo no encontrado")
			return
		}
		w.WriteHeader(http.StatusNoContent)
	default:
		writeJSONError(w, http.StatusMethodNotAllowed, "Método no permitido")
	}
}
