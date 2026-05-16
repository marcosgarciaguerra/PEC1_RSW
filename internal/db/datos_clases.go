package db

import (
	"log"
	"pec2/internal/models"
)

func ObtenerClasePorID(id int) (*models.Clases, error) {
	var c models.Clases
	err := DB.QueryRow("SELECT id, nombre, entrenador, aforo, horario, descripcion, lugar FROM clases WHERE id = ?", id).
		Scan(&c.ID, &c.NombreClase, &c.Entrenador, &c.Aforo, &c.Horario, &c.Descripcion, &c.Lugar)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func CrearClase(c models.Clases) (int, error) {
	res, err := DB.Exec("INSERT INTO clases (nombre, entrenador, aforo, horario, descripcion, lugar) VALUES (?, ?, ?, ?, ?, ?)",
		c.NombreClase, c.Entrenador, c.Aforo, c.Horario, c.Descripcion, c.Lugar)
	if err != nil {
		log.Println("Error creando clase:", err)
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func ActualizarClase(c models.Clases) (bool, error) {
	res, err := DB.Exec("UPDATE clases SET nombre=?, entrenador=?, aforo=?, horario=?, descripcion=?, lugar=? WHERE id=?",
		c.NombreClase, c.Entrenador, c.Aforo, c.Horario, c.Descripcion, c.Lugar, c.ID)
	if err != nil {
		log.Println("Error actualizando clase:", err)
		return false, err
	}
	n, err := res.RowsAffected()
	if err != nil {
		return false, err
	}
	return n > 0, nil
}

func EliminarClase(id int) (bool, error) {
	res, err := DB.Exec("DELETE FROM clases WHERE id = ?", id)
	if err != nil {
		log.Println("Error eliminando clase:", err)
		return false, err
	}
	n, err := res.RowsAffected()
	if err != nil {
		return false, err
	}
	return n > 0, nil
}
