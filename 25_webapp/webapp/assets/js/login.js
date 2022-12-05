(function($) {
    $(function() {
        /**
         *
         * @param {Event} event
         * @returns {undefined}
         */
        function fazerLogin(event) {
            event.preventDefault();

            var form = $(event.target);

            $.ajax({
                url: "/login",
                method: "post",
                data: {
                    email: form.find("#email").val(),
                    senha: form.find("#senha").val()
                }
            }).done(function () {
                window.location = "/home";
            }).fail(function () {
                Swal.fire("Ops...", "Usuário ou senha inválidos.", "error");
            });
        }

        $("#login").on("submit", fazerLogin);
    });
}(jQuery));
