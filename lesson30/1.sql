CREATE TABLE users ( 
id SERIAL PRIMARY KEY, 
username VARCHAR(50) NOT NULL,
email VARCHAR(100) NOT NULL UNIQUE, 
password VARCHAR(100) NOT NULL );

CREATE TABLE products ( 
id SERIAL PRIMARY KEY, 
name VARCHAR(100) NOT NULL, 
description TEXT, 
price NUMERIC(10, 2) NOT NULL, 
stock_quantity INT NOT NULL );

CREATE TABLE user_products(
id SERIAL PRIMARY KEY,
user_id int references users(id),
product_id int references products(id)
);



-- select up.id,u.username,p.name 
-- from user_products as up
-- join users as u
-- on u.id=up.user_id
-- join products as p 
-- on p.id=up.product_id;


-- drop table products,users,user_products;