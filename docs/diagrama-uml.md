# Diagrama UML Del Proyecto

Este diagrama representa la arquitectura de clases y paquetes principales de Platinium Gym. El proyecto está organizado en capas: controladores HTTP (`handlers`), lógica de negocio (`services`), acceso a datos (`db`) y entidades (`models`).

```mermaid
classDiagram
direction LR

namespace models {
    class Usuario {
        +int ID
        +string Nombre
        +string Apellidos
        +string FechaNacimiento
        +string Direccion
        +string Telefono
        +string Correo
        +string Documento
        +string Plan
        +string MetodoPago
        +string NumeroPago
        +string Password
    }

    class Socio {
        +int ID
        +string Nombre
        +string Contrasena
        +bool SuscripcionActiva
    }

    class Clases {
        +int ID
        +string NombreClase
        +string Entrenador
        +int Aforo
        +string Horario
        +string Descripcion
        +string Lugar
    }

    class Reserva {
        +int ID
        +int SocioID
        +int ActividadID
        +string FechaAsist
    }

    class Articulo {
        +int ID
        +string Nombre
        +string Categoria
        +float64 Precio
        +string Imagen
        +string Etiqueta
    }

    class Pedido {
        +int ID
        +int UsuarioID
        +string Fecha
        +string DireccionEnvio
        +string MetodoPago
        +float64 Total
        +string Estado
        +PedidoItem[] Items
    }

    class PedidoItem {
        +int ID
        +int PedidoID
        +int ArticuloID
        +string Nombre
        +float64 PrecioUnitario
        +int Cantidad
        +float64 Subtotal
    }

    class Maquina {
        +int ID
        +string Nombre
        +string Marca
        +string Zona
        +string Descripcion
        +string Beneficios
        +string Imagen
    }

    class Servicio {
        +int ID
        +string Nombre
        +string Icono
        +string DescripcionBreve
        +string DescripcionLarga
        +string Beneficios
    }

    class Resena {
        +string Autor
        +int Puntuacion
        +string Texto
    }
}

namespace handlers {
    class VistasHandler {
        +IndexHandler()
        +TiendaHandler()
        +MaquinariaHandler()
        +ServiciosHandler()
        +EquipoHandler()
        +ReglasHandler()
        +ApuntateHandler()
    }

    class RegistroHandler {
        +RegistroHandler()
    }

    class LoginHandler {
        +LoginHandler()
        +LogoutHandler()
    }

    class ReservasHandler {
        +ReservasHandler()
        +ProcesarReservaHandler()
        +ProcesarCancelacionHandler()
    }

    class TiendaHandler {
        +TramitarPedidoHandler()
        +ConfirmarPedidoHandler()
    }

    class RenderTemplate {
        +InitTemplates()
        +RenderTemplate()
    }
}

namespace services {
    class RegistroService {
        +RegistrarUsuario(usuario, repetirPassword) error
    }

    class ReservaService {
        +ObtenerReservasDetalleSocio(socioID) 
        +CrearReservaSocio(socioID, claseID) bool
        +CancelarReservaSocio(reservaID, socioID) bool
    }

    class TiendaService {
        +CatalogoTienda() Articulo[]
        +CrearPedido(usuario, direccionEnvio, carritoJSON) Pedido
    }

    class ReservaDetalle {
        +Reserva Reserva
        +string NombreClase
    }
}

namespace db {
    class RepositorioSQLite {
        +InitDB()
        +GuardarUsuario(usuario) error
        +ObtenerUsuarioPorCorreo(correo) Usuario
        +ObtenerSocioPorCorreo(correo) Socio
        +ObtenerClases() Clases[]
        +CrearReserva(socioID, actividadID, fecha) bool
        +EliminarReserva(reservaID, socioID) bool
        +GuardarPedido(pedido) error
        +ObtenerResenas() Resena[]
        +GuardarResena(resena) error
    }

    class CatalogoDatos {
        +Maquinas Maquina[]
        +Servicios Servicio[]
        +ArticulosTienda Articulo[]
        +ObtenerArticuloPorID(id) Articulo
    }
}

Pedido "1" *-- "1..*" PedidoItem : contiene
Usuario "1" --> "0..*" Pedido : realiza
Socio "1" --> "0..*" Reserva : crea
Clases "1" --> "0..*" Reserva : recibe
Articulo "1" --> "0..*" PedidoItem : referencia
ReservaDetalle *-- Reserva : extiende

VistasHandler ..> RenderTemplate : renderiza
VistasHandler ..> CatalogoDatos : consulta
RegistroHandler ..> RegistroService : delega
LoginHandler ..> RepositorioSQLite : autentica
ReservasHandler ..> ReservaService : delega
TiendaHandler ..> TiendaService : delega

RegistroService ..> RepositorioSQLite : guarda usuario
ReservaService ..> RepositorioSQLite : gestiona reservas
TiendaService ..> RepositorioSQLite : guarda pedido
TiendaService ..> CatalogoDatos : valida catalogo

RepositorioSQLite ..> Usuario : persiste
RepositorioSQLite ..> Socio : consulta
RepositorioSQLite ..> Reserva : persiste
RepositorioSQLite ..> Pedido : persiste
RepositorioSQLite ..> Resena : persiste
CatalogoDatos ..> Maquina : carga
CatalogoDatos ..> Servicio : carga
CatalogoDatos ..> Articulo : carga
```

## Lectura Del Diagrama

- `handlers` recibe las peticiones HTTP y decide qué vista renderizar o qué acción ejecutar.
- `services` concentra la lógica de negocio: validación de registro, cálculo de reservas y creación de pedidos.
- `db` encapsula SQLite y los catálogos iniciales de maquinaria, servicios y tienda.
- `models` contiene las entidades usadas por las plantillas, servicios y base de datos.
