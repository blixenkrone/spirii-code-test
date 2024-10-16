-- name: GetFoo :one
SELECT
    *
FROM
    foo
WHERE
    id = $1
LIMIT
    1;

-- name: GetFoos :many
SELECT
    *
FROM
    foo;

-- name: WriteFoo :one
INSERT INTO foo (id, value) VALUES ($1, $2) RETURNING id, value;
