CREATE TABLE users (
    id int primary key auto_increment,
    name varchar(255) not null,
    phone_number varchar(255) not null unique,
    created_at datetime DEFAULT CURRENT_TIMESTAMP
);