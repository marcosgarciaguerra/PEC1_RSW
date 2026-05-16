package db

import (
	"log"
	"pec2/internal/models"
)

// Inicializa datos básicos si no los hay
func seedArticulos() {
	var count int
	DB.QueryRow("SELECT COUNT(*) FROM articulos").Scan(&count)
	if count == 0 {
		_, _ = DB.Exec(`INSERT INTO articulos (nombre, categoria, precio, imagen, etiqueta) VALUES 
		('100% Platinum Whey Isolate', 'sup', 59.99, '/img/shop/protein.jpg', 'Top Ventas'),
		('Pre-Entreno Extremo', 'sup', 34.99, '/img/shop/prenetreno.avif', 'Nuevo'),
		('Creatina Monohidrato Creapure', 'sup', 24.99, '/img/shop/creatina.avif', ''),
		('Camiseta Stringer Platinum', 'rop', 22.00, '/img/shop/tirantes.jpg', 'Atleta'),
		('Chaqueta Cortavientos', 'rop', 45.00, '/img/shop/chaqueta.webp', 'Nuevo'),
		('Shaker Mezclador 700ml', 'acc', 12.00, '/img/shop/botella.jpg', 'Básico'),
		('Cinturón Levantamiento', 'acc', 45.00, '/img/shop/cinturon.jfif', 'Pro'),
		('Straps de Elevación Premium', 'acc', 15.00, '/img/shop/straps.jpg', 'Esencial'),
		('Camiseta F*ck Your Idols', 'rop', 65.00, '/img/shop/fuckidols.jfif', '')`)
	}
}

func ObtenerArticulosTienda() []models.Articulo {
	// Llamamos a seed aquí por seguridad si aún no ha insertado
	seedArticulos()
	var articulos []models.Articulo
	rows, err := DB.Query("SELECT id, nombre, categoria, precio, imagen, etiqueta FROM articulos")
	if err != nil {
		log.Println("Error obteniendo artículos:", err)
		return articulos
	}
	defer rows.Close()

	for rows.Next() {
		var a models.Articulo
		if err := rows.Scan(&a.ID, &a.Nombre, &a.Categoria, &a.Precio, &a.Imagen, &a.Etiqueta); err == nil {
			articulos = append(articulos, a)
		}
	}
	return articulos
}

func ObtenerArticuloPorID(id int) (models.Articulo, bool) {
	var a models.Articulo
	err := DB.QueryRow("SELECT id, nombre, categoria, precio, imagen, etiqueta FROM articulos WHERE id = ?", id).
		Scan(&a.ID, &a.Nombre, &a.Categoria, &a.Precio, &a.Imagen, &a.Etiqueta)
	if err != nil {
		return a, false
	}
	return a, true
}

func CrearArticulo(a models.Articulo) (int, error) {
	res, err := DB.Exec("INSERT INTO articulos (nombre, categoria, precio, imagen, etiqueta) VALUES (?, ?, ?, ?, ?)",
		a.Nombre, a.Categoria, a.Precio, a.Imagen, a.Etiqueta)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func ActualizarArticulo(a models.Articulo) (bool, error) {
	res, err := DB.Exec("UPDATE articulos SET nombre=?, categoria=?, precio=?, imagen=?, etiqueta=? WHERE id=?",
		a.Nombre, a.Categoria, a.Precio, a.Imagen, a.Etiqueta, a.ID)
	if err != nil {
		return false, err
	}
	n, err := res.RowsAffected()
	if err != nil {
		return false, err
	}
	return n > 0, nil
}

func EliminarArticulo(id int) (bool, error) {
	res, err := DB.Exec("DELETE FROM articulos WHERE id = ?", id)
	if err != nil {
		return false, err
	}
	n, err := res.RowsAffected()
	if err != nil {
		return false, err
	}
	return n > 0, nil
}
