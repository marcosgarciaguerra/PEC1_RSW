# -*- coding: utf-8 -*-
"""Genera copia solo frontend (HTML/CSS/SCSS) en carpeta hermana Proyecto-frontend."""
from __future__ import annotations

import shutil
from pathlib import Path

ROOT = Path(__file__).resolve().parents[1]
DEST = ROOT.parent / "Proyecto-frontend"

MAQUINAS = [
    {"id": 1, "nombre": "Chest Press Biomecánico", "marca": "Panatta Sport", "zona": "Tren Superior",
     "descripcion": "Máquina diseñada para trabajar el pectoral de forma segura y biomecánicamente perfecta. Aisla el músculo y permite cargar peso sin comprometer los hombros.",
     "beneficios": "Desarrollo máximo del pectoral mayor, deltoides anterior y tríceps. Ideal para hipertrofia segura.",
     "imagen": "img/maquinaria/ChestPress.webp"},
    {"id": 2, "nombre": "Iso-Lateral Row", "marca": "Hammer Strength", "zona": "Tren Superior",
     "descripcion": "Permite trabajar la espalda de forma unilateral, corrigiendo asimetrías y logrando una contracción máxima de los dorsales y romboides.",
     "beneficios": "Mejora la densidad de la espalda, corrige desequilibrios musculares y fortalece el core al requerir estabilización.",
     "imagen": "img/maquinaria/IsoLateralRow.webp"},
    {"id": 3, "nombre": "Leg Press 45°", "marca": "Cybex", "zona": "Tren Inferior",
     "descripcion": "Prensa de piernas a 45 grados con rieles ultrasuaves. Permite mover grandes cargas con total seguridad para la columna lumbar.",
     "beneficios": "Desarrollo masivo de cuádriceps, glúteos e isquiosurales sin tensión en la espalda baja.",
     "imagen": "img/maquinaria/LegPress.webp"},
    {"id": 4, "nombre": "Hack Squat Lineal", "marca": "Nautilus", "zona": "Tren Inferior",
     "descripcion": "Simula el movimiento de una sentadilla pero proporcionando un apoyo total a la espalda. Ideal para enfocar el esfuerzo en los cuádriceps.",
     "beneficios": "Hipertrofia extrema en las piernas, menos riesgo de lesión de rodilla gracias a la biomecánica guiada.",
     "imagen": "img/maquinaria/HackSquat.jpg"},
    {"id": 5, "nombre": "Abdominal Crunch", "marca": "Technogym", "zona": "Core (Abdomen)",
     "descripcion": "Máquina de aislamiento para la pared abdominal. Su diseño permite una flexión del tronco perfecta aislando el recto abdominal.",
     "beneficios": "Fortalecimiento del core, prevención de dolores lumbares y desarrollo estético del abdomen.",
     "imagen": "img/maquinaria/AbdominalCrunch.jpg"},
    {"id": 6, "nombre": "Cable Crossover", "marca": "Panatta Sport", "zona": "Tren Superior",
     "descripcion": "Estación de poleas cruzadas de alta precisión con múltiples alturas ajustables. Perfecta para ejercicios de aislamiento.",
     "beneficios": "Trabajo constante durante todo el rango de movimiento, ideal para bombear sangre al final del entrenamiento.",
     "imagen": "img/maquinaria/CableCrossover.png"},
    {"id": 7, "nombre": "Lat Pulldown Convergente", "marca": "Hammer Strength", "zona": "Tren Superior",
     "descripcion": "Máquina para jalón al pecho con movimiento convergente, lo que permite un estiramiento y contracción superior a los jalones tradicionales.",
     "beneficios": "Ensanchamiento de la espalda (dorsal ancho) y mejora de la fuerza de tracción.",
     "imagen": "img/maquinaria/LatPulldown.avif"},
    {"id": 8, "nombre": "Glute Drive", "marca": "Nautilus", "zona": "Tren Inferior",
     "descripcion": "Máquina diseñada específicamente para el Hip Thrust. Elimina la necesidad de usar barra libre e incomodidad en la pelvis.",
     "beneficios": "Desarrollo óptimo de los glúteos, mejora el rendimiento atlético sin sobrecargar la espalda.",
     "imagen": "img/maquinaria/GluteDrive.webp"},
]

SERVICIOS = [
    {"id": 1, "nombre": "Sala de Máquinas", "icono": "fas fa-dumbbell",
     "breve": "Bio-mechanics focus. Equipamiento seleccionado científicamente para aislar cada grupo muscular de forma estricta.",
     "larga": "Nuestra sala de musculación cuenta con más de 100 máquinas de marcas premium (Panatta, Hammer Strength, Cybex). Cada zona está optimizada para diferentes grupos musculares, con luz controlada y ambiente enfocado al rendimiento puro. Perfecto tanto para principiantes como para culturistas de élite.",
     "beneficios": "Equipamiento de máxima calidad, cero tiempos de espera gracias a duplicados de máquinas populares y entorno motivador."},
    {"id": 2, "nombre": "Fisioterapia Premium", "icono": "fas fa-user-md",
     "breve": "Recuperación profunda, punción seca, liberación miofascial y osteopatía al servicio de tu rendimiento.",
     "larga": "Un equipo de fisioterapeutas deportivos a tu disposición. Tratamos lesiones agudas, crónicas y realizamos labores de prevención. Utilizamos técnicas avanzadas de terapia manual, electrolisis percutánea, neuromodulación y trabajo activo.",
     "beneficios": "Recuperación acelerada, prevención de lesiones y tratamientos personalizados."},
    {"id": 3, "nombre": "Nutricionista", "icono": "fas fa-apple-alt",
     "breve": "Planes dietéticos personalizados, periodización nutricional y análisis corporal por pliegues y DEXA.",
     "larga": "Alcanza tus objetivos más rápido con nuestro servicio de nutrición clínica y deportiva. Incluye valoración antropométrica detallada, diseño de planes de alimentación adaptables y seguimiento continuo para garantizar que tu dieta apoye tu entrenamiento.",
     "beneficios": "Pérdida de grasa eficiente, aumento de masa muscular limpio y mejora del rendimiento."},
    {"id": 4, "nombre": "Sauna Finlandesa", "icono": "fas fa-hot-tub",
     "breve": "Relajación térmica para eliminación de toxinas, mejora cardiovascular y recuperación del sistema nervioso.",
     "larga": "Termina tu sesión de entrenamiento pesado relajándote en nuestras saunas finlandesas de madera de pino a 85°C. Un entorno perfecto para la recuperación muscular y el bienestar mental.",
     "beneficios": "Mejora del flujo sanguíneo, relajación muscular profunda y beneficios cardiovasculares."},
    {"id": 5, "nombre": "Piscina Olímpica", "icono": "fas fa-swimmer",
     "breve": "50 metros de longitud, temperatura controlada y uso reservado por carriles para entreno de impacto cero.",
     "larga": "Piscina de 50 metros con 8 calles disponibles para nado libre o entrenamientos dirigidos. El agua está tratada con sal, evitando los químicos agresivos del cloro tradicional, lo que beneficia a tu piel y respiración.",
     "beneficios": "Cardio sin impacto articular, mejora de la capacidad pulmonar y recuperación activa."},
    {"id": 6, "nombre": "Crioterapia", "icono": "fas fa-snowflake",
     "breve": "Terapias bajo cero (-110°C) para acelerar recuperación muscular y reducir niveles de inflamación.",
     "larga": "Las cabinas de crioterapia de cuerpo entero someten a tu cuerpo a temperaturas extremas por corto tiempo (2-3 minutos). Este shock térmico desencadena potentes respuestas antiinflamatorias y analgésicas en todo el cuerpo.",
     "beneficios": "Reducción del estrés oxidativo, disminución drástica del dolor post-entrenamiento e incremento de energía."},
]

