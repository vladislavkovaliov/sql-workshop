/*
 * Week 3 - Day 1 — GROUP BY & Aggregations
 *
 * Covers: COUNT, SUM, AVG, GROUP BY, HAVING, ORDER BY with aggregates.
 * Prerequisite: order_items table linking products to orders.
 */

/* ─── Setup: create order_items table if not exists ────────── */

CREATE TABLE IF NOT EXISTS order_items (
    order_id    BIGINT  NOT NULL,
    product_id  BIGINT  NOT NULL,
    quantity    INT     NOT NULL DEFAULT 1,
    PRIMARY KEY (order_id, product_id),
    FOREIGN KEY (order_id)   REFERENCES orders(id)   ON DELETE CASCADE,
    FOREIGN KEY (product_id) REFERENCES products(id)  ON DELETE CASCADE
);

/* Seed some mock order_items so we have data to aggregate */
INSERT INTO order_items (order_id, product_id, quantity)
SELECT
    o.id,
    p.id,
    floor(random() * 3 + 1)::int AS quantity
FROM orders o
CROSS JOIN LATERAL (
    SELECT id FROM products ORDER BY random() LIMIT floor(random() * 3 + 1)::int
) p
ON CONFLICT DO NOTHING;

/* ─── 1. Count the total number of orders ──────────────────── */

SELECT COUNT(*) AS total_orders FROM orders;

/* ─── 2. Total revenue of all sold products (using SUM) ────── */

SELECT
    SUM(p.price * oi.quantity) AS total_revenue
FROM order_items oi
JOIN products p ON p.id = oi.product_id;

/*
 * Equivalent — revenue broken down by product:
 * SELECT p.title, SUM(p.price * oi.quantity) AS revenue
 * FROM order_items oi
 * JOIN products p ON p.id = oi.product_id
 * GROUP BY p.id, p.title
 * ORDER BY revenue DESC;
 */

/* ─── 3. Average product price in a specific category ──────── */

SELECT
    c.title AS category,
    ROUND(AVG(p.price)::numeric, 2) AS avg_price
FROM products p
JOIN product_categories pc ON pc.product_id = p.id
JOIN categories c           ON c.id = pc.category_id
GROUP BY c.id, c.title;

/* ─── 4. Orders grouped by day — purchases per date ────────── */

SELECT
    created_at::date AS order_date,
    COUNT(*)         AS purchases
FROM orders
GROUP BY order_date
ORDER BY order_date;

/* ─── 5. Top 3 buyers (most orders placed) ─────────────────── */

SELECT u.id, u.email, u.name, COUNT(*) as "purchases" FROM users u
JOIN orders o ON o.user_id = u.id
GROUP BY u.id, u.email, u.name
ORDER BY purchases DESC
LIMIT 3;
