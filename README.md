# Platinium Gym - Servidor Web Dinámico en Go

Proyecto de la PEC 2 de Redes y Sistemas Web. Migra la web estática de Platinium Gym a una aplicación web dinámica con servidor HTTP en Go, plantillas HTML, formularios procesados en servidor y persistencia en SQLite.

## Características

- Servidor web en Go con `net/http`.
- Plantillas dinámicas con `html/template`.
- Arquitectura por capas:
  - `cmd/app`: punto de entrada y registro de rutas.
  - `internal/handlers`: controladores HTTP y renderizado de vistas.
  - `internal/services`: lógica de negocio y validaciones.
  - `internal/db`: acceso a datos, catálogo inicial y persistencia.
  - `internal/models`: entidades del dominio.
- Formularios procesados por Go: registro, login, reseñas, búsqueda, calculadora, reservas y confirmación de pedidos.
- Persistencia local en `internal/db/platinum.db` mediante SQLite.
- Interfaz responsive con HTML, CSS y SCSS.

## Estructura

```text
├── cmd/
│   └── app/
│       └── main.go
├── internal/
│   ├── db/
│   ├── handlers/
│   ├── models/
│   └── services/
├── web/
│   ├── static/
│   │   ├── css/
│   │   ├── img/
│   │   └── scss/
│   └── templates/
├── go.mod
└── go.sum
```

## Ejecución

Requisitos:

- Go instalado.
- Navegador web moderno.

Comandos:

```bash
go mod download
go run ./cmd/app
```

Después, abrir:

```text
http://localhost:8080
```

## Rutas Principales

- `/`: inicio con reseñas dinámicas.
- `/apuntate`: alta de socios y selección de plan.
- `/login`: inicio de sesión.
- `/reservas`: reserva y cancelación de clases para socios autenticados.
- `/tienda`: catálogo renderizado desde Go.
- `/tienda/tramitar`: revisión del carrito para usuarios autenticados.
- `/tienda/confirmar`: confirmación de pedido por `POST`, validación en servidor y guardado en SQLite.
- `/buscar`: buscador del sitio.
- `/calculadora`: cálculo de métricas fitness.

## Notas PEC 2

La aplicación separa presentación, lógica y datos siguiendo una estructura habitual en proyectos Go. La tienda no confía en precios enviados desde el cliente: el servidor recibe IDs y cantidades, recupera el catálogo definido en Go, calcula el total y guarda el pedido con sus líneas en SQLite.

## Autores

- Grupo 23 - RSW
- Marcos García Guerra
- Eric Specht de la Torre
