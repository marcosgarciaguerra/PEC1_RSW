package handlers

import (
	"encoding/json"
	"net/http"
	"pec2/internal/db"
	"pec2/internal/models"
	"strconv"
	"strings"
)

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
			if err := json.NewDecoder(r.Body).Decode(&nuevo); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			id, err := db.CrearMaquina(nuevo)
			if err != nil {
				http.Error(w, "Error creando maquinaria", http.StatusInternalServerError)
				return
			}
			nuevo.ID = int(id)
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
		maquina, ok := db.BuscarMaquinaPorID(id)
		if !ok {
			http.Error(w, "Maquinaria no encontrada", http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(maquina)
	case http.MethodPut:
		var actualizado models.Maquina
		if err := json.NewDecoder(r.Body).Decode(&actualizado); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		actualizado.ID = id
		if err := db.ActualizarMaquina(actualizado); err != nil {
			http.Error(w, "Error actualizando maquinaria", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(actualizado)
	case http.MethodDelete:
		if err := db.EliminarMaquina(id); err != nil {
			http.Error(w, "Error eliminando maquinaria", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	default:
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}
