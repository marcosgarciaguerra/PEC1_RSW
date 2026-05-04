package models

type Usuario struct {
	ID              int    `json:"id"`

	Nombre          string `json:"nombre"`
	Apellidos       string `json:"apellidos"`
	FechaNacimiento string `json:"fecha_nacimiento"`
	Direccion       string `json:"direccion"`
	Telefono        string `json:"telefono"`
	Correo          string `json:"correo"`
	Documento       string `json:"documento"`
	MetodoPago      string `json:"metodo_pago"`
	NumeroPago      string `json:"numero_pago"`
	Password        string `json:"password"`
}
