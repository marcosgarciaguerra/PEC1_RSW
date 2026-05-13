package handlers

import (
	"errors"
	"net/http"
	"strconv"
)

func GuardarResenaHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	socio := obtenerSocioLogueado(r)
	if socio == nil {
		http.Redirect(w, r, "/login?return=/", http.StatusSeeOther)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error al procesar el formulario", http.StatusBadRequest)
		return
	}

	autor := socio.Nombre
	puntuacionStr := r.FormValue("puntuacion")
	puntuacion, err := strconv.Atoi(puntuacionStr)
	if err != nil {
		http.Error(w, "Puntuación inválida", http.StatusBadRequest)
		return
	}
	texto := r.FormValue("texto")

	if err := GuardarResena(autor, puntuacion, texto); err != nil {
		if errors.Is(err, ErrResenaInvalida) {
			http.Error(w, "Reseña inválida", http.StatusBadRequest)
			return
		}
		http.Error(w, "No se pudo guardar la reseña", http.StatusBadRequest)
		return
	}

	// Redirigir de vuelta al index
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
