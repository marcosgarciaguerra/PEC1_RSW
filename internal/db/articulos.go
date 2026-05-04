package db

import (
	"log"
	"pec2/internal/models"
)

func CrearTablaArticulos() {
	query := `
	CREATE TABLE IF NOT EXISTS articulos (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		nombre TEXT NOT NULL,
		categoria TEXT NOT NULL,
		precio REAL NOT NULL,
		imagen TEXT NOT NULL,
		etiqueta TEXT
	);`

	_, err := DB.Exec(query)
	if err != nil {
		log.Fatal("Error creando tabla articulos: ", err)
	}

	// Insertar datos iniciales si está vacía
	var count int
	DB.QueryRow("SELECT COUNT(*) FROM articulos").Scan(&count)
	if count == 0 {
		_, _ = DB.Exec(`INSERT INTO articulos (nombre, categoria, precio, imagen, etiqueta) VALUES 
		('100% Platinum Whey Isolate', 'Suplemento', 59.99, '/img/shop/protein.jpg', 'Top Ventas'),
		('Pre-Entreno Extremo', 'Suplemento', 34.99, '/img/shop/prenetreno.avif', 'Nuevo'),
		('Creatina Monohidrato Creapure', 'Suplemento', 24.99, '/img/shop/creatina.avif', ''),
		('Camiseta Stringer Platinum', 'Ropa', 22.00, '/img/shop/tirantes.jpg', 'Atleta'),
		('Chaqueta Cortavientos', 'Ropa', 45.00, '/img/shop/chaqueta.webp', 'Nuevo'),
		('Shaker Mezclador 700ml', 'Accesorios', 12.00, '/img/shop/botella.jpg', 'Básico')
		`)
	}
}

func ObtenerArticulos() ([]models.Articulo, error) {
	var articulos []models.Articulo
	rows, err := DB.Query("SELECT id, nombre, categoria, precio, imagen, etiqueta FROM articulos ORDER BY id DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var a models.Articulo
		if err := rows.Scan(&a.ID, &a.Nombre, &a.Categoria, &a.Precio, &a.Imagen, &a.Etiqueta); err == nil {
			articulos = append(articulos, a)
		}
	}
	return articulos, nil
}

func ObtenerArticulo(id int) (*models.Articulo, error) {
	var a models.Articulo
	err := DB.QueryRow("SELECT id, nombre, categoria, precio, imagen, etiqueta FROM articulos WHERE id = ?", id).
		Scan(&a.ID, &a.Nombre, &a.Categoria, &a.Precio, &a.Imagen, &a.Etiqueta)
	if err != nil {
		return nil, err
	}
	return &a, nil
}

func CrearArticulo(a models.Articulo) (int, error) {
	stmt, err := DB.Prepare("INSERT INTO articulos (nombre, categoria, precio, imagen, etiqueta) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(a.Nombre, a.Categoria, a.Precio, a.Imagen, a.Etiqueta)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func ActualizarArticulo(id int, a models.Articulo) error {
	stmt, err := DB.Prepare("UPDATE articulos SET nombre=?, categoria=?, precio=?, imagen=?, etiqueta=? WHERE id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(a.Nombre, a.Categoria, a.Precio, a.Imagen, a.Etiqueta, id)
	return err
}

func EliminarArticulo(id int) error {
	stmt, err := DB.Prepare("DELETE FROM articulos WHERE id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	return err
}
