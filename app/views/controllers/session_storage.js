class SSHandler {
    constructor(){
        this.started = false,
        this.logged = false,
        this.account = null,
        this.token = "";
    }
}

function starter() {
    let aux = new SSHandler();

    // Verifica si 'SSH' no está definido en sessionStorage
    if (!sessionStorage.SSH) {
        // Inicializa las propiedades si 'SSH' no existe
        aux.started = false;
        aux.logged = false;
        aux.account = null;
        aux.token = null;
    } else {
        // Si 'SSH' existe, recupera los valores
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
function get_email(){
    let aux = new SSHandler();
    aux.started = JSON.parse(sessionStorage.SSH).started;
    aux.logged = JSON.parse(sessionStorage.SSH).logged;
    if (aux.started === true && aux.logged === true) {
        let email = JSON.parse(sessionStorage.SSH).account.email
        return email
    }
}

function islogged(){
    return JSON.parse(sessionStorage.SSH).logged;

}