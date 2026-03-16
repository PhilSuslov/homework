package user

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	orderRepoModel "github.com/PhilSuslov/homework/iam/internal/repository/model"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"
)

func (r *IAMRepo) Create(ctx context.Context, user orderRepoModel.UserRedis) (string, error) {
	if r.conn == nil {
		return "", fmt.Errorf("pgx pool is nil")
	}

	// Если пустые значения, создается сгенерированный User
	if user.User_uuid == "" {
		user.User_uuid = uuid.NewString()
	}
	if user.Login == "" {
		user.Login = gofakeit.LastName()
	}
	if user.Password == "" {
		user.Password = gofakeit.Password(true, true, true, true, false, 8)
	}
	if user.Email == "" {
		user.Email = gofakeit.Email()
	}
	if len(user.Notification_methods) == 0 {
		providerName := "telegram"
		user.Notification_methods = append(user.Notification_methods, orderRepoModel.NotificationMethods{
			ProviderName: &providerName,
			Target:       &user.Email,
		})
	}

	// Проверка соединения
	if err := r.conn.Ping(ctx); err != nil {
		return "", fmt.Errorf("failed to ping DB: %w", err)
	}

	// Конвертируем массив UUID для pgx
	notifBytes, err := json.Marshal(user.Notification_methods)
	if err != nil {
		return "", fmt.Errorf("failed to marshal notification methods: %w", err)
	}

	res, err := r.conn.Exec(ctx, `
        INSERT INTO iam (user_uuid, login, password, email, notification_methods)
        VALUES ($1, $2, $3, $4, $5)`,
		user.User_uuid, user.Login, user.Password, user.Email, notifBytes)

	if err != nil {
		return "", fmt.Errorf("failed to create order: %w", err)
	}

	log.Printf("Created %d rows\n", res.RowsAffected())
	return user.User_uuid, nil
}
