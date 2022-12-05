(function ($) {
    $(function() {

        /**
         * @param {Event} event
         * @returns {undefined}
         */
        function criarUsuario(event) {
            event.preventDefault();

            var form = $(event.target);

            if (form.find('#senha').val() != form.find('#confirmar-senha').val()) {
                Swal.fire("Ops...", "As senhas não coincidem", "error");
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
                Swal.fire("Sucesso!", "Usuário cadastrado com sucesso!", "success")
                    .then(function() {
                        window.location = "/";
                    });
            }).fail(function(erro) {
                Swal.fire("Ops...", "Erro ao cadastrar usuário.", "error");
                console.log(erro)
            });
        }

        $('#formulario-cadastro').on('submit', criarUsuario);
    });
}(jQuery));
