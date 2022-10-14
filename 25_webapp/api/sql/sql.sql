CREATE DATABASE IF NOT EXISTS social;
USE social;

DROP TABLE IF EXISTS seguidores;
DROP TABLE IF EXISTS usuarios;

CREATE TABLE usuarios (
    id int auto_increment primary key,
    nome varchar(127) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    senha varchar(100) not null,
    criadoEm timestamp default current_timestamp()
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
