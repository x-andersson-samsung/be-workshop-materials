---
title: "Devbook "
level: basic
tags: []
created_at: 2026-01-26 07-12-12
modified_at: 2026-03-12 18-47-41
slideNumber: "true"
---

 %% Required for proper codeblock width %%
<style>

li,p {
	font-size: 32px;
}

code {
    font-size: 16px;
    line-height: normal;
}

/* left-align all content in Slides */
.reveal .slides {
    text-align: left;
}

</style>

%% Start of slides %%

# Devbook 
## Meeting 3

---

<grid drag="100 10" drop="0 0" align="left">
### Plan
</grid>

<grid drag="100 85" drop="0 10" align="left" justify-content="center">
- Current application architecture
- Database
- Database configuration for our project
</grid>

---

## Current application architecture

--

<grid drag="100 10" drop="0 0" align="left" >
### How the application works today?
</grid>

<grid drag="100 50" drop="0 30" align="center" justify-content="center">
```mermaid  
flowchart TB

	A[Browser Frontend UI]  
	B[Go HTTP Server REST API]  
	C[In-memory Store
	Go map]  
	  
	A <--> B  
	B <--> C
```
</grid>

--

<grid drag="100 10" drop="0 0" align="left" >
### Current store
</grid>

<grid drag="100 85" drop="0 10" align="left" justify-content="center">

```go[]
type Store struct {  
	nextID int  
	data map[int]Item  
}


store := devbook.NewStore()  
  
store.Add(devbook.Item{  
Name: "item1",  
URL: "https://google.pl",  
})
```
</grid>

--

<grid drag="100 10" drop="0 0" align="left">  
  
### Problems of In-Memory Storage  
  
</grid>  
  
<grid drag="70 60" drop="5 25" align="left">  
  
- Data disappears after application restart  
  
- Data cannot be shared between multiple servers  
  
- Limited querying capabilities  
  
- Difficult to scale  
  
- No data persistence  
  
</grid>

---
## Database
--

<grid drag="100 10" drop="0 0" align="left">

### What is a Database?

</grid>

<grid drag="100 85" drop="0 10" align="left">

 A database is a system designed to:

- store data
- retrieve data efficiently
- manage large amounts of information
- ensure data consistency

</grid>
--
<grid drag="100 20" drop="0 0" align="left">

### Example data structure

</grid>
<grid drag="100 85" drop="0 15" align="center">

| id | name | description | url |
|----|------|-------------|-----|
| 1 | Google | Search engine | google.com |
| 2 | Bing | Search engine | bing.com |

</grid>

--
<grid drag="100 20" drop="0 0" align="left">

### Relational Databases

</grid>

<grid drag="100 15" drop="0 15" align="left">

A relational database stores data in **tables**.

Each table consists of:

- columns (data attributes)
- rows (records)

</grid>

<grid drag="100 15" drop="0 55" align="left">

Example table

| id | name | url |
|----|------|-----|
| 1 | Google | google.com |
| 2 | Bing | bing.com |

</grid>
--
<grid drag="100 20" drop="0 0" align="left">

### Primary Keys

</grid>

<grid drag="100 35" drop="5 20" align="left">

A primary key is a **unique identifier** for each row in a table.

Properties:

- must be unique
- cannot be NULL
- identifies a single record

</grid>
--
<grid drag="100 20" drop="0 0" align="left">

### Why primary key should be unique?
</grid>

<grid drag="100 70" drop="0 20" align="center">
![[meme.png]]
</grid>


--
<grid drag="100 20" drop="0 0" align="left">
### Foreign Keys and Relationships
</grid>

<grid drag="90 35" drop="5 20" align="left">

A foreign key is a column that **references the primary key of another table**.

It creates a relationship between tables.

</grid>
--
<grid drag="70 50" drop="15 25" align="center">  
  
```mermaid  
erDiagram  
  
	USERS {  
	int id PK  
	string name  
	}  
	  
	ORDERS {  
	int id PK  
	int user_id FK  
	string product  
	}  
	  
	USERS ||--o{ ORDERS : places
```
</grid>
%% --
<grid drag="70 35" drop="15 50" align="center">

Users

| id | name |
|----|------|
| 1 | Alice |
| 2 | Bob |

Orders

| id | user_id | product |
|----|---------|--------|
| 1 | 1 | Laptop |
| 2 | 1 | Mouse |
| 3 | 2 | Keyboard |

