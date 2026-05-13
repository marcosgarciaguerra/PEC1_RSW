package db

import (
	"log"
	"pec2/internal/models"
)

func ListarMaquinas() []models.Maquina {
	var maquinarias []models.Maquina
	rows, err := DB.Query("SELECT id, nombre, marca, zona, descripcion, beneficios, imagen FROM maquinarias")
	if err != nil {
		log.Println("Error obteniendo maquinarias:", err)
		return maquinarias
	}
	defer rows.Close()

	for rows.Next() {
		var m models.Maquina
		if err := rows.Scan(&m.ID, &m.Nombre, &m.Marca, &m.Zona, &m.Descripcion, &m.Beneficios, &m.Imagen); err == nil {
			maquinarias = append(maquinarias, m)
		} else {
			log.Println("Error escaneando maquinaria:", err)
		}
	}
	return maquinarias
}

func BuscarMaquinaPorID(id int) (models.Maquina, bool) {
	var m models.Maquina
	err := DB.QueryRow("SELECT id, nombre, marca, zona, descripcion, beneficios, imagen FROM maquinarias WHERE id = ?", id).
		Scan(&m.ID, &m.Nombre, &m.Marca, &m.Zona, &m.Descripcion, &m.Beneficios, &m.Imagen)
	if err != nil {
		return m, false
	}
	return m, true
}

func CrearMaquina(m models.Maquina) (int64, error) {
	res, err := DB.Exec("INSERT INTO maquinarias (nombre, marca, zona, descripcion, beneficios, imagen) VALUES (?, ?, ?, ?, ?, ?)",
		m.Nombre, m.Marca, m.Zona, m.Descripcion, m.Beneficios, m.Imagen)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func ActualizarMaquina(m models.Maquina) error {
	_, err := DB.Exec("UPDATE maquinarias SET nombre=?, marca=?, zona=?, descripcion=?, beneficios=?, imagen=? WHERE id=?",
		m.Nombre, m.Marca, m.Zona, m.Descripcion, m.Beneficios, m.Imagen, m.ID)
	return err
}

func EliminarMaquina(id int) error {
	_, err := DB.Exec("DELETE FROM maquinarias WHERE id=?", id)
	return err
}
