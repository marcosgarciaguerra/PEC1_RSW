document.addEventListener('DOMContentLoaded', () => {
    const API_URL = '/api/maquinarias';
    const maquinariasList = document.getElementById('maquinariasList');
    const maquinariaTemplate = document.getElementById('maquinariaTemplate');
    const maquinariaForm = document.getElementById('maquinariaForm');
    const btnCancelar = document.getElementById('btnCancelar');
    const btnGuardar = document.getElementById('btnGuardar');
    const formTitle = document.getElementById('formTitle');
    const formMessage = document.getElementById('formMessage');
    const countBadge = document.getElementById('countBadge');

    cargarMaquinarias();

    function obtenerDatosFormulario() {
        const nombre = document.getElementById('nombre').value.trim();
        const marca = document.getElementById('marca').value.trim();
        const zona = document.getElementById('zona').value.trim();
        const descripcion = document.getElementById('descripcion').value.trim();
        const beneficios = document.getElementById('beneficios').value.trim();
        const imagen = document.getElementById('imagen').value.trim();

        if (!nombre || !marca || !zona) {
            mostrarMensaje('Datos inválidos o incompletos', 'danger');
            return null;
        }

        return { nombre, marca, zona, descripcion, beneficios, imagen };
    }

    maquinariaForm.addEventListener('submit', (e) => {
        e.preventDefault();
        const data = obtenerDatosFormulario();
        if (!data) {
            return;
        }
        const id = document.getElementById('maquinariaId').value;
        if (id) {
            actualizarMaquinaria(id, data);
        } else {
            crearMaquinaria(data);
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
        clearElement(maquinariasList);
        const div = document.createElement('motion');
        div.className = 'msg-danger grid-full';
        div.textContent = mensaje;
        maquinariasList.appendChild(div);
    }

    function mostrarEstadoCarga() {
        clearElement(maquinariasList);
        const div = document.createElement('motion');
        div.className = 'loading-state grid-full';
        const icon = document.createElement('i');
        icon.className = 'fas fa-circle-notch';
        div.appendChild(icon);
        div.appendChild(document.createTextNode(' Cargando maquinarias...'));
        maquinariasList.appendChild(div);
    }

    function mostrarEstadoVacio() {
        clearElement(maquinariasList);
        const div = document.createElement('motion');
        div.className = 'empty-state';
        const icon = document.createElement('i');
        icon.className = 'fas fa-dumbbell';
        const p = document.createElement('p');
        p.textContent = 'No hay maquinarias registradas aún.';
        div.appendChild(icon);
        div.appendChild(p);
        maquinariasList.appendChild(div);
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

    async function validarRespuestaFormulario(response, fallback) {
        if (response.status === 422) {
            throw new Error('Datos inválidos o incompletos');
        }
        if (!response.ok) {
            const msg = await mensajeErrorRespuesta(response, fallback);
            throw new Error(msg);
        }
    }

    async function cargarMaquinarias() {
        mostrarEstadoCarga();
        try {
            const response = await fetch(API_URL);
            if (!response.ok) {
                const msg = await mensajeErrorRespuesta(response, 'Error al cargar la maquinaria');
                throw new Error(msg);
            }
            const data = await response.json();
            clearElement(maquinariasList);

            if (!data || data.length === 0) {
                mostrarEstadoVacio();
                countBadge.textContent = '0';
                return;
            }

            countBadge.textContent = data.length;
            data.forEach(m => {
                const clone = maquinariaTemplate.content.cloneNode(true);
                const img = clone.querySelector('.maquinaria-imagen');
                img.src = m.imagen || '/img/placeholder.jpg';
                img.alt = m.nombre;

                clone.querySelector('.maquinaria-zona').textContent = m.zona;
                clone.querySelector('.maquinaria-marca').textContent = m.marca;
                clone.querySelector('.maquinaria-nombre').textContent = m.nombre;

                clone.querySelector('.btn-editar').addEventListener('click', () => llenarFormulario(m));
                clone.querySelector('.btn-eliminar').addEventListener('click', () => eliminarMaquinaria(m.id));
                maquinariasList.appendChild(clone);
            });
        } catch (error) {
            mostrarErrorLista(error.message);
        }
    }

    async function crearMaquinaria(data) {
        btnGuardar.disabled = true;
        try {
            const response = await fetch(API_URL, {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify(data)
            });
            await validarRespuestaFormulario(response, 'Error al crear la maquinaria');
            await response.json();
            mostrarMensaje('Maquinaria creada con éxito', 'success');
            limpiarFormulario();
            await cargarMaquinarias();
        } catch (error) {
            mostrarMensaje(error.message, 'danger');
        } finally {
            btnGuardar.disabled = false;
        }
    }

    async function actualizarMaquinaria(id, data) {
        btnGuardar.disabled = true;
        try {
            const response = await fetch(`${API_URL}/${id}`, {
                method: 'PUT',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify(data)
            });
            await validarRespuestaFormulario(response, 'Error al actualizar la maquinaria');
            await response.json();
            mostrarMensaje('Maquinaria actualizada con éxito', 'success');
            limpiarFormulario();
            await cargarMaquinarias();
        } catch (error) {
            mostrarMensaje(error.message, 'danger');
        } finally {
            btnGuardar.disabled = false;
        }
    }

    async function eliminarMaquinaria(id) {
        if (!confirm('¿Eliminar esta máquina? Esta acción no se puede deshacer.')) {
            return;
        }
        try {
            const response = await fetch(`${API_URL}/${id}`, { method: 'DELETE' });
            if (!response.ok) {
                const msg = await mensajeErrorRespuesta(response, 'Error al eliminar la maquinaria');
                throw new Error(msg);
            }
            mostrarMensaje('Máquina eliminada', 'success');
            await cargarMaquinarias();
        } catch (error) {
            mostrarMensaje(error.message, 'danger');
        }
    }

    function llenarFormulario(m) {
        document.getElementById('maquinariaId').value = m.id;
        document.getElementById('nombre').value = m.nombre;
        document.getElementById('marca').value = m.marca;
        document.getElementById('zona').value = m.zona;
        document.getElementById('descripcion').value = m.descripcion;
        document.getElementById('beneficios').value = m.beneficios;
        document.getElementById('imagen').value = m.imagen;

        formTitle.textContent = 'Editar Máquina';
        setGuardarLabel('Actualizar Máquina');
        btnCancelar.classList.remove('hidden');
        window.scrollTo({ top: 0, behavior: 'smooth' });
    }

    function limpiarFormulario() {
        maquinariaForm.reset();
        document.getElementById('maquinariaId').value = '';
        formTitle.textContent = 'Nueva Máquina';
        setGuardarLabel('Guardar Máquina');
        btnCancelar.classList.add('hidden');
        clearElement(formMessage);
    }
});
