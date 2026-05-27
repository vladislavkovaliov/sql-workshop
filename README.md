# SQL and Databases for Frontend Developers
## 4-Week Interactive Learning Plan (Go Edition)

**Format:**
* 3 sessions per week
* 1–1.5 hours per session
* Practice: Local PostgreSQL running in Docker

**Goal:** Understand backend architecture and SQL at a professional, working level.

---

## What You Will Achieve
* Read and write SQL queries of any complexity.
* Understand database structures and design principles.
* Master JOINs, aggregations, and CTEs.
* Understand indexes, optimization, and how to read `EXPLAIN ANALYZE`.
* Build a Backend API in **Go** connected to a SQL database.

---

## Tech Stack
* **Core:** PostgreSQL, Docker, pgAdmin / DBeaver.
* **Additional:** Go (Golang), `database/sql` + `pgx` (PostgreSQL driver).

---

## Week 1 — SQL & PostgreSQL Basics
*Goal: Set up the environment, learn CRUD operations, and understand basic relationships.*

### Day 1 — Setting Up PostgreSQL
* **Theory:** What is a database, table, row, and column? Primary Key (PK) and Foreign Key (FK).
* **Practice:**
    * [x] Install Docker and Docker Compose (if not already installed).
    * [x] Spin up a PostgreSQL container via the terminal.
    * [x] Spin up a pgAdmin container for a visual GUI management tool.
    * [x] Create a new database named `shop`.
    * [x] Create a `users` table (fields: id, name, email).
    * [x] Create a `products` table (fields: id, title, price).
    * [x] Create an `orders` table (fields: id, user_id, created_at) and link it to the `users` table.

### Day 2 — CRUD Operations
* **Theory:** `SELECT`, `INSERT`, `UPDATE`, `DELETE` statements, filtering with `WHERE`, sorting, and limits.
* **Practice:**
    * [x] Insert 5 test users into the `users` table using `INSERT`.
    * [x] Insert 5 items into the `products` table with different prices.
    * [x] Write a `SELECT` query to fetch absolutely all users.
    * [x] Write a query to select products that cost strictly more than 100.
    * [x] Fetch the last 5 orders from the `orders` table, sorted by date.
    * [x] Update the price of a specific product by its `id` using `UPDATE`.
    * [x] Delete one test user by their `id` using `DELETE`.

### Day 3 — JOINs and Relations
* **Theory:** Relationship types (1:N), `INNER JOIN` vs. `LEFT JOIN`.
* **Practice:**
    * [x] Create several test orders in the `orders` table, binding them to existing `user_id`s.
    * [x] Write an `INNER JOIN` query to get a list of user names alongside their order IDs.
    * [x] Write a `LEFT JOIN` query to display all users (even those without any orders).
    * [x] Formulate a query to find only the users who have never placed an order.
    * [x] Count how many orders each individual user has made.

---

## Week 2 — DB Design & Optimization
*Goal: Learn how to design complex structures and speed up queries.*

### Day 1 — Normalization and M:N
* **Theory:** Normal forms, many-to-many relationships, Pivot tables (junction tables).
* **Practice:**
    * [x] Create a new `categories` table (id, title) to group products.
    * [x] Create a pivot table `product_categories` with foreign keys pointing to products and categories.
    * [x] Add test categories (e.g., Electronics, Apparel, Books).
    * [x] Link a single product to two categories simultaneously.
    * [x] Write a query using two `JOIN`s to display the product title along with all its categories.

### Day 2 — Indexes
* **Theory:** How indexes work (B-tree), why queries slow down (Full/Sequential Scan).
* **Practice:**
    * [x] Generate and insert over 10,000 mock rows into the `products` table (you can use loops or the `generate_series` function).
    * [x] Write a standard `SELECT` query to find a product by price and measure execution time (without an index).
    * [x] Create a B-tree index for the `price` column in the `products` table.
    * [x] Run the same search query again and record the performance difference.

### Day 3 — EXPLAIN ANALYZE
* **Theory:** Query execution plans, `Seq Scan`, `Index Scan`, and query `cost`.
* **Practice:**
    * [x] Run the `EXPLAIN` command for a search query on a text field (without an index) and study the output.
    * [x] Run `EXPLAIN ANALYZE` for the price query (where the index exists) and find the `Index Scan` line in the logs.
    * [x] Identify the `cost` value and the actual execution time in milliseconds.
    * [x] Try writing a query that intentionally forces the database to ignore the index and trigger a `Seq Scan`.

---

## Week 3 — SQL for Backend
*Goal: Aggregations, complex subqueries, and building an API.*

