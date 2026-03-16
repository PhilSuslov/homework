package user

import (
	"context"
	"log"

	orderRepoModel "github.com/PhilSuslov/homework/iam/internal/repository/model"
)

func (i *IAMRepo) Get(ctx context.Context, userUuid string) (orderRepoModel.UserRedis, error) {
	var user orderRepoModel.UserRedis
	rows, err := i.conn.Query(ctx, `SELECT user_uuid, login, password, email,
	notification_methods FROM iam WHERE user_uuid = $1`, userUuid)

	if err != nil {
		log.Printf("failed to select orders by uuid: %v. Error: %v\n", userUuid, err)
		return orderRepoModel.UserRedis{}, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&user.User_uuid, &user.Login, &user.Password, &user.Email,
			&user.Notification_methods)
		if err != nil {
			log.Printf("failed to scan order: %v\n", err)
			return user, err
		}
		log.Println(user.User_uuid, user.Login, user.Password, user.Email,
			user.Notification_methods)
	}
	return user, nil
}

func (i *IAMRepo) Login(ctx context.Context, login, password string) (string, error) {
	var user orderRepoModel.UserRedis
	rows, err := i.conn.Query(ctx, `SELECT user_uuid FROM iam WHERE login = $1 AND password = $2`,
		login, password)

	if err != nil {
		log.Printf("failed to authorization. Error: %v\n", err)
		return "", err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&user.User_uuid)
		if err != nil {
			log.Printf("failed to scan order: %v\n", err)
			return "", err
		}
		log.Println(user.User_uuid)
	}
	return user.User_uuid, nil
}
