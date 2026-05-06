package models

type Articulo struct {
	ID        int     `json:"id"`
	Nombre    string  `json:"nombre"`
	Categoria string  `json:"categoria"` // Ejemplo: "Suplemento", "Ropa"
	Precio    float64 `json:"precio"`
	Imagen    string  `json:"imagen"`
	Etiqueta  string  `json:"etiqueta"` // Ejemplo: "Top Ventas", "Nuevo", "Pro"
}
