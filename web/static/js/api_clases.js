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

    claseForm.addEventListener('submit', (e) => {
        e.preventDefault();
        const id = document.getElementById('claseId').value;
        const claseData = {
            nombre_clase: document.getElementById('nombre').value,
            entrenador: document.getElementById('entrenador').value,
            aforo: parseInt(document.getElementById('aforo').value),
            horario: document.getElementById('horario').value,
            descripcion: document.getElementById('descripcion').value,
            lugar: document.getElementById('lugar').value
        };
        if (id) {
            actualizarClase(id, claseData);
        } else {
            crearClase(claseData);
        }
    });

    btnCancelar.addEventListener('click', limpiarFormulario);

    function cargarClases() {
        clasesList.innerHTML = '<div class="loading-state"><i class="fas fa-circle-notch"></i> Cargando clases...</div>';

        fetch(API_URL)
            .then(response => {
                if (!response.ok) throw new Error('Error al cargar las clases');
                return response.json();
            })
            .then(data => {
                clasesList.innerHTML = '';

                if (!data || data.length === 0) {
                    clasesList.innerHTML = '<div class="empty-state"><i class="fas fa-calendar-times"></i><p>No hay clases registradas aún.</p></div>';
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
            })
            .catch(error => {
                clasesList.innerHTML = `<div class="msg-danger">${error.message}</div>`;
            });
    }

    function crearClase(claseData) {
        fetch(API_URL, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(claseData)
        })
        .then(response => {
            if (!response.ok) throw new Error('Error al crear la clase');
            return response.json();
        })
        .then(() => {
            mostrarMensaje('Clase creada con éxito', 'success');
            limpiarFormulario();
            cargarClases();
        })
        .catch(error => mostrarMensaje(error.message, 'danger'));
    }

    function actualizarClase(id, claseData) {
        fetch(`${API_URL}/${id}`, {
            method: 'PUT',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(claseData)
        })
        .then(response => {
            if (!response.ok) throw new Error('Error al actualizar la clase');
            return response.json();
        })
        .then(() => {
            mostrarMensaje('Clase actualizada con éxito', 'success');
            limpiarFormulario();
            cargarClases();
        })
        .catch(error => mostrarMensaje(error.message, 'danger'));
    }

    function eliminarClase(id) {
        if (!confirm('¿Eliminar esta clase? Esta acción no se puede deshacer.')) return;
        fetch(`${API_URL}/${id}`, { method: 'DELETE' })
            .then(response => {
                if (!response.ok) throw new Error('Error al eliminar la clase');
                mostrarMensaje('Clase eliminada', 'success');
                cargarClases();
            })
            .catch(error => mostrarMensaje(error.message, 'danger'));
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
        btnGuardar.innerHTML = '<i class="fas fa-save"></i> Actualizar Clase';
        btnCancelar.style.display = 'inline-block';
        window.scrollTo({ top: 0, behavior: 'smooth' });
    }

    function limpiarFormulario() {
        claseForm.reset();
        document.getElementById('claseId').value = '';
        formTitle.textContent = 'Nueva Clase';
        btnGuardar.innerHTML = '<i class="fas fa-save"></i> Guardar Clase';
        btnCancelar.style.display = 'none';
        formMessage.innerHTML = '';
    }

    function mostrarMensaje(mensaje, tipo) {
        formMessage.innerHTML = `<div class="msg-${tipo}">${tipo === 'success' ? '✓' : '✗'} ${mensaje}</div>`;
        setTimeout(() => { formMessage.innerHTML = ''; }, 3000);
    }
});