</grid> %%
--
<grid drag="100 20" drop="0 0" align="left">
### The N+1 Query Problem
</grid>

<grid drag="100 35" drop="5 20" align="left">
We want to fetch **users** and their **orders**

**Bad approach:**

- 1 query for users  
- N queries for orders
</grid>

<grid drag="100 0" drop="5 60" align="left">
Example queries

```sql[]
SELECT * FROM users;

SELECT * FROM orders WHERE user_id = 1;
SELECT * FROM orders WHERE user_id = 2;
SELECT * FROM orders WHERE user_id = 3;
```
</grid>

--

<grid drag="100 20" drop="0 0" align="left">
### Using JOIN to Solve the N+1 Problem
</grid>


<grid drag="100 45" drop="0 25" align="left">
Single query with JOIN

```sql[]
SELECT users.name, orders.product
FROM users
JOIN orders
ON users.id = orders.user_id;
```
</grid>

<grid drag="60 15" drop="20 80" align="center">

1 query instead of N+1

</grid>

--
![[joins.avif]]
--
<grid drag="100 20" drop="5 0" align="left">

### SQL Operations (CRUD)

</grid>

<grid drag="70 50" drop="15 25" align="center">

| Operation | SQL | API |
|-----------|-----|-----|
| Create | INSERT | POST |
| Read | SELECT | GET |
| Update | UPDATE | PUT / PATCH |
| Delete | DELETE | DELETE |

</grid>

<grid drag="80 20" drop="10 85" align="center">

These four operations are known as **CRUD**

</grid>
--
<grid drag="100 20" drop="5 0" align="left">

### Basic SELECT Query

</grid>

<grid drag="70 40" drop="5 20" align="left">

```sql[]
SELECT *
FROM items;
```

</grid>

<grid drag="70 25" drop="5 60" align="left">

- SELECT – choose columns to retrieve
- `*` – select all columns
- FROM – specify the table  

</grid>

<grid drag="80 15" drop="10 85" align="center">

Returns all rows from the table

</grid>

--

<grid drag="100 20" drop="5 0" align="left">

### Selecting Specific Columns

</grid>

<grid drag="70 40" drop="5 10" align="left">
```sql[]
SELECT name, url
FROM items;
```
</grid>

<grid drag="70 25" drop="5 50" align="left">
- SELECT can specify individual columns  
- reduces amount of returned data  
- improves readability of queries
</grid>

<grid drag="80 15" drop="5 80" align="left">
Avoid using SELECT * in production systems
</grid>

--

<grid drag="100 5" drop="0 0" align="left">
### Filtering Data with WHERE
</grid>

<grid drag="70 25" drop="5 10" align="left">
```sql[]
SELECT name, url
FROM items
WHERE name = 'Google';
```
</grid>

<grid drag="70 25" drop="5 50" align="left">

- WHERE filters rows returned by the query  
- it applies a condition to the data  
- only rows matching the condition are returned

</grid>

<grid drag="80 15" drop="10 80" align="center">
Example conditions: `=`, `>`, `<`, `LIKE`, `IN`
</grid>

--

<grid drag="100 5" drop="0 0" align="left">
### Common WHERE Conditions
</grid>

<grid drag="70 25" drop="5 10" align="left">
```sql[]
SELECT * FROM items
WHERE name = 'Google';

SELECT * FROM items
WHERE id > 10;

SELECT * FROM items
WHERE name LIKE '%search%';

SELECT * FROM items
WHERE id IN (1, 2, 3);
```
</grid>

<grid drag="70 25" drop="5 50" align="left">
Common operators:

- `=` equal  
- `>` `<` `>=` `<=` comparisons  
- `LIKE` pattern matching  
- `IN` match multiple values  
- `AND` / `OR` combine conditions
</grid>

--

<grid drag="100 5" drop="0 0" align="left">
### Sorting Results with ORDER BY
</grid>

<grid drag="70 25" drop="5 10" align="left">
```sql[]
SELECT name, url
FROM items
ORDER BY name ASC;

SELECT name, url
FROM items
ORDER BY name DESC;
```
</grid>

<grid drag="70 25" drop="5 40" align="left">
Sorting options:

- `ASC` – ascending order (default)  
- `DESC` – descending order  


Example use cases:

- newest records → `ORDER BY created_at DESC`  
- alphabetical list → `ORDER BY name ASC`
</grid>

--
<grid drag="100 5" drop="0 0" align="left">
### Limiting Results with LIMIT and OFFSET
</grid>

