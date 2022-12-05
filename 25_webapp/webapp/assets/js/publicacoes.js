(function($) {
    $(function() {
        /**
         * @param {Event} event
         */
        function criarPublicacao(event) {
            event.preventDefault();
            const form = $(event.target);
            const tituloField = form.find("#titulo");
            const conteudoField = form.find("#conteudo");

            $.ajax({
                url: "/publicacoes",
                method: "POST",
                data: {
                    titulo: tituloField.val(),
                    conteudo: conteudoField.val()
                }
            }).done(function(){
                tituloField.val("");
                conteudoField.val("");
                window.location.reload();
            }).fail(function(){
                Swal.fire("Ops...", "Erro ao criar publicação", "error");
            });
        }

        /**
         * @param {Event} event
         */
        function atualizarPublicacao(event) {
            event.preventDefault();
            const form = $(event.target);
            const btn = form.find("button[type=submit]");
            const publicacaoId = btn.data("publicacao-id");
            btn.prop("disabled", true);

            $.ajax({
                url: `/publicacoes/${publicacaoId}`,
                method: "PUT",
                data: {
                    titulo: form.find("#titulo").val(),
                    conteudo: form.find("#conteudo").val()
                }
            }).done(function() {
                Swal.fire(
                    "Sucesso!",
                    "Publicação criada com sucesso!",
                    "success"
                ).then(function() {
                    window.location = "/home";
                })
            }).fail(function() {
                Swal.fire("Ops...", "Erro ao editar publicação!", "error");
            }).always(function() {
                btn.prop("disabled", false);
            })
        }

        /**
         *
         * @param {Event} event
         */
        function curtirPublicacao(event) {
            event.preventDefault();

            const target = $(event.target);
            const publicacaoContainer = target.closest(".publicacao-container");
            const publicacaoId = publicacaoContainer.data('publicacao-id')

            target.prop("disabled", true)
            $.ajax({
                url: `/publicacoes/${publicacaoId}/curtir`,
                method: "POST"
            }).done(function() {
                const contadorDeCurtidas = target.next("span");
                const totalCurtidas = parseInt(contadorDeCurtidas.text());
                contadorDeCurtidas.text(totalCurtidas + 1);

                target
                    .removeClass("curtir-publicacao")
                    .addClass("descurtir-publicacao")
                    .addClass("text-danger")
                ;
            }).fail(function() {
                Swal.fire("Ops...", "Erro ao curtir publicação", "error");
            }).always(function() {
                target.prop("disabled", false)
            });
        }

        /**
         *
         * @param {Event} event
         */
        function descurtirPublicacao(event) {
            event.preventDefault();

            const target = $(event.target);
            const publicacaoContainer = target.closest(".publicacao-container");
            const publicacaoId = publicacaoContainer.data('publicacao-id')

            target.prop("disabled", true)
            $.ajax({
                url: `/publicacoes/${publicacaoId}/descurtir`,
                method: "POST"
            }).done(function() {
                const contadorDeCurtidas = target.next("span");
                const totalCurtidas = parseInt(contadorDeCurtidas.text());
                contadorDeCurtidas.text(totalCurtidas - 1);

                target
                    .addClass("curtir-publicacao")
                    .removeClass("descurtir-publicacao")
                    .removeClass("text-danger")
                ;
            }).fail(function() {
                Swal.fire("Ops...", "Erro ao descurtir publicação", "error");
            }).always(function() {
                target.prop("disabled", false)
            });
        }

        /**
         *
         * @param {Event} event
         */
        function deletarPublicacao(event) {
            event.preventDefault();

            Swal.fire({
                title: "Atenção!",
                text: "Esta ação é irreverssível. Deseja realmente excluir a publicação?",
                showCancelButton: true,
                cancelButtonText: "Cancelar",
                icon: "warning"
            }).then(function(confirmacao) {
                if (!confirmacao.value) return;

                const target = $(event.target);
                const publicacaoContainer = target.closest(".publicacao-container")
                const publicacaoId = publicacaoContainer.data("publicacao-id");

                target.prop("disabled", true);
                $.ajax({
                    url: `/publicacoes/${publicacaoId}`,
                    method: "DELETE"
                }).done(function() {
                    publicacaoContainer.animate({opacity: 0}, 200, function() {
                        $(this).animate({height: 0, paddingBottom: 0, paddingTop: 0}, 200, function() {
                            $(this).remove();
                            Swal.fire({
                                // position: 'top-end',
                                icon: 'success',
                                title: 'A publicação foi apagada.',
                                showConfirmButton: false,
                                timer: 1500
                            });
                        });
                    });
                }).fail(function(){
                    Swal.fire("Ops...", "Erro ao excluir publicação", "error");
                }).always(function() {
                    target.prop("disabled", false);
                });
            });
        }

        $("#nova-publicacao").on("submit", criarPublicacao);
        $("#editar-publicacao").on("submit", atualizarPublicacao);
        $(".deletar-publicacao").on("click", deletarPublicacao);

        $(document).on("click", ".curtir-publicacao", curtirPublicacao);
        $(document).on("click", ".descurtir-publicacao", descurtirPublicacao);
    });
}(jQuery));
