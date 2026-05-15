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


INSERT INTO users (name, email) VALUES
('Alice Johnson', 'alice@example.com'),
('Bob Smith', 'bob@example.com'),
('Charlie Brown', 'charlie@example.com'),
('Diana Prince', 'diana@example.com'),
('Ethan Hunt', 'ethan@example.com');

INSERT INTO products (title, price) VALUES
('Keyboard', 45.99),
('Mouse', 25.50),
('Monitor', 199.99),
('Laptop Stand', 120.00),
('Mechanical Keyboard', 149.90);

INSERT INTO orders (user_id, created_at) VALUES
(1, NOW() - INTERVAL '5 days'),
(2, NOW() - INTERVAL '4 days'),
(3, NOW() - INTERVAL '3 days'),
(1, NOW() - INTERVAL '2 days'),
(5, NOW() - INTERVAL '1 day'),
(2, NOW());


select * from users;

select * from products;

select * from products where price > 100;

select * from orders order by created_at DESC limit 5;

update products p set price = 121 where p.id = 3;

insert into products (title, price) values 
('fake', 1);

delete from products p where id = 6;


