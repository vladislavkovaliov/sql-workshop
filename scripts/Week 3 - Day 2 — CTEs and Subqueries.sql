/*
 * Week 3 - Day 2 — CTEs and Subqueries
 *
 * Covers: subqueries in WHERE, CTEs (WITH), window functions with CTEs,
 *         NOT IN subqueries.
 * Prerequisite: order_items table with mock data.
 */

/* ─── 1. Subquery in WHERE — users who bought the most expensive product ── */

SELECT DISTINCT u.id, u.email, u.name FROM users u
JOIN orders o ON o.user_id = u.id
JOIN order_items oi ON oi.order_id = o.id
WHERE oi.product_id = (
    SELECT id FROM products
    ORDER BY price DESC
    LIMIT 1
);

/* ─── Debug: check if that product actually exists in any order ─────────── */

SELECT * from order_items WHERE order_items.product_id = 7008;
SELECT * from order_items;

/* ─── Debug: confirm which product is the most expensive ────────────────── */

SELECT * FROM products
ORDER BY price DESC
LIMIT 1;

/* ─── 2. Same query, rewritten with a CTE ──────────────────────────────── */

WITH one_top_product_price AS (
    SELECT * FROM products
    ORDER BY price DESC
    LIMIT 1
) 
SELECT DISTINCT u.id, u.email, u.name FROM users u
JOIN orders o ON o.user_id = u.id
JOIN order_items oi ON oi.order_id = o.id
JOIN one_top_product_price ON one_top_product_price.id = oi.product_id;

/* ─── 3. CTE with ROW_NUMBER — most recent order per user ──────────────── */

WITH ranked_order AS (
    SELECT 
        o.id AS order_id,
        o.user_id,
        o.created_at,
        ROW_NUMBER() OVER (
            PARTITION BY o.user_id
            ORDER BY o.created_at DESC
        ) AS rn
    FROM orders o
)
SELECT 
    u.id, u.name, u.email
FROM users u
LEFT JOIN ranked_order ro ON ro.user_id = u.id AND ro.rn = 1;

/* ─── 4. NOT IN subquery — products that have never been purchased ──────── */

SELECT * FROM products
WHERE id NOT IN (
    SELECT DISTINCT product_id FROM order_items
);
