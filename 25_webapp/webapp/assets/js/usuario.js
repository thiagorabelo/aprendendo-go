(function($){
    $(function() {

        /**
         *
         * @param {Event} event
         */
        function pararDeSeguir(event) {
            event.preventDefault();
            const target = $(event.target);
            const usuarioId = target.data('usuario-id');

            target.prop("disabled", true);

            $.ajax({
                url: `/usuarios/${usuarioId}/parar-de-seguir`,
                method: "POST"
            }).done(function() {
                window.location.reload();
            }).fail(function() {
                Swal.fire("Ops...", "Erro ao parar de seguir o usuário!", "error");
                target.prop("disabled", false);
            });
        }

        /**
         *
         * @param {Event} event
         */
        function seguir(event) {
            event.preventDefault();
            const target = $(event.target);
            const usuarioId = target.data('usuario-id');

            target.prop("disabled", true);

            $.ajax({
                url: `/usuarios/${usuarioId}/seguir`,
                method: "POST"
            }).done(function() {
                window.location.reload();
            }).fail(function() {
                Swal.fire("Ops...", "Erro ao seguir o usuário!", "error");
                target.prop("disabled", false);
            });
        }

        /**
         *
         * @param {Event} event
         */
        function editarUsuario(event) {
            event.preventDefault();

            const form = $(event.target);
            const nome = form.find("#nome");
            const email = form.find("#email");
            const nick = form.find("#nick");
            const btn = form.find("button[type=submit]");

            btn.prop("disabled", true);

            $.ajax({
                url: "/editar-usuario",
                method: "PUT",
                data: {
                    nome: nome.val(),
                    email: email.val(),
                    nick: nick.val()
                }
            }).done(function() {
                Swal.fire("Sucesso!", "Usuário atualizado com sucesso", "success")
                    .then(function() {
                        window.location = "/perfil";
                    });
            }).fail(function() {
                Swal.fire("Ops...", "Erro ao atualizar o usuário", "error");
            }).always(function() {
                btn.prop("disabled", false);
            });
        }

        /**
         *
         * @param {Event} event
         */
        function atualizarSenha(event) {
            event.preventDefault();

            const form = $(event.target);
            const senhaAtual = form.find("#senha-atual");
            const novaSenha = form.find("#nova-senha");
            const confirmarSenha = form.find("#confirmar-senha");
            const btn = form.find("button[type=submit]");


            if (novaSenha.val() != confirmarSenha.val()) {
                Swal.fire("Ops...", "Os campos Nova Senha e Confirmar Senha devem coincidirem", "warning");
                return;
            }

            btn.prop("disabled", true);

            $.ajax({
                url: "/atualizar-senha",
                method: "POST",
                data: {
                    atual: senhaAtual.val(),
                    nova: novaSenha.val()
                }
            }).done(function() {
                Swal.fire("Sucesso!", "Senha atualizada", "success")
                    .then(function() {
                        window.location = "/perfil";
                    });
            }).fail(function() {
                Swal.fire("Ops...", "Erro ao atualizar a senha", "error");
            }).always(function() {
                btn.prop("disabled", false);
            });
        }

        /**
         *
         * @param {Event} event
         */
        function deletarUsuario(event) {
            event.preventDefault();

            Swal.fire({
                title: "Atenção!",
                text: "Tem certeza que deseja apagar a sua conta? Esta é uma ação irreversível.",
                showCancelButton: true,
                cancelButtonText: "Cancelar",
                icon: "warning"
            }).then(function(confirmacao) {
                if (confirmacao.value) {
                    $.ajax({
                        url: "/deletar-usuario",
                        method: "DELETE"
                    }).done(function() {
                        Swal.fire("Sucesso!", "Sua conta foi excluída", "success")
                            .then(function() {
                                window.location = "/logout";
                            });
                    }).fail(function() {
                        Swal.fire("Ops...", "Ocorreu um erro ao excluir a sua conta", "error");
                    });
                }
            });
        }

        $("#seguir").on("click", seguir);
        $("#parar-de-seguir").on("click", pararDeSeguir);
        $("#editar-usuario").on("submit", editarUsuario);
        $("#atualizar-senha").on("submit", atualizarSenha);
        $("#deletar-usuario").on("click", deletarUsuario);
    });
}(jQuery));
