package session

import (
	"context"
	"fmt"
	"time"
)

// AddSessionToUserSet добавляет sessionUUID в множество активных сессий пользователя
func (r *Repository) AddSessionToUserSet(ctx context.Context, userUUID, sessionUUID string, ttl time.Duration) error {
	key := fmt.Sprintf("iam:user:sessions:%s", userUUID)

	// Добавляем sessionUUID в Set
	if err := r.cache.SAdd(ctx, key, sessionUUID); err != nil {
		return fmt.Errorf("failed to add session to user set: %w", err)
	}

	// TTL на Set (чтобы автоматически удалялись старые сессии, если нужно)
	if err := r.cache.Expire(ctx, key, ttl); err != nil {
		return fmt.Errorf("failed to set expiration on user session set: %w", err)
	}

	return nil
}