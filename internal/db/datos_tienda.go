package db

import "pec2/internal/models"

var ArticulosTienda = []models.Articulo{
	{ID: 1, Nombre: "100% Platinum Whey Isolate", Categoria: "sup", Precio: 59.99, Imagen: "/img/shop/protein.jpg", Etiqueta: "Top Ventas"},
	{ID: 2, Nombre: "Pre-Entreno Extremo", Categoria: "sup", Precio: 34.99, Imagen: "/img/shop/prenetreno.avif", Etiqueta: "Nuevo"},
	{ID: 3, Nombre: "Creatina Monohidrato Creapure", Categoria: "sup", Precio: 24.99, Imagen: "/img/shop/creatina.avif"},
	{ID: 4, Nombre: "Camiseta Stringer Platinum", Categoria: "rop", Precio: 22.00, Imagen: "/img/shop/tirantes.jpg", Etiqueta: "Atleta"},
	{ID: 5, Nombre: "Chaqueta Cortavientos", Categoria: "rop", Precio: 45.00, Imagen: "/img/shop/chaqueta.webp", Etiqueta: "Nuevo"},
	{ID: 6, Nombre: "Shaker Mezclador 700ml", Categoria: "acc", Precio: 12.00, Imagen: "/img/shop/botella.jpg", Etiqueta: "Básico"},
	{ID: 7, Nombre: "Cinturón Levantamiento", Categoria: "acc", Precio: 45.00, Imagen: "/img/shop/cinturon.jfif", Etiqueta: "Pro"},
	{ID: 8, Nombre: "Straps de Elevación Premium", Categoria: "acc", Precio: 15.00, Imagen: "/img/shop/straps.jpg", Etiqueta: "Esencial"},
	{ID: 9, Nombre: "Camiseta F*ck Your Idols", Categoria: "rop", Precio: 65.00, Imagen: "/img/shop/fuckidols.jfif"},
}

func ObtenerArticuloPorID(id int) (models.Articulo, bool) {
	for _, articulo := range ArticulosTienda {
		if articulo.ID == id {
			return articulo, true
		}
	}
	return models.Articulo{}, false
}
