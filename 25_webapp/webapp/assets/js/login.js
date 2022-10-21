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
            alert("Usuário ou senha inválidos");
        });
    }

    $("#login").on("submit", fazerLogin);
});
