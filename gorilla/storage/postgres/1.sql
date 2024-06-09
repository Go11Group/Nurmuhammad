-- Create the student table
CREATE TABLE student (
    student_id SERIAL PRIMARY KEY,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    date_of_birth DATE,
    enrollment_date DATE NOT NULL DEFAULT CURRENT_DATE
);

CREATE TABLE lesson (
    lesson_id SERIAL PRIMARY KEY,
    lesson_name VARCHAR(100) NOT NULL,
    description TEXT
);

Create table student_lesson(
    id SERIAL PRIMARY KEY,
    student_id int references student(student_id),
    lesson_id int references lesson(lesson_id)
);

INSERT INTO student (first_name, last_name, email, date_of_birth, enrollment_date) VALUES
('Alice', 'Smith', 'alice.smith1@example.com', '2001-02-15', '2024-01-10'),
('Bob', 'Johnson', 'bob.johnson@example.com', '2002-03-20', '2024-01-11'),
('Charlie', 'Brown', 'charlie.brown@example.com', '2000-07-30', '2024-01-12'),
('David', 'Wilson', 'david.wilson@example.com', '1999-05-25', '2024-01-13'),
('Eve', 'Davis', 'eve.davis@example.com', '2003-11-10', '2024-01-14'),
('Frank', 'Miller', 'frank.miller@example.com', '2001-12-05', '2024-01-15'),
('Grace', 'Lee', 'grace.lee@example.com', '2002-04-25', '2024-01-16'),
('Hannah', 'Taylor', 'hannah.taylor@example.com', '2000-09-14', '2024-01-17'),
('Ivy', 'Anderson', 'ivy.anderson@example.com', '2003-02-08', '2024-01-18'),
('Jack', 'Thomas', 'jack.thomas@example.com', '2001-10-19', '2024-01-19');

INSERT INTO lesson (lesson_name, description) VALUES
('Math 101', 'Introduction to Mathematics'),
('Physics 101', 'Introduction to Physics'),
('Chemistry 101', 'Introduction to Chemistry'),
('Biology 101', 'Introduction to Biology'),
('History 101', 'Introduction to History'),
('Geography 101', 'Introduction to Geography'),
('English 101', 'Introduction to English Literature'),
('Art 101', 'Introduction to Art'),
('Music 101', 'Introduction to Music'),
('Computer Science 101', 'Introduction to Computer Science'),
('Economics 101', 'Introduction to Economics'),
('Philosophy 101', 'Introduction to Philosophy');

INSERT INTO student_lesson (student_id, lesson_id) VALUES
(1, 1), (1, 2), (1, 3),
(2, 4), (2, 5), (2, 6),
(3, 7), (3, 8), (3, 9),
(4, 10), (4, 11), (4, 12),
(5, 1), (5, 2), (5, 3);
