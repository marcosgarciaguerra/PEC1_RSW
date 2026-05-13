package handlers

import (
	"errors"
	"log"
	"net/http"

	"pec2/internal/models"
	"pec2/internal/services"
	"pec2/internal/validation"
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
		Plan:            r.FormValue("plan"),
		MetodoPago:      r.FormValue("metodo-pago"),
		NumeroPago:      r.FormValue("numero-pago"),
		Password:        r.FormValue("password"),
	}

	repeatPassword := r.FormValue("repeat-password")
	err = services.RegistrarUsuario(usuario, repeatPassword)
	if errors.Is(err, services.ErrPasswordsNoCoinciden) {
		http.Error(w, "Las contraseñas ingresadas no coinciden", http.StatusBadRequest)
		return
	}
	if errors.Is(err, services.ErrDatosRegistroInvalidos) {
		http.Error(w, "Faltan datos obligatorios del formulario", http.StatusBadRequest)
		return
	}
	if errors.Is(err, validation.ErrEmailInvalido) ||
		errors.Is(err, validation.ErrDocumentoInvalido) ||
		errors.Is(err, validation.ErrTelefonoInvalido) ||
		errors.Is(err, validation.ErrPasswordDebil) ||
		errors.Is(err, validation.ErrPlanInvalido) ||
		errors.Is(err, validation.ErrMetodoPagoInvalido) ||
		errors.Is(err, validation.ErrNumeroPagoInvalido) ||
		errors.Is(err, validation.ErrDireccionInvalida) {
		http.Error(w, "Datos de formulario inválidos", http.StatusBadRequest)
		return
	}
	if err != nil {
		log.Println("Error guardando usuario (posible DNI o Correo duplicado):", err)
		http.Error(w, "El correo electrónico o el DNI/Pasaporte ya se encuentra registrado en el sistema.", http.StatusConflict)
		return
	}

	log.Printf("Nuevo usuario registrado: %s %s", usuario.Nombre, usuario.Apellidos)
	http.Redirect(w, r, "/?registro=exito", http.StatusSeeOther)
}
