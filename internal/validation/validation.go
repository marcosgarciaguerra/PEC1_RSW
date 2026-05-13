package validation

import (
	"errors"
	"net/mail"
	"regexp"
	"strings"
	"unicode"
)

var (
	ErrEmailInvalido      = errors.New("email invalido")
	ErrDocumentoInvalido  = errors.New("documento invalido")
	ErrTelefonoInvalido   = errors.New("telefono invalido")
	ErrPasswordDebil      = errors.New("password debil")
	ErrPlanInvalido       = errors.New("plan invalido")
	ErrMetodoPagoInvalido = errors.New("metodo de pago invalido")
	ErrNumeroPagoInvalido = errors.New("numero de pago invalido")
	ErrDireccionInvalida  = errors.New("direccion invalida")
	ErrPuntuacionInvalida = errors.New("puntuacion invalida")
	ErrTextoResenaInvalido = errors.New("texto de resena invalido")
)

var (
	dniRegex      = regexp.MustCompile(`^[0-9]{8}[A-Za-z]$`)
	pasaporteRegex = regexp.MustCompile(`^[A-Za-z0-9]{6,15}$`)
	telefonoRegex = regexp.MustCompile(`^[0-9]{9,15}$`)
	tarjetaRegex  = regexp.MustCompile(`^[0-9]{12,19}$`)
	ibanRegex     = regexp.MustCompile(`^[A-Za-z]{2}[0-9A-Za-z]{13,32}$`)
)

func ValidarEmail(email string) error {
	email = strings.TrimSpace(email)
	if email == "" {
		return ErrEmailInvalido
	}
	_, err := mail.ParseAddress(email)
	if err != nil {
		return ErrEmailInvalido
	}
	return nil
}

func ValidarDocumento(documento string) error {
	doc := strings.ToUpper(strings.TrimSpace(documento))
	if dniRegex.MatchString(doc) || pasaporteRegex.MatchString(doc) {
		return nil
	}
	return ErrDocumentoInvalido
}

func ValidarTelefono(telefono string) error {
	tel := strings.TrimSpace(telefono)
	if telefonoRegex.MatchString(tel) {
		return nil
	}
	return ErrTelefonoInvalido
}

func ValidarPasswordFuerte(password string) error {
	if len(password) < 8 {
		return ErrPasswordDebil
	}
	var hasUpper, hasLower, hasDigit bool
	for _, r := range password {
		switch {
		case unicode.IsUpper(r):
			hasUpper = true
		case unicode.IsLower(r):
			hasLower = true
		case unicode.IsDigit(r):
			hasDigit = true
		}
	}
	if !hasUpper || !hasLower || !hasDigit {
		return ErrPasswordDebil
	}
	return nil
}

func ValidarPlan(plan string) error {
	switch strings.TrimSpace(plan) {
	case "basico", "premium", "vip":
		return nil
	default:
		return ErrPlanInvalido
	}
}

func ValidarMetodoPago(metodo string) error {
	switch strings.TrimSpace(metodo) {
	case "tarjeta", "domiciliacion":
		return nil
	default:
		return ErrMetodoPagoInvalido
	}
}

func ValidarDireccion(direccion string) error {
	dir := strings.TrimSpace(direccion)
	if len(dir) < 5 || len(dir) > 150 {
		return ErrDireccionInvalida
	}
	return nil
}

func ValidarPuntuacion(puntuacion int) error {
	if puntuacion < 1 || puntuacion > 5 {
		return ErrPuntuacionInvalida
	}
	return nil
}

func ValidarTextoResena(texto string) error {
	value := strings.TrimSpace(texto)
	if len(value) < 10 || len(value) > 500 {
		return ErrTextoResenaInvalido
	}
	return nil
}

func EnmascararNumeroPago(metodoPago string, numeroPago string) (string, error) {
	clean := strings.ReplaceAll(strings.TrimSpace(numeroPago), " ", "")
	switch metodoPago {
	case "tarjeta":
		if !tarjetaRegex.MatchString(clean) {
			return "", ErrNumeroPagoInvalido
		}
		last4 := clean[len(clean)-4:]
		return "****" + last4, nil
	case "domiciliacion":
		if !ibanRegex.MatchString(clean) {
			return "", ErrNumeroPagoInvalido
		}
		last4 := clean[len(clean)-4:]
		return "IBAN-****" + strings.ToUpper(last4), nil
	default:
		return "", ErrMetodoPagoInvalido
	}
}
