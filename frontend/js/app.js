const API_URL = "/api";

function listarCategorias() {
    fetch(`${API_URL}/categorias`)
        .then(res => res.json())
        .then(data => {
            const ul = document.getElementById("categorias");
            ul.innerHTML = "";

            if (!Array.isArray(data)) {
                ul.innerHTML = `<li>Error: ${JSON.stringify(data)}</li>`;
                return;
            }

            data.forEach(cat => {
                ul.innerHTML += `<li>${cat.nombre} - ${cat.descripcion}</li>`;
            });
        })
        .catch(err => console.error("Error:", err));
}
