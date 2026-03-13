package converter

import (
	"github.com/PhilSuslov/homework/iam/internal/model"
	// repoModel "github.com/PhilSuslov/homework/iam/internal/repository/model"
	authV1 "github.com/PhilSuslov/homework/shared/pkg/proto/auth/v1"

)

func StringToLoginResponse(in model.Session) model.LoginResponse {
	return model.LoginResponse{in.Uuid}
}

func UserAuthToWhoamiResponse(in model.Session) *model.WhoamiResponse{
	return &model.WhoamiResponse{
		UserUuid: in.User.User_uuid,
		Login: in.User.Login,
		Email: in.User.Email,
	}
}

func ProtoToLoginRequestModel(in *authV1.LoginRequest) model.LoginRequest {
	return model.LoginRequest{
		Login: in.Login,
		Password: in.Password,
	}
}

func ProtoToWhoamiRequestModel(in *authV1.WhoamiRequest) model.WhoamiRequest {
	return model.WhoamiRequest{
		SessionUuid: in.SessionUuid,
	}
}