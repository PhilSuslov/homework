package converter

import (
	"github.com/PhilSuslov/homework/iam/internal/model"
	repoModel "github.com/PhilSuslov/homework/iam/internal/repository/model"
)

func UserFromRedis(user repoModel.UserRedis) model.User {
	return model.User{
		User_uuid:            user.User_uuid,
		Login:                user.Login,
		Email:                user.Email,
		Notification_methods: NotificationMethodsFromRedis(user.Notification_methods),
	}
}

func NotificationMethodsFromRedis(repo []repoModel.NotificationMethods) []model.NotificationMethods {
	ans := make([]model.NotificationMethods, 0, len(repo))
	var out model.NotificationMethods

	for _, f := range repo {

		out = model.NotificationMethods{
			ProviderName: *f.ProviderName,
			Target:       *f.Target,
		}
		ans = append(ans, out)
	}

	return ans
}

func UserFromRepo(user model.User) repoModel.UserRedis {
	return repoModel.UserRedis{
		User_uuid:            user.User_uuid,
		Login:                user.Login,
		Email:                user.Email,
		Notification_methods: NotificationMethodsFromRepo(user.Notification_methods),
	}
}

func NotificationMethodsFromRepo(model []model.NotificationMethods) []repoModel.NotificationMethods {
	ans := make([]repoModel.NotificationMethods, 0, len(model))
	var out repoModel.NotificationMethods

	for _, f := range model {

		out = repoModel.NotificationMethods{
			ProviderName: &f.ProviderName,
			Target:       &f.Target,
		}
		ans = append(ans, out)
	}

	return ans
}
