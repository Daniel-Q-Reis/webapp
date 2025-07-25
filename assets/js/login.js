$('#login').on('submit', fazerLogin);

function fazerLogin(evento) {
    evento.preventDefault(); //Aqui estamos previnindo o comportamento padrão do formulário

    $.ajax({
        url: "/login",
        method: "POST",
        data: {
            email: $('#email').val(), //está pegando os valores la de login.html
            senha: $('#senha').val(),//está pegando os valores la de login.html
        }
    }).done(function() {
        window.location = "/home"; //redirecionar para a pagina que será a pagina principal da pessoa (o feed), pois obteve sucesso no login
    }).fail(function() {
        Swal.fire("Ops...", "Usuário ou senha incorretos!", "error");
    });
}