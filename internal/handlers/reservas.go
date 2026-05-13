package handlers

import (
	"errors"
	"net/http"
	"pec2/internal/services"
	"strconv"
)

func ReservasHandler(w http.ResponseWriter, r *http.Request) {
	socio := obtenerSocioLogueado(r)
	if socio == nil {
		http.Redirect(w, r, "/login?return="+r.URL.Path, http.StatusSeeOther)
		return
	}

	clasesLista, misReservasDetalle := services.ObtenerReservasDetalleSocio(socio.ID)

	errorParam := r.URL.Query().Get("error")
	var errorMsg string
	switch errorParam {
	case "clase_invalida":
		errorMsg = "ID de clase inválido."
	case "reserva_fallida":
		errorMsg = "No se pudo realizar la reserva. Es posible que el aforo esté completo o ya tengas esta clase reservada."
	case "reserva_invalida":
		errorMsg = "No se pudo cancelar la reserva."
	}

	RenderTemplate(w, "reservas", map[string]interface{}{
		"Socio":       socio,
		"Clases":      clasesLista,
		"MisReservas": misReservasDetalle,
		"ErrorMsg":    errorMsg,
	})
}

func ProcesarReservaHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/reservas", http.StatusSeeOther)
		return
	}

	socio := obtenerSocioLogueado(r)
	if socio == nil {
		http.Redirect(w, r, "/login?return="+r.URL.Path, http.StatusSeeOther)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Redirect(w, r, "/reservas?error=clase_invalida", http.StatusSeeOther)
		return
	}
	claseIDStr := r.FormValue("clase_id")
	claseID, err := strconv.Atoi(claseIDStr)
	if err != nil || claseID <= 0 {
		http.Redirect(w, r, "/reservas?error=clase_invalida", http.StatusSeeOther)
		return
	}
	if err := services.CrearReservaSocio(socio.ID, claseID); errors.Is(err, services.ErrClaseInvalida) {
		http.Redirect(w, r, "/reservas?error=clase_invalida", http.StatusSeeOther)
		return
	} else if err != nil {
		http.Redirect(w, r, "/reservas?error=reserva_fallida", http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/reservas", http.StatusSeeOther)
}

func ProcesarCancelacionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/reservas", http.StatusSeeOther)
		return
	}

	socio := obtenerSocioLogueado(r)
	if socio == nil {
		http.Redirect(w, r, "/login?return="+r.URL.Path, http.StatusSeeOther)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Redirect(w, r, "/reservas?error=reserva_invalida", http.StatusSeeOther)
		return
	}
	reservaIDStr := r.FormValue("reserva_id")
	reservaID, err := strconv.Atoi(reservaIDStr)
	if err != nil || reservaID <= 0 {
		http.Redirect(w, r, "/reservas?error=reserva_invalida", http.StatusSeeOther)
		return
	}

	if err := services.CancelarReservaSocio(reservaID, socio.ID); err != nil {
		http.Redirect(w, r, "/reservas?error=reserva_invalida", http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/reservas", http.StatusSeeOther)
}
