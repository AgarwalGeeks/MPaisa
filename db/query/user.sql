-- name: CreateUser :one
INSERT INTO "Finance"."users" (
    email,
    user_password,
    username
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetUserById :one
SELECT * FROM "Finance"."users"
WHERE id = $1;

-- name: GetUserByEmail :one
SELECT * FROM "Finance"."users"
WHERE email = $1;

-- name: DeleteUserByEmail :exec
DELETE FROM "Finance"."users"
WHERE email = $1;