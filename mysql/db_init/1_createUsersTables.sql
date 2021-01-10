use game_db;

create table users
(
    id int auto_increment primary key,
    name varchar(64) not null,
    token varchar(256) not null
);

insert into users (name, token) values ('bob', 'bob_token');