<grid drag="100 5" drop="0 17" align="left">
```sql[]
SELECT name, url
FROM items
LIMIT 10;

SELECT name, url
FROM items
ORDER BY name
LIMIT 10 OFFSET 20;
```
</grid>

<grid drag="70 25" drop="5 46" align="left">
Pagination concepts:

- `LIMIT` – restrict number of returned rows  
- `OFFSET` – skip a number of rows before returning results  

Common use case:

- API pagination  
- displaying results page by page
</grid>

--

<grid drag="100 5" drop="0 0" align="left">
### INSERT – Adding Data
</grid>

<grid drag="70 25" drop="0 10" align="left">
```sql[]
INSERT INTO items (name, description, url)
VALUES (
  'Google',
  'Search engine',
  'https://google.com'
);
```
</grid>

<grid drag="70 25" drop="5 35" align="left">
This operation creates a **new row** in the table
</grid>

<grid drag="70 25" drop="5 60" align="left">
Key points:

- `INSERT INTO` specifies the table  
- columns define where values go  
- `VALUES` contains the data to insert  
</grid>

--

<grid drag="100 5" drop="0 0" align="left">
### UPDATE – Modifying Data
</grid>

<grid drag="70 25" drop="0 20" align="left">
```sql[]
UPDATE items
SET name = 'Google Search'
WHERE id = 1;
```
</grid>

<grid drag="70 25" drop="5 45" align="left">
Key points:

- `UPDATE` specifies the table to modify  
- `SET` defines new values for columns  
- `WHERE` selects which rows will be updated  
</grid>

<grid drag="70 25" drop="5 80" align="left">
⚠️ Without `WHERE`, **all rows will be updated**
</grid>

--

<grid drag="100 5" drop="0 0" align="left">
### DELETE – Removing Data
</grid>

<grid drag="70 25" drop="0 20" align="left">
```sql[]
DELETE FROM items
WHERE id = 1;
```
</grid>

<grid drag="70 25" drop="5 50" align="left">
Key points:

- `DELETE FROM` specifies the table  
- `WHERE` selects which rows to remove  
</grid>

--

<grid drag="100 100" drop="25 -3" align="left">
![[meme_delete.png]]
</grid>

<grid drag="70 25" drop="17 80" align="center">
⚠️ Without `WHERE`, **all rows will be deleted**
</grid>

--
<grid drag="100 5" drop="0 0" align="left">
### Database Transactions
</grid>

<grid drag="70 30" drop="0 14" align="left">
```sql[]
BEGIN;

UPDATE accounts
SET balance = balance - 100
WHERE id = 1;

UPDATE accounts
SET balance = balance + 100
WHERE id = 2;

COMMIT;
```
</grid>

<grid drag="100 25" drop="5 45" align="left">
Used when multiple queries must succeed together
</grid>

<grid drag="100 25" drop="5 65" align="left">
Transactions allow grouping multiple operations:

- `BEGIN` – start transaction  
- `COMMIT` – save changes  
- `ROLLBACK` – undo changes
</grid>
--
<grid drag="100 5" drop="5 0" align="left">
### ACID Properties
</grid>

<grid drag="70 25" drop="5 20" align="left">
ACID ensures reliable database transactions.
</grid>

<grid drag="70 35" drop="5 45" align="left">
- **Atomicity** – all operations succeed or none  
- **Consistency** – database stays in a valid state  
- **Isolation** – concurrent transactions do not interfere  
- **Durability** – committed data is permanently saved
</grid>

<grid drag="100 20" drop="5 80" align="left">

Example: bank transfer must complete both updates or none.

</grid>
--

<grid drag="100 5" drop="0 0" align="left">
### Creating Tables with CREATE TABLE
</grid>

<grid drag="70 30" drop="0 10" align="left">
```sql[]
CREATE TABLE items (
  id SERIAL PRIMARY KEY,
  name TEXT NOT NULL,
  description TEXT,
  url TEXT,
  created_at TIMESTAMP
);
```
</grid>

<grid drag="70 25" drop="5 50" align="left">
Key elements:

- `CREATE TABLE` defines a new table  
- columns specify the structure of the data  
- `PRIMARY KEY` uniquely identifies each row  
- data types define what kind of values can be stored
</grid>

--

<grid drag="100 5" drop="0 0" align="left">
### Database Indexes
</grid>

