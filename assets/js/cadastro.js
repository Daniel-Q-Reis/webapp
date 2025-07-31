//comando reconhecido pelo jquery, quando um item tiver um id chamado formulario-cadastro, eu vou atrelar um evento a ele, chamar a func criarUsuario
$('#formulario-cadastro').on('submit', criarUsuario)


function criarUsuario(evento) {
    evento.preventDefault();

    if ( $('#senha').val() != $('#confirmar-senha').val()) { //o jquery vai procurar as infomações la no html que possuam esses nomes, e vai pegar o valor deles
        Swal.fire("Ops...", "As senhas não coincidem!", "error"); //vai aparecer um pop-up na tela, dizendo isso
        return;
    }

    //agora vamos procurar dentro da nossa aplicação web uma url = URI que possua /usuarios
    $.ajax({
        url: "/usuarios",
        method: "POST",
        data: {//data passa todos os campos que vamos mandar para a nossa rota (nome, email, nick, senha)
            nome: $('#nome').val(),
            email: $('#email').val(),
            nick: $('#nick').val(),
            senha: $('#senha').val(),
        }
 }).done(function() {
        Swal.fire("Sucesso!", "Usuário cadastrado com sucesso!", "success")
            .then(function() {
                $.ajax({
                    url: "/login",
                    method: "POST",
                    data: {
                        email: $('#email').val(),//Aqui como email e senha acabaram de ser cadastrados, sempre vai dar certo
                        senha: $('#senha').val()
                    }
                }).done(function() {
                    window.location = "/home";//logo depois de criar o usuário, já somos redirecionados para tela principal
                }).fail(function() {
                    Swal.fire("Ops...", "Erro ao autenticar o usuário!", "error");
                })
            })
    }).fail(function() {
        Swal.fire("Ops...", "Erro ao cadastrar o usuário!", "error");
    });
}