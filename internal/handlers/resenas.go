package handlers

import (
	"net/http"
	"pec2/internal/db"
	"pec2/internal/models"
	"strconv"
)

func GuardarResenaHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	// Verificar sesión (la cookie contiene el correo directamente)
	cookie, err := r.Cookie("session_id")
	if err != nil {
		http.Redirect(w, r, "/login?return=/index", http.StatusSeeOther)
		return
	}
	usuario := db.ObtenerSocioPorCorreo(cookie.Value)
	if usuario == nil {
		http.Redirect(w, r, "/login?return=/index", http.StatusSeeOther)
		return
	}

	err = r.ParseForm()
	if err != nil {
		http.Error(w, "Error al procesar el formulario", http.StatusBadRequest)
		return
	}

	autor := usuario.Nombre
	puntuacionStr := r.FormValue("puntuacion")
	puntuacion, _ := strconv.Atoi(puntuacionStr)
	texto := r.FormValue("texto")

	if texto != "" && puntuacion >= 1 && puntuacion <= 5 {
		resena := models.Resena{
			Autor:      autor,
			Puntuacion: puntuacion,
			Texto:      texto,
		}
		db.GuardarResena(resena)
	}

	// Redirigir de vuelta al index
	http.Redirect(w, r, "/index", http.StatusSeeOther)
}
