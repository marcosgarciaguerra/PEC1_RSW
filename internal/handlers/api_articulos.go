package handlers

import (
	"encoding/json"
	"net/http"
	"pec2/internal/db"
	"pec2/internal/models"
	"strconv"
	"strings"
)

// Helper para enviar respuestas JSON
func JSONResponse(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

// APIArticulosHandler maneja las peticiones a /api/articulos y /api/articulos/{id}
func APIArticulosHandler(w http.ResponseWriter, r *http.Request) {
	// Extraer ID de la URL si existe (ej. /api/articulos/5)
	path := strings.TrimPrefix(r.URL.Path, "/api/articulos")
	path = strings.TrimPrefix(path, "/")
	
	var id int
	var err error
	if path != "" {
		id, err = strconv.Atoi(path)
		if err != nil {
			JSONResponse(w, http.StatusBadRequest, map[string]string{"error": "ID inválido"})
			return
		}
	}

	switch r.Method {
	case http.MethodGet:
		if id > 0 {
			getArticulo(w, r, id)
		} else {
			getArticulos(w, r)
		}
	case http.MethodPost:
		if id > 0 {
			JSONResponse(w, http.StatusMethodNotAllowed, map[string]string{"error": "Método no permitido en esta ruta"})
			return
		}
		crearArticulo(w, r)
	case http.MethodPut:
		if id <= 0 {
			JSONResponse(w, http.StatusMethodNotAllowed, map[string]string{"error": "Se requiere un ID para actualizar"})
			return
		}
		actualizarArticulo(w, r, id)
	case http.MethodDelete:
		if id <= 0 {
			JSONResponse(w, http.StatusMethodNotAllowed, map[string]string{"error": "Se requiere un ID para eliminar"})
			return
		}
		eliminarArticulo(w, r, id)
	default:
		JSONResponse(w, http.StatusMethodNotAllowed, map[string]string{"error": "Método no soportado"})
	}
}

func getArticulos(w http.ResponseWriter, r *http.Request) {
	articulos, err := db.ObtenerArticulos()
	if err != nil {
		JSONResponse(w, http.StatusInternalServerError, map[string]string{"error": "Error al obtener artículos"})
		return
	}
	JSONResponse(w, http.StatusOK, articulos)
}

func getArticulo(w http.ResponseWriter, r *http.Request, id int) {
	articulo, err := db.ObtenerArticulo(id)
	if err != nil {
		JSONResponse(w, http.StatusNotFound, map[string]string{"error": "Artículo no encontrado"})
		return
	}
	JSONResponse(w, http.StatusOK, articulo)
}

func crearArticulo(w http.ResponseWriter, r *http.Request) {
	var a models.Articulo
	if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
		JSONResponse(w, http.StatusBadRequest, map[string]string{"error": "Cuerpo de petición inválido"})
		return
	}

	id, err := db.CrearArticulo(a)
	if err != nil {
		JSONResponse(w, http.StatusInternalServerError, map[string]string{"error": "Error al crear artículo"})
		return
	}

	a.ID = id
	JSONResponse(w, http.StatusCreated, a)
}

func actualizarArticulo(w http.ResponseWriter, r *http.Request, id int) {
	var a models.Articulo
	if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
		JSONResponse(w, http.StatusBadRequest, map[string]string{"error": "Cuerpo de petición inválido"})
		return
	}

	err := db.ActualizarArticulo(id, a)
	if err != nil {
		JSONResponse(w, http.StatusInternalServerError, map[string]string{"error": "Error al actualizar artículo"})
		return
	}

	a.ID = id
	JSONResponse(w, http.StatusOK, a)
}

func eliminarArticulo(w http.ResponseWriter, r *http.Request, id int) {
	err := db.EliminarArticulo(id)
	if err != nil {
		JSONResponse(w, http.StatusInternalServerError, map[string]string{"error": "Error al eliminar artículo"})
		return
	}
	JSONResponse(w, http.StatusOK, map[string]string{"mensaje": "Artículo eliminado correctamente"})
}