<grid drag="70 30" drop="0 10" align="left">
```sql[]
CREATE INDEX idx_items_name
ON items(name);
```
</grid>

<grid drag="70 25" drop="5 35" align="left">
What indexes do:

- speed up data retrieval  
- help the database find rows faster  
- often used on columns in `WHERE`, `JOIN`, `ORDER BY`
</grid>
<grid drag="70 25" drop="5 70" align="left">
Trade-offs:

- faster reads  
- slightly slower writes (`INSERT`, `UPDATE`, `DELETE`)
</grid>

---
## Our project
--
<grid drag="100 5" drop="0 0" align="left">
### Database Schema for Our Project
</grid>

<grid drag="70 30" drop="5 25" align="left">
```
items
-----
id
name
description
url
created_at
```
</grid>

<grid drag="70 25" drop="5 55" align="left">
Future extensions:

```
categories
----------
id
name

items.category_id → categories.id
```
</grid>

<grid drag="70 15" drop="5 90" align="left">
This schema will replace our in-memory storage.
</grid>

--

<grid drag="100 5" drop="0 0" align="left">
### Connecting to PostgreSQL
</grid>

<grid drag="70 25" drop="5 20" align="left">
Connection parameters:

- host
- port
- database
- user
- password
</grid>

<grid drag="70 25" drop="5 65" align="left">
Example connection string:

<pre>
postgres://user:password@localhost:5432/workshop_db
</pre>
</grid>

--
<grid drag="100 5" drop="0 0" align="left">
### Connecting to PostgreSQL
</grid>

<grid drag="90 30" drop="5 20" align="left">
Connect to PostgreSQL using the system postgres user:
<pre>
sudo -u postgres psql
</pre>
</grid>

<grid drag="90 25" drop="5 55" align="left">
After connecting you will see the PostgreSQL prompt:

<pre>
postgres=#
</pre>
</grid>

<grid drag="70 20" drop="5 80" align="left">
From this console we can execute SQL commands to manage databases and users.
</grid>

--

<grid drag="100 5" drop="0 0" align="left">
### Creating a Database
</grid>

<grid drag="70 25" drop="5 10" align="left">
Create a new database:
<pre>
CREATE DATABASE workshop_db;
</pre>
</grid>

<grid drag="70 20" drop="5 35" align="left">
Verify that the database was created:
<pre>
\l
</pre>
</grid>

<grid drag="70 25" drop="5 55" align="left">
Example output:

<pre>
      List of databases
   Name        | Owner
----------------+----------
 postgres       | postgres
 workshop_db    | postgres
 template0      | postgres
 template1      | postgres
</pre>
</grid>

--

<grid drag="100 5" drop="0 0" align="left">
### Connecting to PostgreSQL
</grid>

<grid drag="70 25" drop="5 20" align="left">
Create a user with a password:
```sql[]
CREATE USER workshop_user
WITH PASSWORD 'Secret1!';
```
</grid>

<grid drag="70 25" drop="5 50" align="left">
Grant access to the database:
```sql[]
GRANT ALL PRIVILEGES
ON DATABASE workshop_db
TO workshop_user;
```
</grid>

<grid drag="90 20" drop="5 80" align="left">
This user will be used by our Go application to connect to PostgreSQL.
</grid>

--

%% <grid drag="100 5" drop="0 0" align="left">
### Create .ENV file
</grid>

<grid drag="70 40" drop="5 15" align="left">
Our project structure:

<pre>
meeting3/code/
├ cmd/
│  ├ static/
│  └ main.go
├ pkg/
│  ├ api/
│  ├ devbook/
│  └ tui/
└ .env
</pre>
</grid>

<grid drag="90 25" drop="5 70" align="left">
In this step we create the `.env` file  in the root directory of the project.

This file will store the database configuration.
</grid>

-- %%

%% <grid drag="100 5" drop="0 0" align="left">
### Database Configuration (.env)
</grid>

<grid drag="70 35" drop="5 30" align="left">
Example `.env` file:

<pre>
DB_HOST=localhost
DB_PORT=5432
DB_USER=workshop_user
DB_PASSWORD=Secret1!
DB_NAME=workshop_db
DB_SSLMODE=disable
</pre>
</grid>

<grid drag="90 25" drop="5 75" align="left">
These variables will be used by our Go application
to build the database connection.
</grid> %%
<grid drag="100 5" drop="0 0" align="left">
## Environment Configuration (env.sh)
</grid>

