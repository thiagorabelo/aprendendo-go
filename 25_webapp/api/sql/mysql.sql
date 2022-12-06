CREATE DATABASE IF NOT EXISTS social;
USE social;

DROP TABLE IF EXISTS publicacoes;
DROP TABLE IF EXISTS seguidores;
DROP TABLE IF EXISTS usuarios;

CREATE TABLE usuarios (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nome VARCHAR(127) NOT NULL,
    nick VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(50) NOT NULL UNIQUE,
    senha VARCHAR(100) NOT NULL,
    criadoEm TIMESTAMP DEFAULT CURRENT_TIMESTAMP()
) ENGINE=INNODB;

/*

GRANT ALL PRIVILEGES ON <database> TO '<myuser>'@'%' IDENTIFIED BY '<mypaassword>';
GRANT ALL PRIVILEGES ON <database> TO '<myuser>'@'localhost' IDENTIFIED BY '<mypaassword>';
GRANT ALL PRIVILEGES ON <database> TO '<myuser>'@'%' WITH GRANT OPTION;

FLUSH PRIVILEGES;

*/

CREATE TABLE seguidores (
	usuario_id INT NOT NULL,
	FOREIGN KEY (usuario_id)
	REFERENCES usuarios(id)
	ON DELETE CASCADE,

	seguidor_id INT NOT NULL,
	FOREIGN KEY (seguidor_id)
	REFERENCES usuarios(id)
	ON DELETE CASCADE,

	PRIMARY KEY(usuario_id, seguidor_id)
) ENGINE=INNODB;


CREATE TABLE publicacoes (
    id INT AUTO_INCREMENT PRIMARY KEY,
    titulo VARCHAR(50) NOT NULL,
    conteudo VARCHAR(300) NOT NULL,

    autor_id INT NOT NULL,
    FOREIGN KEY (autor_id)
    REFERENCES usuarios(id)
    ON DELETE CASCADE,

    curtidas INT DEFAULT 0,

    criadaEm TIMESTAMP DEFAULT CURRENT_TIMESTAMP
) ENGINE=INNODB;
