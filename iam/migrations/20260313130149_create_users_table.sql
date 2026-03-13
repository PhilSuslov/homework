-- +goose Up
CREATE TABLE iam(
    id SERIAL PRIMARY KEY,
    user_uuid TEXT NOT NULL,
    login TEXT,
    password TEXT,
    email TEXT,
    notification_methods TEXT []
);
-- +goose Down
DROP TABLE iam;




/*
 type User struct {
 User_uuid string // UUID пользователя
 Login string // Логин пользователя
 Password string // Пароль пользователя
 Email string // Email пользователя
 Notification_methods []NotificationMethods // Список каналов уведомлений пользователя
 }
 
 type NotificationMethods struct {
 ProviderName string // Имя провайдера (например, telegram, email, push)
 Target string // Адрес получателя — email, ID чата и т.д.
 }
 /*