CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    enrollment_date DATE NOT NULL DEFAULT CURRENT_DATE
);

CREATE TABLE Problems (
    Problem_id SERIAL PRIMARY KEY,
    Problem_name VARCHAR(100) NOT NULL,
    description TEXT
);

Create table Solved_problems(
    id SERIAL PRIMARY KEY,
    User_id int references Users(User_id),
    Problem_id int references Problems(Problem_id)
);

-- Inserting 5 records into the User table
INSERT INTO Users (first_name, last_name, email) VALUES
('John', 'Doe', 'john.doe@example.com'),
('Jane', 'Smith', 'jane.smith@example.com'),
('Emily', 'Johnson', 'emily.johnson@example.com'),
('Michael', 'Brown', 'michael.brown@example.com'),
('Sarah', 'Davis', 'sarah.davis@example.com');

-- Inserting 20 records into the Problems table
INSERT INTO Problems (Problem_name, description) VALUES
('Problem 1', 'Description for Problem 1'),
('Problem 2', 'Description for Problem 2'),
('Problem 3', 'Description for Problem 3'),
('Problem 4', 'Description for Problem 4'),
('Problem 5', 'Description for Problem 5'),
('Problem 6', 'Description for Problem 6'),
('Problem 7', 'Description for Problem 7'),
('Problem 8', 'Description for Problem 8'),
('Problem 9', 'Description for Problem 9'),
('Problem 10', 'Description for Problem 10'),
('Problem 11', 'Description for Problem 11'),
('Problem 12', 'Description for Problem 12'),
('Problem 13', 'Description for Problem 13'),
('Problem 14', 'Description for Problem 14'),
('Problem 15', 'Description for Problem 15'),
('Problem 16', 'Description for Problem 16'),
('Problem 17', 'Description for Problem 17'),
('Problem 18', 'Description for Problem 18'),
('Problem 19', 'Description for Problem 19'),
('Problem 20', 'Description for Problem 20');

-- Inserting 30 records into the Solved_problems table
INSERT INTO Solved_problems (User_id, Problem_id) VALUES
(1, 1), (1, 2), (1, 3), (1, 4), (1, 5),
(2, 6), (2, 7), (2, 8), (2, 9), (2, 10),
(3, 11), (3, 12), (3, 13), (3, 14), (3, 15),
(4, 16), (4, 17), (4, 18), (4, 19), (4, 20),
(5, 1), (5, 2), (5, 3), (5, 4), (5, 5),
(1, 6), (2, 11), (3, 16), (4, 1), (5, 6);

select u.first_name,u.last_name,u.email,p.Problem_name,p.description 
from users as u
join Solved_problems as s
on s.user_id=u.user_id
join Problems as p 
on p.Problem_id=s.Problem_id;