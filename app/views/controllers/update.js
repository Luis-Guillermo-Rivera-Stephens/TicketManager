starter()

if (!islogged()){
    window.location.href = `${URL_SERVER}/login`;
}


document.getElementById('submit').addEventListener("click", () => {
    let currentpass = document.getElementById("password").value;
    let newpass1 = document.getElementById("newpassword1").value;
    let newpass2 = document.getElementById("newpassword2").value;
    if (newpass1 === "" || newpass2 === "") {
        console.log("favor de llenar los campos")
        return
    }
    if (newpass1 != newpass2) {
        console.log("las contraseñas no coinciden")
        return
    }
    let newpass = new NewPassword(currentpass, newpass1)
    let xhr = new XMLHttpRequest();
    xhr.open('PUT', `${URL_SERVER}/account`);
    xhr.setRequestHeader('email', get_email())
    xhr.setRequestHeader('token', get_token());

    xhr.send(JSON.stringify(newpass));
    xhr.onload= (req, res) => {
        if (xhr.status == 202) {
            let data = JSON.parse(xhr.responseText)
            let token = xhr.getResponseHeader('token')
            console.log(data);
            unlogger()
            logger(data, token)
            console.log("Redirect to home")
            window.location.href = 'http://localhost:8080/home'
            window.location.href = 'http://localhost:8080/home';
        }
        else {
            console.log("error trying to update the account");
        }
    }
})

class NewPassword {
    constructor(password, newpassword) {
        this.password = password;
        this.new_password = newpassword;
    }
}