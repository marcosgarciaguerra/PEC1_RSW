package db

import "pec2/internal/models"

func ListarServicios() []models.Servicio {
	return Servicios
}

func ListarEquipo() []models.MiembroEquipo {
	return EquipoCompleto
}

func BuscarServicioPorID(id int) (models.Servicio, bool) {
	for _, s := range Servicios {
		if s.ID == id {
			return s, true
		}
	}
	return models.Servicio{}, false
}
