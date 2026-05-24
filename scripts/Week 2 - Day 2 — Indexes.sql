/*
 * Week 2 - Day 2 — Indexes
 *
 * Demonstrates how B-tree indexes affect query performance.
 * 1. Insert 10,000+ mock products
 * 2. Run a price lookup WITHOUT an index (Seq Scan)
 * 3. Create a B-tree index on price
 * 4. Run the same lookup WITH an index (Index Scan)
 * 5. Drop the index
 */

/* 1. Insert 10,000 products with randomized names and prices */
INSERT INTO products (title, price)
SELECT
    (ARRAY['Wireless','Bluetooth','Portable','Premium','Smart','Compact','Ergonomic','Heavy-Duty','Lightweight','Eco-Friendly'])[floor(random() * 10) + 1]
    || ' ' ||
    (ARRAY['Keyboard','Mouse','Monitor','Laptop Stand','Headphones','Webcam','Speaker','Charger','Cable','Desk Lamp','Mouse Pad','USB Hub','Microphone','Tablet Case','Phone Stand'])[floor(random() * 15) + 1] AS title,
    round((random() * 990 + 10)::numeric, 2) AS price
FROM generate_series(1, 10000);

/* 2. Select by price WITHOUT an index — expects a Sequential Scan */
EXPLAIN ANALYZE
SELECT * FROM products WHERE price = 744.60;

/* 3. Create a B-tree index on the price column */
CREATE INDEX idx_products_price ON products (price);

/* 4. Same query WITH the index — expects an Index Scan */
EXPLAIN ANALYZE
SELECT * FROM products WHERE price = 500;

/* 5. Clean up — drop the index */
DROP INDEX IF EXISTS idx_products_price;

