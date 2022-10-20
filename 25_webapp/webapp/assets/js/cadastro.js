
/**
 * @param {Event} event
 * @returns {undefined}
 */
function criarUsuario(event) {
    event.preventDefault();

    var form = $(event.target);

    if (form.find('#senha').val() != form.find('#confirmar-senha').val()) {
        alert('As senhas não coincidem');
        return;
    }

    $.ajax({
        url: "/usuarios",
        method: "POST",
        data: {
            nome: form.find("#nome").val(),
            email: form.find("#email").val(),
            nick: form.find("#nick").val(),
            senha: form.find("#senha").val()
        }
    }).done(function() {
        alert("Usuário cadastrado com sucesso!");
    }).fail(function(erro) {
        alert("Erro ao cadastrar usuário.");
        console.log(erro)
    });
}


$('#formulario-cadastro').on('submit', criarUsuario);