EQUIPO = [
    {"nombre": "Arnold Schwarzenegger", "cargo": "Director Deportivo & Leyenda",
     "desc": "El roble austriaco. 7 veces Mr. Olympia. Supervisa los estándares de entrenamiento de Platinum Gym.",
     "img": "img/equipo/arnold.png"},
    {"nombre": "Chris Bumstead", "cargo": "Head Coach Classic Physique",
     "desc": "Dominante en la categoría Classic. Especialista en estética, vacío abdominal y poses clásicas.",
     "img": "img/equipo/cbum.png"},
    {"nombre": "Ronnie Coleman", "cargo": "Especialista en Fuerza Extrema",
     "desc": "The King. 8 veces Mr. Olympia. Su lema: 'Light weight baby!'. Gestiona la zona de peso libre pesado.",
     "img": "img/equipo/ronnie_coleman.png"},
    {"nombre": "Joan Pradells", "cargo": "Coach de Powerlifting & Masa",
     "desc": "Referente nacional en fuerza. Experto en sentadilla, banca y peso muerto con cargas masivas.",
     "img": "img/equipo/joan_pradells.png"},
    {"nombre": "Andoni Fitness", "cargo": "Especialista en Estética Natural",
     "desc": "Experto en transformación física y entrenamiento de alta intensidad enfocado a la estética.",
     "img": "img/equipo/andoni.png"},
    {"nombre": "Coach Greg", "cargo": "Asesor Nutricional & Cardio",
     "desc": "Doctor en sentido común. Especialista en nutrición eficiente, suplementación y entrenamiento inteligente.",
     "img": "img/equipo/coach_greg.png"},
    {"nombre": "Alex Eubank", "cargo": "Instructor de Fitness Moderno",
     "desc": "Enfocado en la estética 'Greek God'. Especialista en entrenamiento de hombro y brazo para físico fitness.",
     "img": "img/equipo/alex_eubank.png"},
    {"nombre": "The Tren Twins", "cargo": "Directores de Intensidad",
     "desc": "Especialistas en la cultura del 'hardcore gym'. Llevan la intensidad de tus entrenos al límite absoluto.",
     "img": "img/equipo/tren_twins.png"},
    {"nombre": "Andreas Shaw", "cargo": "Coach de Strongman",
     "desc": "Experto en movimientos de fuerza absoluta y transporte de cargas pesadas.",
     "img": "img/equipo/andreashaw.jpeg"},
]

RESENAS = [
    {"autor": "Carlos M.", "puntuacion": 5, "texto": "Mejor gimnasio de la zona. Máquinas de lujo y ambiente brutal."},
    {"autor": "Laura G.", "puntuacion": 4, "texto": "La piscina olímpica es una maravilla. Un poco lleno a las 18h."},
    {"autor": "David R.", "puntuacion": 5, "texto": "El equipo de entrenadores es de otro nivel. He ganado 5kg de músculo en 3 meses."},
]

CLASES_DEMO = [
    {"id": 1, "nombre": "HIIT Metabolic", "horario": "Lunes y Miércoles 19:00", "lugar": "Sala 2", "entrenador": "Alex Eubank",
     "desc": "Sesión de alta intensidad para quemar grasa y mejorar el condicionamiento."},
    {"id": 2, "nombre": "Powerlifting Técnico", "horario": "Martes 20:00", "lugar": "Zona Fuerza", "entrenador": "Joan Pradells",
     "desc": "Trabajo de banca, sentadilla y peso muerto con corrección de técnica."},
    {"id": 3, "nombre": "Yoga Recovery", "horario": "Domingos 10:00", "lugar": "Sala Multipurpose", "entrenador": "Laura G.",
     "desc": "Movilidad y recuperación activa para complementar la musculación."},
]

DEMO_USER = {
    "nombre": "Socio", "apellidos": "Demo",
    "direccion": "Calle Mayor 12, Meco, Madrid",
    "telefono": "+34 600 000 000",
    "correo": "socio.demo@platinumgym.com",
    "metodo_pago": "tarjeta",
}


def stars(n: int) -> str:
    m = {5: "★★★★★", 4: "★★★★☆", 3: "★★★☆☆", 2: "★★☆☆☆", 1: "★☆☆☆☆"}
    return m.get(n, "★☆☆☆☆")


