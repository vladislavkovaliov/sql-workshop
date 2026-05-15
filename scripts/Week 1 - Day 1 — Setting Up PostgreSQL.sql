/* Create table and user */

CREATE USER shop_user WITH PASSWORD 'password';

CREATE DATABASE shop OWNER shop_user;

GRANT ALL PRIVILEGES ON DATABASE shop TO shop_user;

/* Create tables */

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL
);

CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    price NUMERIC(10,2) NOT NULL
);

CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
