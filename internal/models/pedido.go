package models

type Pedido struct {
	ID             int          `json:"id"`
	UsuarioID      int          `json:"usuario_id"`
	Fecha          string       `json:"fecha"`
	DireccionEnvio string       `json:"direccion_envio"`
	MetodoPago     string       `json:"metodo_pago"`
	Total          float64      `json:"total"`
	Estado         string       `json:"estado"`
	Items          []PedidoItem `json:"items"`
}

type PedidoItem struct {
	ID             int     `json:"id"`
	PedidoID       int     `json:"pedido_id"`
	ArticuloID     int     `json:"articulo_id"`
	Nombre         string  `json:"nombre"`
	PrecioUnitario float64 `json:"precio_unitario"`
	Cantidad       int     `json:"cantidad"`
	Subtotal       float64 `json:"subtotal"`
}
