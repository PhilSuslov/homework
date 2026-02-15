package part

import (
	"errors"

	model "github.com/PhilSuslov/homework/inventory/internal/model"
	repoModel "github.com/PhilSuslov/homework/inventory/internal/repository/model"
	repo "github.com/PhilSuslov/homework/inventory/internal/repository/part"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"
)

func (s ServiceSuite) TestServiceGet() {

	var (
		uuid = gofakeit.UUID()
	)

	info := model.GetPartRequest{
		Uuid: uuid,
	}

	expected := model.GetPartResponse{
		Part: model.Part{
			Uuid: uuid,
		},
	}

	s.serviceInventory.
		On("GetPart", s.ctx, info).
		Return(expected, nil)

	part, err := s.service.GetPart(s.ctx, info)
	s.NoError(err)
	s.Equal(part.Part.Uuid, expected.Part.Uuid)
}

func (s ServiceSuite) TestServiceGetErr() {

	var (
		uuid = gofakeit.UUID()
	)

	info := model.GetPartRequest{
		Uuid: uuid,
	}

	s.serviceInventory.
		On("GetPart", s.ctx, info).
		Return(model.GetPartResponse{}, errors.New("not found"))

	_, err := s.service.GetPart(s.ctx, info)
	s.Error(err)
	s.ErrorIs(err, repoModel.ErrNotFound)
}

func (s *ServiceSuite) TestGetPart_Exists() {
    r := repo.NewRepository()
    partUUID := uuid.New().String()
    r.Parts[partUUID] = &repoModel.Part{Uuid: partUUID}

    resp, err := r.GetPart(s.ctx, model.GetPartRequest{Uuid: partUUID})

    s.NoError(err)
    s.Equal(partUUID, resp.Part.Uuid)
}

func (s *ServiceSuite) TestGetPart_NotFound() {
    r := repo.NewRepository()
    partUUID := uuid.New().String()

    _, err := r.GetPart(s.ctx, model.GetPartRequest{Uuid: partUUID})
    s.ErrorIs(err, repoModel.ErrNotFound)
}