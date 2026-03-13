package model

type Session struct {
	Uuid        string    `redis:"uuid"`
	User        UserRedis `redis:"user"`
	CreatedAtNS int64     `redis:"created_at"`
	UpdatedAtNS *int64    `redis:"updated_at,omitempty"`
	DeletedAtNS *int64    `redis:"deleted_at,omitempty"`
}
