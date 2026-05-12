package handlers

import (
	"net/http"
	"pec2/internal/db"
)

func TramitarPedidoHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		http.Redirect(w, r, "/login?return=/tienda/tramitar", http.StatusSeeOther)
		return
	}
	usuario := db.ObtenerUsuarioPorCorreo(cookie.Value)
	if usuario == nil {
		http.Redirect(w, r, "/login?return="+r.URL.Path, http.StatusSeeOther)
		return
	}

	RenderTemplate(w, "tramitar-pedido", map[string]interface{}{
		"Usuario": usuario,
	})
}
