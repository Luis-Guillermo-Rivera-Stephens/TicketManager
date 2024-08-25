class SSHandler {
    constructor(){
        this.started = false,
        this.logged = false,
        this.user = null;
    }
}

function starter(){
    let aux = new SSHandler();
    aux.started = JSON.parse(sessionStorage.SSH).started;
    aux.logged = JSON.parse(sessionStorage.SSH).logged;
    aux.user = JSON.parse(sessionStorage.SSH).user;
    if(!aux.started){
        aux.started = true;
        sessionStorage.SSH = JSON.stringify(aux);
    }
}

function logger(usuario){
    let aux = new SSHandler();
    aux.started = JSON.parse(sessionStorage.SSH).started;
    aux.logged = JSON.parse(sessionStorage.SSH).logged;
    aux.user = JSON.parse(sessionStorage.SSH).user;
    if (!aux.logged) {
        aux.logged = true;
        aux.user = usuario;
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
        aux.user = null;
        sessionStorage.SSH = JSON.stringify(aux);
    }
}