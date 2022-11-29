$(function() {
    /**
     * @param {Event} event
     */
    function criarPublicacao(event) {
        event.preventDefault();
        const form = $(event.target);

        $.ajax({
            url: "/publicacoes",
            method: "POST",
            data: {
                titulo: form.find("#titulo").val(),
                conteudo: form.find("#conteudo").val()
            }
        }).done(function(){
            window.location.reload();
        }).fail(function(){
            alert("Erro ao criar publicação");
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
            alert("Publicação editada com sucesso!");
        }).fail(function() {
            alert("Erro ao editar publicação!");
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
            alert("Erro ao curtir publicação");
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
            alert("Erro ao descurtir publicação");
        }).always(function() {
            target.prop("disabled", false)
        });
    }

    $("#nova-publicacao").on("submit", criarPublicacao);
    $("#editar-publicacao").on("submit", atualizarPublicacao);

    $(document).on("click", ".curtir-publicacao", curtirPublicacao);
    $(document).on("click", ".descurtir-publicacao", descurtirPublicacao);
});
