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
            console.log(data);
        }
    }
})