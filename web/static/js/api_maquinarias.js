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

    maquinariaForm.addEventListener('submit', (e) => {
        e.preventDefault();
        const id = document.getElementById('maquinariaId').value;
        const data = {
            nombre: document.getElementById('nombre').value,
            marca: document.getElementById('marca').value,
            zona: document.getElementById('zona').value,
            descripcion: document.getElementById('descripcion').value,
            beneficios: document.getElementById('beneficios').value,
            imagen: document.getElementById('imagen').value
        };
        if (id) {
            actualizarMaquinaria(id, data);
        } else {
            crearMaquinaria(data);
        }
    });

    btnCancelar.addEventListener('click', limpiarFormulario);

    function cargarMaquinarias() {
        maquinariasList.innerHTML = '<div class="loading-state" style="grid-column:1/-1"><i class="fas fa-circle-notch"></i> Cargando maquinarias...</div>';

        fetch(API_URL)
            .then(response => {
                if (!response.ok) throw new Error('Error al cargar la maquinaria');
                return response.json();
            })
            .then(data => {
                maquinariasList.innerHTML = '';

                if (!data || data.length === 0) {
                    maquinariasList.innerHTML = '<div class="empty-state"><i class="fas fa-dumbbell"></i><p>No hay maquinarias registradas aún.</p></div>';
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
            })
            .catch(error => {
                maquinariasList.innerHTML = `<div class="msg-danger" style="grid-column:1/-1">${error.message}</div>`;
            });
    }

    function crearMaquinaria(data) {
        fetch(API_URL, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(data)
        })
        .then(response => {
            if (!response.ok) throw new Error('Error al crear la maquinaria');
            return response.json();
        })
        .then(() => {
            mostrarMensaje('Maquinaria creada con éxito', 'success');
            limpiarFormulario();
            cargarMaquinarias();
        })
        .catch(error => mostrarMensaje(error.message, 'danger'));
    }

    function actualizarMaquinaria(id, data) {
        fetch(`${API_URL}/${id}`, {
            method: 'PUT',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(data)
        })
        .then(response => {
            if (!response.ok) throw new Error('Error al actualizar la maquinaria');
            return response.json();
        })
        .then(() => {
            mostrarMensaje('Maquinaria actualizada con éxito', 'success');
            limpiarFormulario();
            cargarMaquinarias();
        })
        .catch(error => mostrarMensaje(error.message, 'danger'));
    }

    function eliminarMaquinaria(id) {
        if (!confirm('¿Eliminar esta máquina? Esta acción no se puede deshacer.')) return;
        fetch(`${API_URL}/${id}`, { method: 'DELETE' })
            .then(response => {
                if (!response.ok) throw new Error('Error al eliminar la maquinaria');
                mostrarMensaje('Máquina eliminada', 'success');
                cargarMaquinarias();
            })
            .catch(error => mostrarMensaje(error.message, 'danger'));
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
        btnGuardar.innerHTML = '<i class="fas fa-save"></i> Actualizar Máquina';
        btnCancelar.style.display = 'inline-block';
        window.scrollTo({ top: 0, behavior: 'smooth' });
    }

    function limpiarFormulario() {
        maquinariaForm.reset();
        document.getElementById('maquinariaId').value = '';
        formTitle.textContent = 'Nueva Máquina';
        btnGuardar.innerHTML = '<i class="fas fa-save"></i> Guardar Máquina';
        btnCancelar.style.display = 'none';
        formMessage.innerHTML = '';
    }

    function mostrarMensaje(mensaje, tipo) {
        formMessage.innerHTML = `<div class="msg-${tipo}">${tipo === 'success' ? '✓' : '✗'} ${mensaje}</div>`;
        setTimeout(() => { formMessage.innerHTML = ''; }, 3000);
    }
});
