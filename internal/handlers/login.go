package handlers

import (
	"net/http"
	"pec2/internal/auth"
	"pec2/internal/db"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	returnPath := r.URL.Query().Get("return")

	switch r.Method {
	case http.MethodGet:
		RenderTemplate(w, "login", map[string]interface{}{
			"ReturnPath": returnPath,
		})
	case http.MethodPost:
		r.ParseForm()
		correo := r.FormValue("correo")
		contrasena := r.FormValue("contrasena")
		returnPath = r.FormValue("return") // Leer desde el input hidden

		socio := db.ObtenerSocioPorCorreo(correo)
		if socio != nil && bcrypt.CompareHashAndPassword([]byte(socio.Contrasena), []byte(contrasena)) == nil {
			// Login exitoso - generate secure session token
			token := uuid.NewString()
			auth.Set(token, correo)
			http.SetCookie(w, &http.Cookie{
				Name:     "session_id",
				Value:    token,
				Expires:  time.Now().Add(24 * time.Hour),
				Path:     "/",
				HttpOnly: true,
				Secure:   true,
				SameSite: http.SameSiteLaxMode,
			})
			http.SetCookie(w, &http.Cookie{
				Name:    "session_nombre",
				Value:   socio.Nombre,
				Expires: time.Now().Add(24 * time.Hour),
				Path:    "/",
			})

			// Redirigir al sitio original si existe, si no a /reservas
			if returnPath != "" {
				http.Redirect(w, r, returnPath, http.StatusSeeOther)
			} else {
				http.Redirect(w, r, "/reservas", http.StatusSeeOther)
			}
			return
		}

		// Login fallido
		RenderTemplate(w, "login", map[string]interface{}{
			"ErrorMsg":   "Credenciales inválidas. Por favor, inténtelo de nuevo.",
			"ReturnPath": returnPath,
		})
	}
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
		if cookie, err := r.Cookie("session_id"); err == nil {
			auth.Delete(cookie.Value)
			http.SetCookie(w, &http.Cookie{
				Name:    "session_id",
				Value:   "",
				Expires: time.Unix(0, 0),
				Path:    "/",
			})
		}
	http.SetCookie(w, &http.Cookie{
		Name:    "session_nombre",
		Value:   "",
		Expires: time.Unix(0, 0),
		Path:    "/",
	})
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