def navbar() -> str:
    return """        <header class="barra-navegacion unified-header">
            <div class="header-logo">
                <a href="index.html" class="logo-enlace" aria-label="Volver al inicio">
                    <div class="logo-brand"></div>
                </a>
            </div>

            <nav class="header-nav">
                <ul class="lista-enlaces">
                    <li class="elemento-lista"><a href="index.html" class="enlace"><i class="fas fa-home" aria-hidden="true"></i><span>Inicio</span></a></li>
                    <li class="elemento-lista dropdown">
                        <a href="#" class="enlace">
                            <i class="fas fa-dumbbell" aria-hidden="true"></i>
                            <span>El Gimnasio</span>
                            <i class="fas fa-chevron-down" style="font-size: 0.7rem; margin-left: 0.5rem;"></i>
                        </a>
                        <ul class="dropdown-menu">
                            <li><a href="maquinaria.html"><i class="fas fa-cogs"></i> Maquinaria</a></li>
                            <li><a href="servicios.html"><i class="fas fa-crown"></i> Servicios</a></li>
                            <li><a href="equipo.html"><i class="fas fa-users"></i> Nuestro Equipo</a></li>
                            <li><a href="reglas.html"><i class="fas fa-scroll"></i> Reglas del Centro</a></li>
                        </ul>
                    </li>
                    <li class="elemento-lista"><a href="tienda.html" class="enlace"><i class="fas fa-shopping-bag" aria-hidden="true"></i><span>Tienda</span></a></li>
                    <li class="elemento-lista dropdown">
                        <a href="#" class="enlace">
                            <i class="fas fa-user-circle" aria-hidden="true"></i>
                            <span>Socios</span>
                            <i class="fas fa-chevron-down" style="font-size: 0.7rem; margin-left: 0.5rem;"></i>
                        </a>
                        <ul class="dropdown-menu">
                            <li><a href="reservas.html"><i class="fas fa-calendar-check"></i> Área Socios / Reservas</a></li>
                            <li><a href="calculadora.html"><i class="fas fa-calculator"></i> Calculadora FFMI</a></li>
                            <li><a href="apuntate.html"><i class="fas fa-user-plus"></i> ¡Apúntate!</a></li>
                        </ul>
                    </li>
                </ul>
            </nav>

            <div class="header-actions">
                <div class="barra-navegacion__buscador">
                    <form action="buscar.html" method="GET" class="search-form">
                        <button type="submit" class="search-icon-btn" aria-label="Buscar"><i class="fas fa-search"></i></button>
                        <input type="text" name="q" class="search-input-with-icon" placeholder="¿Qué buscas?">
                    </form>
                </div>
                <div class="barra-navegacion__iconos">
                    <div class="user-dropdown-container">
                        <a href="login.html" class="user-dropdown-trigger" style="text-decoration:none;color:inherit;display:inline-flex;align-items:center;gap:0.35rem;">
                            <span>Login</span>
                            <i class="fas fa-user-circle"></i>
                        </a>
                    </div>
                </div>
                <label for="menu-toggle" class="mobile-menu-btn" aria-label="Alternar menú móvil"><i class="fas fa-bars"></i></label>
            </div>
        </header>
"""


def footer() -> str:
    return """<footer class="pie-pagina">
    <div class="pie-pagina__contenedor">
        <section class="pie-pagina__columna">
            <div class="pie-pagina__logo">
                <div class="pie-pagina__logo-brand"></div>
            </div>
            <h4>Platinum Gym</h4>
            <p>Expertos en ponerte mamadismo</p>
        </section>
        <section class="pie-pagina__columna">
            <h4>Enlaces Rápidos</h4>
            <nav>
                <ul class="pie-pagina__lista">
                    <li><a href="maquinaria.html">Maquinaria</a></li>
                    <li><a href="servicios.html">Servicios</a></li>
                    <li><a href="tienda.html">Tienda</a></li>
                    <li><a href="reglas.html">Reglas</a></li>
                    <li><a href="equipo.html">Equipo</a></li>
                    <li><a href="apuntate.html">¡Apúntate!</a></li>
                </ul>
            </nav>
        </section>
        <section class="pie-pagina__columna">
            <h4>Contacto</h4>
            <ul class="pie-pagina__lista">
                <li><i class="fas fa-phone" aria-hidden="true"></i> +34 912 345 678</li>
                <li><i class="fas fa-envelope" aria-hidden="true"></i> info@platinumgym.com</li>
            </ul>
        </section>
        <section class="pie-pagina__columna">
            <h4>Gimnasios</h4>
            <ul class="pie-pagina__lista">
                <li><i class="fas fa-map-marker-alt" aria-hidden="true"></i> Seattle</li>
                <li><i class="fas fa-map-marker-alt" aria-hidden="true"></i> Tel-Aviv</li>
                <li><i class="fas fa-map-marker-alt" aria-hidden="true"></i> Isla de Epstein</li>
                <li><i class="fas fa-map-marker-alt" aria-hidden="true"></i> Vallecas</li>
                <li><i class="fas fa-map-marker-alt" aria-hidden="true"></i> Meco</li>
                <li><i class="fas fa-map-marker-alt" aria-hidden="true"></i> Berlín</li>
            </ul>
        </section>
    </div>
    <div class="pie-pagina__copyright">
        <p>&copy; 2026 Platinum Gym. Todos los derechos reservados.</p>
    </div>
</footer>
"""


def layout(title: str, main: str, extra_head: str = "") -> str:
    return f"""<!DOCTYPE html>
<html lang="es">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Platinum Gym | {title}</title>
    <link rel="icon" type="image/png" href="img/logo_platinium.png">
    <link rel="stylesheet" href="css/estilos.css">
    <link rel="preconnect" href="https://fonts.googleapis.com">
    <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
    <link href="https://fonts.googleapis.com/css2?family=Montserrat:ital,wght@0,300;0,400;0,600;0,800;0,900;1,900&display=swap" rel="stylesheet">
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.4.0/css/all.min.css">
    {extra_head}
</head>
<body>
    <input type="checkbox" id="menu-toggle" class="menu-toggle">
    <div class="aplicacion">
        {navbar()}
        <main class="contenido-principal" id="contenido">
{main}
        </main>
        {footer()}
    </div>
</body>
</html>
"""


