package model

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

type GetUserRequest struct {
	UserUuid string
}

type GetUserResponse struct{
	User User
}

type RegisterRequest struct {
	User User
}

type RegisterResponse struct {
	Uuid string
}