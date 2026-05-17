package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

const maxAPIBodyBytes = 1 << 20 // 1 MiB

// WithCORS envuelve handlers API con cabeceras CORS y responde preflight OPTIONS con 204.
func WithCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next(w, r)
	}
}

// APIRoute combina CORS y autenticación de socio para rutas REST.
func APIRoute(handler http.HandlerFunc) http.HandlerFunc {
	return WithCORS(RequireSocioAuth(handler))
}

func setJSONContentType(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

// writeJSONError envía un error en formato JSON manteniendo Content-Type application/json.
func writeJSONError(w http.ResponseWriter, code int, msg string) {
	setJSONContentType(w)
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(map[string]string{"error": msg})
}

// writeJSONCreated responde 201 con cabecera Location y cuerpo JSON.
func writeJSONCreated(w http.ResponseWriter, location string, payload interface{}) {
	setJSONContentType(w)
	w.Header().Set("Location", location)
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(payload)
}

// parsePathID extrae y valida el parámetro de ruta {id} (Go 1.22+).
func parsePathID(w http.ResponseWriter, r *http.Request) (int, bool) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id <= 0 {
		writeJSONError(w, http.StatusBadRequest, "ID inválido")
		return 0, false
	}
	return id, true
}

func writeJSONOK(w http.ResponseWriter, payload interface{}) {
	setJSONContentType(w)
	_ = json.NewEncoder(w).Encode(payload)
}

func resourceLocation(basePath string, id int) string {
	return fmt.Sprintf("%s/%d", basePath, id)
}

// decodeJSONBody lee el cuerpo con límite de tamaño y rechaza campos desconocidos.
func decodeJSONBody(w http.ResponseWriter, r *http.Request, v interface{}) error {
	r.Body = http.MaxBytesReader(w, r.Body, maxAPIBodyBytes)
	defer r.Body.Close()

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(v); err != nil {
		var maxErr *http.MaxBytesError
		if errors.As(err, &maxErr) {
			writeJSONError(w, http.StatusRequestEntityTooLarge, "cuerpo de petición demasiado grande")
			return err
		}
		writeJSONError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return err
	}
	return nil
}
