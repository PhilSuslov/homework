package user

import (
	"context"
	"log"

	orderRepoModel "github.com/PhilSuslov/homework/iam/internal/repository/model"

)

func (i *IAMRepo) Get(ctx context.Context, userUuid string) (orderRepoModel.UserRedis, error) {
	// user, ok := s.db.[orderUUID.String()]
	var user orderRepoModel.UserRedis
	rows, err := i.conn.Query(ctx, `SELECT user_uuid, login, email,
	notification_methods FROM iam WHERE user_uuid = $1`, userUuid)
	if err != nil {
		log.Printf("failed to select orders by uuid: %v. Error: %v\n", userUuid, err)
		return orderRepoModel.UserRedis{}, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&user.User_uuid, &user.Login, &user.Email,
			&user.Notification_methods)
		if err != nil {
			log.Printf("failed to scan order: %v\n", err)
			return user, err
		}
		log.Println(user.User_uuid, user.Login, user.Email,
			user.Notification_methods)
	}
	return user, nil
}