def page_index() -> str:
    res_html = "\n".join(
        f"""                <li class="resena-card">
                    <h3 class="resena-autor">{r["autor"]}</h3>
                    <div class="resena-stars" aria-label="Puntuación: {r["puntuacion"]}">{stars(r["puntuacion"])}</div>
                    <p class="resena-text">"{r["texto"]}"</p>
                </li>"""
        for r in RESENAS
    )
    main = f"""<nav class="migas-pan" aria-label="Breadcrumb">
    <span class="actual">Inicio</span>
</nav>

<section class="portada">
    <div class="texto-portada">
        <h1>Bienvenido a <span class="texto-destacado">Platinum Gym</span></h1>
        <a href="apuntate.html" class="boton-principal">¡Hazte Socio Ahora!</a>
    </div>
</section>

<section class="seccion">
    <h2 class="titulo-seccion">Nuestros Locales</h2>
    <ul class="contenido-seccion">
        <li class="tarjeta">
            <img src="img/ubicaciones/meco.png" alt="Local Meco" class="imagen-tarjeta">
            <div class="cuerpo-tarjeta">
                <h3 class="nombre-tarjeta">Meco</h3>
                <details class="detalles-tarjeta">
                    <summary class="ver-mas">Ver información</summary>
                    <div class="descripcion-tarjeta">
                        <p>Gimnasio completo con zona de peso libre, máquinas de cardio y sala de clases dirigidas. Abierto de 6:00 a 23:00.</p>
                        <p><strong>Dirección:</strong> C/ Calle Mayor 12, Meco, Madrid</p>
                    </div>
                </details>
            </div>
        </li>
        <li class="tarjeta">
            <img src="img/ubicaciones/vallecas.avif" alt="Local Vallecas" class="imagen-tarjeta">
            <div class="cuerpo-tarjeta">
                <h3 class="nombre-tarjeta">Vallecas</h3>
                <details class="detalles-tarjeta">
                    <summary class="ver-mas">Ver información</summary>
                    <div class="descripcion-tarjeta">
                        <p>Ubicación emblemática con equipamiento Hammer Strength y zona de powerlifting. Abierto 24h para socios VIP.</p>
                        <p><strong>Dirección:</strong> Av. de la Albufera 45, Vallecas, Madrid</p>
                    </div>
                </details>
            </div>
        </li>
        <li class="tarjeta">
            <img src="img/ubicaciones/Chicago.jpg" alt="Local Seattle" class="imagen-tarjeta">
            <div class="cuerpo-tarjeta">
                <h3 class="nombre-tarjeta">Seattle (Central)</h3>
                <details class="detalles-tarjeta">
                    <summary class="ver-mas">Ver información</summary>
                    <div class="descripcion-tarjeta">
                        <p>Nuestra sede central. 5 plantas de puro acero, piscina olímpica y zona de crioterapia.</p>
                        <p><strong>Dirección:</strong> 1201 3rd Ave, Seattle, WA 98101, USA</p>
                    </div>
                </details>
            </div>
        </li>
        <li class="tarjeta">
            <img src="img/ubicaciones/tel-aviv.webp" alt="Local Tel-Aviv" class="imagen-tarjeta">
            <div class="cuerpo-tarjeta">
                <h3 class="nombre-tarjeta">Tel-Aviv</h3>
                <details class="detalles-tarjeta">
                    <summary class="ver-mas">Ver información</summary>
                    <div class="descripcion-tarjeta">
                        <p>Un gimnasio costero de última generación. Especialidad en calistenia deportiva y CrossTraining al aire libre.</p>
                        <p><strong>Dirección:</strong> 12 Namal Tel Aviv St, Tel Aviv-Yafo, Israel</p>
                    </div>
                </details>
            </div>
        </li>
        <li class="tarjeta">
            <img src="img/ubicaciones/epsteinisla.jfif" alt="Local Isla de Epstein" class="imagen-tarjeta">
            <div class="cuerpo-tarjeta">
                <h3 class="nombre-tarjeta">Isla de Epstein</h3>
                <details class="detalles-tarjeta">
                    <summary class="ver-mas">Ver información</summary>
                    <div class="descripcion-tarjeta">
                        <p>Complejo exclusivo para entrenamientos de extrema privacidad y alto rendimiento. Acceso restringido.</p>
                        <p><strong>Dirección:</strong> Little St. James, US Virgin Islands</p>
                    </div>
                </details>
            </div>
        </li>
        <li class="tarjeta">
            <img src="img/ubicaciones/berlin.jfif" alt="Local Berlín" class="imagen-tarjeta">
            <div class="cuerpo-tarjeta">
                <h3 class="nombre-tarjeta">Berlín</h3>
                <details class="detalles-tarjeta">
                    <summary class="ver-mas">Ver información</summary>
                    <div class="descripcion-tarjeta">
                        <p>Espacio industrial reciclado, ideal para rutinas de Powerlifting underground y Strongman.</p>
                        <p><strong>Dirección:</strong> Friedrichstraße 43, 10117 Berlin, Alemania</p>
                    </div>
                </details>
            </div>
        </li>
    </ul>
</section>

<section class="seccion">
    <h2 class="titulo-seccion">Opiniones de nuestros Guerreros</h2>
    <div class="resenas-container">
        <div class="resenas-form-wrapper">
            <div class="resena-form-card resena-form-card--bloqueado">
                <i class="fas fa-lock lock-icon"></i>
                <h3>Acceso Restringido</h3>
                <p>Versión estática: reseñas de ejemplo abajo. En la app completa solo socios autenticados publican.</p>
                <a href="login.html" class="boton-formulario mt-1 inline-block no-underline">Iniciar Sesión</a>
            </div>
        </div>
        <div class="resenas-list-wrapper">
            <ul class="resenas-grid">
{res_html}
            </ul>
        </div>
    </div>
</section>
"""
    return layout("Inicio", main)


def page_maquinaria() -> str:
    cards = []
    for m in MAQUINAS:
        cards.append(f"""            <li class="tarjeta">
                <img src="{m["imagen"]}" alt="{m["nombre"]}" class="imagen-tarjeta">
                <div class="cuerpo-tarjeta">
                    <span class="item-resultado__tipo">{m["marca"]} | {m["zona"]}</span>
                    <h3 class="nombre-tarjeta">{m["nombre"]}</h3>
                    <p class="descripcion-tarjeta">{m["descripcion"]}</p>
                    <a href="maquinaria-detalle-{m["id"]}.html" class="boton-principal boton-bloque">Ficha Técnica</a>
                </div>
            </li>""")
    main = f"""    <nav class="migas-pan" aria-label="Breadcrumb">
        <a href="index.html">Inicio</a>
        <span class="separador"><i class="fas fa-chevron-right"></i></span>
        <span class="actual">Maquinaria</span>
    </nav>

    <section id="maquinaria-list" class="seccion">
        <div class="cabecera-seccion">
            <h2>NUESTRA <span class="texto-destacado">MAQUINARIA</span></h2>
            <p>Ingeniería de vanguardia para el máximo aislamiento muscular.</p>
        </div>

        <ul class="contenido-seccion">
{chr(10).join(cards)}
        </ul>
    </section>
"""
    return layout("Maquinaria", main)


