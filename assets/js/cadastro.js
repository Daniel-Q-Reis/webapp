//comando reconhecido pelo jquery, quando um item tiver um id chamado formulario-cadastro, eu vou atrelar um evento a ele, chamar a func criarUsuario
$('#formulario-cadastro').on('submit', criarUsuario)


function criarUsuario(evento) {
    evento.preventDefault();
    console.log("Dentro da funcao usuario"); //fmt.Println() -> so que loga la no console do browser (vai imprimir la) e não no VsCode

    if ( $('#senha').val() != $('#confirmar-senha').val()) { //o jquery vai procurar as infomações la no html que possuam esses nomes, e vai pegar o valor deles
        alert("As senhas não coincidem!"); //vai aparecer um pop-up na tela, dizendo isso
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
    });
}