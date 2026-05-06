package models

type Usuario struct {
    ID              int    `json:"id"`
    Nombre          string `json:"nombre" validate:"required"`
    Apellidos       string `json:"apellidos" validate:"required"`
    FechaNacimiento string `json:"fecha_nacimiento" validate:"required,datetime=2006-01-02"`
    Direccion       string `json:"direccion" validate:"required"`
    Telefono        string `json:"telefono" validate:"required"`
    Correo          string `json:"correo" validate:"required,email"`
    Documento       string `json:"documento" validate:"required"`
    MetodoPago      string `json:"metodo_pago" validate:"required,oneof=tarjeta domiciliacion"`
    NumeroPago      string `json:"numero_pago" validate:"required"`
    Password        string `json:"password" validate:"required,min=8"`
}
