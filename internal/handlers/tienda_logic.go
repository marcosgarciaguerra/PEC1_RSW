package handlers

import (
	"encoding/json"
	"errors"
	"strings"
	"time"

	"pec2/internal/db"
	"pec2/internal/models"
	"pec2/internal/validation"
)

var (
	ErrCarritoVacio       = errors.New("el carrito esta vacio")
	ErrArticuloInvalido   = errors.New("el carrito contiene un articulo invalido")
	ErrDireccionInvalida  = errors.New("la direccion de envio no es valida")
	ErrCantidadInvalida   = errors.New("cantidad de articulo invalida")
)

type itemCarrito struct {
	ID       int `json:"id"`
	Quantity int `json:"quantity"`
}

func CatalogoTienda() []models.Articulo {
	return db.ArticulosTienda
}

func CrearPedido(usuario models.Usuario, direccionEnvio string, carritoJSON string) (*models.Pedido, error) {
	direccionEnvio = strings.TrimSpace(direccionEnvio)
	if err := validation.ValidarDireccion(direccionEnvio); err != nil {
		return nil, ErrDireccionInvalida
	}

	var itemsCarrito []itemCarrito
	if err := json.Unmarshal([]byte(carritoJSON), &itemsCarrito); err != nil {
		return nil, ErrArticuloInvalido
	}
	if len(itemsCarrito) == 0 {
		return nil, ErrCarritoVacio
	}

	pedido := &models.Pedido{
		UsuarioID:      usuario.ID,
		Fecha:          time.Now().Format("2006-01-02 15:04:05"),
		DireccionEnvio: direccionEnvio,
		MetodoPago:     usuario.MetodoPago,
		Estado:         "confirmado",
		Items:          make([]models.PedidoItem, 0, len(itemsCarrito)),
	}

	for _, item := range itemsCarrito {
		if item.Quantity <= 0 || item.Quantity > 20 {
			return nil, ErrCantidadInvalida
		}

		articulo, ok := db.ObtenerArticuloPorID(item.ID)
		if !ok {
			return nil, ErrArticuloInvalido
		}

		subtotal := articulo.Precio * float64(item.Quantity)
		pedido.Total += subtotal
		pedido.Items = append(pedido.Items, models.PedidoItem{
			ArticuloID:     articulo.ID,
			Nombre:         articulo.Nombre,
			PrecioUnitario: articulo.Precio,
			Cantidad:       item.Quantity,
			Subtotal:       subtotal,
		})
	}

	if err := db.GuardarPedido(pedido); err != nil {
		return nil, err
	}

	return pedido, nil
}
