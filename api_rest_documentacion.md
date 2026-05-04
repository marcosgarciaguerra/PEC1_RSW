# Documentación para la Memoria (PEC 3)

*Copia este contenido a tu documento Word o PDF de la memoria de la PEC 3.*

---

## 2. Creación de API REST

Para la implementación de la API REST se ha desarrollado un sistema CRUD (Crear, Leer, Actualizar, Eliminar) para la gestión del inventario de **Artículos** de la tienda. Esta API permite que la web solicite y modifique la información de los productos en la base de datos de manera asíncrona mediante peticiones HTTP.

A continuación se detalla la especificación de la API REST implementada:

### Endpoint: `/api/articulos`

#### 1. Obtener todos los artículos
- **Ruta de acceso:** `/api/articulos`
- **Método:** `GET`
- **Códigos de estado:** 
  - `200 OK`: Éxito al devolver los artículos.
  - `500 Internal Server Error`: Error al acceder a la base de datos.
- **Descripción:** Obtiene una lista en formato JSON con todos los artículos almacenados en el sistema (suplementos, ropa, accesorios).

#### 2. Obtener un artículo específico
- **Ruta de acceso:** `/api/articulos/{id}`
- **Método:** `GET`
- **Códigos de estado:** 
  - `200 OK`: Éxito, devuelve el objeto JSON del artículo.
  - `400 Bad Request`: El ID proporcionado no es válido.
  - `404 Not Found`: No existe un artículo con ese ID.
- **Descripción:** Obtiene los detalles de un artículo específico identificando su ID único en la ruta.

#### 3. Crear un nuevo artículo
- **Ruta de acceso:** `/api/articulos`
- **Método:** `POST`
- **Códigos de estado:** 
  - `201 Created`: Artículo creado satisfactoriamente.
  - `400 Bad Request`: El cuerpo de la petición no contiene un JSON válido.
  - `500 Internal Server Error`: Error interno al guardar en la base de datos.
- **Descripción:** Crea un nuevo artículo en el inventario. Se debe enviar un objeto JSON en el cuerpo de la petición con los campos requeridos (`nombre`, `categoria`, `precio`, `imagen`, `etiqueta`).

#### 4. Actualizar un artículo existente
- **Ruta de acceso:** `/api/articulos/{id}`
- **Método:** `PUT`
- **Códigos de estado:** 
  - `200 OK`: Artículo actualizado correctamente.
  - `400 Bad Request`: El JSON es inválido o el ID no es correcto.
  - `405 Method Not Allowed`: No se ha especificado el ID en la URL.
  - `500 Internal Server Error`: Fallo al actualizar el registro.
- **Descripción:** Sobrescribe la información de un artículo existente por la provista en el cuerpo de la petición en formato JSON.

#### 5. Eliminar un artículo
- **Ruta de acceso:** `/api/articulos/{id}`
- **Método:** `DELETE`
- **Códigos de estado:** 
  - `200 OK`: Artículo eliminado correctamente.
  - `405 Method Not Allowed`: No se ha especificado el ID en la URL.
  - `500 Internal Server Error`: Fallo al borrar el registro.
- **Descripción:** Elimina permanentemente del inventario el artículo indicado en la URL.

---

## 3. Explotación de la API

La explotación de la API desde la web se realiza en la página de administración de artículos (`admin_articulos.html`).
- **Promesas:** Se ha utilizado la API nativa de JavaScript `fetch`, la cual devuelve promesas (`.then()`, `.catch()`) para procesar las respuestas del servidor de forma asíncrona y actualizar la interfaz sin recargar la página.
- **Plantillas HTML (`<template>`):** Para renderizar la tabla de artículos dinámicamente, se hace uso de una etiqueta `<template id="articulo-row-template">`. Mediante JavaScript, esta plantilla es clonada y sus valores internos (`textContent`, `src`) son poblados en base a la información extraída de la API REST antes de insertar los nodos en el DOM.
- Todo esto se ha logrado usando única y exclusivamente **JavaScript puro** (Vanilla JS), absteniéndose de usar librerías externas como jQuery, Axios, etc.

---

## 4. Adaptaciones a partir de la PEC anterior

- **Análisis de requisitos:** Se ha modificado el sistema para añadir un panel de administración del inventario. El rol de administrador (o sistema) requiere la capacidad de realizar operaciones de creación, lectura, actualización y borrado (CRUD) sobre los productos de la tienda sin depender de una base de datos estática incrustada en el código y logrando una fluidez dinámica.
- **Diseño de la aplicación (UML):** En el diagrama UML se ha de incluir una interfaz `APIArticulos` o `Controlador REST` con los métodos (`GET`, `POST`, `PUT`, `DELETE`) que se comunican asíncronamente con el `Gestor de Base de Datos SQLite` a través de Promesas lanzadas por el Cliente.
- **Diseño de la Interfaz:** Se ha creado una nueva vista HTML (`admin_articulos.html`) que integra un formulario y una tabla. La manipulación visual se realiza de forma limpia mediante JS modificando contenido o alterando el atributo `display` / inyectando clases de CSS.
- **Implementación:** El servidor de Golang fue adaptado con un nuevo enrutador de multiplexación para separar las rutas HTML tradicionales de las rutas de la API bajo el prefijo `/api/...`, centralizando la comunicación en respuestas en formato JSON a través del paquete `encoding/json`.

---

## 5. Bibliografía y Uso de IA
*(No olvides añadir los enlaces de ChatGPT / Gemini o la IA que hayas consultado, libros o apuntes de la asignatura en este apartado, de lo contrario la nota máxima será de un 5).*
