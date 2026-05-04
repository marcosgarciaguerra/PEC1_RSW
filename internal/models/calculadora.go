package models

type DatosFFMI struct {
	Altura              float64
	Peso                float64
	IndiceGrasaCorporal float64
	Resultado           float64
}


func (d *DatosFFMI) ValidarDatos() bool {
	if d.Altura <= 0 || d.Peso <= 0 || d.IndiceGrasaCorporal <= 0 || d.IndiceGrasaCorporal >= 100 {
		return false
	}
	// Asumimos que la altura se da en centímetros o metros. Si es > 3, asumimos cm y la convertimos a metros.
	if d.Altura > 3.0 {
		d.Altura = d.Altura / 100.0
	}
	return true
}
