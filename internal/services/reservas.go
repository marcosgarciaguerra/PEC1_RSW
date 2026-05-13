package services

import (
	"errors"
	"time"

	"pec2/internal/db"
	"pec2/internal/models"
)

type ReservaDetalle struct {
	models.Reserva
	NombreClase string
}

var (
	ErrClaseInvalida   = errors.New("clase invalida")
	ErrReservaFallida  = errors.New("reserva fallida")
	ErrReservaInvalida = errors.New("reserva invalida")
)

func ObtenerReservasDetalleSocio(socioID int) ([]models.Clases, []ReservaDetalle) {
	misReservas := db.ObtenerReservasDeSocio(socioID)
	clasesLista := db.ObtenerClases()
	nombresClases := make(map[int]string, len(clasesLista))

	for _, clase := range clasesLista {
		nombresClases[clase.ID] = clase.NombreClase
	}

	reservasDetalle := make([]ReservaDetalle, 0, len(misReservas))
	for _, reserva := range misReservas {
		reservasDetalle = append(reservasDetalle, ReservaDetalle{
			Reserva:     reserva,
			NombreClase: nombresClases[reserva.ActividadID],
		})
	}

	return clasesLista, reservasDetalle
}

func CrearReservaSocio(socioID int, claseID int) error {
	if claseID <= 0 {
		return ErrClaseInvalida
	}

	claseExiste := false
	for _, clase := range db.ObtenerClases() {
		if clase.ID == claseID {
			claseExiste = true
			break
		}
	}
	if !claseExiste {
		return ErrClaseInvalida
	}

	fecha := time.Now().Add(24 * time.Hour).Format("2006-01-02")
	if ok := db.CrearReserva(socioID, claseID, fecha); !ok {
		return ErrReservaFallida
	}
	return nil
}

func CancelarReservaSocio(reservaID int, socioID int) error {
	if reservaID <= 0 {
		return ErrReservaInvalida
	}
	if ok := db.EliminarReserva(reservaID, socioID); !ok {
		return ErrReservaInvalida
	}
	return nil
}
