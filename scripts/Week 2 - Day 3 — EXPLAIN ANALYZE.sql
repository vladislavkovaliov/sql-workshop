EXPLAIN
SELECT * FROM products WHERE title = 'Premium Keyboard';

/* 7. Re-create the index and run EXPLAIN ANALYZE — look for "Index Scan" */
CREATE INDEX IF NOT EXISTS idx_products_price ON products (price);

EXPLAIN ANALYZE
SELECT * FROM products WHERE price BETWEEN 100 AND 200;

/*
 * 8. Read the plan output above and note:
 *    - "cost=…"  — estimated startup and total cost
 *    - "actual time=…" — real execution time in milliseconds
 *    - "rows=…"  — actual vs estimated row count
 *    - "Index Scan" — confirms the index was used
 */

/* 9. Force a Seq Scan by disabling index usage for this query */
EXPLAIN ANALYZE
SELECT * FROM products
WHERE price BETWEEN 100 AND 200
AND title LIKE '%Keyboard%';

