package handlers

import (
	"strconv"
	"strings"

	"pec2/internal/db"
	"pec2/internal/models"
)

type ResultadoBusqueda struct {
	Titulo string
	Tipo   string
	URL    string
}

func ObtenerResenasInicio() []models.Resena {
	return db.ObtenerResenas()
}

func ObtenerMaquinas() []models.Maquina {
	return db.ListarMaquinas()
}

func ObtenerServicios() []models.Servicio {
	return db.ListarServicios()
}

func ObtenerEquipoCompleto() []models.MiembroEquipo {
	return db.ListarEquipo()
}

func ObtenerMaquinaPorID(id int) (models.Maquina, bool) {
	return db.BuscarMaquinaPorID(id)
}

func ObtenerServicioPorID(id int) (models.Servicio, bool) {
	return db.BuscarServicioPorID(id)
}

func BuscarContenido(query string) []ResultadoBusqueda {
	query = strings.ToLower(strings.TrimSpace(query))
	if query == "" {
		return nil
	}

	resultados := make([]ResultadoBusqueda, 0)

	for _, m := range db.ListarMaquinas() {
		if strings.Contains(strings.ToLower(m.Nombre), query) ||
			strings.Contains(strings.ToLower(m.Descripcion), query) ||
			strings.Contains(strings.ToLower(m.Marca), query) {
			resultados = append(resultados, ResultadoBusqueda{
				Titulo: m.Nombre + " (" + m.Marca + ")",
				Tipo:   "Maquinaria",
				URL:    "/maquinaria-detalle?id=" + strconv.Itoa(m.ID),
			})
		}
	}

	for _, s := range db.ListarServicios() {
		if strings.Contains(strings.ToLower(s.Nombre), query) ||
			strings.Contains(strings.ToLower(s.DescripcionBreve), query) ||
			strings.Contains(strings.ToLower(s.DescripcionLarga), query) {
			resultados = append(resultados, ResultadoBusqueda{
				Titulo: s.Nombre,
				Tipo:   "Servicio",
				URL:    "/servicio-detalle?id=" + strconv.Itoa(s.ID),
			})
		}
	}

	return resultados
}
