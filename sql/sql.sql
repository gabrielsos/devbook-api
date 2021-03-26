create database if not exists devbook;

use devbook;

drop TABLE if exists user;

create table user(
  id int AUTO_INCREMENT PRIMARY KEY,
  name varchar(50) not null,
  nick varchar(50) not null unique,
  email varchar(50) not null unique,
  password varchar(20) not null,
  createdAt timestamp default CURRENT_TIMESTAMP
) ENGINE=INNODB; 