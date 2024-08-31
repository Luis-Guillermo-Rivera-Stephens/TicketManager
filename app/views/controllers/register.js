var department = 0
var departmentname = array[department]

function updateDepartment(elementId) {
    department = parseInt(document.getElementById(elementId).getAttribute('data-value'));
    departmentname = array[department];
    document.getElementById('dep').innerHTML = departmentname;
}

// Agrega los event listeners para cada item
document.getElementById('departmentF').addEventListener('click', () => updateDepartment('departmentF'));
document.getElementById('departmentB').addEventListener('click', () => updateDepartment('departmentB'));
document.getElementById('departmentA').addEventListener('click', () => updateDepartment('departmentA'));
document.getElementById('departmentS').addEventListener('click', () => updateDepartment('departmentS'));
document.getElementById('departmentN').addEventListener('click', () => updateDepartment('departmentN'));
document.getElementById('departmentT').addEventListener('click', () => updateDepartment('departmentT'));
document.getElementById('departmentSy').addEventListener('click', () => updateDepartment('departmentSy'));
document.getElementById('departmentP').addEventListener('click', () => updateDepartment('departmentP'));
document.getElementById('departmentG').addEventListener('click', () => updateDepartment('departmentG'));

document.getElementById('register').addEventListener('click', () => {
    var name = document.getElementById('name').value;
    var email = document.getElementById('email').value;
    var password = document.getElementById('password').value;
    if (validInfo(name, email, password, department) === 0 ){
        console.log("error en algun campo")
        return
    }
    
    let account = new Account(0,name, email, password, false, department, false);
    let accountJSON = JSON.stringify(account);
    console.log(accountJSON)

    let xhr = new XMLHttpRequest();
    xhr.open("POST", "http://localhost:8080/account");
    xhr.setRequestHeader('Content-Type','application/json');  
    xhr.send(accountJSON);

    xhr.onload = (req, res) => {
        if (xhr.status == 201){
            console.log("usuario registrado");
        } else if (xhr.status == 401){
            console.log("usuario ya existe");
        }
        else {
            console.log(xhr.responseText);
        }
            
    }


})

function validInfo(name, email, password, department) {
    if (name === "") {
        console.log("Nombre invalido")
        return 0
    }
    if (email === "") {
        console.log("Correo invalido")
        return 0
    }
    if (password === "") {
        console.log("Contraseña invalida")
        return 0
    }
    if (department === 0) {
        console.log("Departamento invalido")
        return 0
    }

    return 1

}