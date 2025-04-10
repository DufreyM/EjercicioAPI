const API_URL = 'http://localhost:8083/incidentes';

document.addEventListener("DOMContentLoaded", () => {
    cargarIncidentes();

    document.getElementById("incidenteForm").addEventListener("submit", async (e) => {
        e.preventDefault();

        const reportero = document.getElementById("reportero").value;
        const descripcion = document.getElementById("descripcion").value;
        const status = document.getElementById("status").value;

        const data = { reportero, descripcion, status };

        const response = await fetch(API_URL, {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify(data)
        });

        if (response.ok) {
            document.getElementById("incidenteForm").reset();
            cargarIncidentes();
        } else {
            alert("Error al crear el incidente");
        }
    });
});

async function cargarIncidentes() {
    const container = document.getElementById("listaIncidentes");
    container.innerHTML = "";

    const res = await fetch(API_URL);
    const incidentes = await res.json();

    incidentes.forEach(inc => {
        const card = document.createElement("div");
        card.className = `card ${inc.status.replace(" ", "-")}`;

        card.innerHTML = `
            <h3>${inc.descripcion}</h3>
            <p><strong>Reportero:</strong> ${inc.reportero}</p>
            <p><strong>Estado:</strong> ${inc.status}</p>
            <small>${new Date(inc.fecha).toLocaleString()}</small>

            <select onchange="actualizarEstado(${inc.id}, this.value)">
                <option value="">--Actualizar estado--</option>
                <option value="pendiente">Pendiente</option>
                <option value="en proceso">En proceso</option>
                <option value="resuelto">Resuelto</option>
            </select>

            <button onclick="eliminarIncidente(${inc.id})">Eliminar</button>
        `;

        container.appendChild(card);
    });
}

async function eliminarIncidente(id) {
    if (confirm("¿Seguro que deseas eliminar este incidente?")) {
        const res = await fetch(`${API_URL}/${id}`, { method: "DELETE" });
        if (res.ok) {
            cargarIncidentes();
        } else {
            alert("Error al eliminar el incidente");
        }
    }
}

async function actualizarEstado(id, nuevoEstado) {
    if (!nuevoEstado) return;

    const res = await fetch(`${API_URL}/${id}`, {
        method: "PUT",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ status: nuevoEstado })
    });

    if (res.ok) {
        cargarIncidentes();
    } else {
        alert("Error al actualizar el estado");
    }
}
