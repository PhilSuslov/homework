package converter

import (
	"github.com/PhilSuslov/homework/iam/internal/model"
	repoModel "github.com/PhilSuslov/homework/iam/internal/repository/model"
	userV1 "github.com/PhilSuslov/homework/shared/pkg/proto/common/v1"
)

func UserToRepo(in model.User) repoModel.UserRedis {
	return repoModel.UserRedis{
		User_uuid:            in.User_uuid,
		Login:                in.Login,
		Password:             in.Password,
		Email:                in.Email,
		Notification_methods: NotificationMethodsToRepo(in.Notification_methods),
	}
}

func NotificationMethodsToRepo(in []model.NotificationMethods) []repoModel.NotificationMethods {
	out := make([]repoModel.NotificationMethods, len(in))
	for _, f := range in {
		out = append(out, repoModel.NotificationMethods{
			ProviderName: &f.ProviderName,
			Target:       &f.Target,
		})
	}

	return out
}

func StringToRegisterResponse(in string) model.RegisterResponse {
	return model.RegisterResponse{Uuid: in}
}

func UserToGetUserResponse(in repoModel.UserRedis) model.GetUserResponse {
	out := model.User{
		User_uuid:            in.User_uuid,
		Login:                in.Login,
		Password:             in.Password,
		Email:                in.Email,
		Notification_methods: NotificationMethodsToModel(in.Notification_methods),
	}

	return model.GetUserResponse{
		User: out,
	}

}

func NotificationMethodsToModel(in []repoModel.NotificationMethods) []model.NotificationMethods {
	out := make([]model.NotificationMethods, len(in))
	for _, f := range in {
		out = append(out, model.NotificationMethods{
			ProviderName: *f.ProviderName,
			Target:       *f.ProviderName,
		})
	}

	return out
}

func ProtoToGetUserRequest(in *userV1.GetUserRequest) model.GetUserRequest {
	return model.GetUserRequest{
		UserUuid: in.UserUuid,
	}
}

func GetUserResponseToProto(in model.User) *userV1.User {
	user := userV1.User{
		UserUuid:            in.User_uuid,
		Login:               in.Login,
		Password:            in.Password,
		Email:               in.Email,
		NotificationMethods: NotificationMethodsToProto(in.Notification_methods),
	}

	return &user
}

func NotificationMethodsToProto(in []model.NotificationMethods) []*userV1.NotificationMethods {
	out := make([]*userV1.NotificationMethods, len(in))
	for _, f := range in {
		out = append(out, &userV1.NotificationMethods{
			ProviderName: &f.ProviderName,
			Target:       &f.ProviderName,
		})
	}

	return out
}


func ProtoToRegister(in *userV1.RegisterRequest) model.RegisterRequest {
	user := model.User{
		User_uuid: in.User.UserUuid,
		Login: in.User.Login,
		Password: in.User.Password,
		Email: in.User.Email,
		Notification_methods: ProtoToNotificationMethods(in.User.NotificationMethods),
	}
	
	return model.RegisterRequest{
		User: user,
	}
}

func ProtoToNotificationMethods(in []*userV1.NotificationMethods) []model.NotificationMethods {
	out := make([]model.NotificationMethods, len(in))
	for _, f := range in {
		out = append(out, model.NotificationMethods{
			ProviderName: *f.ProviderName,
			Target:       *f.ProviderName,
		})
	}

	return out
}
