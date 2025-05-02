CREATE TABLE IF NOT EXISTS authors (
    id BIGSERIAL PRIMARY KEY,
    name text NOT NULL,
    bio text,
    birth_year int NOT NULL
);
