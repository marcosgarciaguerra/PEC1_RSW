package handlers

import (
	"net/http"
	"pec2/internal/db"
	"pec2/internal/models"
	"pec2/internal/validation"
	"strconv"
)

const articulosBasePath = "/api/articulos"

// ListArticulos GET /api/articulos — admite ?pagina=&tamano= para paginación.
func ListArticulos(w http.ResponseWriter, r *http.Request) {
	paginaStr := r.URL.Query().Get("pagina")
	tamanoStr := r.URL.Query().Get("tamano")

	var articulos []models.Articulo
	if paginaStr != "" && tamanoStr != "" {
		pagina, err1 := strconv.Atoi(paginaStr)
		tamano, err2 := strconv.Atoi(tamanoStr)
		if err1 != nil || err2 != nil || pagina < 1 || tamano < 1 {
			writeJSONError(w, http.StatusBadRequest, "Parámetros de paginación inválidos")
			return
		}
		articulos = db.ObtenerArticulosPaginados(pagina, tamano)
	} else {
		articulos = db.ObtenerArticulosTienda()
	}

	writeJSONOK(w, articulos)
}

// CreateArticulo POST /api/articulos
func CreateArticulo(w http.ResponseWriter, r *http.Request) {
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
	writeJSONCreated(w, resourceLocation(articulosBasePath, id), nuevo)
}

// GetArticulo GET /api/articulos/{id}
func GetArticulo(w http.ResponseWriter, r *http.Request) {
	id, ok := parsePathID(w, r)
	if !ok {
		return
	}
	articulo, found := db.ObtenerArticuloPorID(id)
	if !found {
		writeJSONError(w, http.StatusNotFound, "Artículo no encontrado")
		return
	}
	writeJSONOK(w, articulo)
}

// UpdateArticulo PUT /api/articulos/{id}
func UpdateArticulo(w http.ResponseWriter, r *http.Request) {
	id, ok := parsePathID(w, r)
	if !ok {
		return
	}
	var actualizado models.Articulo
	if err := decodeJSONBody(w, r, &actualizado); err != nil {
		return
	}
	actualizado.ID = id
	if err := validation.ValidarArticulo(actualizado); err != nil {
		writeJSONError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}
	updated, err := db.ActualizarArticulo(actualizado)
	if err != nil {
		writeJSONError(w, http.StatusInternalServerError, "Error actualizando artículo")
		return
	}
	if !updated {
		writeJSONError(w, http.StatusNotFound, "Artículo no encontrado")
		return
	}
	writeJSONOK(w, actualizado)
}

// DeleteArticulo DELETE /api/articulos/{id}
func DeleteArticulo(w http.ResponseWriter, r *http.Request) {
	id, ok := parsePathID(w, r)
	if !ok {
		return
	}
	deleted, err := db.EliminarArticulo(id)
	if err != nil {
		writeJSONError(w, http.StatusInternalServerError, "Error eliminando artículo")
		return
	}
	if !deleted {
		writeJSONError(w, http.StatusNotFound, "Artículo no encontrado")
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
