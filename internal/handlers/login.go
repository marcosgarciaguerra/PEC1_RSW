package handlers

import (
	"errors"
	"net/http"
	"strings"
	"time"
	"pec2/internal/validation"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	returnPath := SanitizarReturnPath(r.URL.Query().Get("return"))

	switch r.Method {
	case http.MethodGet:
		RenderTemplate(w, "login", map[string]interface{}{
			"ReturnPath": returnPath,
		})
	case http.MethodPost:
		if err := r.ParseForm(); err != nil {
			http.Error(w, "No se pudo procesar el login", http.StatusBadRequest)
			return
		}
		correo := strings.TrimSpace(r.FormValue("correo"))
		contrasena := strings.TrimSpace(r.FormValue("contrasena"))
		returnPath = SanitizarReturnPath(r.FormValue("return"))

		if err := validation.ValidarEmail(correo); err != nil || contrasena == "" {
			RenderTemplate(w, "login", map[string]interface{}{
				"ErrorMsg":   "Introduce credenciales válidas.",
				"ReturnPath": returnPath,
			})
			return
		}

		socio, err := Authenticate(correo, contrasena)
		if errors.Is(err, ErrCuentaInactiva) {
			RenderTemplate(w, "login", map[string]interface{}{
				"ErrorMsg":   "Tu cuenta está inactiva. Contacta con el gimnasio.",
				"ReturnPath": returnPath,
			})
			return
		}
		if err != nil {
			RenderTemplate(w, "login", map[string]interface{}{
				"ErrorMsg":   "Credenciales inválidas. Por favor, inténtelo de nuevo.",
				"ReturnPath": returnPath,
			})
			return
		}

		sessionToken, err := CreateUserSession(socio.ID, 24*time.Hour)
		if err != nil {
			http.Error(w, "No se pudo iniciar sesión", http.StatusInternalServerError)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:     sessionCookieName,
			Value:    sessionToken,
			Expires:  time.Now().Add(24 * time.Hour),
			Path:     "/",
			HttpOnly: true,
			SameSite: http.SameSiteLaxMode,
			Secure:   r.TLS != nil,
		})

		http.SetCookie(w, &http.Cookie{
			Name:     "gym_logged_in",
			Value:    "true",
			Expires:  time.Now().Add(24 * time.Hour),
			Path:     "/",
			HttpOnly: false,
			SameSite: http.SameSiteLaxMode,
			Secure:   r.TLS != nil,
		})

		http.SetCookie(w, &http.Cookie{
			Name:     "gym_user_name",
			Value:    socio.Nombre,
			Expires:  time.Now().Add(24 * time.Hour),
			Path:     "/",
			HttpOnly: false,
			SameSite: http.SameSiteLaxMode,
			Secure:   r.TLS != nil,
		})

		if returnPath != "" {
			http.Redirect(w, r, returnPath, http.StatusSeeOther)
		} else {
			http.Redirect(w, r, "/reservas", http.StatusSeeOther)
		}
	default:
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	if cookie, err := r.Cookie(sessionCookieName); err == nil {
		RevokeSession(cookie.Value)
	}
	http.SetCookie(w, &http.Cookie{
		Name:     sessionCookieName,
		Value:    "",
		Expires:  time.Unix(0, 0),
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Secure:   r.TLS != nil,
	})
	http.SetCookie(w, &http.Cookie{
		Name:     "gym_logged_in",
		Value:    "",
		Expires:  time.Unix(0, 0),
		Path:     "/",
		HttpOnly: false,
		SameSite: http.SameSiteLaxMode,
		Secure:   r.TLS != nil,
	})
	http.SetCookie(w, &http.Cookie{
		Name:     "gym_user_name",
		Value:    "",
		Expires:  time.Unix(0, 0),
		Path:     "/",
		HttpOnly: false,
		SameSite: http.SameSiteLaxMode,
		Secure:   r.TLS != nil,
	})
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