<grid drag="90 35" drop="5 20" align="left">
Create a file `env.sh` in the project root:

```bash[]
#!/bin/bash

export DB_HOST=localhost
export DB_PORT=5432
export DB_USER=workshop_user
export DB_PASSWORD=workshop_password
export DB_NAME=workshop_db
export DB_SSLMODE=disable

export DB_URL=
"postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?
sslmode=${DB_SSLMODE}"
```
</grid>

<grid drag="70 20" drop="5 70" align="left">
Load environment variables before running the application:
</grid>

<grid drag="70 20" drop="5 80" align="left">
```bash
source env.sh
```
</grid>
--

<grid drag="100 5" drop="0 0" align="left">
### Building the Database Connection String
</grid>

<grid drag="70 25" drop="5 20" align="left">
PostgreSQL connection string format:

<pre>
postgres://USER:PASSWORD@HOST:PORT/DATABASE?sslmode=disable
</pre>
</grid>

<grid drag="70 25" drop="5 50" align="left">
Example using our `.env` values:

<pre>
postgres://workshop_user:workshop_password
@localhost:5432/workshop_db?sslmode=disable
</pre>
</grid>

<grid drag="90 20" drop="5 80" align="left">
This connection string will be used by our Go application
to connect to PostgreSQL.
</grid>

--

<grid drag="100 5" drop="0 0" align="left">
### Connecting to PostgreSQL in Go
</grid>

<grid drag="85 60" drop="5 9" align="left">
Example `connectDB()` function in `main.go`:

```go[]
func connectDB() *sql.DB {
    connStr := os.Getenv("DB_URL")
    
    db, err := sql.Open("pgx", connStr)
    
    if err != nil {
        panic(err)
    }
    
    if err := db.Ping(); err != nil {
        panic(err)
    }
    
    return db
}
```
</grid>

<grid drag="85 20" drop="5 85" align="left">
The function reads environment variables and
creates a PostgreSQL connection using the pgx driver.
</grid>

--

<grid drag="100 5" drop="0 0" align="left">
### Using the Database Connection in main()
</grid>

<grid drag="85 55" drop="5 20" align="left">
Initialize the connection when the application starts:

```go[]
func main() {

    db := connectDB()
    defer db.Close()

	...
}
```
</grid>

<grid drag="85 25" drop="5 75" align="left">
The database connection is created once during startup
and reused by the application.
</grid>

--

<grid drag="100 5" drop="0 0" align="left">
### Installing Goose
</grid>

<grid drag="70 30" drop="5 20" align="left">
Install the Goose migration tool:

<pre>
go install github.com/pressly/goose/v3/cmd/goose@latest
</pre>
</grid>

<grid drag="70 25" drop="5 55" align="left">
Verify installation:

<pre>
goose -version
</pre>
</grid>

<grid drag="70 20" drop="5 80" align="left">
Goose is used to manage database schema migrations.
</grid>

--

<grid drag="100 5" drop="0 0" align="left">
### Creating a Migration
</grid>

<grid drag="70 30" drop="5 20" align="left">
Create a new migration file inside "migrations" directory:

<pre>
goose create create_items_table sql
</pre>
</grid>

<grid drag="70 25" drop="5 50" align="left">
This command creates a file:
</grid>

<grid drag="70 25" drop="5 70" align="left">
<pre>
migrations/
  00001_create_items_table.sql
</pre>
</grid>

--

<grid drag="100 5" drop="0 0" align="left">
### Example Migration
</grid>

<grid drag="70 50" drop="5 20" align="left">
```sql[]
-- +goose Up
-- +goose StatementBegin
CREATE TABLE items (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT,
    url TEXT,
    created_at TIMESTAMP DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE items;
-- +goose StatementEnd
```
</grid>

--

<grid drag="100 5" drop="0 0" align="left">
### Running Migrations
</grid>

<grid drag="70 25" drop="5 20" align="left">
Run migrations from the project root:

<pre>
goose -dir migrations postgres "$DB_URL" up
</pre>
</grid>

<grid drag="70 25" drop="5 55" align="left">
Check migration status:

<pre>
goose -dir migrations postgres "$DB_URL" status
</pre>
</grid>

<grid drag="70 20" drop="5 80" align="left">
This creates the tables defined in migration files.
</grid>

--

<grid drag="100 5" drop="0 0" align="left">
### Verifying Database Schema
</grid>

<grid drag="70 25" drop="5 10" align="left">
Connect to the database:

