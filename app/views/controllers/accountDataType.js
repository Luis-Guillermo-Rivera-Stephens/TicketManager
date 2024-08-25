class Account {
    constructor(id, name, email, password, isstarted, department_fk, ispm) {
        this.id_account= id
        this.setName(name)
        this.setEmail(email)
        this.setPassword(password)
        this.setIsstarted(isstarted)
        this.setDepartmentFk(department_fk)
        this.setIsPM(ispm)
    }

    setName(name) {
        if (name === "") {
            console.log("Name cannot be empty");
            return false;
        }
        this.name = name;
        return true;
    }

    setEmail(email) {
        if (email === "") {
            console.log("Email cannot be empty");
            return false;
        }
        this.email = email;
        return true;
    }

    setPassword(password) {
        if (password === "") {
            console.log("Password cannot be empty");
            return false;
        }
        this.password = password;
        return true;
    }

    setIsstarted(isstarted) {
        if (typeof isstarted !== "boolean") {
            console.log("isStarted must be a boolean");
            return false;
        }
        this.isstarted = isstarted;
        return true;
    }

    setDepartmentFk(department_fk) {
        if (!Number.isInteger(department_fk)) {
            console.log("departmentFk must be an integer");
            return false;
        }
        this.department_fk = department_fk;
        return true;
    }

    setIsPM(IsPM) {
        if (typeof IsPM !== "boolean") {
            console.log("IsPM must be a boolean");
            return false;
        }
        this.ispm = IsPM;
        return true;
    }

    create_account(name, email, password, departmentFk) {
        this.id = 0

        if (!this.setName(name)) {
            return false;
        }
        if (!this.setEmail(email)) {
            return false;
        }
        if (!this.setPassword(password)) {
            return false;
        }
        if (!this.setDepartmentFk(departmentFk)) {
            return false;
        }
        if (!this.setIsstarted(false)){
            return false;
        }
        if (!this.setIsPM(false)){
            return false;
        }
        return true
    }
}
