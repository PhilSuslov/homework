package model

import "errors"

var (
	ErrSessionNotFound = errors.New("session not found")
	ErrUnauthorized = errors.New("User isn't authorization")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrFailRegister = errors.New("failed to register/create user")
	ErrFailGetUser = errors.New("failed to get user")
)