drop database if exists social;

create database social
with
    template = template0
    encoding = 'UTF8'
    lc_collate = 'pt_BR.UTF-8'
    lc_ctype = 'pt_BR.UTF-8';

\connect social

create table usuarios (
	id serial primary key,
	nome varchar(127) not null,
	nick varchar(50) not null unique,
	email varchar(50) not null unique,
	senha varchar(100) not null,
	criadoEm timestamp default now()
);


CREATE TABLE seguidores (
	usuario_id bigint NOT NULL,
	seguidor_id bigint NOT NULL,

	constraint seguidores_usuario_id_fk foreign key (usuario_id) references usuarios(id),
	constraint seguidores_seguidor_id_fk foreign key (seguidor_id) references usuarios(id),
	primary key (usuario_id, seguidor_id)
);


CREATE TABLE publicacoes (
    id serial primary key,
    titulo varchar(50) not null,
    conteudo varchar(300) not null,
    autor_id bigint not null,
    curtidas INT default 0,
    criadoEm timestamp default now(),

    constraint publicacoes_autor_id_fk foreign key (autor_id) references usuarios(id) on delete cascade
);
