function isUpper(symbol) {
    return((symbol >= 'A' && symbol <= 'Z') || (symbol >= 'А' && symbol <= 'Я'))
}

function isLower(symbol) {
    return((symbol >= 'a' && symbol <= 'z') || (symbol >= 'а' && symbol <= 'я'))
}

function isDigit(symbol) {
    return(symbol >= '0' && symbol <= '9')
}

function passwordComplexityCheck() {
    var passwordString = document.getElementById('password').value;

    let password = {
        upperCase : false,
        lowerCase : false,
        digits : false,
        lengthOf10 : false,
        specialSymbols : false
    }
    password.result = function() {
        return(password.upperCase + password.lowerCase + password.digits + password.lengthOf10 + password.specialSymbols)
    }

    var n = passwordString.length;
    for(let i = 0; i < n; i++) {
        if(isUpper(passwordString[i])) password.upperCase = true;
        else if(isLower(passwordString[i])) password.lowerCase = true;
        else if(isDigit(passwordString[i])) password.digits = true;
        else password.specialSymbols = true;
    }
    if(n > 9) password.lengthOf10 = true;

    var complexity = password.result();
    var indicator = document.getElementById('passwordIndicator');
    var status = document.getElementById('passwordStatus')
    switch(complexity) {
        case 1:
            indicator.style.width = "5%";
            indicator.style.backgroundColor = "red";
            status.style.color = "red";
            status.innerHTML = "Weak";
            break;
        case 2:
            indicator.style.width = "20%";
            indicator.style.backgroundColor = "orange";
            status.style.color = "orange";
            status.innerHTML = "Insecure";
            break;
        case 3:
            indicator.style.width = "45%"
            indicator.style.backgroundColor = "#e1e729";
            status.style.color = "#e1e729";
            status.innerHTML = "Moderate";
            break;
        case 4:
            indicator.style.width = "72%";
            indicator.style.backgroundColor = "#76ff00b5";
            status.style.color = "#76ff00b5";
            status.innerHTML = "Good";
            break;
        case 5:
            indicator.style.width = "100%";
            indicator.style.backgroundColor = "#00ff00a6";
            status.style.color = "#00ff00a6";
            status.innerHTML = "Strong";
            break;
        default:
            indicator.style.width = "0%"
            indicator.style.backgroundColor = "#919191";
            status.style.color = "#919191";
            status.innerHTML = "";
            break;
    }
}

var input = document.getElementById("password");
var warning = document.getElementById("warn");
warn.style.display = "none";
warn.innerHTML = "CapsLock is ON!";
warn.style.color = "orange";
warning.style.marginLeft = "17px"
input.addEventListener("keyup", function(event) {
    if (event.getModifierState("CapsLock")) {
        warning.style.display = "block";
    } else {
        warning.style.display = "none"
    }
})

document.addEventListener("input", function () {
    const name = document.getElementById("name").value;
    const surname = document.getElementById("surname").value;
    const email = document.getElementById("email").value;
    const dob = document.getElementById("dob").value;
    const password = document.getElementById("password").value;
    const confirmPassword = document.getElementById("confirm-password").value;
    const registerButton = document.getElementById("registerButton");

    if (name && surname && isValidEmail(email) && dob && password && confirmPassword && password === confirmPassword) {
        registerButton.disabled = false;
    } else {
        registerButton.disabled = true;
    }
});

var today = new Date();

// Format today's date as YYYY-MM-DD
var yyyy = today.getFullYear();
var mm = String(today.getMonth() + 1).padStart(2, '0'); // January is 0!
var dd = String(today.getDate()).padStart(2, '0');
var maxDate = yyyy + '-' + mm + '-' + dd;

document.getElementById("dob").max = maxDate;

function isValidEmail(email) {
    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;


    return emailRegex.test(email);
}