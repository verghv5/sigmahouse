//Validation function
function validate(){
    var request = require(['request']);
    var username = document.getElementById("username").value;
    var password = document.getElementById("password").value;

    var encrypted = CryptoJS.AES.encrypt(password, "BOB JONES HEH");
    var decrypted = CryptoJS.AES.decrypt(encrypted, "BOB JONES HEH");
    console.log(decrypted, )

    var credentials = [{ "username": username }, { "password": encrypted }];

    $.ajax({
        type: "POST",
        url: "http://localhost:8000/login",
        // The key needs to match your method's input parameter (case-sensitive).
        data: JSON.stringify({ Credentials: credentials }),
        contentType: "application/json; charset=utf-8",
        dataType: "json",
        success: function(data){alert(data);},
        failure: function(errMsg) {
            alert(errMsg);
        }
    });
}
