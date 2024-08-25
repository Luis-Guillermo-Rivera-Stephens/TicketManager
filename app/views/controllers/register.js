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
document.getElementById('departmenTt').addEventListener('click', () => updateDepartment('departmenTt'));
document.getElementById('departmentSy').addEventListener('click', () => updateDepartment('departmentSy'));
document.getElementById('departmentP').addEventListener('click', () => updateDepartment('departmentP'));
document.getElementById('departmentG').addEventListener('click', () => updateDepartment('departmentG'));

document.getElementById('register').addEventListener('click', () => {
    var name = document.getElementById('name').value;
    var email = document.getElementById('email').value;
    var password = document.getElementById('password').value;
    let account = new Account(0,name, email, password, false, department, false);
    let accountsJSON = JSON.stringify(account);
    console.log(accountsJSON)

    

})