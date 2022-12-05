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

        $("#seguir").on("click", seguir);
        $("#parar-de-seguir").on("click", pararDeSeguir);
    });
}(jQuery));