def page_maquinaria_detalle(m: dict) -> str:
    main = f"""    <nav class="migas-pan" aria-label="Breadcrumb">
        <a href="index.html">Inicio</a>
        <span class="separador"><i class="fas fa-chevron-right"></i></span>
        <a href="maquinaria.html">Maquinaria</a>
        <span class="separador"><i class="fas fa-chevron-right"></i></span>
        <span class="actual">{m["nombre"]}</span>
    </nav>

    <div class="detalle-entidad">
        <div class="detalle-grid">
            <div class="detalle-imagen">
                <img src="{m["imagen"]}" alt="{m["nombre"]}" class="img-fluida rounded shadow">
            </div>
            <div class="detalle-info">
                <span class="badget text-muted">{m["marca"]} | {m["zona"]}</span>
                <h2 class="detalle-titulo">{m["nombre"]}</h2>
                <div class="detalle-seccion">
                    <h3>Descripción</h3>
                    <p>{m["descripcion"]}</p>
                </div>
                <div class="detalle-seccion">
                    <h3>Beneficios Clave</h3>
                    <p>{m["beneficios"]}</p>
                </div>
                <div class="detalle-acciones">
                    <a href="maquinaria.html" class="boton-secundario"><i class="fas fa-arrow-left"></i> Volver al listado</a>
                    <a href="login.html" class="boton-principal">Probar en Sala</a>
                </div>
            </div>
        </div>
    </div>
"""
    return layout(m["nombre"], main)


def page_servicios() -> str:
    cards = []
    for s in SERVICIOS:
        cards.append(f"""            <li class="tarjeta">
                <div class="cuerpo-tarjeta">
                    <div class="regla-icon"><i class="{s["icono"]}"></i></div>
                    <h3 class="nombre-tarjeta">{s["nombre"]}</h3>
                    <p class="descripcion-tarjeta">{s["breve"]}</p>
                    <a href="servicio-detalle-{s["id"]}.html" class="boton-principal boton-bloque">Saber Más</a>
                </div>
            </li>""")
    main = f"""    <nav class="migas-pan" aria-label="Breadcrumb">
        <a href="index.html">Inicio</a>
        <span class="separador"><i class="fas fa-chevron-right"></i></span>
        <span class="actual">Servicios</span>
    </nav>

    <section id="servicios-list" class="seccion">
        <div class="cabecera-seccion">
            <h2>SERVICIOS <span class="texto-destacado">PLATINUM</span></h2>
            <p>Más que un gimnasio, un centro de alto rendimiento integral.</p>
        </div>

        <ul class="contenido-seccion">
{chr(10).join(cards)}
        </ul>
    </section>
"""
    return layout("Servicios", main)


def page_servicio_detalle(s: dict) -> str:
    main = f"""    <nav class="migas-pan" aria-label="Breadcrumb">
        <a href="index.html">Inicio</a>
        <span class="separador"><i class="fas fa-chevron-right"></i></span>
        <a href="servicios.html">Servicios</a>
        <span class="separador"><i class="fas fa-chevron-right"></i></span>
        <span class="actual">{s["nombre"]}</span>
    </nav>

    <div class="detalle-entidad">
        <div class="detalle-grid">
            <div class="detalle-icono-wrapper">
                <div class="grande-icon"><i class="{s["icono"]}"></i></div>
            </div>
            <div class="detalle-info">
                <h2 class="detalle-titulo">{s["nombre"]}</h2>
                <div class="detalle-seccion">
                    <h3>Descripción Detallada</h3>
                    <p>{s["larga"]}</p>
                </div>
                <div class="detalle-seccion">
                    <h3>Ventajas para Ti</h3>
                    <p>{s["beneficios"]}</p>
                </div>
                <div class="detalle-acciones">
                    <a href="servicios.html" class="boton-secundario"><i class="fas fa-arrow-left"></i> Volver a servicios</a>
                    <a href="apuntate.html" class="boton-principal">Solicitar Información</a>
                </div>
            </div>
        </div>
    </div>
"""
    return layout(s["nombre"], main)


def page_equipo() -> str:
    cards = []
    for e in EQUIPO:
        cards.append(f"""            <li class="tarjeta tarjeta-equipo">
                <div class="caja-imagen-equipo">
                    <img src="{e["img"]}" alt="{e["nombre"]}" class="imagen-tarjeta imagen-equipo">
                </div>
                <div class="cuerpo-tarjeta cuerpo-equipo">
                    <h3 class="nombre-tarjeta mb-05">{e["nombre"]}</h3>
                    <div class="etiqueta-cargo">{e["cargo"]}</div>
                    <p class="descripcion-tarjeta lh-15">{e["desc"]}</p>
                </div>
            </li>""")
    main = f"""    <nav class="migas-pan" aria-label="Breadcrumb">
        <a href="index.html">Inicio</a>
        <span class="separador"><i class="fas fa-chevron-right"></i></span>
        <span class="actual">Nuestro Equipo</span>
    </nav>

    <section class="seccion">
        <div class="cabecera-seccion">
            <h2>EQUIPO DE <span class="texto-destacado">ÉLITE</span></h2>
            <p>Conoce al corazón de Platinum Gym: profesionales altamente cualificados dedicados a tu éxito.</p>
        </div>

        <ul class="contenido-seccion equipo-grid">
{chr(10).join(cards)}
        </ul>
    </section>
"""
    return layout("Equipo", main)


def page_reglas() -> str:
    main = """        <nav class="migas-pan" aria-label="Breadcrumb">
            <a href="index.html">Inicio</a>
            <span class="separador"><i class="fas fa-chevron-right"></i></span>
            <span class="actual">Reglas del Centro</span>
        </nav>

        <section id="reglas-centro">
            <div class="cabecera-seccion">
                <h2>NORMAS DE <span class="texto-destacado">PLATINUM GYM</span></h2>
                <p>La disciplina es el puente entre las metas y los logros. Respeta estas normas por tu seguridad y la de los demás.</p>
            </div>

            <div class="reglas-grid">
                <article class="regla-card">
                    <div class="regla-icon"><i class="fas fa-tshirt"></i></div>
                    <h3>Ropa Adecuada</h3>
                    <p>Es obligatorio el uso de ropa deportiva y calzado cerrado. No se permite entrenar sin camiseta o con calzado de calle.</p>
                </article>
                <article class="regla-card">
                    <div class="regla-icon"><i class="fas fa-hands-wash"></i></div>
                    <h3>Higiene Personal</h3>
                    <p>El uso de toalla es obligatorio en todas las máquinas y bancos. Por favor, limpia el sudor después de cada uso con el desinfectante disponible.</p>
                </article>
                <article class="regla-card">
                    <div class="regla-icon"><i class="fas fa-sync-alt"></i></div>
                    <h3>Orden en Sala</h3>
                    <p>Descarga las máquinas y devuelve las pesas y barras a su lugar correspondiente al terminar el ejercicio. No abandones material en el suelo.</p>
                </article>
                <article class="regla-card">
                    <div class="regla-icon"><i class="fas fa-mobile-alt"></i></div>
                    <h3>Uso de Dispositivos</h3>
                    <p>Se permite el uso de móviles para música, pero no para llamadas largas o grabaciones que incomoden a otros socios sin permiso.</p>
                </article>
                <article class="regla-card">
                    <div class="regla-icon"><i class="fas fa-clock"></i></div>
                    <h3>Horarios y Turnos</h3>
                    <p>Respeta los turnos de las máquinas en horas punta. No acapares material si no estás realizando el ejercicio activamente.</p>
                </article>
                <article class="regla-card">
                    <div class="regla-icon"><i class="fas fa-user-shield"></i></div>
                    <h3>Comportamiento</h3>
                    <p>Mantén un clima de respeto. Los gritos excesivos o el lanzamiento violento de pesas (salvo en zonas de plataforma) están prohibidos.</p>
                </article>
            </div>
        </section>
"""
    return layout("Reglas del Centro", main)


