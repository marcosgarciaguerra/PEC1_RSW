package handlers

import (
	"encoding/json"
	"net/http"
	"pec2/internal/db"
	"pec2/internal/models"
	"strconv"
	"strings"
)

// HandleAPIClases maneja el CRUD de clases para la API REST.
func HandleAPIClases(w http.ResponseWriter, r *http.Request) {
	// CORS si es necesario, o simplemente Content-Type
	w.Header().Set("Content-Type", "application/json")

	path := strings.TrimPrefix(r.URL.Path, "/api/clases")
	path = strings.TrimPrefix(path, "/")

	if path == "" {
		switch r.Method {
		case http.MethodGet:
			// Obtener todas las clases
			clases := db.ObtenerClases()
			json.NewEncoder(w).Encode(clases)
		case http.MethodPost:
			// Crear una nueva clase
			var nuevaClase models.Clases
			if err := json.NewDecoder(r.Body).Decode(&nuevaClase); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			id, err := db.CrearClase(nuevaClase)
			if err != nil {
				http.Error(w, "Error creando clase", http.StatusInternalServerError)
				return
			}
			nuevaClase.ID = id
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(nuevaClase)
		default:
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
		return
	}

	// Manejar /api/clases/{id}
	id, err := strconv.Atoi(path)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		clase, err := db.ObtenerClasePorID(id)
		if err != nil {
			http.Error(w, "Clase no encontrada", http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(clase)
	case http.MethodPut:
		var claseActualizada models.Clases
		if err := json.NewDecoder(r.Body).Decode(&claseActualizada); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		claseActualizada.ID = id
		if err := db.ActualizarClase(claseActualizada); err != nil {
			http.Error(w, "Error actualizando clase", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(claseActualizada)
	case http.MethodDelete:
		if err := db.EliminarClase(id); err != nil {
			http.Error(w, "Error eliminando clase", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	default:
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}
