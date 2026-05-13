package handlers

import (
	"errors"
	"strings"

	"pec2/internal/db"
	"pec2/internal/models"
	"pec2/internal/validation"
)

var ErrResenaInvalida = errors.New("resena invalida")

func GuardarResena(autor string, puntuacion int, texto string) error {
	autor = strings.TrimSpace(autor)
	texto = strings.TrimSpace(texto)
	if autor == "" {
		return ErrResenaInvalida
	}
	if err := validation.ValidarPuntuacion(puntuacion); err != nil {
		return err
	}
	if err := validation.ValidarTextoResena(texto); err != nil {
		return err
	}
	return db.GuardarResena(models.Resena{
		Autor:      autor,
		Puntuacion: puntuacion,
		Texto:      texto,
	})
}