def page_apuntate() -> str:
    main = """<section class="seccion auth-flex-container">
    <div class="auth-card">
        <div class="cabecera-seccion">
            <h2>Únete al <span class="texto-destacado">Olimpo</span></h2>
            <p>Versión estática: formulario maquetado sin envío a servidor.</p>
        </div>

        <form action="#" method="get" class="formulario">
            <div class="grupo-formulario">
                <label for="nombre" class="etiqueta-formulario">Nombre:</label>
                <input type="text" id="nombre" name="nombre" class="campo-formulario" placeholder="Ej: Arnold" required>
            </div>
            <div class="grupo-formulario">
                <label for="apellidos" class="etiqueta-formulario">Apellidos:</label>
                <input type="text" id="apellidos" name="apellidos" class="campo-formulario" placeholder="Ej: Schwarzenegger" required>
            </div>
            <div class="grupo-formulario">
                <label for="fecha-nacimiento" class="etiqueta-formulario">Fecha de Nacimiento:</label>
                <input type="date" id="fecha-nacimiento" name="fecha-nacimiento" class="campo-formulario" required>
            </div>
            <div class="grupo-formulario">
                <label for="direccion" class="etiqueta-formulario">Dirección:</label>
                <input type="text" id="direccion" name="direccion" class="campo-formulario" placeholder="Calle Mayor, 12" required>
            </div>
            <div class="grupo-formulario">
                <label for="telefono" class="etiqueta-formulario">Teléfono:</label>
                <input type="tel" id="telefono" name="telefono" class="campo-formulario" placeholder="Ej: 600123456" required>
            </div>
            <div class="grupo-formulario">
                <label for="correo" class="etiqueta-formulario">Email:</label>
                <input type="email" id="correo" name="correo" class="campo-formulario" placeholder="tu@email.com" required>
            </div>
            <div class="grupo-formulario">
                <label for="documento" class="etiqueta-formulario">DNI/Pasaporte:</label>
                <input type="text" id="documento" name="documento" class="campo-formulario" placeholder="12345678X" required>
            </div>
            <div class="grupo-formulario">
                <label for="plan" class="etiqueta-formulario">Plan deseado:</label>
                <select id="plan" name="plan" class="campo-formulario">
                    <option value="basico">Básico (Solo Sala)</option>
                    <option value="premium">Premium (Sala + Piscina + Sauna)</option>
                    <option value="vip">VIP Platinum (Todo Incluido + Fisio)</option>
                </select>
            </div>
            <div class="grupo-formulario">
                <label for="metodo-pago" class="etiqueta-formulario">Método de pago:</label>
                <select id="metodo-pago" name="metodo-pago" class="campo-formulario">
                    <option value="tarjeta">Tarjeta de Crédito</option>
                    <option value="domiciliacion">Domiciliación Bancaria</option>
                </select>
            </div>
            <div class="grupo-formulario">
                <label for="numero-pago" class="etiqueta-formulario">Número de Tarjeta / IBAN:</label>
                <input type="text" id="numero-pago" name="numero-pago" class="campo-formulario" placeholder="ESXX..." required>
            </div>
            <div class="grupo-formulario">
                <label for="password" class="etiqueta-formulario">Crea una contraseña:</label>
                <input type="password" id="password" name="password" class="campo-formulario" placeholder="••••••••" required>
            </div>
            <div class="grupo-formulario">
                <label for="repeat-password" class="etiqueta-formulario">Repite la contraseña:</label>
                <input type="password" id="repeat-password" name="repeat-password" class="campo-formulario" placeholder="••••••••" required>
            </div>
            <button type="submit" class="boton-principal boton-bloque">Confirmar Registro (demo)</button>
        </form>
        <p class="auth-footer">¿Ya eres socio? <a href="login.html" class="texto-destacado">Inicia sesión</a></p>
    </div>
</section>
"""
    return layout("Apúntate", main)


def page_login() -> str:
    main = """    <section class="seccion auth-flex-container">
        <div class="auth-card">
            <div class="cabecera-seccion">
                <h2>Iniciar <span class="texto-destacado">Sesión</span></h2>
                <p>Versión estática: maquetación sin backend.</p>
            </div>
            <form action="#" method="get" class="formulario">
                <div class="grupo-formulario">
                    <label for="correo" class="etiqueta-formulario">Correo Electrónico:</label>
                    <input type="email" id="correo" name="correo" class="campo-formulario" placeholder="tu@email.com" required>
                </div>
                <div class="grupo-formulario">
                    <label for="contrasena" class="etiqueta-formulario">Contraseña:</label>
                    <input type="password" id="contrasena" name="contrasena" class="campo-formulario" placeholder="••••••••" required>
                </div>
                <button type="submit" class="boton-principal boton-bloque">Entrar (demo)</button>
            </form>
            <p class="auth-footer">¿Aún no tienes cuenta? <a href="apuntate.html" class="texto-destacado">Regístrate aquí</a></p>
        </div>
    </section>
"""
    return layout("Login", main)


