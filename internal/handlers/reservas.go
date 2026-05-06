package handlers

import (
	"net/http"
	"pec2/internal/auth"
	"pec2/internal/db"
	"pec2/internal/models"
	"strconv"
	"time"
)

// Devuelve el socio logueado o nil si no hay cookie/socio
func obtenerSocioLogueado(r *http.Request) *models.Socio {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		return nil
	}
	email, ok := auth.Get(cookie.Value)
	if !ok {
		return nil
	}
	return db.ObtenerSocioPorCorreo(email)
}

func ReservasHandler(w http.ResponseWriter, r *http.Request) {
	socio := obtenerSocioLogueado(r)
	if socio == nil {
		http.Redirect(w, r, "/login?return="+r.URL.Path, http.StatusSeeOther)
		return
	}

	misReservas := db.ObtenerReservasDeSocio(socio.ID)
	clasesLista := db.ObtenerClases()
	
	type ReservaDetalle struct {
		models.Reserva
		NombreClase string
	}
	
	var misReservasDetalle []ReservaDetalle
	for _, res := range misReservas {
		var nombre string
		for _, c := range clasesLista {
			if c.ID == res.ActividadID {
				nombre = c.NombreClase
				break
			}
		}
		misReservasDetalle = append(misReservasDetalle, ReservaDetalle{
			Reserva:     res,
			NombreClase: nombre,
		})
	}

	errorParam := r.URL.Query().Get("error")
	var errorMsg string
	if errorParam == "clase_invalida" {
		errorMsg = "ID de clase inválido."
	} else if errorParam == "reserva_fallida" {
		errorMsg = "No se pudo realizar la reserva. Es posible que el aforo esté completo o ya tengas esta clase reservada."
	} else if errorParam == "reserva_invalida" {
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

	r.ParseForm()
	claseIDStr := r.FormValue("clase_id")
	claseID, err := strconv.Atoi(claseIDStr)
	if err != nil || claseID <= 0 {
		http.Redirect(w, r, "/reservas?error=clase_invalida", http.StatusSeeOther)
		return
	}
	
	// Usamos la fecha de mañana por defecto para simplificar
	fecha := time.Now().Add(24 * time.Hour).Format("2006-01-02")
	
	exito := db.CrearReserva(socio.ID, claseID, fecha)
	if !exito {
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

	r.ParseForm()
	reservaIDStr := r.FormValue("reserva_id")
	reservaID, err := strconv.Atoi(reservaIDStr)
	if err != nil || reservaID <= 0 {
		http.Redirect(w, r, "/reservas?error=reserva_invalida", http.StatusSeeOther)
		return
	}

	db.EliminarReserva(reservaID, socio.ID)

	http.Redirect(w, r, "/reservas", http.StatusSeeOther)
}


