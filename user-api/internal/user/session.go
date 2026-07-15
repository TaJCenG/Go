package user

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type SessionStore struct {
	rdb *redis.Client
}

func NewSessionStore(rdb *redis.Client) *SessionStore {
	return &SessionStore{rdb: rdb}
}

// Save refresh token with expiration
func (s *SessionStore) SaveRefreshToken(ctx context.Context, userID int, token string) error {
	key := "refresh_token:" + token
	return s.rdb.Set(ctx, key, userID, 7*24*time.Hour).Err()
}

// Validate refresh token
func (s *SessionStore) ValidateRefreshToken(ctx context.Context, token string) (int, error) {
	key := "refresh_token:" + token
	userID, err := s.rdb.Get(ctx, key).Int()
	return userID, err
}

// Revoke refresh token
func (s *SessionStore) RevokeRefreshToken(ctx context.Context, token string) error {
	key := "refresh_token:" + token
	return s.rdb.Del(ctx, key).Err()
}
