starter()

if (!islogged()){
    window.location.href = 'http://localhost:8080/login';
}

let TicketMap = new Map();
var ticket_info;

let name_department = JSON.parse(sessionStorage.SSH).account.name + " - " + array[JSON.parse(sessionStorage.SSH).account.department_fk]
console.log(name_department)
//document.getElementById("navbar_info").innerHTML = "Hola buenas tardes"
document.getElementById('navbar_info').innerHTML = name_department

document.getElementById('logout').addEventListener("click", ()=>{
    unlogger()
    window.location.href = 'http://localhost:8080/login'
})


function ticket_formateado(ticket) {
    let row = document.createElement("tr");
    row.classList.add("ticket-row");
    row.id = `ticket-${ticket.id_ticket}`;
    row.setAttribute('data-value', ticket.id_ticket);
    row.setAttribute("data-bs-toggle","modal")
    row.setAttribute("data-bs-target","#modalIdTicket")
                

    row.innerHTML = `
        <td class="id_td">${ticket.id_ticket}</td>            
        <td class="title_td">${ticket.title}</td>   
        <td class="owner_td">${ticket.owner !== "" ? ticket.owner : "-"}</td>   
        <td class="department_td">${ticket.department}</td>
        <td class="status_td">${ticket.status}</td>  
        <td class="priority_td">${ticket.priority}</td>  
    `;

    row.addEventListener("click", () => {
        let ticket_id = parseInt(row.getAttribute('data-value')); 
        let ticket_flag = TicketMap.has(ticket_id);
        if (ticket_flag) {
            ticket_info = TicketMap.get(ticket_id)
            console.log(ticket_info);
            document.getElementById("ticketTitle").innerHTML = "Ticket "+ ticket_id + `: ${ticket_info.title}`;
            document.getElementById("ticketDescription").innerHTML = ticket_info.t_description;
            document.getElementById("ticketOwner").innerHTML = `${ticket_info.owner == ""? "-": ticket_info.owner}`;
            document.getElementById("ticketDepartment").innerHTML = ticket_info.department;
            document.getElementById("ticketPriority").innerHTML = ticket_info.priority;
            document.getElementById("ticketCurrentStatus").innerHTML = ticket_info.status;
        
            //document.getElementById("modalIdBody").innerHTML = ticket_info.t_description
        } else {
            console.error(`Ticket con ID ${ticket_id} no encontrado en TicketMap`);
        }
    });

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

function getTickets(map, route){
    let xhr = new XMLHttpRequest();

    let url = 'http://localhost:8080/tickets' + route

    xhr.open('GET', url);

    let id = JSON.parse(sessionStorage.SSH).account.id_account;
    let token = get_token()

    xhr.setRequestHeader('id', id);
    xhr.setRequestHeader('token', token);
    

    xhr.send();

    xhr.onload = () => {
        if (xhr.status == 200) {
            set_token(xhr.getResponseHeader('token'));
            on_load_tickets(map, xhr)
            setLastCall(url, id, token)
        }
        else if (xhr.status == 418) {
            console.log('redireccionando a login por token invalido')
            unlogger()
            window.location.href = 'http://localhost:8080/login'
        }
    }
}

function getTicketsWithDeparment(map, route, dep){
    let xhr = new XMLHttpRequest();
    route += ("/" + dep)

    let url = 'http://localhost:8080/tickets' + route

    xhr.open('GET', url);

    let id = JSON.parse(sessionStorage.SSH).account.id_account;
    let token = get_token()

    xhr.setRequestHeader('id', id);
    xhr.setRequestHeader('token', token);

    xhr.send();

    xhr.onload = () => {
        if (xhr.status == 200) {
            set_token(xhr.getResponseHeader('token'));
            on_load_tickets(map, xhr)
            setLastCall(url, id, token)
        }
        else if (xhr.status == 418) {
            console.log('redireccionando a login por token invalido')
            unlogger()
            window.location.href = 'http://localhost:8080/login'
        }
    }
}

function doLastCall() {
    // Obtener el objeto SSH y lastCall una sola vez
    let sshData = JSON.parse(sessionStorage.SSH);
    let lastCall = sshData.lastCall;

    // Crear el XMLHttpRequest
    let xhr = new XMLHttpRequest();
    xhr.open('GET', lastCall.url);

    // Configurar los encabezados con id y token
    xhr.setRequestHeader('id', lastCall.id);
    xhr.setRequestHeader('token', lastCall.token);

    // Enviar la solicitud
    xhr.send();

    // Manejar la respuesta
    xhr.onload = () => {
        if (xhr.status === 200) {
            // Actualizar el token
            set_token(xhr.getResponseHeader('token'));
            
            // Cargar los tickets
            on_load_tickets(TicketMap, xhr);
        } else if (xhr.status === 418) {
            console.log('Redireccionando a login por token inválido');
            unlogger();
            window.location.href = 'http://localhost:8080/login';
        }
    };
}


document.getElementById('get-tickets-all-open').addEventListener("click", () => getTickets(TicketMap,'/open'));
document.getElementById("get-tickets-all").addEventListener("click",() => getTickets(TicketMap,'/all'));
document.getElementById("get-tickets-all-unassigned").addEventListener("click", () => getTickets(TicketMap,'/unassigned'));
document.getElementById("get-tickets-mytickets").addEventListener("click", () => getTickets(TicketMap,'/owner'));
document.getElementById("get-tickets-closed").addEventListener("click", () => getTickets(TicketMap,'/closed'));

document.getElementById("get-tickets-mydepartment").addEventListener("click", () => getTicketsWithDeparment(TicketMap,'/department', JSON.parse(sessionStorage.SSH).account.department_fk));
document.getElementById("get-tickets-unassigned-bydepartment").addEventListener("click", () => getTicketsWithDeparment(TicketMap,'/department', JSON.parse(sessionStorage.SSH).account.department_fk));

document.getElementById("get-tickets-front").addEventListener("click", () => getTicketsWithDeparment(TicketMap, `/department`, 1));
document.getElementById("get-tickets-back").addEventListener("click", () => getTicketsWithDeparment(TicketMap, `/department`, 2));
document.getElementById("get-tickets-app").addEventListener("click", () => getTicketsWithDeparment(TicketMap, `/department`, 3));
document.getElementById("get-tickets-sql").addEventListener("click", () =>getTicketsWithDeparment(TicketMap, `/department`, 4));
document.getElementById("get-tickets-neo4j").addEventListener("click", () => getTicketsWithDeparment(TicketMap, `/department`, 5));
document.getElementById("get-tickets-testing").addEventListener("click", () => getTicketsWithDeparment(TicketMap, `/department`, 8));
document.getElementById("get-tickets-sysadmin").addEventListener("click", () => getTicketsWithDeparment(TicketMap, `/department`, 9));
document.getElementById("get-tickets-PM").addEventListener("click", () => getTicketsWithDeparment(TicketMap, `/department`, 6));

doLastCall()