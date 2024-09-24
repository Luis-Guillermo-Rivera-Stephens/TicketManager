class SSHandler {
    constructor(){
        this.started = false,
        this.logged = false,
        this.account = null,
        this.token = "";
        this.lastCall = new LastCall('http://localhost:8080/tickets/open', this.account ? this.account.id_account : null, this.token);
    }
}

class LastCall {
    constructor(url, id, token) {
        this.url = url;
        this.id = id;
        this.token = token;
    }
}

function starter() {
    let aux = new SSHandler();

    // Verifica si 'SSH' no está definido en sessionStorage
    if (!sessionStorage.SSH) {
        console.log("Starting the SSH");
        // Inicializa las propiedades si 'SSH' no existe
        aux.started = false;
        aux.logged = false;
        aux.account = null;
        aux.token = null;
    } else {
        // Si 'SSH' existe, recupera los valores
        console.log("Getting the SSH info")
        aux.started = JSON.parse(sessionStorage.SSH).started;
        aux.logged = JSON.parse(sessionStorage.SSH).logged;
        aux.account = JSON.parse(sessionStorage.SSH).account;
        aux.token = JSON.parse(sessionStorage.SSH).token;
    }

    // Si no ha sido iniciado, lo marca como iniciado
    if (!aux.started) {
        aux.started = true;
        sessionStorage.SSH = JSON.stringify(aux);
    }
}

function logger(usuario, token){
    let aux = new SSHandler();
    aux.started = JSON.parse(sessionStorage.SSH).started;
    aux.logged = JSON.parse(sessionStorage.SSH).logged;
    aux.user = JSON.parse(sessionStorage.SSH).user;
    if (!aux.logged) {
        aux.logged = true;
        aux.account = usuario;
        aux.token = token;
        sessionStorage.SSH = JSON.stringify(aux);
    }
}

function unlogger(){
    let aux = new SSHandler();
    aux.started = JSON.parse(sessionStorage.SSH).started;
    aux.logged = JSON.parse(sessionStorage.SSH).logged;
    aux.user = JSON.parse(sessionStorage.SSH).user;
    if (aux.logged) {
        aux.logged = false;
        aux.account = null;
        aux.token = "";
        sessionStorage.SSH = JSON.stringify(aux);
    }
}

function get_token(){
    let aux = new SSHandler();
    aux.started = JSON.parse(sessionStorage.SSH).started;
    aux.logged = JSON.parse(sessionStorage.SSH).logged;
    if (aux.started === true && aux.logged === true) {
        let token = JSON.parse(sessionStorage.SSH).token
        return token
    }
}

function set_token(token){
    let aux = JSON.parse(sessionStorage.SSH)
    if (aux.started === true && aux.logged === true) {
        aux.token = token
        sessionStorage.SSH = JSON.stringify(aux)
    }
}

function get_email(){
    let aux = new SSHandler();
    aux.started = JSON.parse(sessionStorage.SSH).started;
    aux.logged = JSON.parse(sessionStorage.SSH).logged;
    if (aux.started === true && aux.logged === true) {
        let email = JSON.parse(sessionStorage.SSH).account.email
        return email
    }
}

function get_department(){
    let aux = new SSHandler();
    aux.started = JSON.parse(sessionStorage.SSH).started;
    aux.logged = JSON.parse(sessionStorage.SSH).logged;
    if (aux.started === true && aux.logged === true) {
        let department = JSON.parse(sessionStorage.SSH).account.department_fk
        return department
    }
}

function get_id_account(){
    let aux = new SSHandler();
    aux.started = JSON.parse(sessionStorage.SSH).started;
    aux.logged = JSON.parse(sessionStorage.SSH).logged;
    if (aux.started === true && aux.logged === true) {
        let id_account = JSON.parse(sessionStorage.SSH).account.id_account
        return id_account
    }
}

function islogged(){
    return JSON.parse(sessionStorage.SSH).logged;

}

function setLastCall(url, id, token) {
    let sshData = JSON.parse(sessionStorage.SSH);

    let aux = new SSHandler();
    aux.started = sshData.started;
    aux.logged = sshData.logged;


    if (aux.started === true && aux.logged === true) {
        sshData.lastCall.url = url;
        sshData.lastCall.id = id;
        sshData.lastCall.token = token;

        sessionStorage.SSH = JSON.stringify(sshData);
    }
}

starter()
