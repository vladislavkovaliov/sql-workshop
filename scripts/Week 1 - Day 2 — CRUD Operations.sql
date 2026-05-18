create table users (
    id SERIAL primary key,
    name TEXT not null,
    email TEXT unique not null
);

create table products (
    id SERIAL primary key,
    title TEXT not null,
    price numeric(10, 2) not null
);

create table orders (
    id SERIAL primary key,
    user_id INTEGER not null,
    created_at TIMESTAMP default NOW(),
    
    foreign key (user_id) references users(id) on
delete
	cascade
);

insert
	into
	users (name,
	email)
values
('Alice Johnson',
'alice@example.com'),
('Bob Smith',
'bob@example.com'),
('Charlie Brown',
'charlie@example.com'),
('Diana Prince',
'diana@example.com'),
('Ethan Hunt',
'ethan@example.com');

insert
	into
	products (title,
	price)
values
('Keyboard',
45.99),
('Mouse',
25.50),
('Monitor',
199.99),
('Laptop Stand',
120.00),
('Mechanical Keyboard',
149.90);

insert
	into
	orders (user_id,
	created_at)
values
(1,
NOW() - interval '5 days'),
(2,
NOW() - interval '4 days'),
(3,
NOW() - interval '3 days'),
(1,
NOW() - interval '2 days'),
(5,
NOW() - interval '1 day'),
(2,
NOW());

select
	*
from
	users;

select
	*
from
	products;

select
	*
from
	products
where
	price > 100;

select
	*
from
	orders
order by
	created_at desc
limit 5;

update
	products p
set
	price = 121
where
	p.id = 3;

insert
	into
	products (title,
	price)
values 
('fake',
1);

delete
from
	products p
where
	id = 6;
