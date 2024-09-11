starter()

if (!islogged()){
    window.location.href = 'http://localhost:8080/login';
}

if (get_department() != 6 && get_department() != 7) {
    console.log("You're not a PM")
    window.location.href = 'http://localhost:8080/home';
}

var department = 0
var departmentname = array[department]
var prioridad = "Baja"


function updateDepartment(elementId) {
    department = parseInt(document.getElementById(elementId).getAttribute('data-value'));
    departmentname = array[department];
    document.getElementById('dep').innerHTML = departmentname;
}
function updatePriority(elementId) {
    prioridad = document.getElementById(elementId).getAttribute('data-value');
    document.getElementById('pri').innerHTML = prioridad;
}

document.getElementById('departmentF').addEventListener('click', () => updateDepartment('departmentF'));
document.getElementById('departmentB').addEventListener('click', () => updateDepartment('departmentB'));
document.getElementById('departmentA').addEventListener('click', () => updateDepartment('departmentA'));
document.getElementById('departmentS').addEventListener('click', () => updateDepartment('departmentS'));
document.getElementById('departmentN').addEventListener('click', () => updateDepartment('departmentN'));
document.getElementById('departmentT').addEventListener('click', () => updateDepartment('departmentT'));
document.getElementById('departmentSy').addEventListener('click', () => updateDepartment('departmentSy'));
document.getElementById('departmentP').addEventListener('click', () => updateDepartment('departmentP'));

document.getElementById('priorityB').addEventListener('click', () => updatePriority('priorityB'));
document.getElementById('priorityM').addEventListener('click', () => updatePriority('priorityM'));
document.getElementById('priorityA').addEventListener('click', () => updatePriority('priorityA'));
document.getElementById('priorityU').addEventListener('click', () => updatePriority('priorityU'));

document.getElementById('create-ticket').addEventListener('click', () => {
    event.preventDefault();
    let title = document.getElementById('title').value
    let description = document.getElementById('description').value
    let department = departmentname
    let priority = prioridad

    let new_ticket = Create_Ticket(title, description, department, priority)
    console.log(new_ticket)
    let TicketJSON = JSON.stringify(new_ticket);

    let xhr = new XMLHttpRequest();
    xhr.open("POST", "http://localhost:8080/tickets");
    xhr.setRequestHeader('Content-Type','application/json');  
    xhr.setRequestHeader('token', get_token());
    xhr.setRequestHeader('id', get_id_account());
    xhr.send(TicketJSON);

    xhr.onload = (req, res) => {
        if (xhr.status == 201){
            alerta("Accion exitosa","ticket creado");
        } else if (xhr.status == 418) {
            alerta("Ups algo salio mal",'redireccionando a login por token invalido')
            unlogger()
            window.location.href = 'http://localhost:8080/login'
        }
        else {
            alerta("Ups hubo un problema",xhr.responseText);
        }
            
    }

})