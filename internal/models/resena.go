package models

type Resena struct {
	Autor      string `json:"autor"`
	Puntuacion int    `json:"puntuacion"`
	Texto      string `json:"texto"`
}
