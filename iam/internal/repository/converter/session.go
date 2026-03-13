package converter

import (
	"time"

	"github.com/PhilSuslov/homework/iam/internal/model"
	repoModel "github.com/PhilSuslov/homework/iam/internal/repository/model"
	"github.com/samber/lo"
)

func SessionFromRedis(session repoModel.Session) model.Session {
	var updatedAt *time.Time
	if session.UpdatedAtNS != nil {
		tmp := time.Unix(0, *session.UpdatedAtNS)
		updatedAt = &tmp
	}

	var deletedAt *time.Time
	if session.DeletedAtNS != nil {
		tmp := time.Unix(0, *session.DeletedAtNS)
		deletedAt = &tmp
	}

	return model.Session{
		Uuid:      session.Uuid,
		User:      UserFromRedis(session.User),
		CreatedAt: time.Unix(0, session.CreatedAtNS),
		UpdatedAt: updatedAt,
		DeletedAt: deletedAt,
	}
}

func SessionFromRepo(session model.Session) repoModel.Session {
	var updatedAt *int64
	if session.UpdatedAt != nil {
		updatedAt = lo.ToPtr(session.UpdatedAt.UnixNano())
	}

	var deletedAt *int64
	if session.DeletedAt != nil {
		deletedAt = lo.ToPtr(session.DeletedAt.UnixNano())
	}

	createdAt := lo.ToPtr(session.CreatedAt.UnixNano())

	return repoModel.Session{
		Uuid:        session.Uuid,
		User:        UserFromRepo(session.User),
		CreatedAtNS: *createdAt,
		UpdatedAtNS: updatedAt,
		DeletedAtNS: deletedAt,
	}
}
