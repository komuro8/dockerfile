const API_URL = "http://localhost:8080";

function listarCategorias() {
    fetch(`${API_URL}/categorias`)
        .then(res => res.json())
        .then(data => {
            const ul = document.getElementById("categorias");
            ul.innerHTML = "";
            ul.innerHTML = `<li>${JSON.stringify(data)}</li>`;
        })
        .catch(err => console.error("Error:", err));
}
