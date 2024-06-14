CREATE TABLE users (
    user_id uuid primary key default gen_random_uuid() not null,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL,
    birthday TIMESTAMP NOT NULL,
    password VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at BIGINT default 0
);
CREATE TABLE courses (
    course_id uuid primary key default gen_random_uuid() not null,
    title VARCHAR(100) NOT NULL,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at BIGINT default 0
);
CREATE TABLE lessons (
    lesson_id uuid primary key default gen_random_uuid() not null,
    course_id UUID REFERENCES courses(course_id),
    title VARCHAR(100) NOT NULL,
    content TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at BIGINT default 0
);
CREATE TABLE enrollments (
    enrollment_id uuid primary key default gen_random_uuid() not null,
    user_id UUID REFERENCES users(user_id),
    course_id UUID REFERENCES courses(course_id),
    enrollment_date TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at BIGINT default 0,
    UNIQUE(user_id,course_id)
);



-- SELECT 
--     user_id,
--     name,
--     email,
--     birthday,
--     EXTRACT(YEAR FROM age(birthday)) AS age
-- FROM 
--     users;


