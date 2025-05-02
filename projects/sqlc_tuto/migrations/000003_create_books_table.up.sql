create table if not exists books (
    id bigserial primary key,
    name text not null,
    summary text,
    author_id bigint references authors (id)
);
