CREATE TABLE if not exists orders (
    id SERIAL PRIMARY KEY,
    name VARCHAR,
    user_id int references users(id)
);