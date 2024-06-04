CREATE TABLE users (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    username VARCHAR(50) NOT NULL,
    usersurname VARCHAR(50) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(100) NOT NULL,
    phonenumber VARCHAR(15) NOT NULL
);

CREATE index indx_main on users (phonenumber,username,usersurname);

CREATE index indx_not_main on users using hash(phonenumber);

CREATE index indx_not on users (phonenumber);



drop index indx_main;