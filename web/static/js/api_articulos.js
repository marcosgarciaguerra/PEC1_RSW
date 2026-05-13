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

    function cargarArticulos() {
        articulosList.innerHTML = '<div class="loading-state" style="grid-column:1/-1"><i class="fas fa-circle-notch"></i> Cargando artículos...</div>';

        fetch(API_URL)
            .then(response => {
                if (!response.ok) throw new Error('Error al cargar los artículos');
                return response.json();
            })
            .then(data => {
                articulosList.innerHTML = '';

                if (!data || data.length === 0) {
                    articulosList.innerHTML = '<div class="empty-state"><i class="fas fa-box-open"></i><p>No hay artículos en el catálogo.</p></div>';
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
                        tagEl.style.display = 'inline-block';
                    }

                    clone.querySelector('.btn-editar').addEventListener('click', () => llenarFormulario(articulo));
                    clone.querySelector('.btn-eliminar').addEventListener('click', () => eliminarArticulo(articulo.id));

                    articulosList.appendChild(clone);
                });
            })
            .catch(error => {
                articulosList.innerHTML = `<div class="msg-danger" style="grid-column:1/-1">${error.message}</div>`;
            });
    }

    function crearArticulo(articuloData) {
        fetch(API_URL, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(articuloData)
        })
        .then(response => {
            if (!response.ok) throw new Error('Error al crear el artículo');
            return response.json();
        })
        .then(() => {
            mostrarMensaje('Artículo creado con éxito', 'success');
            limpiarFormulario();
            cargarArticulos();
        })
        .catch(error => mostrarMensaje(error.message, 'danger'));
    }

    function actualizarArticulo(id, articuloData) {
        fetch(`${API_URL}/${id}`, {
            method: 'PUT',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(articuloData)
        })
        .then(response => {
            if (!response.ok) throw new Error('Error al actualizar el artículo');
            return response.json();
        })
        .then(() => {
            mostrarMensaje('Artículo actualizado con éxito', 'success');
            limpiarFormulario();
            cargarArticulos();
        })
        .catch(error => mostrarMensaje(error.message, 'danger'));
    }

    function eliminarArticulo(id) {
        if (!confirm('¿Eliminar este artículo? Esta acción no se puede deshacer.')) return;
        fetch(`${API_URL}/${id}`, { method: 'DELETE' })
            .then(response => {
                if (!response.ok) throw new Error('Error al eliminar el artículo');
                mostrarMensaje('Artículo eliminado', 'success');
                cargarArticulos();
            })
            .catch(error => mostrarMensaje(error.message, 'danger'));
    }

    function llenarFormulario(articulo) {
        document.getElementById('articuloId').value = articulo.id;
        document.getElementById('nombre').value = articulo.nombre;
        document.getElementById('categoria').value = articulo.categoria;
        document.getElementById('precio').value = articulo.precio;
        document.getElementById('imagen').value = articulo.imagen;
        document.getElementById('etiqueta').value = articulo.etiqueta || '';
        formTitle.textContent = 'Editar Artículo';
        btnGuardar.innerHTML = '<i class="fas fa-save"></i> Actualizar Artículo';
        btnCancelar.style.display = 'inline-block';
        window.scrollTo({ top: 0, behavior: 'smooth' });
    }

    function limpiarFormulario() {
        articuloForm.reset();
        document.getElementById('articuloId').value = '';
        formTitle.textContent = 'Nuevo Artículo';
        btnGuardar.innerHTML = '<i class="fas fa-save"></i> Guardar Artículo';
        btnCancelar.style.display = 'none';
        formMessage.innerHTML = '';
    }

    function mostrarMensaje(mensaje, tipo) {
        formMessage.innerHTML = `<div class="msg-${tipo}">${tipo === 'success' ? '✓' : '✗'} ${mensaje}</div>`;
        setTimeout(() => { formMessage.innerHTML = ''; }, 3000);
    }
});
