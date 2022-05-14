CREATE TABLE products(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    guid VARCHAR(55) UNIQUE NOT NULL,
    name VARCHAR(255) UNIQUE NOT NULL,
    price REAL NOT NULL,
    description TEXT,
    createdAt TEXT NOT NULL

);

-- insert into products (guid,name,price,description,createdAt)
-- values ('guid','appple',1.98,'description','2020-02-03');