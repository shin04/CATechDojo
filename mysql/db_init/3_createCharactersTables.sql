create table characters
(
    id int auto_increment primary key,
    name varchar(64) not null,
    rarity_id int not null,
    foreign key (rarity_id) references rarities(id)
);