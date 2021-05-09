create database if not exists devbooks;

use devbooks;

drop table if exists usuarios;

create table usuarios(
    id int auto_increment primary key,
    nome varchar(255),
    nick varchar(255)  not null unique,
    email varchar(255)  not null unique,
    senha varchar(255)  not null unique,
    criadoEm timestamp default current_timestamp()
);