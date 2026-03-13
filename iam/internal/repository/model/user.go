package model

type UserRedis struct {
	User_uuid            string                `redis:"user_uuid"`
	Login                string                `redis:"login"`
	Password             string                `redis:"password"`
	Email                string                `redis:"email"`
	Notification_methods []NotificationMethods `redis:"notification_methods,omitempty"`
}

type NotificationMethods struct {
	ProviderName *string
	Target       *string
}
