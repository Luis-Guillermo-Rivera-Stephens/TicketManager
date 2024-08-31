starter()

if (!islogged()){
    window.location.href = 'http://localhost:8080/login';
}

let name_department = JSON.parse(sessionStorage.SSH).account.name + " - " + array[JSON.parse(sessionStorage.SSH).account.department_fk]
console.log(name_department)
//document.getElementById("navbar_info").innerHTML = "Hola buenas tardes"
document.getElementById('navbar_info').innerHTML = name_department

document.getElementById('logout').addEventListener("click", ()=>{
    unlogger()
    window.location.href = 'http://localhost:8080/login'
})

