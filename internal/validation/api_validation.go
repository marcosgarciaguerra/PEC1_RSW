package validation

import (
	"errors"
	"pec2/internal/models"
	"strings"
)

var (
	ErrNombreRequerido     = errors.New("el nombre es obligatorio")
	ErrPrecioInvalido      = errors.New("el precio debe ser mayor o igual a cero")
	ErrCategoriaInvalida   = errors.New("categoría inválida")
	ErrAforoInvalido       = errors.New("el aforo debe ser mayor que cero")
	ErrCampoRequerido      = errors.New("campo obligatorio vacío")
)

func ValidarArticulo(a models.Articulo) error {
	if strings.TrimSpace(a.Nombre) == "" {
		return ErrNombreRequerido
	}
	if a.Precio < 0 {
		return ErrPrecioInvalido
	}
	switch strings.TrimSpace(a.Categoria) {
	case "sup", "rop", "acc":
	default:
		return ErrCategoriaInvalida
	}
	return nil
}

func ValidarClase(c models.Clases) error {
	if strings.TrimSpace(c.NombreClase) == "" ||
		strings.TrimSpace(c.Entrenador) == "" ||
		strings.TrimSpace(c.Horario) == "" ||
		strings.TrimSpace(c.Lugar) == "" {
		return ErrCampoRequerido
	}
	if c.Aforo <= 0 {
		return ErrAforoInvalido
	}
	return nil
}

func ValidarMaquina(m models.Maquina) error {
	if strings.TrimSpace(m.Nombre) == "" ||
		strings.TrimSpace(m.Marca) == "" ||
		strings.TrimSpace(m.Zona) == "" {
		return ErrCampoRequerido
	}
	return nil
}
