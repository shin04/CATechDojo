create table rarities
(
    id int auto_increment primary key,
    value int not null,
    weight int not null
);

insert into rarities (value, weight) values (1, 80);