function alerta(title, message) {
    
    let modalBody = document.getElementById('modalIdAlertaBody');
    modalBody.innerHTML = message;
    document.getElementById("modalTitleIdAlert").innerHTML = title 

    let miModal = new bootstrap.Modal(document.getElementById("modalIdAlerta"));
    miModal.show();
}

function cerrarAlerta() {
    let miModal = new bootstrap.Modal(document.getElementById("modalIdAlerta"));
    miModal.hide();
}
