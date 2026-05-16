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

// HandleAPIMaquinarias maneja el CRUD REST de maquinaria del gimnasio.
func HandleAPIMaquinarias(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	path := strings.TrimPrefix(r.URL.Path, "/api/maquinarias")
	path = strings.TrimPrefix(path, "/")

	if path == "" {
		switch r.Method {
		case http.MethodGet:
			maquinarias := db.ListarMaquinas()
			json.NewEncoder(w).Encode(maquinarias)
		case http.MethodPost:
			var nuevo models.Maquina
			if err := decodeJSONBody(w, r, &nuevo); err != nil {
				return
			}
			if err := validation.ValidarMaquina(nuevo); err != nil {
				writeJSONError(w, http.StatusUnprocessableEntity, err.Error())
				return
			}
			id, err := db.CrearMaquina(nuevo)
			if err != nil {
				writeJSONError(w, http.StatusInternalServerError, "Error creando maquinaria")
				return
			}
			nuevo.ID = int(id)
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
		maquina, ok := db.BuscarMaquinaPorID(id)
		if !ok {
			writeJSONError(w, http.StatusNotFound, "Maquinaria no encontrada")
			return
		}
		json.NewEncoder(w).Encode(maquina)
	case http.MethodPut:
		var actualizado models.Maquina
		if err := decodeJSONBody(w, r, &actualizado); err != nil {
			return
		}
		actualizado.ID = id
		if err := validation.ValidarMaquina(actualizado); err != nil {
			writeJSONError(w, http.StatusUnprocessableEntity, err.Error())
			return
		}
		ok, err := db.ActualizarMaquina(actualizado)
		if err != nil {
			writeJSONError(w, http.StatusInternalServerError, "Error actualizando maquinaria")
			return
		}
		if !ok {
			writeJSONError(w, http.StatusNotFound, "Maquinaria no encontrada")
			return
		}
		json.NewEncoder(w).Encode(actualizado)
	case http.MethodDelete:
		ok, err := db.EliminarMaquina(id)
		if err != nil {
			writeJSONError(w, http.StatusInternalServerError, "Error eliminando maquinaria")
			return
		}
		if !ok {
			writeJSONError(w, http.StatusNotFound, "Maquinaria no encontrada")
			return
		}
		w.WriteHeader(http.StatusNoContent)
	default:
		writeJSONError(w, http.StatusMethodNotAllowed, "Método no permitido")
	}
}
