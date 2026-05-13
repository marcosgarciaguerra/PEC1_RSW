package handlers

import (
	"errors"
	"log"
	"net/http"

	"pec2/internal/models"
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
	err = RegistrarUsuario(usuario, repeatPassword)
	if err != nil {
		var errorMsg string
		switch {
		case errors.Is(err, ErrPasswordsNoCoinciden):
			errorMsg = "Las contraseñas ingresadas no coinciden."
		case errors.Is(err, ErrDatosRegistroInvalidos):
			errorMsg = "Faltan datos obligatorios del formulario."
		case errors.Is(err, validation.ErrEmailInvalido):
			errorMsg = "Datos de formulario inválidos: El correo electrónico es inválido."
		case errors.Is(err, validation.ErrDocumentoInvalido):
			errorMsg = "Datos de formulario inválidos: El DNI o Pasaporte es inválido."
		case errors.Is(err, validation.ErrTelefonoInvalido):
			errorMsg = "Datos de formulario inválidos: El teléfono es inválido."
		case errors.Is(err, validation.ErrPasswordDebil):
			errorMsg = "Datos de formulario inválidos: La contraseña debe tener al menos 8 caracteres, mayúscula, minúscula y número."
		case errors.Is(err, validation.ErrPlanInvalido):
			errorMsg = "Datos de formulario inválidos: El plan seleccionado es inválido."
		case errors.Is(err, validation.ErrMetodoPagoInvalido):
			errorMsg = "Datos de formulario inválidos: El método de pago es inválido."
		case errors.Is(err, validation.ErrNumeroPagoInvalido):
			errorMsg = "Datos de formulario inválidos: El número de tarjeta o IBAN es inválido."
		case errors.Is(err, validation.ErrDireccionInvalida):
			errorMsg = "Datos de formulario inválidos: La dirección es inválida."
		default:
			log.Println("Error guardando usuario (posible DNI o Correo duplicado):", err)
			errorMsg = "El correo electrónico o el DNI/Pasaporte ya se encuentra registrado en el sistema."
			RenderTemplate(w, "apuntate", map[string]interface{}{
				"ErrorMsg": errorMsg,
			})
			return
		}
		
		RenderTemplate(w, "apuntate", map[string]interface{}{
			"ErrorMsg": errorMsg,
		})
		return
	}

	log.Printf("Nuevo usuario registrado: %s %s", usuario.Nombre, usuario.Apellidos)
	http.Redirect(w, r, "/?registro=exito", http.StatusSeeOther)
}
