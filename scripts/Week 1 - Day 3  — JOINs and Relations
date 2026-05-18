
/* Create several test orders in the orders table, binding them to existing user_ids. */
insert
	into
	orders (user_id,
	created_at)
select
	u.id as user_id,
	-- Генерируем случайную дату в диапазоне последних 30 дней
	NOW() - (random() * interval '30 days') as created_at
from
	users u
	-- generate_series(1, 3) означает, что для КАЖДОГО пользователя создастся по 3 заказа
cross join 
    generate_series(1, 3);

insert
	into
	users (name,
	email)
values ('Ghost User',
'ghost@example.com');

/* Write an INNER JOIN query to get a list of user names alongside their order IDs. */
select
	u."name" "Name",
	u.email "Email",
	orders.id as "Order ID"
from
	users u
inner join orders on
	u.id = orders.user_id;


/* Write a LEFT JOIN query to display all users (even those without any orders). */
select
	u."name" "Name",
	u.email "Email",
	orders.id as "Order ID"
from
	users u
left join orders on
	u.id = orders.user_id;

select
	u.id as user_id,
	u.name as "Name",
	u.email as "Email"
from
	users u
left join 
    orders o on
	u.id = o.user_id
where
	o.id is null;

/* Count how many orders each individual user has made. */
select
	u.id as user_id,
	u.name as "Name",
	COUNT(o.id) as "Total Orders"
from
	users u
left join 
    orders o on
	u.id = o.user_id
group by
	u.id,
	u.name
order by
	"Total Orders" desc;
