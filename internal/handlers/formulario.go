package handlers

import (
	"log"
	"net/http"
	"pec2/internal/db"
	"pec2/internal/models"
	"golang.org/x/crypto/bcrypt"
	"github.com/go-playground/validator/v10"
)

func RegistroHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/apuntate", http.StatusSeeOther)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error al procesar el formulario", http.StatusBadRequest)
		return
	}

	usuario := models.Usuario{
		Nombre:          r.FormValue("nombre"),
		Apellidos:       r.FormValue("apellidos"),
		FechaNacimiento: r.FormValue("fecha-nacimiento"),
		Direccion:       r.FormValue("direccion"),
		Telefono:        r.FormValue("telefono"),
		Correo:          r.FormValue("correo"),
		Documento:       r.FormValue("documento"),
		MetodoPago:      r.FormValue("metodo-pago"),
		NumeroPago:      r.FormValue("numero-pago"),
		Password:        r.FormValue("password"),
	}

	repeatPassword := r.FormValue("repeat-password")
	if usuario.Password != repeatPassword {
		http.Error(w, "Las contraseñas ingresadas no coinciden", http.StatusBadRequest)
		return
	}

	// Validate fields using go-playground/validator
	validate := validator.New()
	if err := validate.Struct(usuario); err != nil {
		log.Println("Validation error:", err)
		http.Error(w, "Datos de registro no válidos o faltan campos obligatorios", http.StatusBadRequest)
		return
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(usuario.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error hashing password:", err)
		http.Error(w, "Error interno del servidor", http.StatusInternalServerError)
		return
	}
	usuario.Password = string(hashed)

	hashedDNI, err := bcrypt.GenerateFromPassword([]byte(usuario.Documento), bcrypt.DefaultCost)
	if err == nil {
		usuario.Documento = string(hashedDNI)
	}

	hashedPago, err := bcrypt.GenerateFromPassword([]byte(usuario.NumeroPago), bcrypt.DefaultCost)
	if err == nil {
		usuario.NumeroPago = string(hashedPago)
	}

	err = db.GuardarUsuario(usuario)
	if err != nil {
		log.Println("Error guardando usuario (posible DNI o Correo duplicado):", err)
		http.Error(w, "El correo electrónico o el DNI/Pasaporte ya se encuentra registrado en el sistema.", http.StatusConflict)
		return
	}

	log.Printf("Nuevo usuario registrado: %s %s", usuario.Nombre, usuario.Apellidos)
	http.Redirect(w, r, "/?registro=exito", http.StatusSeeOther)
}
