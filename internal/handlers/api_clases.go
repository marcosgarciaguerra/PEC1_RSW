package handlers

import (
	"net/http"
	"pec2/internal/db"
	"pec2/internal/models"
	"pec2/internal/validation"
)

const clasesBasePath = "/api/clases"

// ListClases GET /api/clases
func ListClases(w http.ResponseWriter, r *http.Request) {
	clases := db.ObtenerClases()
	writeJSONOK(w, clases)
}

// CreateClase POST /api/clases
func CreateClase(w http.ResponseWriter, r *http.Request) {
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
	writeJSONCreated(w, resourceLocation(clasesBasePath, id), nuevaClase)
}

// GetClase GET /api/clases/{id}
func GetClase(w http.ResponseWriter, r *http.Request) {
	id, ok := parsePathID(w, r)
	if !ok {
		return
	}
	clase, err := db.ObtenerClasePorID(id)
	if err != nil {
		writeJSONError(w, http.StatusNotFound, "Clase no encontrada")
		return
	}
	writeJSONOK(w, clase)
}

// UpdateClase PUT /api/clases/{id}
func UpdateClase(w http.ResponseWriter, r *http.Request) {
	id, ok := parsePathID(w, r)
	if !ok {
		return
	}
	var claseActualizada models.Clases
	if err := decodeJSONBody(w, r, &claseActualizada); err != nil {
		return
	}
	claseActualizada.ID = id
	if err := validation.ValidarClase(claseActualizada); err != nil {
		writeJSONError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}
	updated, err := db.ActualizarClase(claseActualizada)
	if err != nil {
		writeJSONError(w, http.StatusInternalServerError, "Error actualizando clase")
		return
	}
	if !updated {
		writeJSONError(w, http.StatusNotFound, "Clase no encontrada")
		return
	}
	writeJSONOK(w, claseActualizada)
}

// DeleteClase DELETE /api/clases/{id}
func DeleteClase(w http.ResponseWriter, r *http.Request) {
	id, ok := parsePathID(w, r)
	if !ok {
		return
	}
	deleted, err := db.EliminarClase(id)
	if err != nil {
		writeJSONError(w, http.StatusInternalServerError, "Error eliminando clase")
		return
	}
	if !deleted {
		writeJSONError(w, http.StatusNotFound, "Clase no encontrada")
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
