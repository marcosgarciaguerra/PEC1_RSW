package handlers

import (
	"net/http"

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
