/*
 * Week 2 - Day 3 — EXPLAIN ANALYZE
 *
 * Learn to read PostgreSQL query plans.
 * 1. EXPLAIN (estimate only) on a text search — no index on title
 * 2. EXPLAIN ANALYZE on a price range query — with index
 * 3. Understand cost, actual time, rows, and scan type in the plan
 * 4. Force a Seq Scan when an index exists
 */

/* 1. Estimated plan only — no index on title, expects Seq Scan */
EXPLAIN
SELECT * FROM products WHERE title = 'Premium Keyboard';

/* 2. Execute and measure — create index first, then run */
CREATE INDEX IF NOT EXISTS idx_products_price ON products (price);

EXPLAIN ANALYZE
SELECT * FROM products WHERE price BETWEEN 100 AND 200;

/*
 * 3. In the plan output above, identify:
 *    - "cost=…"       — estimated startup..total cost (arbitrary units)
 *    - "actual time=…" — real execution time in milliseconds
 *    - "rows=…"       — actual vs estimated row count
 *    - "Index Scan"   — confirms the index was used
 *    - "width=…"      — average row width in bytes
 */

/* 4. Force Seq Scan — use a pattern that cannot use the index */
EXPLAIN ANALYZE
SELECT *
FROM products
WHERE price BETWEEN 100 AND 200
  AND title LIKE '%Keyboard%';

/*
 * If the planner still picks Index Scan, disable it temporarily:
 *    SET enable_indexscan = OFF;
 * Re-run the query above, then re-enable:
 *    SET enable_indexscan = ON;
 */
