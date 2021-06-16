create database if not exists `devbook`;

create table if not exists `user` (
	id int primary key not null auto_increment,
    name varchar(100) not null,
    email varchar(100) not null unique,
    nick varchar(50),
    password varchar(50),
    created_at datetime default current_timestamp
);

insert into `user` (name, email, nick, password) values ('erislandio soares', 'erislandiosoares@gmail.com', 'eris', '12345678');

select * from user;