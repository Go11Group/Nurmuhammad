-- author tableni yaratish

create table author(
    id serial not null,
    name varchar not null,
    age int,
    primary key (name)
);

-- book tableni yaratish 

create table book(
id serial not null,
name varchar(50) not null,
author_name varchar not null, 
foreign key (author_name) references author(name)
);

-- authorga qiymat berish 

INSERT INTO author (name, age) VALUES
('J.K. Rowling', 55),
('George R.R. Martin', 72),
('J.R.R. Tolkien', 81),
('Stephen King', 73),
('Agatha Christie', 85),
('Isaac Asimov', 72),
('Ernest Hemingway', 61),
('Mark Twain', 74);

-- bookga qiymat berish 

INSERT INTO book (name, author_name) VALUES
('Harry Potter and the Sorcerers Stone', 'J.K. Rowling'),
('Harry Potter and the Chamber of Secrets', 'J.K. Rowling'),
('A Game of Thrones', 'George R.R. Martin'),
('A Clash of Kings', 'George R.R. Martin'),
('The Hobbit', 'J.R.R. Tolkien'),
('The Lord of the Rings', 'J.R.R. Tolkien'),
('The Shining', 'Stephen King'),
('It', 'Stephen King'),
('Murder on the Orient Express', 'Agatha Christie'),
('The ABC Murders', 'Agatha Christie'),
('Foundation', 'Isaac Asimov'),
('I, Robot', 'Isaac Asimov'),
('The Old Man and the Sea', 'Ernest Hemingway'),
('For Whom the Bell Tolls', 'Ernest Hemingway'),
('The Adventures of Tom Sawyer', 'Mark Twain');