### Day 1 — GROUP BY & Aggregations
* **Theory:** `COUNT`, `SUM`, `AVG`, `GROUP BY`, and filtering groups using `HAVING`.
* **Practice:**
    * [x] Count the total number of orders in the shop using `COUNT`.
    * [x] Find the total revenue of all sold products using `SUM`.
    * [x] Calculate the average product price in a specific category using `AVG`.
    * [x] Group orders by day and display the number of purchases for each date.
    * [x] Find the top 3 buyers who have placed the most orders (use `HAVING` or `ORDER BY`).

### Day 2 — CTEs and Subqueries
* **Theory:** Nested queries, Common Table Expressions (`WITH` syntax).
* **Practice:**
    * [x] Write a subquery inside a `WHERE` clause to find users who bought the most expensive product.
    * [x] Rewrite that exact query using the `WITH` (CTE) syntax to compare readability.
    * [x] Use a CTE to fetch the single most recent order for every registered user.
    * [x] Write a query returning products that have never been purchased (using a `NOT IN` subquery).

### Day 3 — Backend API in Go
* **Theory:** Core interaction principles: Frontend → Go API → SQL. Understanding connection pools (`sql.DB`) and row scanning (`rows.Scan`).
* **Practice:**
    * [x] Initialize a new Go module in a clean directory (`go mod init shop-api`).
    * [x] Install the official pgx driver for PostgreSQL (`go get github.com/jackc/pgx/v5/stdlib`).
    * [x] Create a `Product` struct in Go that matches the fields in your database table.
    * [x] Set up the DB connection using `sql.Open("pgx", "postgres://postgres:postgres@localhost:5432/shop")`.
    * [x] Create an HTTP handler for the `/api/products` route that executes `SELECT id, title, price FROM products`.
    * [x] Implement a `rows.Next()` loop to parse the data and return it to the client as JSON using `json.NewEncoder`.
    * [x] Start the server on port `:8080` and test the endpoint using a browser or Postman.

---

## Week 4 — Production Mindset
*Goal: Architecture, performance, and the final project.*

### Day 1 — Real-World Performance
* **Theory:** The N+1 query problem, pagination strategies (Offset vs. Cursor), heavy JOINs.
* **Practice:**
    * [x] Modify your Go HTTP handler for `/api/products` to accept `limit` and `offset` URL parameters.
    * [x] Add input validation (handling default values if parameters are missing).
    * [x] Integrate `LIMIT` and `OFFSET` into the SQL query inside your Go code using placehoders (`$1`, `$2`) to prevent SQL injections.
    * [x] Write a mock query for Cursor-based pagination based on an element ID in Go.
    * [x] Run `EXPLAIN` on a query with a massive offset (`OFFSET 5000`) to visualize why it degrades performance.

### Day 2 — Database Design Challenge
* **Practice:**
    * [ ] Choose a theme for a standalone project (e.g., Blog, Task Tracker, CRM, or Social Network clone).
    * [ ] Draw a physical data schema (ER Diagram) using a free tool like dbdiagram.io.
    * [ ] Define relationships: explicitly mark where 1:N connections belong and where M:N junction tables are needed.
    * [ ] Write a raw `.sql` file with DDL commands (`CREATE TABLE`, `ALTER TABLE`) to spin up this structure from scratch.

### Day 3 — Final Project
* **Assignment:** Build a mini-backend application in Go.
* **Practice:**
    * [ ] Set up a clean project repository and structure your directories (e.g., separating DB logic into its own package).
    * [ ] Implement your 5–6 table database schema in PostgreSQL.
    * [ ] Write a Go script (or a SQL seed file) to populate the tables with initial mock data.
    * [ ] Build a REST API using Go's native `net/http` package (or a lightweight router like `chi` or `gin`) for basic CRUD operations.
    * [ ] Add at least one analytical/dashboard endpoint using complex `GROUP BY` and `JOIN` logic.
    * [ ] Index the database columns that are heavily targetted by your API's search or filtering routes.

---

## Useful Resources
* **Interactive Practice:** [SQLBolt](https://sqlbolt.com/), [Select Star SQL](https://selectstarsql.com/), [LeetCode SQL](https://leetcode.com/problemset/database/).
* **Tools:** [Explain Visualizer](https://explain.dalibo.com/) (visualizes execution plans).
* **Documentation:** [PostgreSQL Tutorial](https://www.postgresqltutorial.com/), [Go database/sql tutorial](https://go.dev/doc/database/).

---
## What to Learn Next
* Go ORMs (GORM) or code generators (SQLC) — *SQLC is highly recommended to maintain raw SQL control*.
* Database migration tools (golang-migrate, Goose).
* Advanced transactions in Go (`db.BeginTx`).
* Caching layers (Redis) and Message Queues.
