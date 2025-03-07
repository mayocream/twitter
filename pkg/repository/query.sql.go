// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: query.sql

package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (username, email, password_hash) VALUES ($1, $2, $3) RETURNING id, username, email, password_hash, name, bio, avatar_url, banner_url, email_verified, created_at, updated_at
`

type CreateUserParams struct {
	Username     string
	Email        string
	PasswordHash string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser, arg.Username, arg.Email, arg.PasswordHash)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.PasswordHash,
		&i.Name,
		&i.Bio,
		&i.AvatarUrl,
		&i.BannerUrl,
		&i.EmailVerified,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteUser, id)
	return err
}

const getUser = `-- name: GetUser :one
SELECT id, username, email, password_hash, name, bio, avatar_url, banner_url, email_verified, created_at, updated_at FROM users WHERE id = $1
`

func (q *Queries) GetUser(ctx context.Context, id int32) (User, error) {
	row := q.db.QueryRow(ctx, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.PasswordHash,
		&i.Name,
		&i.Bio,
		&i.AvatarUrl,
		&i.BannerUrl,
		&i.EmailVerified,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, username, email, password_hash, name, bio, avatar_url, banner_url, email_verified, created_at, updated_at FROM users WHERE email = $1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRow(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.PasswordHash,
		&i.Name,
		&i.Bio,
		&i.AvatarUrl,
		&i.BannerUrl,
		&i.EmailVerified,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserByUsername = `-- name: GetUserByUsername :one
SELECT id, username, email, password_hash, name, bio, avatar_url, banner_url, email_verified, created_at, updated_at FROM users WHERE username = $1
`

func (q *Queries) GetUserByUsername(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRow(ctx, getUserByUsername, username)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.PasswordHash,
		&i.Name,
		&i.Bio,
		&i.AvatarUrl,
		&i.BannerUrl,
		&i.EmailVerified,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUsers = `-- name: GetUsers :many
SELECT id, username, email, password_hash, name, bio, avatar_url, banner_url, email_verified, created_at, updated_at FROM users ORDER BY id
`

func (q *Queries) GetUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.Query(ctx, getUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Username,
			&i.Email,
			&i.PasswordHash,
			&i.Name,
			&i.Bio,
			&i.AvatarUrl,
			&i.BannerUrl,
			&i.EmailVerified,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
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
UPDATE users SET username = $2, email = $3, password_hash = $4, name = $5, bio = $6, avatar_url = $7, banner_url = $8, email_verified = $9, updated_at = CURRENT_TIMESTAMP WHERE id = $1 RETURNING id, username, email, password_hash, name, bio, avatar_url, banner_url, email_verified, created_at, updated_at
`

type UpdateUserParams struct {
	ID            int32
	Username      string
	Email         string
	PasswordHash  string
	Name          pgtype.Text
	Bio           pgtype.Text
	AvatarUrl     pgtype.Text
	BannerUrl     pgtype.Text
	EmailVerified pgtype.Bool
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, updateUser,
		arg.ID,
		arg.Username,
		arg.Email,
		arg.PasswordHash,
		arg.Name,
		arg.Bio,
		arg.AvatarUrl,
		arg.BannerUrl,
		arg.EmailVerified,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.PasswordHash,
		&i.Name,
		&i.Bio,
		&i.AvatarUrl,
		&i.BannerUrl,
		&i.EmailVerified,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
