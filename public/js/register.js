function verifyPassword(){
    var basePass = $("#password");
    var confPass = $("confPassword");
    if (basePass === confPassword){
        return false
    }
    return true
}
$(document).ready(function(){


    $("#registerForm").submit(function(e){
        var registerForm ={
            login: $("#login").val(),
            mail: $("#mail").val(),
            password: $("#password").val(),
            confPassword: $("#confPassword").val()
        };
        console.log(registerForm);
        var registerUrl = "/register";
        $.ajax({
			url: registerUrl,
            type:"POST",
            data: JSON.stringify(registerForm),
            async: true,
            contentType: 'application/json; charset=UTF-8',
			success: function(data)
			{
                console.log(data);
            },
            error: function(XMLHttpRequest, textStatus, errorThrown){
                console.log("Error : "+ textStatus + errorThrown);
            }
        });


        //e.preventDefault();



    });




});
