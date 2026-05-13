package models

import "time"

type Session struct {
	ID         int
	UsuarioID  int
	TokenHash  string
	ExpiresAt  time.Time
	RevokedAt  *time.Time
	CreatedAt  time.Time
}
