package db

import "pec2/internal/models"

func ListarMaquinas() []models.Maquina {
	return Maquinas
}

func ListarServicios() []models.Servicio {
	return Servicios
}

func ListarEquipo() []models.MiembroEquipo {
	return EquipoCompleto
}

func BuscarMaquinaPorID(id int) (models.Maquina, bool) {
	for _, m := range Maquinas {
		if m.ID == id {
			return m, true
		}
	}
	return models.Maquina{}, false
}

func BuscarServicioPorID(id int) (models.Servicio, bool) {
	for _, s := range Servicios {
		if s.ID == id {
			return s, true
		}
	}
	return models.Servicio{}, false
}
