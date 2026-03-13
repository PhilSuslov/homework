package model

import "time"

type Session struct {
	Uuid      string
	User      User
	CreatedAt time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}
