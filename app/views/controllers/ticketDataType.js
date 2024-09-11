class Tickets {
    constructor(id, title, t_description, creation_date, id_account, owner, department, status, priority) {
        this.id_ticket= id
        this.title = title
        this.t_description = t_description
        this.creation_date = creation_date
        this.id_account = id_account
        this.owner = owner
        this.department = department
        this.status = status
        this.priority = priority

    }

    setTitle(title) {
        if (title === "") {
            console.log("Title cannot be empty");
            return false;
        }
        this.title = title;
        return true;
    }

    setT_Description(description) {
        if (description === "") {
            console.log("Description cannot be empty");
            return false;
        }
        this.t_description = description;
        return true;
    }

    setID_Account(id_account) {
        if (id_account > 0) {
            console.log("ID_Account invalido");
            return false;
        }
        this.id_account = id_account;
        return true;
    }

    setOwner(Owner) {
        if (Owner =="") {
            console.log("Owner invalido");
            return false;
        }
        this.owner = Owner;
        return true;
    }

    setDepartment(department) {
        if (!hashMap.has(department)) {
            console.log("Invalid department");
            return false;
        }
        this.department = department;
        return true;
    }
    
    setStatus(status) {
        if (status =="") {
            console.log("Status invalido");
            return false;
        }
        this.status = status;
        return true;
    }

    setPriority(priority) {
        if (priority =="") {
            console.log("Prioridad invalido");
            return false;
        }
        this.priority = priority;
        return true;
    }

    

    create_ticket(title, description, department, priority) {
        this.id_ticket = 0

        if (!this.setTitle(title)) {s
            return false;
        }
        if (!this.setT_Description(description)) {
            return false;
        }
        if (!this.setDepartment(department)) {
            return false;
        }
        if (!this.setPriority(priority)) {
            return false;
        }
        this.id_account = 0
        this.owner = ""
        this.creation_date = ""
        this.status = ""
        return true
    }
}

function Create_Ticket(title, description, department, priority) {
    let Ticket = new Tickets(0, "", "", "", 0, "", "", "", "")
    Ticket.create_ticket(title, description, department, priority)
    return Ticket
}