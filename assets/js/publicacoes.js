$('#nova-publicacao').on('submit', criarPublicacao);

$(document).on('click', '.curtir-publicacao', curtirPublicacao);//referenciar com o document é porque haverá mudança na class no html
$(document).on('click', '.descurtir-publicacao', descurtirPublicacao);//essa é uma classe que não vai ser carregada inicialmente na pagina, somente será chamada apois curtir-mos uma vez

$('#atualizar-publicacao').on('click', atualizarPublicacao);
$('.deletar-publicacao').on('click', deletarPublicacao);

//Todas as funções Swal. que é um pop-up personalizado de interação com o usuário, vem do seguinte local: templates/scrips.html -> <script src="https://cdn.jsdelivr.net/npm/sweetalert2@11"></script>

function criarPublicacao(evento) {
    evento.preventDefault(); //parar o comportamento padrao do envio de formulario

    $.ajax({
        url: "/publicacoes",
        method: "POST",
        data: {//vamos pegar os dados do titulo e do conteudo da publicação e por no data do ajax
            titulo: $('#titulo').val(),//titulo e conteudo também estão la no html como input e text area (sao form-groups)
            conteudo: $('#conteudo').val(),//hashtab é para pegar por id
        }
    }).done(function() {
        window.location = "/home";
    }).fail(function() {
        Swal.fire("Ops...", "Erro ao criar a publicação!", "error");
    })
}

function curtirPublicacao(evento) {
    evento.preventDefault();

    const elementoClicado = $(evento.target);//fazendo isso consigo pegar especificamente o coraçãozinho que foi clicado
    const publicacaoId = elementoClicado.closest('div').data('publicacao-id');//vai subindo na arvore da div mais proxima, até encontrar o publicacao-id

    elementoClicado.prop('disabled', true); //vai desabilitar o botão curtir termporariamente, que será liberado apos a requisição abaixo rodar
    $.ajax({
        url: `/publicacoes/${publicacaoId}/curtir`, //aqui temos uma URL com variavel dentro ${publicacaoID}, usamos crase
        method: "POST"
    }).done(function() {
        const contadorDeCurtidas = elementoClicado.next('span');
        const quantidadeDeCurtidas = parseInt(contadorDeCurtidas.text());

        contadorDeCurtidas.text(quantidadeDeCurtidas + 1);

        elementoClicado.addClass('descurtir-publicacao');
        elementoClicado.addClass('text-danger');//sempre vai deixar o coraçãozinho vermelho depois que o usuário curtir
        elementoClicado.removeClass('curtir-publicacao');

    }).fail(function() {
        Swal.fire("Ops...", "Erro ao curtir a publicação!", "error");
    }).always(function() {
        elementoClicado.prop('disabled', false);//aqui libera o botão curtir denovo
    });
}

function descurtirPublicacao(evento) {//exatamente a mesma função que curtir, mas vai subitrari a quantidade de curtidas e vai chamar a url descurtir ao inves de curtir
    evento.preventDefault();

    const elementoClicado = $(evento.target);
    const publicacaoId = elementoClicado.closest('div').data('publicacao-id');

    elementoClicado.prop('disabled', true);
    $.ajax({
        url: `/publicacoes/${publicacaoId}/descurtir`,
        method: "POST"
    }).done(function() {
        const contadorDeCurtidas = elementoClicado.next('span');
        const quantidadeDeCurtidas = parseInt(contadorDeCurtidas.text());

        contadorDeCurtidas.text(quantidadeDeCurtidas - 1);

        elementoClicado.removeClass('descurtir-publicacao');//aqui damos removeClass ao invés de addClass usado em curtir
        elementoClicado.removeClass('text-danger');
        elementoClicado.addClass('curtir-publicacao');//aqui resetamos o processo, e se clicar no coração dnv chama o curtir publicacao novamente (cria tipo um loop entre as duas)

    }).fail(function() {
        Swal.fire("Ops...", "Erro ao descurtir a publicação!", "error");
    }).always(function() {
        elementoClicado.prop('disabled', false);
    });
}

function atualizarPublicacao() {//nao precisa passar o evento, pois vamos usar o this para se referir ao botão que foi clicado
    $(this).prop('disabled', true);//desabilita o botão que foi clicado (this = referencia ao botão)

    const publicacaoId = $(this).data('publicacao-id');//no nosso botão em atualizar-publicacao.html temos data-publicacao-id
    
    $.ajax({
        url: `/publicacoes/${publicacaoId}`,
        method: "PUT",
        data: {
            titulo: $('#titulo').val(),
            conteudo: $('#conteudo').val()
        }
    }).done(function() {
        Swal.fire('Sucesso!', 'Publicação criada com sucesso!', 'success')
            .then(function() {
                window.location = "/home";//aqui apos o pop-up swal ele irá nos direncionar devolta para /home, apos clicarmos no retorno dele (OK)
            })
    }).fail(function() {
        Swal.fire("Ops...", "Erro ao editar a publicação!", "error");
    }).always(function() {
        $('#atualizar-publicacao').prop('disabled', false);
    //aqui não podemos usar o this, pois estamos usando uma função nova dentro da função então devemos referenciar novamente o atualizar publicacao para reativar o botao
    })
}

function deletarPublicacao(evento) {
    evento.preventDefault();

    //aqui seria como se fosse uma verificação de duas etapas, pois é um caminho sem volta, então exige uma maior atenção do usuário
    Swal.fire({
        title: "Atenção!",
        text: "Tem certeza que deseja excluir essa publicação? Essa ação é irreversível!",
        showCancelButton: true,
        cancelButtonText: "Cancelar",
        icon: "warning"
    }).then(function(confirmacao) {
        if (!confirmacao.value) return;

        const elementoClicado = $(evento.target);
        const publicacao = elementoClicado.closest('div')
        const publicacaoId = publicacao.data('publicacao-id');
    
        elementoClicado.prop('disabled', true);
    
        $.ajax({
            url: `/publicacoes/${publicacaoId}`,
            method: "DELETE"
        }).done(function() {
            publicacao.fadeOut("slow", function() { //aqui estamos ocultamos o jumbotron (se o delete for um sucesso), não queremos mais ele na tela
                $(this).remove();//aqui remove o jumbotron além de termos ocultado-o, não queremos mais ele em nosso html apos o delete
            });
        }).fail(function() {
            Swal.fire("Ops...", "Erro ao excluir a publicação!", "error");
        });
    })

}