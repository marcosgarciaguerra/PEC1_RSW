package db

import "pec2/internal/models"

var Servicios = []models.Servicio{
	{
		ID:               1,
		Nombre:           "Sala de Máquinas",
		Icono:            "fas fa-dumbbell",
		DescripcionBreve: "Bio-mechanics focus. Equipamiento seleccionado científicamente para aislar cada grupo muscular de forma estricta.",
		DescripcionLarga: "Nuestra sala de musculación cuenta con más de 100 máquinas de marcas premium (Panatta, Hammer Strength, Cybex). Cada zona está optimizada para diferentes grupos musculares, con luz controlada y ambiente enfocado al rendimiento puro. Perfecto tanto para principiantes como para culturistas de élite.",
		Beneficios:       "Equipamiento de máxima calidad, cero tiempos de espera gracias a duplicados de máquinas populares y entorno motivador.",
	},
	{
		ID:               2,
		Nombre:           "Fisioterapia Premium",
		Icono:            "fas fa-user-md",
		DescripcionBreve: "Recuperación profunda, punción seca, liberación miofascial y osteopatía al servicio de tu rendimiento.",
		DescripcionLarga: "Un equipo de fisioterapeutas deportivos a tu disposición. Tratamos lesiones agudas, crónicas y realizamos labores de prevención. Utilizamos técnicas avanzadas de terapia manual, electrolisis percutánea, neuromodulación y trabajo activo.",
		Beneficios:       "Recuperación acelerada, prevención de lesiones y tratamientos personalizados.",
	},
	{
		ID:               3,
		Nombre:           "Nutricionista",
		Icono:            "fas fa-apple-alt",
		DescripcionBreve: "Planes dietéticos personalizados, periodización nutricional y análisis corporal por pliegues y DEXA.",
		DescripcionLarga: "Alcanza tus objetivos más rápido con nuestro servicio de nutrición clínica y deportiva. Incluye valoración antropométrica detallada, diseño de planes de alimentación adaptables y seguimiento continuo para garantizar que tu dieta apoye tu entrenamiento.",
		Beneficios:       "Pérdida de grasa eficiente, aumento de masa muscular limpio y mejora del rendimiento.",
	},
	{
		ID:               4,
		Nombre:           "Sauna Finlandesa",
		Icono:            "fas fa-hot-tub",
		DescripcionBreve: "Relajación térmica para eliminación de toxinas, mejora cardiovascular y recuperación del sistema nervioso.",
		DescripcionLarga: "Termina tu sesión de entrenamiento pesado relajándote en nuestras saunas finlandesas de madera de pino a 85°C. Un entorno perfecto para la recuperación muscular y el bienestar mental.",
		Beneficios:       "Mejora del flujo sanguíneo, relajación muscular profunda y beneficios cardiovasculares.",
	},
	{
		ID:               5,
		Nombre:           "Piscina Olímpica",
		Icono:            "fas fa-swimmer",
		DescripcionBreve: "50 metros de longitud, temperatura controlada y uso reservado por carriles para entreno de impacto cero.",
		DescripcionLarga: "Piscina de 50 metros con 8 calles disponibles para nado libre o entrenamientos dirigidos. El agua está tratada con sal, evitando los químicos agresivos del cloro tradicional, lo que beneficia a tu piel y respiración.",
		Beneficios:       "Cardio sin impacto articular, mejora de la capacidad pulmonar y recuperación activa.",
	},
	{
		ID:               6,
		Nombre:           "Crioterapia",
		Icono:            "fas fa-snowflake",
		DescripcionBreve: "Terapias bajo cero (-110°C) para acelerar recuperación muscular y reducir niveles de inflamación.",
		DescripcionLarga: "Las cabinas de crioterapia de cuerpo entero someten a tu cuerpo a temperaturas extremas por corto tiempo (2-3 minutos). Este shock térmico desencadena potentes respuestas antiinflamatorias y analgésicas en todo el cuerpo.",
		Beneficios:       "Reducción del estrés oxidativo, disminución drástica del dolor post-entrenamiento e incremento de energía.",
	},
}
