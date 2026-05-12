package services

import (
	"time"

	"pec2/internal/db"
	"pec2/internal/models"
)

type ReservaDetalle struct {
	models.Reserva
	NombreClase string
}

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

func CrearReservaSocio(socioID int, claseID int) bool {
	fecha := time.Now().Add(24 * time.Hour).Format("2006-01-02")
	return db.CrearReserva(socioID, claseID, fecha)
}

func CancelarReservaSocio(reservaID int, socioID int) bool {
	return db.EliminarReserva(reservaID, socioID)
}
