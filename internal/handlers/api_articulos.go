package handlers

import (
	"encoding/json"
	"net/http"
	"pec2/internal/db"
	"pec2/internal/models"
	"strconv"
	"strings"
)

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
			if err := json.NewDecoder(r.Body).Decode(&nuevo); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			id, err := db.CrearArticulo(nuevo)
			if err != nil {
				http.Error(w, "Error creando artículo", http.StatusInternalServerError)
				return
			}
			nuevo.ID = id
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(nuevo)
		default:
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
		return
	}

	id, err := strconv.Atoi(path)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		articulo, ok := db.ObtenerArticuloPorID(id)
		if !ok {
			http.Error(w, "Artículo no encontrado", http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(articulo)
	case http.MethodPut:
		var actualizado models.Articulo
		if err := json.NewDecoder(r.Body).Decode(&actualizado); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		actualizado.ID = id
		if err := db.ActualizarArticulo(actualizado); err != nil {
			http.Error(w, "Error actualizando artículo", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(actualizado)
	case http.MethodDelete:
		if err := db.EliminarArticulo(id); err != nil {
			http.Error(w, "Error eliminando artículo", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	default:
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}
