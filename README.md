# 🏋️‍♂️ Platinium Gym - Servidor Web Dinámico en Go

Este proyecto es la resolución de la **PEC 2** de la asignatura **Redes y Sistemas Web (RSW)** de la Universidad Camilo José Cela. Consiste en la migración de un sitio web estático (desarrollado en la PEC 1) a una aplicación web dinámica completa, construida con un servidor backend en **Golang**.

## 🚀 Características Principales

* **Servidor HTTP en Go:** Implementación de un servidor web desde cero sin depender de frameworks pesados, aplicando las mejores prácticas de la comunidad.
* **Arquitectura de Capas:** Código altamente estructurado separando las responsabilidades en Presentación (handlers/vistas), Lógica de Negocio (services) y Datos (db).
* **Renderizado Dinámico:** Uso del paquete `html/template` para inyectar datos dinámicos en las vistas HTML.
* **Procesamiento de Formularios:** Captura, validación y procesamiento de datos provenientes del cliente (login, registro de socios, reseñas, buscador y calculadora).
* **Persistencia de Datos:** Almacenamiento seguro y persistente de la información utilizando una base de datos local **SQLite** (`platinum.db`).
* **Diseño Responsivo:** Interfaz construida con **SCSS** modular, asegurando adaptabilidad a dispositivos móviles y de escritorio.

## 🛠️ Stack Tecnológico

* **Backend:** Go (Golang)
* **Frontend:** HTML5, CSS3 (mediante SCSS)
* **Base de Datos:** SQLite 3
* **Arquitectura:** Patrón MVC (Modelo-Vista-Controlador) adaptado a las convenciones de Go (`internal/`, `cmd/servidor/`).

## 📂 Estructura del Proyecto

El proyecto sigue el estándar de distribución de directorios de Go:

```text
├── go.mod                  # Archivo de definición del módulo de Go y dependencias
├── go.sum                  # Hashes de seguridad de las dependencias
├── servidor/               # Punto de entrada de la aplicación
│   └── main.go             # Inicialización del servidor, rutas y conexión a BD
├── internal/               # Código privado de la aplicación (Lógica central)
│   ├── db/                 # Capa de acceso a datos y base de datos SQLite
│   │   ├── platinum.db     # Archivo de base de datos SQLite
│   │   └── ...             # Repositorios (almacenamiento, datos_equipo, etc.)
│   ├── handlers/           # Capa de presentación (Controladores HTTP)
│   │   ├── vistas.go       # Carga de páginas generales
│   │   ├── formulario.go   # Procesamiento de formularios
│   │   └── ...             # Handlers específicos (login, tienda, reservas, etc.)
│   ├── models/             # Estructuras de datos (Entidades del dominio)
│   │   └── usuario.go, maquina.go, servicio.go, etc.
│   └── services/           # Capa de lógica de negocio
│       └── calculadora.go, registro.go
└── web/                    # Archivos del Frontend
    ├── static/             # Archivos estáticos (públicos)
    │   ├── css/            # Hojas de estilo compiladas
    │   ├── scss/           # Código fuente modular de estilos
    │   └── img/            # Imágenes y recursos multimedia (equipo, maquinaria, tienda)
    └── templates/          # Plantillas HTML dinámicas
        ├── layout.html     # Plantilla base (Header, Footer, Meta tags)
        ├── index.html      # Página principal
        └── ...             # Vistas de contenido (tienda, reservas, login, etc.)
```

## ⚙️ Requisitos Previos

Para ejecutar este proyecto en tu máquina local, necesitas tener instalado:

* [Go (versión 1.20 o superior)](https://golang.org/doc/install)
* Un navegador web moderno.
* *(Opcional)* Compilador de SCSS/SASS si deseas modificar los estilos visuales.

## 🏃‍♂️ Instalación y Ejecución

1. **Clonar el repositorio:**
   ```bash
   git clone https://github.com/MarkQWERTY/PEC1_RSW
   cd PEC1_RSW
   ```

2. **Descargar las dependencias:**
   ```bash
   go mod download
   ```

3. **Ejecutar el servidor:**
   El punto de entrada de la aplicación se encuentra en la carpeta `servidor`.
   ```bash
   go run servidor/main.go
   ```

4. **Acceder a la aplicación:**
   Abre tu navegador web y navega a:
   ```text
   http://localhost:8080
   ```
   *(Nota: Si configuraste otro puerto en main.go, utiliza ese).*

## 📖 Rutas Disponibles (Vistas Principales)

* `/` - Inicio (Página principal de Platinium Gym)
* `/equipo` - Información sobre los entrenadores (Andoni, Cbum, Arnold, etc.)
* `/maquinaria` - Catálogo de máquinas del gimnasio
* `/tienda` - Venta de suplementos y ropa (Creatina, pre-entrenos, straps)
* `/reservas` - Formulario para reserva de clases/pistas
* `/calculadora` - Herramienta de cálculo de métricas fitness
* `/login` - Acceso y registro de usuarios
* `/reglas` - Normativa interna del gimnasio

## 📝 Notas de Desarrollo (PEC 2)

* **Cumplimiento de rúbrica:** Todos los requisitos obligatorios como la separación de lógica, comentarios explicativos en funciones exportadas, y almacenamiento persistente (en este caso SQLite) han sido integrados.
* Se ha empleado el paquete nativo `html/template` en lugar de servir archivos estáticos crudos para permitir el paso de estructuras de datos a las vistas.

## 👥 Autores

* **Grupo XX** - RSW
* *Marcos García Guerra*
* *Eric Specht de la Torre*
```

