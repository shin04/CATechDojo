create table usercharacters
(
    id int primary key,
    user_id int not null,
    foreign key (user_id) references users(id),
    character_id int not null,
    foreign key (character_id) references characters(id),
    possessions int not null
);