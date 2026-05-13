package handlers

import (
	"errors"
	"strings"

	"pec2/internal/db"
	"pec2/internal/models"
	"pec2/internal/validation"
)

var (
	ErrPasswordsNoCoinciden = errors.New("las contrasenas no coinciden")
	ErrDatosRegistroInvalidos = errors.New("faltan datos obligatorios del registro")
)

func RegistrarUsuario(usuario models.Usuario, repetirPassword string) error {
	usuario.Nombre = strings.TrimSpace(usuario.Nombre)
	usuario.Apellidos = strings.TrimSpace(usuario.Apellidos)
	usuario.Correo = strings.ToLower(strings.TrimSpace(usuario.Correo))
	usuario.Documento = strings.ToUpper(strings.TrimSpace(usuario.Documento))
	usuario.Telefono = strings.TrimSpace(usuario.Telefono)
	usuario.Direccion = strings.TrimSpace(usuario.Direccion)
	usuario.Password = strings.TrimSpace(usuario.Password)

	if usuario.Nombre == "" || usuario.Correo == "" || usuario.Documento == "" || usuario.Password == "" {
		return ErrDatosRegistroInvalidos
	}

	if usuario.Password != repetirPassword {
		return ErrPasswordsNoCoinciden
	}

	if strings.TrimSpace(usuario.Plan) == "" {
		usuario.Plan = "basico"
	}

	if err := validation.ValidarEmail(usuario.Correo); err != nil {
		return err
	}
	if err := validation.ValidarDocumento(usuario.Documento); err != nil {
		return err
	}
	if err := validation.ValidarTelefono(usuario.Telefono); err != nil {
		return err
	}
	if err := validation.ValidarDireccion(usuario.Direccion); err != nil {
		return err
	}
	if err := validation.ValidarPlan(usuario.Plan); err != nil {
		return err
	}
	if err := validation.ValidarMetodoPago(usuario.MetodoPago); err != nil {
		return err
	}
	if err := validation.ValidarPasswordFuerte(usuario.Password); err != nil {
		return err
	}

	maskedPayment, err := validation.EnmascararNumeroPago(usuario.MetodoPago, usuario.NumeroPago)
	if err != nil {
		return err
	}
	usuario.NumeroPago = maskedPayment

	hash, err := HashPassword(usuario.Password)
	if err != nil {
		return err
	}
	usuario.Password = hash

	return db.GuardarUsuario(usuario)
}
