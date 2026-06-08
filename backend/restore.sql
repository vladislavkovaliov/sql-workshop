-- =====================================================
-- restore.sql — полный сброс БД (схема + данные)
-- =====================================================

-- 1. Таблицы (порядок важен из-за FK)
DROP TABLE IF EXISTS order_items CASCADE;
DROP TABLE IF EXISTS product_categories CASCADE;
DROP TABLE IF EXISTS orders CASCADE;
DROP TABLE IF EXISTS products CASCADE;
DROP TABLE IF EXISTS categories CASCADE;
DROP TABLE IF EXISTS users CASCADE;

CREATE TABLE users (
    id    SERIAL PRIMARY KEY,
    name  TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL
);

CREATE TABLE products (
    id    SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    price NUMERIC(10,2) NOT NULL
);

CREATE TABLE orders (
    id         SERIAL PRIMARY KEY,
    user_id    INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE categories (
    id    SERIAL PRIMARY KEY,
    title TEXT NOT NULL
);

CREATE TABLE product_categories (
    product_id  BIGINT NOT NULL REFERENCES products(id) ON DELETE CASCADE,
    category_id BIGINT NOT NULL REFERENCES categories(id) ON DELETE CASCADE,
    PRIMARY KEY (product_id, category_id)
);

CREATE TABLE order_items (
    order_id   BIGINT NOT NULL REFERENCES orders(id) ON DELETE CASCADE,
    product_id BIGINT NOT NULL REFERENCES products(id) ON DELETE CASCADE,
    quantity   INT NOT NULL DEFAULT 1,
    PRIMARY KEY (order_id, product_id)
);

-- 2. Пользователи
INSERT INTO users (name, email) VALUES
('Alice Johnson', 'alice@example.com'),
('Bob Smith',     'bob@example.com'),
('Charlie Brown', 'charlie@example.com'),
('Diana Prince',  'diana@example.com'),
('Ethan Hunt',    'ethan@example.com'),
('Ghost User',    'ghost@example.com');

-- 3. Товары
INSERT INTO products (title, price) VALUES
('Keyboard',              45.99),
('Mouse',                 25.50),
('Monitor',              199.99),
('Laptop Stand',         120.00),
('Mechanical Keyboard',  149.90),
('Smartphone Alpha',     100),
('Wireless Headphones',   50),
('Winter Jacket',         25),
('Running Shoes',       1000),
('SQL for Beginners Book', 10);

-- 4. Категории
INSERT INTO categories (id, title) VALUES
(1, 'Electronics'),
(2, 'Apparel'),
(3, 'Books');

-- 5. Связи товаров с категориями
INSERT INTO product_categories (product_id, category_id) VALUES
(6, 1),  -- Smartphone  → Electronics
(7, 1),  -- Headphones  → Electronics
(8, 2),  -- Jacket      → Apparel
(9, 2),  -- Shoes       → Apparel
(10, 3); -- Book        → Books

-- 6. Заказы
INSERT INTO orders (user_id, created_at) VALUES
(1, NOW() - interval '5 days'),
(2, NOW() - interval '4 days'),
(3, NOW() - interval '3 days'),
(1, NOW() - interval '2 days'),
(5, NOW() - interval '1 day'),
(2, NOW());

INSERT INTO orders (user_id, created_at)
SELECT u.id, NOW() - (random() * interval '30 days')
FROM users u
CROSS JOIN generate_series(1, 3);

-- 7. Элементы заказов (случайные товары в заказах)
INSERT INTO order_items (order_id, product_id, quantity)
SELECT o.id, p.id, floor(random() * 3 + 1)::int
FROM orders o
CROSS JOIN LATERAL (
    SELECT id FROM products ORDER BY random() LIMIT floor(random() * 3 + 1)::int
) p
ON CONFLICT DO NOTHING;

-- 8. Индекс для обучения
CREATE INDEX IF NOT EXISTS idx_products_price ON products (price);
