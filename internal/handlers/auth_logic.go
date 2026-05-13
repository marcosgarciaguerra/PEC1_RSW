package handlers

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"

	"pec2/internal/db"
	"pec2/internal/models"
	"pec2/internal/validation"
)

var (
	ErrCredencialesInvalidas = errors.New("credenciales invalidas")
	ErrCuentaInactiva        = errors.New("cuenta inactiva")
	ErrSesionInvalida        = errors.New("sesion invalida")
)

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func normalizeEmail(email string) string {
	return strings.ToLower(strings.TrimSpace(email))
}

func Authenticate(correo string, contrasena string) (*models.Socio, error) {
	correo = normalizeEmail(correo)
	contrasena = strings.TrimSpace(contrasena)
	if err := validation.ValidarEmail(correo); err != nil || contrasena == "" {
		return nil, ErrCredencialesInvalidas
	}

	socio := db.ObtenerSocioPorCorreo(correo)
	if socio == nil {
		return nil, ErrCredencialesInvalidas
	}
	if !socio.SuscripcionActiva {
		return nil, ErrCuentaInactiva
	}

	if strings.HasPrefix(socio.ContrasenaHash, "$2") {
		if err := bcrypt.CompareHashAndPassword([]byte(socio.ContrasenaHash), []byte(contrasena)); err != nil {
			return nil, ErrCredencialesInvalidas
		}
	} else {
		if socio.ContrasenaHash != contrasena {
			return nil, ErrCredencialesInvalidas
		}
		hash, err := HashPassword(contrasena)
		if err == nil {
			_ = db.ActualizarPasswordHash(socio.ID, hash)
		}
	}

	return socio, nil
}

func NewSessionToken() (rawToken string, tokenHash string, err error) {
	b := make([]byte, 32)
	if _, err = rand.Read(b); err != nil {
		return "", "", err
	}
	rawToken = base64.RawURLEncoding.EncodeToString(b)
	sum := sha256.Sum256([]byte(rawToken))
	tokenHash = hex.EncodeToString(sum[:])
	return rawToken, tokenHash, nil
}

func CreateUserSession(userID int, ttl time.Duration) (string, error) {
	rawToken, tokenHash, err := NewSessionToken()
	if err != nil {
		return "", err
	}
	if err := db.CrearSesion(userID, tokenHash, time.Now().Add(ttl)); err != nil {
		return "", err
	}
	return rawToken, nil
}

func GetUserFromSessionToken(rawToken string) (*models.Usuario, error) {
	rawToken = strings.TrimSpace(rawToken)
	if rawToken == "" {
		return nil, ErrSesionInvalida
	}
	sum := sha256.Sum256([]byte(rawToken))
	tokenHash := hex.EncodeToString(sum[:])
	usuario := db.ObtenerUsuarioPorTokenSesion(tokenHash)
	if usuario == nil {
		return nil, ErrSesionInvalida
	}
	return usuario, nil
}

func GetSocioFromSessionToken(rawToken string) (*models.Socio, error) {
	rawToken = strings.TrimSpace(rawToken)
	if rawToken == "" {
		return nil, ErrSesionInvalida
	}
	sum := sha256.Sum256([]byte(rawToken))
	tokenHash := hex.EncodeToString(sum[:])
	socio := db.ObtenerSocioPorTokenSesion(tokenHash)
	if socio == nil {
		return nil, ErrSesionInvalida
	}
	return socio, nil
}

func RevokeSession(rawToken string) {
	rawToken = strings.TrimSpace(rawToken)
	if rawToken == "" {
		return
	}
	sum := sha256.Sum256([]byte(rawToken))
	tokenHash := hex.EncodeToString(sum[:])
	_ = db.RevocarSesion(tokenHash)
}

func SanitizarReturnPath(returnPath string) string {
	path := strings.TrimSpace(returnPath)
	if path == "" {
		return ""
	}
	if strings.HasPrefix(path, "//") || strings.HasPrefix(strings.ToLower(path), "http://") || strings.HasPrefix(strings.ToLower(path), "https://") {
		return ""
	}
	if !strings.HasPrefix(path, "/") {
		return ""
	}
	return path
}
