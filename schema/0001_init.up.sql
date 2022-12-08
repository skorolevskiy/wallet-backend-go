CREATE TABLE users (
    id serial not null unique,
    email varchar(255) not null unique,
    username varchar(255) not null unique,
    password varchar(255) not null,
    register_at TIMESTAMP default now()
);

CREATE TABLE wallets (
    id serial PRIMARY KEY,
    user_id int not null,
    name varchar(255) not null,
    balance float not null,
    currency varchar(255) not null,
    register_at TIMESTAMP default now()
);

CREATE TABLE transactions (
    id serial PRIMARY KEY,
    wallet_id int not null,
    amount float not null,
    balance_after float not null,
    commission_amount float,
    currency varchar(255) not null,
    created_at TIMESTAMP default now()
);