<pre>
psql -h localhost -U workshop_user -d workshop_db
</pre>
</grid>

<grid drag="70 25" drop="5 40" align="left">
List tables in the database:

<pre>
\dt
</pre>
</grid>

<grid drag="70 25" drop="5 65" align="left">
Example output:

<pre>
          List of relations
 Schema |  Name  | Type  | Owner
--------+--------+-------+--------
 public | items  | table | postgres
</pre>
</grid>

--

<grid drag="100 5" drop="0 0" align="left">
### Inspecting Table Structure
</grid>

<grid drag="70 25" drop="5 10" align="left">
Describe the table structure in psql:

<pre>
\d items
</pre>
</grid>

<grid drag="70 30" drop="5 35" align="left">
Example output:

<pre>
Column      | Type      | Nullable
------------+-----------+---------
id          | integer   | not null
name        | text      | not null
description | text      |
url         | text      |
created_at  | timestamp |
</pre>
</grid>

<grid drag="70 20" drop="5 80" align="left">
This command shows columns, types and constraints
defined in the migration.
</grid>

--

<grid drag="100 5" drop="0 0" align="left">
### Replacing In-Memory Storage
</grid>

<grid drag="70 30" drop="5 10" align="left">
Previously our application stored data in memory:

```go[]
type Store struct {
	nextID int
    data map[int]Item
}
```
</grid>

<grid drag="70 30" drop="5 40" align="left">
Now the store will use a database connection:

```go[]
type Store struct {
    db *sql.DB
}
```
</grid>

<grid drag="70 20" drop="5 75" align="left">
All data operations will now be executed using SQL queries.
</grid>

--

<grid drag="100 5" drop="0 0" align="left">
### Store Initialization
</grid>

<grid drag="85 45" drop="5 20" align="left">
The store now receives a database connection:

```go[]
type Store struct {
    db *sql.DB
}

func NewStore(db *sql.DB) *Store {
    return &Store{
        db: db,
    }
}
```
</grid>

<grid drag="85 20" drop="5 75" align="left">
The database connection is injected into the store.
</grid>

--

<grid drag="100 5" drop="0 0" align="left">
### Updating main.go
</grid>

<grid drag="85 60" drop="5 10" align="left">
Initialize the database connection and pass it to the Store:

```go[]
func main() {

    db := connectDB()
    defer db.Close()

    store := devbook.NewStore(db)
    
    if err := setupAndRunServer(store, indexData); err != nil {
	    panic(err)
    }
}
```
</grid>

<grid drag="85 20" drop="5 75" align="left">
The database connection is created in main()
and injected into the Store.
</grid>

--

<grid drag="100 5" drop="0 0" align="left">
### Refactoring Store Methods
</grid>

<grid drag="85 25" drop="5 20" align="left">
After (database storage):

```go[]
func (store *Store) Add(item Item) error {
    query := `
        INSERT INTO items (name, description, url)
        VALUES ($1, $2, $3)
    `
    _, err := store.db.Exec(query,
        item.Name,
        item.Description,
        item.URL,
    )
    return err
}
```
</grid>

--

<grid drag="100 5" drop="0 0" align="left">
### Refactoring List() Method
</grid>

<grid drag="85 55" drop="5 10" align="left">
New implementation using SQL:

```go[]
func (store *Store) List() ([]Item, error) {
    rows, _ := store.db.Query(`
        SELECT id, name, description, url
        FROM items
        ORDER BY id
    `)
    defer rows.Close()
    var items []Item
    for rows.Next() {
        var item Item
        err := rows.Scan(
            &item.ID,
            &item.Name,
            &item.Description,
            &item.URL,
        )
        if err != nil {
            return nil, err
        }
        items = append(items, item)
    }
    return items, nil
}
```
</grid>

<grid drag="85 20" drop="5 85" align="left">
The List() method now retrieves items from the database
instead of reading from an in-memory map.
</grid>

---
## Homework

--

<grid drag="100 5" drop="0 0" align="left">
### Homework Assignment
</grid>

<grid drag="70 45" drop="5 10" align="left">
Complete the database integration in the project.
</grid>

<grid drag="70 50" drop="5 40" align="left">

Tasks:

- Implement **UPDATE** operation for items
- Implement **DELETE** operation
- Add a **categories table**
- Create a **foreign key** between items and categories
- Implement a function to **list items by category**
- Add proper **database error handling**

</grid>