def page_calculadora() -> str:
    main = """    <nav class="migas-pan" aria-label="Breadcrumb">
        <a href="index.html">Inicio</a>
        <span class="separador"><i class="fas fa-chevron-right"></i></span>
        <span class="actual">Calculadora FFMI</span>
    </nav>

    <section class="seccion">
        <div class="cabecera-seccion">
            <h2>Calculadora <span class="texto-destacado">FFMI</span></h2>
            <p>Mide tu índice de masa libre de grasa para determinar tu potencial genético.</p>
        </div>

        <div class="calculadora-card">
            <form action="#" method="get" class="formulario">
                <div class="grupo-formulario">
                    <label for="altura" class="etiqueta-formulario">Altura (cm):</label>
                    <input type="number" id="altura" name="altura" step="0.1" class="campo-formulario" placeholder="Ej: 180" required>
                </div>
                <div class="grupo-formulario">
                    <label for="peso" class="etiqueta-formulario">Peso (kg):</label>
                    <input type="number" id="peso" name="peso" step="0.1" class="campo-formulario" placeholder="Ej: 85" required>
                </div>
                <div class="grupo-formulario">
                    <label for="grasa" class="etiqueta-formulario">Grasa Corporal (%):</label>
                    <input type="number" id="grasa" name="grasa" step="0.1" class="campo-formulario" placeholder="Ej: 12" required>
                </div>
                <button type="submit" class="boton-principal boton-bloque">Calcular (demo estático)</button>
            </form>
            <p class="auth-footer" style="margin-top:1rem;">En la app completa el cálculo se hace en el servidor.</p>
        </div>
    </section>
"""
    return layout("Calculadora FFMI", main)


def page_reservas() -> str:
    cl = []
    for c in CLASES_DEMO:
        cl.append(f"""                <article class="tarjeta">
                    <div class="cuerpo-tarjeta">
                        <h3 class="nombre-tarjeta">{c["nombre"]}</h3>
                        <div class="reserva-info">
                            <p><i class="fas fa-calendar-alt icon-margin"></i> <strong>{c["horario"]}</strong></p>
                            <p><i class="fas fa-map-marker-alt icon-margin"></i> {c["lugar"]}</p>
                            <p><i class="fas fa-user-circle icon-margin"></i> {c["entrenador"]}</p>
                        </div>
                        <p class="reserva-desc">{c["desc"]}</p>
                        <form action="#" method="get" class="mt-auto">
                            <input type="hidden" name="clase_id" value="{c["id"]}">
                            <button type="submit" class="boton-principal boton-bloque">Reservar (demo)</button>
                        </form>
                    </div>
                </article>""")
    main = f"""    <nav class="migas-pan" aria-label="Breadcrumb">
        <a href="index.html">Inicio</a>
        <span class="separador"><i class="fas fa-chevron-right"></i></span>
        <span class="actual">Panel de Reservas</span>
    </nav>

    <div class="contenedor-reservas">
        <section class="seccion">
            <div class="reservas-header">
                <h2>Hola, <span class="texto-destacado">{DEMO_USER["nombre"]}</span></h2>
                <a href="login.html" class="boton-secundario">Cerrar Sesión (demo)</a>
            </div>
            <p>Versión estática con clases de ejemplo.</p>
        </section>

        <section class="seccion">
            <h2 class="titulo-seccion">Clases Disponibles</h2>
            <div class="grid-reservas">
{chr(10).join(cl)}
            </div>
        </section>

        <section class="seccion">
            <h2 class="titulo-seccion">Mis Reservas Activas</h2>
            <div class="grid-reservas">
                <p style="text-align: center; grid-column: 1/-1; padding: 3rem; background: var(--color-secundario); border-radius: 8px;">Aún no tienes reservas activas (demo).</p>
            </div>
        </section>
    </div>
"""
    return layout("Reservas", main)


def page_buscar() -> str:
    main = """    <section class="seccion">
        <div class="cabecera-seccion">
            <h2>Búsqueda <span class="texto-destacado">(demo)</span></h2>
            <p>Enlaces de ejemplo. En la app completa los resultados vienen del servidor.</p>
        </div>

        <ul class="lista-semantica">
            <li class="item-resultado">
                <div class="item-resultado__info">
                    <span class="item-resultado__tipo">Maquinaria</span>
                    <h3 class="item-resultado__titulo">Leg Press 45°</h3>
                </div>
                <a href="maquinaria-detalle-3.html" class="boton-secundario">Ver Detalles</a>
            </li>
            <li class="item-resultado">
                <div class="item-resultado__info">
                    <span class="item-resultado__tipo">Servicio</span>
                    <h3 class="item-resultado__titulo">Sauna Finlandesa</h3>
                </div>
                <a href="servicio-detalle-4.html" class="boton-secundario">Ver Detalles</a>
            </li>
        </ul>
    </section>
"""
    return layout("Buscar", main)


def page_tienda() -> str:
    # Sin JS: mismas tarjetas, botones como enlaces a tramitar demo
    main = """        <nav class="migas-pan" aria-label="Breadcrumb">
            <a href="index.html">Inicio</a>
            <span class="separador"><i class="fas fa-chevron-right"></i></span>
            <span class="actual">Tienda</span>
        </nav>

        <section id="tienda">
            <div class="cabecera-seccion">
                <h2>TIENDA <span class="texto-destacado">OFICIAL</span></h2>
                <p>Nutrición de grado farmacéutico, ropa técnica y accesorios profesionales.</p>
            </div>
            <p class="filtro-actual" style="text-align:center;margin-bottom:1rem;">Versión estática: filtros y carrito requerían JavaScript; enlaces directos a tramitar de demostración.</p>

            <ul class="contenido-seccion">
                <li class="tarjeta sup">
                    <div class="caja-imagen">
                        <span class="etiqueta">Top Ventas</span>
                        <img src="img/shop/protein.jpg" alt="100% Platinum Whey Isolate" class="imagen-tarjeta">
                    </div>
                    <div class="cuerpo-tarjeta">
                        <h3 class="nombre-tarjeta">100% Platinum Whey Isolate</h3>
                        <div class="precio">59.99€</div>
                        <a href="tramitar-pedido.html" class="boton-principal boton-bloque"><i class="fas fa-cart-plus"></i> Ir a tramitar (demo)</a>
                    </div>
                </li>
                <li class="tarjeta sup">
                    <div class="caja-imagen">
                        <span class="etiqueta">Nuevo</span>
                        <img src="img/shop/prenetreno.avif" alt="Pre-Entreno" class="imagen-tarjeta">
                    </div>
                    <div class="cuerpo-tarjeta">
                        <h3 class="nombre-tarjeta">Pre-Entreno Extremo</h3>
                        <div class="precio">34.99€</div>
                        <a href="tramitar-pedido.html" class="boton-principal boton-bloque"><i class="fas fa-cart-plus"></i> Ir a tramitar (demo)</a>
                    </div>
                </li>
                <li class="tarjeta sup">
                    <div class="caja-imagen">
                        <img src="img/shop/creatina.avif" alt="Creatina" class="imagen-tarjeta">
                    </div>
                    <div class="cuerpo-tarjeta">
                        <h3 class="nombre-tarjeta">Creatina Monohidrato Creapure</h3>
                        <div class="precio">24.99€</div>
                        <a href="tramitar-pedido.html" class="boton-principal boton-bloque"><i class="fas fa-cart-plus"></i> Ir a tramitar (demo)</a>
                    </div>
                </li>
            </ul>
        </section>
"""
    return layout("Tienda", main)


