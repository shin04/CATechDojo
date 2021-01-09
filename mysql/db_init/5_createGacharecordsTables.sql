create table gacharecords
(
    id int primary key,
    user_id int not null,
    foreign key (user_id) references users(id),
    character_id int not null,
    foreign key (character_id) references characters(id),
    ts timestamp not null
);