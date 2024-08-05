// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: sessions.sql

package db

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createSession = `-- name: CreateSession :one
INSERT INTO sessions (
  id,
  username,
  refersh_token,
  user_agent,
  client_ip,
  is_blocked,
  expired_at
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
)
RETURNING id, username, refersh_token, user_agent, client_ip, is_blocked, expired_at, created_at
`

type CreateSessionParams struct {
	ID           uuid.UUID `json:"id"`
	Username     string    `json:"username"`
	RefershToken string    `json:"refersh_token"`
	UserAgent    string    `json:"user_agent"`
	ClientIp     string    `json:"client_ip"`
	IsBlocked    bool      `json:"is_blocked"`
	ExpiredAt    time.Time `json:"expired_at"`
}

func (q *Queries) CreateSession(ctx context.Context, arg CreateSessionParams) (Sessions, error) {
	row := q.db.QueryRowContext(ctx, createSession,
		arg.ID,
		arg.Username,
		arg.RefershToken,
		arg.UserAgent,
		arg.ClientIp,
		arg.IsBlocked,
		arg.ExpiredAt,
	)
	var i Sessions
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.RefershToken,
		&i.UserAgent,
		&i.ClientIp,
		&i.IsBlocked,
		&i.ExpiredAt,
		&i.CreatedAt,
	)
	return i, err
}

const getSession = `-- name: GetSession :one
SELECT id, username, refersh_token, user_agent, client_ip, is_blocked, expired_at, created_at FROM sessions
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetSession(ctx context.Context, id uuid.UUID) (Sessions, error) {
	row := q.db.QueryRowContext(ctx, getSession, id)
	var i Sessions
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.RefershToken,
		&i.UserAgent,
		&i.ClientIp,
		&i.IsBlocked,
		&i.ExpiredAt,
		&i.CreatedAt,
	)
	return i, err
}
