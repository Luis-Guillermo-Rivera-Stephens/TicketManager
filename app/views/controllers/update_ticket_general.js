function updateStatus(s_id, t_id) {
    let xhr = new XMLHttpRequest();
    xhr.open('PUT', `http://localhost:8080/tickets/update/status/${s_id}`);
    xhr.setRequestHeader('token', get_token());
    xhr.setRequestHeader('id', get_id_account());
    xhr.setRequestHeader('id_ticket', t_id);


    xhr.send();
    xhr.onload= (req, res) => {
        
    }
}
