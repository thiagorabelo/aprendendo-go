INSERT INTO usuarios (nome, nick, email, senha) /* senha: 123456 */
VALUES
    ('Usuario 1', 'usuario_1', 'usuario1@email.com', '$2a$10$p2fI3tQ23u6GBzvUcq35FO3LUBp.9qbdoWRYXkcZOiofsL6iK1VNa'),
    ('Usuario 2', 'usuario_2', 'usuario2@email.com', '$2a$10$p2fI3tQ23u6GBzvUcq35FO3LUBp.9qbdoWRYXkcZOiofsL6iK1VNa'),
    ('Usuario 3', 'usuario_3', 'usuario3@email.com', '$2a$10$p2fI3tQ23u6GBzvUcq35FO3LUBp.9qbdoWRYXkcZOiofsL6iK1VNa')
;

INSERT INTO seguidores(usuario_id, seguidor_id)
VALUES
    (1, 2),
    (3, 1),
    (1, 3)
;
