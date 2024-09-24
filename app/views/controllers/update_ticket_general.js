function updateStatus(s_id,) {
    if (get_id_account() != ticket_info.id_account){
        alerta("Ups", "No eres dueño de este ticket");
        return;
    }

    let xhr = new XMLHttpRequest();
    xhr.open('PUT', `http://localhost:8080/tickets/update/status/${s_id}`);
    xhr.setRequestHeader('token', get_token());
    xhr.setRequestHeader('id', get_id_account());
    xhr.setRequestHeader('id_ticket', ticket_info.id_ticket);


    xhr.send();
    xhr.onload= (req, res) => {
        if (xhr.status == 202) {
            set_token(xhr.getResponseHeader('token'));
            alerta("Si jalo bro", "Todo bien")
            //TODO actualizar el ticket de manera visual
        }
        else if (xhr.status == 418) {
            console.log('redireccionando a login por token invalido')
            unlogger()
            window.location.href = 'http://localhost:8080/login'
        }
        else [
            alerta(`Error ${xhr.status}`, `${xhr.responseText}`)
        ]
    }
}

function assignMeTicket() {
    let xhr = new XMLHttpRequest();
    xhr.open('PUT', `http://localhost:8080/tickets/update/owner`);
    xhr.setRequestHeader('token', get_token());
    xhr.setRequestHeader('id', get_id_account());
    xhr.setRequestHeader('id_ticket', ticket_info.id_ticket);


    xhr.send();
    xhr.onload= (req, res) => {
        if (xhr.status == 202) {
            set_token(xhr.getResponseHeader('token'));
            //TODO actualizar el ticket de manera visual
            alerta("Si jalo bro", "todo bien pa")
        }
        else if (xhr.status == 418) {
            console.log('redireccionando a login por token invalido')
            unlogger()
            window.location.href = 'http://localhost:8080/login'
        }
        else [
            alerta(`Error ${xhr.status}`, `${xhr.responseText}`)
        ]
    }
}

