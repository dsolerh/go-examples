CREATE TABLE IF NOT EXISTS books (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    isbn VARCHAR ( 255 ) NOT NULL,
    title VARCHAR ( 255 ) NOT NULL,
    image VARCHAR ( 255 ) NOT NULL,
    genre VARCHAR ( 255 ) NOT NULL,
    year_published int NOT NULL
);