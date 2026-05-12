package services

import (
	"errors"
	"strings"

	"pec2/internal/db"
	"pec2/internal/models"
)

var (
	ErrPasswordsNoCoinciden = errors.New("las contrasenas no coinciden")
	ErrDatosRegistroInvalidos = errors.New("faltan datos obligatorios del registro")
)

func RegistrarUsuario(usuario models.Usuario, repetirPassword string) error {
	if strings.TrimSpace(usuario.Nombre) == "" ||
		strings.TrimSpace(usuario.Correo) == "" ||
		strings.TrimSpace(usuario.Documento) == "" ||
		strings.TrimSpace(usuario.Password) == "" {
		return ErrDatosRegistroInvalidos
	}

	if usuario.Password != repetirPassword {
		return ErrPasswordsNoCoinciden
	}

	if strings.TrimSpace(usuario.Plan) == "" {
		usuario.Plan = "basico"
	}

	return db.GuardarUsuario(usuario)
}
