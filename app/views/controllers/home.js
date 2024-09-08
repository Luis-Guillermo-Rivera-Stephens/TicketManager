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

function ticket_formateado(ticket) {
    let row = document.createElement("tr")
    row.classList.add("ticket-row")

    row.innerHTML = `
        <td class="id_td">${ticket.id_ticket}</td>            
        <td class="title_td">${ticket.title}</td>   
        <td class="description_td">${ticket.t_description}</td> 
        <td class="owner_td">${ticket.owner != ""? ticket.owner : "-"}</td>   
        <td class="department_td">${ticket.department}</td>
        <td class="status_td">${ticket.status}</td>  
        <td class="creation_date_td">${ticket.creation_date}</td>
    `

    return row;
}

function on_load_tickets(map, xhr) {
    var data = JSON.parse(xhr.responseText);
    
    let tickets_table = document.getElementById("tickets-body");

            
    tickets_table.innerHTML = '';

    for (let i = 0; i < data.length; i++) {
        map.set(data[i].id_ticket, data[i]);

        tickets_table.appendChild(ticket_formateado(data[i]));
    }
}


document.getElementById('get-tickets-all-open').addEventListener("click", () => {
    let xhr = new XMLHttpRequest();
    let map = new Map();

    xhr.open('GET', 'http://localhost:8080/tickets/open');

    let id = JSON.parse(sessionStorage.SSH).account.id_account;
    let token = get_token()

    xhr.setRequestHeader('id', id);
    xhr.setRequestHeader('token', token);

    xhr.send();

    xhr.onload = () => {
        if (xhr.status == 200) {
            on_load_tickets(map, xhr)
        }
    }
});


