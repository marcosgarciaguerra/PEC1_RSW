package handlers

import (
	"errors"
	"log"
	"net/http"

	"pec2/internal/services"
)

func TramitarPedidoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Redirect(w, r, "/tienda", http.StatusSeeOther)
		return
	}

	usuario := obtenerUsuarioLogueado(r)
	if usuario == nil {
		http.Redirect(w, r, "/login?return=/tienda/tramitar", http.StatusSeeOther)
		return
	}

	RenderTemplate(w, "tramitar-pedido", map[string]interface{}{
		"Usuario": usuario,
	})
}

func ConfirmarPedidoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/tienda/tramitar", http.StatusSeeOther)
		return
	}

	usuario := obtenerUsuarioLogueado(r)
	if usuario == nil {
		http.Redirect(w, r, "/login?return=/tienda/tramitar", http.StatusSeeOther)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error al procesar el pedido", http.StatusBadRequest)
		return
	}

	pedido, err := services.CrearPedido(*usuario, r.FormValue("direccion_envio"), r.FormValue("carrito"))
	if errors.Is(err, services.ErrCarritoVacio) ||
		errors.Is(err, services.ErrArticuloInvalido) ||
		errors.Is(err, services.ErrDireccionInvalida) ||
		errors.Is(err, services.ErrCantidadInvalida) {
		RenderTemplate(w, "tramitar-pedido", map[string]interface{}{
			"Usuario":  usuario,
			"ErrorMsg": "No se pudo confirmar el pedido. Revisa el carrito y la dirección de envío.",
		})
		return
	}
	if err != nil {
		log.Println("Error guardando pedido:", err)
		http.Error(w, "No se pudo guardar el pedido", http.StatusInternalServerError)
		return
	}

	RenderTemplate(w, "tramitar-pedido", map[string]interface{}{
		"Usuario":          usuario,
		"Pedido":           pedido,
		"PedidoConfirmado": true,
	})
}
