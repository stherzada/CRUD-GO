-- name: GetUserById :one
SELECT * FROM users WHERE id = $1;

-- name: GetUsers :many
SELECT * FROM users;

-- name: UpdateUsersName :exec
UPDATE users SET name = $1 WHERE id = $2;

-- name: CreateUser :exec
INSERT INTO users (id, name) VALUES ($1, $2);

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;