def page_tramitar() -> str:
    u = DEMO_USER
    main = f"""    <nav class="migas-pan" aria-label="Breadcrumb">
        <a href="index.html">Inicio</a>
        <span class="separador"><i class="fas fa-chevron-right"></i></span>
        <a href="tienda.html">Tienda</a>
        <span class="separador"><i class="fas fa-chevron-right"></i></span>
        <span class="actual">Tramitar Pedido</span>
    </nav>

    <section class="seccion checkout-container">
        <div class="checkout-card" style="max-width:1100px;margin:40px auto;padding:0 20px;">
            <div class="cabecera-seccion">
                <h2>Resumen del <span class="texto-destacado">Pedido</span> (demo)</h2>
                <p>Datos y artículos fijos para maquetación sin JavaScript.</p>
            </div>
            <div class="checkout-grid" style="display:grid;grid-template-columns:1fr 1fr;gap:2rem;">
                <div class="checkout-section" style="background:rgba(255,255,255,0.03);padding:1.5rem;border:1px solid rgba(255,255,255,0.08);">
                    <h3><i class="fas fa-truck"></i> Datos de Envío</h3>
                    <p><strong>Destinatario:</strong> {u["nombre"]} {u["apellidos"]}</p>
                    <p><strong>Dirección:</strong> {u["direccion"]}</p>
                    <p><strong>Teléfono:</strong> {u["telefono"]}</p>
                    <p><strong>Email:</strong> {u["correo"]}</p>
                    <h3 style="margin-top:1.5rem;"><i class="fas fa-credit-card"></i> Método de Pago</h3>
                    <p><strong>Método:</strong> {u["metodo_pago"]}</p>
                </div>
                <div class="checkout-section" style="background:rgba(255,255,255,0.03);padding:1.5rem;border:1px solid rgba(255,255,255,0.08);">
                    <h3><i class="fas fa-shopping-basket"></i> Artículos (ejemplo)</h3>
                    <div class="cart-item-js" style="display:flex;justify-content:space-between;padding:12px 0;border-bottom:1px solid #333;">
                        <span>Whey Isolate</span><span style="color:var(--color-principal);font-weight:bold;">59.99€</span>
                    </div>
                    <div class="cart-item-js" style="display:flex;justify-content:space-between;padding:12px 0;">
                        <span>Creatina</span><span style="color:var(--color-principal);font-weight:bold;">24.99€</span>
                    </div>
                    <p style="margin-top:1.5rem;font-size:1.5rem;"><strong>Total:</strong> 84.98€</p>
                    <a href="index.html" class="boton-principal boton-bloque" style="margin-top:1rem;display:inline-block;text-align:center;">Volver al inicio</a>
                </div>
            </div>
        </div>
    </section>
"""
    return layout("Tramitar pedido", main)


README = """# Platinum Gym — solo frontend

HTML estático, CSS compilado y fuentes SCSS. Sin Go ni JavaScript.

## Imágenes

Se copian automáticamente desde `web/static/img` del proyecto Go al ejecutar este script.

## SCSS

Para recompilar CSS desde SCSS (requiere sass instalado):

```bash
sass scss/estilos.scss css/estilos.css
```
"""


def main() -> None:
    if DEST.exists():
        shutil.rmtree(DEST)
    DEST.mkdir(parents=True)
    (DEST / "css").mkdir()
    (DEST / "scss").mkdir()

    shutil.copytree(ROOT / "web" / "static" / "css", DEST / "css", dirs_exist_ok=True)
    shutil.copytree(ROOT / "web" / "static" / "scss", DEST / "scss", dirs_exist_ok=True)
    img_src = ROOT / "web" / "static" / "img"
    if img_src.is_dir():
        shutil.copytree(img_src, DEST / "img", dirs_exist_ok=True)
    else:
        (DEST / "img").mkdir(exist_ok=True)
        (DEST / "img" / ".gitkeep").write_text("", encoding="utf-8")

    pages = {
        "index.html": page_index(),
        "maquinaria.html": page_maquinaria(),
        "servicios.html": page_servicios(),
        "equipo.html": page_equipo(),
        "reglas.html": page_reglas(),
        "apuntate.html": page_apuntate(),
        "login.html": page_login(),
        "calculadora.html": page_calculadora(),
        "reservas.html": page_reservas(),
        "buscar.html": page_buscar(),
        "tienda.html": page_tienda(),
        "tramitar-pedido.html": page_tramitar(),
    }
    for name, html in pages.items():
        (DEST / name).write_text(html, encoding="utf-8")

    for m in MAQUINAS:
        (DEST / f'maquinaria-detalle-{m["id"]}.html').write_text(page_maquinaria_detalle(m), encoding="utf-8")
    for s in SERVICIOS:
        (DEST / f'servicio-detalle-{s["id"]}.html').write_text(page_servicio_detalle(s), encoding="utf-8")

    (DEST / "README.md").write_text(README, encoding="utf-8")
    _patch_asset_paths()
    print("OK:", DEST)


def _patch_asset_paths() -> None:
    """Rutas /img/ absolutas → relativas para abrir el sitio desde disco o raíz distinta."""
    for p in (DEST / "css").glob("*.css"):
        t = p.read_text(encoding="utf-8")
        p.write_text(t.replace('url("/img/', 'url("../img/'), encoding="utf-8")
    h = DEST / "scss" / "layout" / "_header.scss"
    if h.exists():
        t = h.read_text(encoding="utf-8")
        h.write_text(
            t.replace("url('/img/logo_platinium.png')", "url('../../img/logo_platinium.png')"),
            encoding="utf-8",
        )
    hero = DEST / "scss" / "components" / "_hero.scss"
    if hero.exists():
        t = hero.read_text(encoding="utf-8")
        hero.write_text(t.replace("url('../img/banner.png')", "url('../../img/banner.png')"), encoding="utf-8")


if __name__ == "__main__":
    main()
