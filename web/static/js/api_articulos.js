document.addEventListener('DOMContentLoaded', () => {
    const API_URL = '/api/articulos';
    const articulosList = document.getElementById('articulosList');
    const articuloTemplate = document.getElementById('articuloTemplate');
    const articuloForm = document.getElementById('articuloForm');
    const btnCancelar = document.getElementById('btnCancelar');
    const btnGuardar = document.getElementById('btnGuardar');
    const formTitle = document.getElementById('formTitle');
    const formMessage = document.getElementById('formMessage');
    const countBadge = document.getElementById('countBadge');

    const CAT_NAMES = { sup: 'Suplemento', rop: 'Ropa', acc: 'Accesorio' };

    cargarArticulos();

    articuloForm.addEventListener('submit', (e) => {
        e.preventDefault();
        const id = document.getElementById('articuloId').value;
        const articuloData = {
            nombre: document.getElementById('nombre').value,
            categoria: document.getElementById('categoria').value,
            precio: parseFloat(document.getElementById('precio').value),
            imagen: document.getElementById('imagen').value,
            etiqueta: document.getElementById('etiqueta').value
        };
        if (id) {
            actualizarArticulo(id, articuloData);
        } else {
            crearArticulo(articuloData);
        }
    });

    btnCancelar.addEventListener('click', limpiarFormulario);

    function clearElement(el) {
        while (el.firstChild) {
            el.removeChild(el.firstChild);
        }
    }

    function setGuardarLabel(texto) {
        const icon = btnGuardar.querySelector('i');
        clearElement(btnGuardar);
        if (icon) {
            btnGuardar.appendChild(icon);
        }
        btnGuardar.appendChild(document.createTextNode(' ' + texto));
    }

    function mostrarMensaje(mensaje, tipo) {
        clearElement(formMessage);
        const div = document.createElement('motion');
        div.className = 'msg-' + tipo;
        div.textContent = (tipo === 'success' ? '✓ ' : '✗ ') + mensaje;
        formMessage.appendChild(div);
        setTimeout(() => clearElement(formMessage), 3000);
    }

    function mostrarErrorLista(mensaje) {
        clearElement(articulosList);
        const div = document.createElement('motion');
        div.className = 'msg-danger grid-full';
        div.textContent = mensaje;
        articulosList.appendChild(div);
    }

    function mostrarEstadoCarga() {
        clearElement(articulosList);
        const div = document.createElement('motion');
        div.className = 'loading-state grid-full';
        const icon = document.createElement('i');
        icon.className = 'fas fa-circle-notch';
        div.appendChild(icon);
        div.appendChild(document.createTextNode(' Cargando artículos...'));
        articulosList.appendChild(div);
    }

    function mostrarEstadoVacio() {
        clearElement(articulosList);
        const div = document.createElement('motion');
        div.className = 'empty-state';
        const icon = document.createElement('i');
        icon.className = 'fas fa-box-open';
        const p = document.createElement('p');
        p.textContent = 'No hay artículos en el catálogo.';
        div.appendChild(icon);
        div.appendChild(p);
        articulosList.appendChild(div);
    }

    async function mensajeErrorRespuesta(response, fallback) {
        try {
            const data = await response.json();
            if (data && data.error) {
                return data.error;
            }
        } catch (_) {
            /* cuerpo no JSON */
        }
        return fallback;
    }

    /** Valida respuestas POST/PUT: 422 → mensaje de validación para el usuario. */
    async function validarRespuestaFormulario(response, fallback) {
        if (response.status === 422) {
            throw new Error('Datos inválidos o incompletos');
        }
        if (!response.ok) {
            const msg = await mensajeErrorRespuesta(response, fallback);
            throw new Error(msg);
        }
    }

    async function cargarArticulos() {
        mostrarEstadoCarga();
        try {
            const response = await fetch(API_URL);
            if (!response.ok) {
                const msg = await mensajeErrorRespuesta(response, 'Error al cargar los artículos');
                throw new Error(msg);
            }
            const data = await response.json();
            clearElement(articulosList);

            if (!data || data.length === 0) {
                mostrarEstadoVacio();
                countBadge.textContent = '0';
                return;
            }

            countBadge.textContent = data.length;
            data.forEach(articulo => {
                const clone = articuloTemplate.content.cloneNode(true);

                clone.querySelector('.articulo-imagen').src = articulo.imagen;
                clone.querySelector('.articulo-imagen').alt = articulo.nombre;
                clone.querySelector('.articulo-nombre').textContent = articulo.nombre;
                clone.querySelector('.articulo-precio').textContent = articulo.precio.toFixed(2);
                clone.querySelector('.articulo-categoria').textContent = CAT_NAMES[articulo.categoria] || articulo.categoria;

                const tagEl = clone.querySelector('.articulo-etiqueta');
                if (articulo.etiqueta && articulo.etiqueta.trim() !== '') {
                    tagEl.textContent = articulo.etiqueta;
                    tagEl.classList.remove('hidden');
                    tagEl.classList.add('visible-inline');
                } else {
                    tagEl.classList.add('hidden');
                    tagEl.classList.remove('visible-inline');
                }

                clone.querySelector('.btn-editar').addEventListener('click', () => llenarFormulario(articulo));
                clone.querySelector('.btn-eliminar').addEventListener('click', () => eliminarArticulo(articulo.id));

                articulosList.appendChild(clone);
            });
        } catch (error) {
            mostrarErrorLista(error.message);
        }
    }

    async function crearArticulo(articuloData) {
        btnGuardar.disabled = true;
        try {
            const response = await fetch(API_URL, {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify(articuloData)
            });
            await validarRespuestaFormulario(response, 'Error al crear el artículo');
            await response.json();
            mostrarMensaje('Artículo creado con éxito', 'success');
            limpiarFormulario();
            await cargarArticulos();
        } catch (error) {
            mostrarMensaje(error.message, 'danger');
        } finally {
            btnGuardar.disabled = false;
        }
    }

    async function actualizarArticulo(id, articuloData) {
        btnGuardar.disabled = true;
        try {
            const response = await fetch(`${API_URL}/${id}`, {
                method: 'PUT',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify(articuloData)
            });
            await validarRespuestaFormulario(response, 'Error al actualizar el artículo');
            await response.json();
            mostrarMensaje('Artículo actualizado con éxito', 'success');
            limpiarFormulario();
            await cargarArticulos();
        } catch (error) {
            mostrarMensaje(error.message, 'danger');
        } finally {
            btnGuardar.disabled = false;
        }
    }

    async function eliminarArticulo(id) {
        if (!confirm('¿Eliminar este artículo? Esta acción no se puede deshacer.')) {
            return;
        }
        try {
            const response = await fetch(`${API_URL}/${id}`, { method: 'DELETE' });
            if (!response.ok) {
                const msg = await mensajeErrorRespuesta(response, 'Error al eliminar el artículo');
                throw new Error(msg);
            }
            mostrarMensaje('Artículo eliminado', 'success');
            await cargarArticulos();
        } catch (error) {
            mostrarMensaje(error.message, 'danger');
        }
    }

    function llenarFormulario(articulo) {
        document.getElementById('articuloId').value = articulo.id;
        document.getElementById('nombre').value = articulo.nombre;
        document.getElementById('categoria').value = articulo.categoria;
        document.getElementById('precio').value = articulo.precio;
        document.getElementById('imagen').value = articulo.imagen;
        document.getElementById('etiqueta').value = articulo.etiqueta || '';
        formTitle.textContent = 'Editar Artículo';
        setGuardarLabel('Actualizar Artículo');
        btnCancelar.classList.remove('hidden');
        window.scrollTo({ top: 0, behavior: 'smooth' });
    }

    function limpiarFormulario() {
        articuloForm.reset();
        document.getElementById('articuloId').value = '';
        formTitle.textContent = 'Nuevo Artículo';
        setGuardarLabel('Guardar Artículo');
        btnCancelar.classList.add('hidden');
        clearElement(formMessage);
    }
});
