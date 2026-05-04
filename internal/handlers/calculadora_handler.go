package handlers

import (
	"net/http"
	"pec2/internal/models"
	"strconv"
)

func CalculadoraHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		MostrarFormulario(w, r)
	} else if r.Method == http.MethodPost {
		ProcesarCalculadora(w, r)
	} else {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}

func MostrarFormulario(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "calculadora", nil)
}

func ProcesarCalculadora(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	alturaStr := r.FormValue("altura")
	pesoStr := r.FormValue("peso")
	grasaStr := r.FormValue("grasa")

	altura, _ := strconv.ParseFloat(alturaStr, 64)
	peso, _ := strconv.ParseFloat(pesoStr, 64)
	grasa, _ := strconv.ParseFloat(grasaStr, 64)

	datos := models.DatosFFMI{
		Altura:              altura,
		Peso:                peso,
		IndiceGrasaCorporal: grasa,
	}

	var resultado float64
	var errorMsg string

	if datos.ValidarDatos() {
		masaMagra := datos.Peso * (1.0 - (datos.IndiceGrasaCorporal / 100.0))
		ffmi := masaMagra / (datos.Altura * datos.Altura)
		resultado = ffmi + 6.1*(1.8-datos.Altura)
		datos.Resultado = resultado
	} else {
		errorMsg = "Por favor, introduce datos válidos."
	}

	RenderTemplate(w, "calculadora", map[string]interface{}{
		"Resultado": resultado,
		"Datos":     datos,
		"ErrorMsg":  errorMsg,
	})
}

