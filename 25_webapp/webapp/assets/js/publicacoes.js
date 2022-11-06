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
        }).fail(function() {
            alert("Erro ao curtir publicação");
        }).always(function() {
            target.prop("disabled", false)
        });
    }

    $("#nova-publicacao").on("submit", criarPublicacao);
    $(".curtir-publicacao").on("click", curtirPublicacao);
});
