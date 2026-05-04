package models

type Maquina struct {
	ID          int    `json:"id"`
	Nombre      string `json:"nombre"`
	Marca       string `json:"marca"`
	Zona        string `json:"zona"`
	Descripcion string `json:"descripcion"`
	Beneficios  string `json:"beneficios"`
	Imagen      string `json:"imagen"`
}
