package handlers

import (
	"net/http"
	"pec2/internal/auth"
	"pec2/internal/db"
)

func TramitarPedidoHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		http.Redirect(w, r, "/login?return=/tienda/tramitar", http.StatusSeeOther)
		return
	}
	email, ok := auth.Get(cookie.Value)
	if !ok {
		http.Redirect(w, r, "/login?return=/tienda/tramitar", http.StatusSeeOther)
		return
	}
	usuario := db.ObtenerUsuarioPorCorreo(email)
	if usuario == nil {
		http.Redirect(w, r, "/login?return="+r.URL.Path, http.StatusSeeOther)
		return
	}

	RenderTemplate(w, "tramitar-pedido", map[string]interface{}{
		"Usuario": usuario,
	})
}
