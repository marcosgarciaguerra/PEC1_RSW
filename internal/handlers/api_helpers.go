package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
)

const maxAPIBodyBytes = 1 << 20 // 1 MiB

// writeJSONError envía un error en formato JSON manteniendo Content-Type application/json.
func writeJSONError(w http.ResponseWriter, code int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(map[string]string{"error": msg})
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
