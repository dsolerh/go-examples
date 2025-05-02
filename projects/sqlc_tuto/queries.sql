-- name: GetAuthor :one
SELECT * FROM authors
WHERE id = $1 LIMIT 1;

-- name: ListAuthors :many
SELECT sqlc.embed(authors), sqlc.embed(books)
FROM authors
inner join books on authors.id = books.author_id
ORDER BY authors.name, books.name;

-- name: CreateAuthor :one
INSERT INTO authors (
  name, bio, birth_year
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: UpdateAuthor :one
UPDATE authors
  set name = $2,
  bio = $3,
  rating = $4
WHERE id = $1
RETURNING id;

-- name: DeleteAuthor :exec
DELETE FROM authors
WHERE id = $1;

-- name: GetBioForAuthor :one
SELECT bio FROM authors
WHERE id = $1;

-- name: GetInfoForAuthor :one
SELECT bio, birth_year FROM authors
WHERE id = $1;

-- name: ListAuthorsByIDs :many
SELECT * FROM authors
WHERE id = ANY($1::int[]);
