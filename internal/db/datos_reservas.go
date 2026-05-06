package db

import (
	"log"
	"pec2/internal/models"
)

func ObtenerClases() []models.Clases {
	var clases []models.Clases
	rows, err := DB.Query("SELECT id, nombre, entrenador, aforo, horario, descripcion, lugar FROM clases")
	if err != nil {
		log.Println("Error obteniendo clases:", err)
		return clases
	}
	defer rows.Close()

	for rows.Next() {
		var c models.Clases
		if err := rows.Scan(&c.ID, &c.NombreClase, &c.Entrenador, &c.Aforo, &c.Horario, &c.Descripcion, &c.Lugar); err == nil {
			clases = append(clases, c)
		}
	}
	return clases
}

func CrearReserva(socioID int, actividadID int, fecha string) bool {
	// Check if already reserved
	var countSocio int
	DB.QueryRow("SELECT COUNT(*) FROM reservas WHERE socio_id = ? AND actividad_id = ? AND fecha_asist = ?", socioID, actividadID, fecha).Scan(&countSocio)
	if countSocio > 0 {
		return false // Usuario ya tiene esta reserva
	}

	// Check aforo
	var aforoMax, aforoActual int
	err := DB.QueryRow("SELECT aforo FROM clases WHERE id = ?", actividadID).Scan(&aforoMax)
	if err != nil {
		return false // Clase no encontrada
	}

	DB.QueryRow("SELECT COUNT(*) FROM reservas WHERE actividad_id = ? AND fecha_asist = ?", actividadID, fecha).Scan(&aforoActual)
	if aforoActual >= aforoMax {
		return false // Aforo completo
	}

	_, err = DB.Exec("INSERT INTO reservas (socio_id, actividad_id, fecha_asist) VALUES (?, ?, ?)", socioID, actividadID, fecha)
	return err == nil
}

func ObtenerReservasDeSocio(socioID int) []models.Reserva {
	var result []models.Reserva
	rows, err := DB.Query("SELECT id, socio_id, actividad_id, fecha_asist FROM reservas WHERE socio_id = ?", socioID)
	if err != nil {
		return result
	}
	defer rows.Close()

	for rows.Next() {
		var r models.Reserva
		if err := rows.Scan(&r.ID, &r.SocioID, &r.ActividadID, &r.FechaAsist); err == nil {
			result = append(result, r)
		}
	}
	return result
}

func EliminarReserva(reservaID int, socioID int) bool {
	res, err := DB.Exec("DELETE FROM reservas WHERE id = ? AND socio_id = ?", reservaID, socioID)
	if err != nil {
		return false
	}
	rows, _ := res.RowsAffected()
	return rows > 0
}

