package model

type Login struct {
	Login    string
	Password string
}

type UserAuth struct {
	User_uuid string
	Login     string
	Email     string
}


type LoginRequest struct {
	Login string
	Password string
}

type LoginResponse struct {
	SessionUuid string
}

type WhoamiRequest struct {
	SessionUuid string
}

type WhoamiResponse struct {
	UserUuid string
	Login string
	Email string

}
