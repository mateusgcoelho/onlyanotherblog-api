package token

import (
	"errors"
	"time"
)

var (
	ErrorExpiredToken error = errors.New("Token expirado.")
)

type Maker interface {
	CreateToken(userId string, duration time.Duration) (string, error)
	VerifyToken(token string) (*Payload, error)
}
