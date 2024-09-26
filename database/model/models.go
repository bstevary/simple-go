// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package model

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type User struct {
	UserID         int64              `json:"user_id"`
	Email          string             `json:"email"`
	HashedPassword string             `json:"hashed_password"`
	CreatedAt      time.Time          `json:"created_at"`
	UpdatedAt      pgtype.Timestamptz `json:"updated_at"`
}
