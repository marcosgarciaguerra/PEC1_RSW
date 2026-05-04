package models

type Servicio struct {
	ID               int    `json:"id"`
	Nombre           string `json:"nombre"`
	Icono            string `json:"icono"`
	DescripcionBreve string `json:"descripcion_breve"`
	DescripcionLarga string `json:"descripcion_larga"`
	Beneficios       string `json:"beneficios"`
}
