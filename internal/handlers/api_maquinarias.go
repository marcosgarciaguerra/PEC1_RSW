package handlers

import (
	"net/http"
	"pec2/internal/db"
	"pec2/internal/models"
	"pec2/internal/validation"
)

const maquinariasBasePath = "/api/maquinarias"

// ListMaquinarias GET /api/maquinarias
func ListMaquinarias(w http.ResponseWriter, r *http.Request) {
	maquinarias := db.ListarMaquinas()
	writeJSONOK(w, maquinarias)
}

// CreateMaquinaria POST /api/maquinarias
func CreateMaquinaria(w http.ResponseWriter, r *http.Request) {
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
	writeJSONCreated(w, resourceLocation(maquinariasBasePath, int(id)), nuevo)
}

// GetMaquinaria GET /api/maquinarias/{id}
func GetMaquinaria(w http.ResponseWriter, r *http.Request) {
	id, ok := parsePathID(w, r)
	if !ok {
		return
	}
	maquina, found := db.BuscarMaquinaPorID(id)
	if !found {
		writeJSONError(w, http.StatusNotFound, "Maquinaria no encontrada")
		return
	}
	writeJSONOK(w, maquina)
}

// UpdateMaquinaria PUT /api/maquinarias/{id}
func UpdateMaquinaria(w http.ResponseWriter, r *http.Request) {
	id, ok := parsePathID(w, r)
	if !ok {
		return
	}
	var actualizado models.Maquina
	if err := decodeJSONBody(w, r, &actualizado); err != nil {
		return
	}
	actualizado.ID = id
	if err := validation.ValidarMaquina(actualizado); err != nil {
		writeJSONError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}
	updated, err := db.ActualizarMaquina(actualizado)
	if err != nil {
		writeJSONError(w, http.StatusInternalServerError, "Error actualizando maquinaria")
		return
	}
	if !updated {
		writeJSONError(w, http.StatusNotFound, "Maquinaria no encontrada")
		return
	}
	writeJSONOK(w, actualizado)
}

// DeleteMaquinaria DELETE /api/maquinarias/{id}
func DeleteMaquinaria(w http.ResponseWriter, r *http.Request) {
	id, ok := parsePathID(w, r)
	if !ok {
		return
	}
	deleted, err := db.EliminarMaquina(id)
	if err != nil {
		writeJSONError(w, http.StatusInternalServerError, "Error eliminando maquinaria")
		return
	}
	if !deleted {
		writeJSONError(w, http.StatusNotFound, "Maquinaria no encontrada")
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
