-- name: CreateUser :one
INSERT INTO users (
    email,
    hashed_password
) VALUES (
    $1,
    $2
) RETURNING  user_id, email,  created_at;


-- name: GetUser :one
SELECT * FROM users WHERE user_id = $1;

-- name: ListUsers :many
SELECT user_id, email,  created_at FROM users 
WHERE  user_id > $1
ORDER BY user_id 
LIMIT $2;

-- name: UpdateUser :one
UPDATE users
SET 
  updated_at = NOW(),
  email = COALESCE(sqlc.narg(email), email),
  hashed_password = COALESCE(sqlc.narg(hashed_password), hashed_password)
WHERE user_id = $1
RETURNING user_id, email,  created_at;
