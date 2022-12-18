CREATE TABLE users (
    id serial not null unique,
    email varchar(255) not null unique,
    username varchar(255) not null unique,
    password varchar(255) not null,
    register_at TIMESTAMP default now()
);

CREATE TABLE wallets (
    id serial PRIMARY KEY,
    user_id int not null REFERENCES users(id) ON DELETE CASCADE,
    name varchar(255) not null,
    balance float not null,
    currency varchar(255) not null,
    register_at TIMESTAMP default now()
);

CREATE TABLE transactions (
    id serial PRIMARY KEY,
    wallet_id int not null REFERENCES wallets(id) ON DELETE CASCADE,
    description varchar(255),
    amount float not null,
    balance_after float not null,
    commission_amount float,
    currency varchar(255) not null,
    created_at TIMESTAMP default now()
);

CREATE TABLE refresh_tokens (
    id serial PRIMARY KEY,
    user_id int not null REFERENCES users(id) ON DELETE CASCADE,
    token varchar(255) not null,
    expires_at TIMESTAMP not null
);
