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

        $("#seguir").on("click", seguir);
        $("#parar-de-seguir").on("click", pararDeSeguir);
        $("#editar-usuario").on("submit", editarUsuario);
    });
}(jQuery));
