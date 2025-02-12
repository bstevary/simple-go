// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: users.sql

package model

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
    email,
    hashed_password
) VALUES (
    $1,
    $2
) RETURNING  user_id, email,  created_at
`

type CreateUserParams struct {
	Email          string `json:"email"`
	HashedPassword string `json:"hashed_password"`
}

type CreateUserRow struct {
	UserID    int64     `json:"user_id"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (CreateUserRow, error) {
	row := q.db.QueryRow(ctx, createUser, arg.Email, arg.HashedPassword)
	var i CreateUserRow
	err := row.Scan(&i.UserID, &i.Email, &i.CreatedAt)
	return i, err
}

const getUser = `-- name: GetUser :one
SELECT user_id, email, hashed_password, created_at, updated_at FROM users WHERE user_id = $1
`

func (q *Queries) GetUser(ctx context.Context, userID int64) (User, error) {
	row := q.db.QueryRow(ctx, getUser, userID)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.Email,
		&i.HashedPassword,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listUsers = `-- name: ListUsers :many
SELECT user_id, email,  created_at FROM users 
WHERE  user_id > $1
ORDER BY user_id 
LIMIT $2
`

type ListUsersParams struct {
	UserID int64 `json:"user_id"`
	Limit  int32 `json:"limit"`
}

type ListUsersRow struct {
	UserID    int64     `json:"user_id"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func (q *Queries) ListUsers(ctx context.Context, arg ListUsersParams) ([]ListUsersRow, error) {
	rows, err := q.db.Query(ctx, listUsers, arg.UserID, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListUsersRow{}
	for rows.Next() {
		var i ListUsersRow
		if err := rows.Scan(&i.UserID, &i.Email, &i.CreatedAt); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateUser = `-- name: UpdateUser :one
UPDATE users
SET 
  updated_at = NOW(),
  email = COALESCE($2, email),
  hashed_password = COALESCE($3, hashed_password)
WHERE user_id = $1
RETURNING user_id, email,  created_at
`

type UpdateUserParams struct {
	UserID         int64       `json:"user_id"`
	Email          pgtype.Text `json:"email"`
	HashedPassword pgtype.Text `json:"hashed_password"`
}

type UpdateUserRow struct {
	UserID    int64     `json:"user_id"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (UpdateUserRow, error) {
	row := q.db.QueryRow(ctx, updateUser, arg.UserID, arg.Email, arg.HashedPassword)
	var i UpdateUserRow
	err := row.Scan(&i.UserID, &i.Email, &i.CreatedAt)
	return i, err
}
