package handlers

import (
	"net/http"
	"strings"

	"pec2/internal/models"
	"pec2/internal/services"
)

const sessionCookieName = "session_token"

func obtenerUsuarioLogueado(r *http.Request) *models.Usuario {
	cookie, err := r.Cookie(sessionCookieName)
	if err != nil {
		return nil
	}
	usuario, err := services.GetUserFromSessionToken(cookie.Value)
	if err != nil {
		return nil
	}
	return usuario
}

func obtenerSocioLogueado(r *http.Request) *models.Socio {
	cookie, err := r.Cookie(sessionCookieName)
	if err != nil {
		return nil
	}
	socio, err := services.GetSocioFromSessionToken(cookie.Value)
	if err != nil {
		return nil
	}
	return socio
}

func RequireSocioAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if obtenerSocioLogueado(r) == nil {
			http.Redirect(w, r, "/login?return="+sanitizarReturnPath(r.URL.Path), http.StatusSeeOther)
			return
		}
		next(w, r)
	}
}

func RequireUsuarioAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if obtenerUsuarioLogueado(r) == nil {
			http.Redirect(w, r, "/login?return="+sanitizarReturnPath(r.URL.Path), http.StatusSeeOther)
			return
		}
		next(w, r)
	}
}

func sanitizarReturnPath(path string) string {
	path = strings.TrimSpace(path)
	if path == "" || !strings.HasPrefix(path, "/") {
		return "/"
	}
	return path
}
