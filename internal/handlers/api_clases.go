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

// HandleAPIClases maneja el CRUD REST de clases del gimnasio.
func HandleAPIClases(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	path := strings.TrimPrefix(r.URL.Path, "/api/clases")
	path = strings.TrimPrefix(path, "/")

	if path == "" {
		switch r.Method {
		case http.MethodGet:
			clases := db.ObtenerClases()
			json.NewEncoder(w).Encode(clases)
		case http.MethodPost:
			var nuevaClase models.Clases
			if err := decodeJSONBody(w, r, &nuevaClase); err != nil {
				return
			}
			if err := validation.ValidarClase(nuevaClase); err != nil {
				writeJSONError(w, http.StatusUnprocessableEntity, err.Error())
				return
			}
			id, err := db.CrearClase(nuevaClase)
			if err != nil {
				writeJSONError(w, http.StatusInternalServerError, "Error creando clase")
				return
			}
			nuevaClase.ID = id
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(nuevaClase)
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
		clase, err := db.ObtenerClasePorID(id)
		if err != nil {
			writeJSONError(w, http.StatusNotFound, "Clase no encontrada")
			return
		}
		json.NewEncoder(w).Encode(clase)
	case http.MethodPut:
		var claseActualizada models.Clases
		if err := decodeJSONBody(w, r, &claseActualizada); err != nil {
			return
		}
		claseActualizada.ID = id
		if err := validation.ValidarClase(claseActualizada); err != nil {
			writeJSONError(w, http.StatusUnprocessableEntity, err.Error())
			return
		}
		ok, err := db.ActualizarClase(claseActualizada)
		if err != nil {
			writeJSONError(w, http.StatusInternalServerError, "Error actualizando clase")
			return
		}
		if !ok {
			writeJSONError(w, http.StatusNotFound, "Clase no encontrada")
			return
		}
		json.NewEncoder(w).Encode(claseActualizada)
	case http.MethodDelete:
		ok, err := db.EliminarClase(id)
		if err != nil {
			writeJSONError(w, http.StatusInternalServerError, "Error eliminando clase")
			return
		}
		if !ok {
			writeJSONError(w, http.StatusNotFound, "Clase no encontrada")
			return
		}
		w.WriteHeader(http.StatusNoContent)
	default:
		writeJSONError(w, http.StatusMethodNotAllowed, "Método no permitido")
	}
}
