document.addEventListener('DOMContentLoaded', () => {
    const API_URL = '/api/clases';
    const clasesList = document.getElementById('clasesList');
    const claseTemplate = document.getElementById('claseTemplate');
    const claseForm = document.getElementById('claseForm');
    const btnCancelar = document.getElementById('btnCancelar');
    const btnGuardar = document.getElementById('btnGuardar');
    const formTitle = document.getElementById('formTitle');
    const formMessage = document.getElementById('formMessage');
    const countBadge = document.getElementById('countBadge');

    cargarClases();

    function obtenerDatosFormulario() {
        const nombre_clase = document.getElementById('nombre').value.trim();
        const entrenador = document.getElementById('entrenador').value.trim();
        const aforoRaw = document.getElementById('aforo').value.trim();
        const horario = document.getElementById('horario').value.trim();
        const descripcion = document.getElementById('descripcion').value.trim();
        const lugar = document.getElementById('lugar').value.trim();

        if (!nombre_clase || !entrenador || !horario || !lugar) {
            mostrarMensaje('Datos inválidos o incompletos', 'danger');
            return null;
        }
        if (aforoRaw === '') {
            mostrarMensaje('Datos inválidos o incompletos', 'danger');
            return null;
        }
        const aforo = Number.parseInt(aforoRaw, 10);
        if (!Number.isInteger(aforo) || aforo <= 0) {
            mostrarMensaje('Datos inválidos o incompletos', 'danger');
            return null;
        }

        return { nombre_clase, entrenador, aforo, horario, descripcion, lugar };
    }

    claseForm.addEventListener('submit', (e) => {
        e.preventDefault();
        const claseData = obtenerDatosFormulario();
        if (!claseData) {
            return;
        }
        const id = document.getElementById('claseId').value;
        if (id) {
            actualizarClase(id, claseData);
        } else {
            crearClase(claseData);
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
        clearElement(clasesList);
        const div = document.createElement('motion');
        div.className = 'msg-danger';
        div.textContent = mensaje;
        clasesList.appendChild(div);
    }

    function mostrarEstadoCarga() {
        clearElement(clasesList);
        const div = document.createElement('motion');
        div.className = 'loading-state';
        const icon = document.createElement('i');
        icon.className = 'fas fa-circle-notch';
        div.appendChild(icon);
        div.appendChild(document.createTextNode(' Cargando clases...'));
        clasesList.appendChild(div);
    }

    function mostrarEstadoVacio() {
        clearElement(clasesList);
        const div = document.createElement('motion');
        div.className = 'empty-state';
        const icon = document.createElement('i');
        icon.className = 'fas fa-calendar-times';
        const p = document.createElement('p');
        p.textContent = 'No hay clases registradas aún.';
        div.appendChild(icon);
        div.appendChild(p);
        clasesList.appendChild(div);
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

    async function cargarClases() {
        mostrarEstadoCarga();
        try {
            const response = await fetch(API_URL);
            if (!response.ok) {
                const msg = await mensajeErrorRespuesta(response, 'Error al cargar las clases');
                throw new Error(msg);
            }
            const data = await response.json();
            clearElement(clasesList);

            if (!data || data.length === 0) {
                mostrarEstadoVacio();
                countBadge.textContent = '0';
                return;
            }

            countBadge.textContent = data.length;
            data.forEach(clase => {
                const clone = claseTemplate.content.cloneNode(true);
                clone.querySelector('.clase-nombre').textContent = clase.nombre_clase;
                clone.querySelector('.clase-entrenador').textContent = clase.entrenador;
                clone.querySelector('.clase-horario').textContent = clase.horario;
                clone.querySelector('.clase-aforo').textContent = clase.aforo;
                clone.querySelector('.clase-lugar').textContent = clase.lugar;
                clone.querySelector('.btn-editar').addEventListener('click', () => llenarFormulario(clase));
                clone.querySelector('.btn-eliminar').addEventListener('click', () => eliminarClase(clase.id));
                clasesList.appendChild(clone);
            });
        } catch (error) {
            mostrarErrorLista(error.message);
        }
    }

    async function crearClase(claseData) {
        btnGuardar.disabled = true;
        try {
            const response = await fetch(API_URL, {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify(claseData)
            });
            await validarRespuestaFormulario(response, 'Error al crear la clase');
            await response.json();
            mostrarMensaje('Clase creada con éxito', 'success');
            limpiarFormulario();
            await cargarClases();
        } catch (error) {
            mostrarMensaje(error.message, 'danger');
        } finally {
            btnGuardar.disabled = false;
        }
    }

    async function actualizarClase(id, claseData) {
        btnGuardar.disabled = true;
        try {
            const response = await fetch(`${API_URL}/${id}`, {
                method: 'PUT',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify(claseData)
            });
            await validarRespuestaFormulario(response, 'Error al actualizar la clase');
            await response.json();
            mostrarMensaje('Clase actualizada con éxito', 'success');
            limpiarFormulario();
            await cargarClases();
        } catch (error) {
            mostrarMensaje(error.message, 'danger');
        } finally {
            btnGuardar.disabled = false;
        }
    }

    async function eliminarClase(id) {
        if (!confirm('¿Eliminar esta clase? Esta acción no se puede deshacer.')) {
            return;
        }
        try {
            const response = await fetch(`${API_URL}/${id}`, { method: 'DELETE' });
            if (!response.ok) {
                const msg = await mensajeErrorRespuesta(response, 'Error al eliminar la clase');
                throw new Error(msg);
            }
            mostrarMensaje('Clase eliminada', 'success');
            await cargarClases();
        } catch (error) {
            mostrarMensaje(error.message, 'danger');
        }
    }

    function llenarFormulario(clase) {
        document.getElementById('claseId').value = clase.id;
        document.getElementById('nombre').value = clase.nombre_clase;
        document.getElementById('entrenador').value = clase.entrenador;
        document.getElementById('aforo').value = clase.aforo;
        document.getElementById('horario').value = clase.horario;
        document.getElementById('descripcion').value = clase.descripcion;
        document.getElementById('lugar').value = clase.lugar;
        formTitle.textContent = 'Editar Clase';
        setGuardarLabel('Actualizar Clase');
        btnCancelar.classList.remove('hidden');
        window.scrollTo({ top: 0, behavior: 'smooth' });
    }

    function limpiarFormulario() {
        claseForm.reset();
        document.getElementById('claseId').value = '';
        formTitle.textContent = 'Nueva Clase';
        setGuardarLabel('Guardar Clase');
        btnCancelar.classList.add('hidden');
        clearElement(formMessage);
    }
});
