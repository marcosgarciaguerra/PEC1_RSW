package handlers

import (
	"log"
	"net/http"
	"pec2/internal/db"
	"pec2/internal/models"
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

	if usuario.Nombre == "" || usuario.Apellidos == "" || usuario.Correo == "" || usuario.Password == "" {
		http.Error(w, "Faltan campos obligatorios", http.StatusBadRequest)
		return
	}
	if len(usuario.Password) < 6 {
		http.Error(w, "La contraseña debe tener al menos 6 caracteres", http.StatusBadRequest)
		return
	}

	err = db.GuardarUsuario(usuario)
	if err != nil {
		log.Println("Error guardando usuario:", err)
		http.Error(w, "Error interno del servidor", http.StatusInternalServerError)
		return
	}

	log.Printf("Nuevo usuario registrado: %s %s", usuario.Nombre, usuario.Apellidos)
	http.Redirect(w, r, "/?registro=exito", http.StatusSeeOther)
}
