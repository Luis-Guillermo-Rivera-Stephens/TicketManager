starter()

document.getElementById("login").addEventListener('click', (event)=>{

    event.preventDefault();
    console.log("click a log in\n")
    let email = document.getElementById("email").value;
    let password = document.getElementById("password").value;

    let xhr = new XMLHttpRequest();
    xhr.open('GET', 'http://localhost:8080/account');
    xhr.setRequestHeader('email', email)
    xhr.setRequestHeader('password', password)
    xhr.send();
    xhr.onload= (req, res) => {
        if (xhr.status == 200) {
            var data = JSON.parse(xhr.responseText);
            var token = xhr.getResponseHeader("token")
            console.log(data, "  token:  ", token);
            logger(data, token)
            console.log(data.isstarted)
            if (data.isstarted == false){
                console.log("Redirect to update")
                window.location.href = 'http://localhost:8080/update';
            } else {
                console.log("Redirect to home")
                window.location.href = 'http://localhost:8080/home';
            }
        }
        else {
            console.log("error trying to get the account");
        }
    }
})
