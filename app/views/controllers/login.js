starter()

document.getElementById("login").addEventListener('click', (event)=>{

    event.preventDefault();
    console.log("click a log in\n")
    let email = document.getElementById("email").value;
    let password = document.getElementById("password").value;

    let xhr = new XMLHttpRequest();
    xhr.open('GET', 'http://localhost:8080/account');
    xhr.setRequestHeader('email', email)
    xhr.setRequestHeader('password', password)
    xhr.send();
    
    xhr.onload= (req, res) => {
        if (xhr.status == 200) {
            var data = JSON.parse(xhr.responseText);
            var token = xhr.getResponseHeader("token");
            
            console.log(data, "token: ", token); // Verificar si los datos y el token se están recibiendo correctamente

            if (data.id_account) {
                console.log("ID Account exists: ", data.id_account); // Confirmar si `id_account` existe en la respuesta
            } else {
                console.error("ID Account not found in response");
            }
            
            if (token) {
                console.log("Token exists: ", token); // Confirmar si el token fue recibido
            } else {
                console.error("Token not found in response headers");
            }

            logger(data, token);
            setLastCall("http://localhost:8080/tickets/open", data.id_account, token);
            
            if (data.isstarted === false){
                alerta("Ups algo salio mal","Redirect to update");
                window.location.href = 'http://localhost:8080/update';
            } else {
                console.log("Redirect to home");
                window.location.href = 'http://localhost:8080/home';
            }
        }
        else if (xhr.status == 401){
            alerta("Ups algo salio mal","Usuario invalido, favor de revisar los campos");
        }
        else if (xhr.status == 400) {
            alerta("Ups algo salio mal","Campos invalidos, favor de revisar el correo");
        }
        else {
            alerta("Ups algo salio mal","error trying to get the account");
        }
    }
});
