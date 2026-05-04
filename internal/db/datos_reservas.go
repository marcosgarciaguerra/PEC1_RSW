package db

import "pec2/internal/models"

var ClasesLista = []models.Clases{
	{ID: 1, NombreClase: "Ciclo Indoor (Spinning)", Entrenador: "Carlos Méndez", Aforo: 20, Horario: "Lunes, 13 de abril de 2026 a las 18:00h", Descripcion: "Entrenamiento cardiovascular de alta intensidad en bicicleta estática, donde se simulan rutas con diferentes resistencias al ritmo de la música.", Lugar: "Sala de Spinning (Planta Baja)"},
	{ID: 2, NombreClase: "Yoga Vinyasa", Entrenador: "Elena Rostova", Aforo: 15, Horario: "Martes, 14 de abril de 2026 a las 09:00h", Descripcion: "Práctica fluida y dinámica que sincroniza la respiración con el movimiento continuo, enfocada en mejorar la flexibilidad, el equilibrio y reducir el estrés.", Lugar: "Sala Zen (Planta Alta)"},
	{ID: 3, NombreClase: "HIIT (High-Intensity Interval Training)", Entrenador: "Marcos Silva", Aforo: 12, Horario: "Miércoles, 15 de abril de 2026 a las 19:30h", Descripcion: "Sesión de entrenamiento cardiovascular que alterna intervalos cortos de ejercicios explosivos con periodos breves de recuperación para maximizar la quema de calorías.", Lugar: "Zona Funcional"},
	{ID: 4, NombreClase: "BodyPump", Entrenador: "Laura Gómez", Aforo: 20, Horario: "Jueves, 16 de abril de 2026 a las 14:00h", Descripcion: "Clase grupal de tonificación muscular y fuerza que utiliza barras y discos con pesos ajustables, trabajando todos los grupos musculares principales.", Lugar: "Sala de Actividades Dirigidas 1"},
	{ID: 5, NombreClase: "Zumba", Entrenador: "Sofía Valdés", Aforo: 25, Horario: "Viernes, 17 de abril de 2026 a las 18:30h", Descripcion: "Entrenamiento cardiovascular divertido y enérgico que combina rutinas de baile aeróbico con ritmos latinos e internacionales (salsa, reggaetón, merengue).", Lugar: "Sala de Actividades Dirigidas 2"},
	{ID: 6, NombreClase: "Pilates Mat", Entrenador: "Javier Ruiz", Aforo: 15, Horario: "Sábado, 18 de abril de 2026 a las 10:00h", Descripcion: "Secuencia de ejercicios en colchoneta diseñados para fortalecer la musculatura profunda (core), mejorar la postura, la flexibilidad y la alineación corporal.", Lugar: "Sala Zen (Planta Alta)"},
	{ID: 7, NombreClase: "Fitboxing", Entrenador: "David Castro", Aforo: 12, Horario: "Domingo, 19 de abril de 2026 a las 11:00h", Descripcion: "Entrenamiento de alta intensidad sin contacto que mezcla movimientos y técnicas de boxeo o kickboxing contra un saco, sincronizados con música y ejercicios funcionales.", Lugar: "Zona de Combate"},
	{ID: 8, NombreClase: "AquaGym", Entrenador: "Ana Morales", Aforo: 20, Horario: "Lunes, 20 de abril de 2026 a las 08:30h", Descripcion: "Gimnasia de bajo impacto realizada en el agua. Aprovecha la resistencia del agua para tonificar los músculos y mejorar la capacidad aeróbica protegiendo las articulaciones.", Lugar: "Piscina Climatizada"},
}

var ReservasActuales = []models.Reserva{}
var contadorReservas = 0

func CrearReserva(socioID int, actividadID int, fecha string) bool {
	// Verificar aforo simple (como simulación)
	var count int
	for _, r := range ReservasActuales {
		if r.ActividadID == actividadID && r.FechaAsist == fecha {
			count++
		}
	}
	
	for _, c := range ClasesLista {
		if c.ID == actividadID {
			if count >= c.Aforo {
				return false // Aforo completo
			}
			break
		}
	}

	contadorReservas++
	nuevaReserva := models.Reserva{
		ID:          contadorReservas,
		SocioID:     socioID,
		ActividadID: actividadID,
		FechaAsist:  fecha,
	}
	ReservasActuales = append(ReservasActuales, nuevaReserva)
	return true
}

func ObtenerReservasDeSocio(socioID int) []models.Reserva {
	var result []models.Reserva
	for _, r := range ReservasActuales {
		if r.SocioID == socioID {
			result = append(result, r)
		}
	}
	return result
}

func EliminarReserva(reservaID int) bool {
	for i, r := range ReservasActuales {
		if r.ID == reservaID {
			ReservasActuales = append(ReservasActuales[:i], ReservasActuales[i+1:]...)
			return true
		}
	}
	return false
}

