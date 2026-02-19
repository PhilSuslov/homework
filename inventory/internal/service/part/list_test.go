package part
// 
// import (
// 	"testing"
// 
// 	"github.com/PhilSuslov/homework/inventory/internal/model"
// 	repoModel "github.com/PhilSuslov/homework/inventory/internal/repository/model"
// 	repoPart "github.com/PhilSuslov/homework/inventory/internal/repository/part"
// 
// 	"github.com/stretchr/testify/suite"
// )
// 
// type ListPartsSuite struct {
// 	suite.Suite
// 	repo    *repoPart.Repository
// 	service *service
// }
// 
// func (s *ListPartsSuite) SetupTest() {
// 	// Создаём реальный репозиторий
// 	s.repo = repoPart.NewRepository()
// 
// 	// Добавляем тестовые части
// 	s.repo.Parts["p1"] = &repoModel.Part{
// 		Uuid: "p1", Name: "PartA", Category: 1,
// 		Manufacturer: repoModel.Manufacturer{Country: "RU"},
// 		Tags:         []string{"tag1"},
// 	}
// 	s.repo.Parts["p2"] = &repoModel.Part{
// 		Uuid: "p2", Name: "PartB", Category: 2,
// 		Manufacturer: repoModel.Manufacturer{Country: "CN"},
// 		Tags:         []string{"tag2"},
// 	}
// 	s.repo.Parts["p3"] = &repoModel.Part{
// 		Uuid: "p3", Name: "PartC", Category: 1,
// 		Manufacturer: repoModel.Manufacturer{Country: "RU"},
// 		Tags:         []string{"tag1", "tag3"},
// 	}
// 
// 	// Создаём сервис с этим репозиторием
// 	s.service = NewService(s.repo)
// }
// 
// func (s *ListPartsSuite) TestListParts_NoFilter() {
// 	resp, err := s.service.ListParts(nil, model.ListPartsRequest{})
// 	s.NoError(err)
// 	s.Len(resp.Parts, len(s.repo.Parts))
// }
// 
// func (s *ListPartsSuite) TestListParts_FilterByUUID() {
// 	resp, err := s.service.ListParts(nil, model.ListPartsRequest{
// 		Filter: model.PartsFilter{Uuids: []string{"p2"}},
// 	})
// 	s.NoError(err)
// 	s.Len(resp.Parts, 1)
// 	s.Equal("p2", resp.Parts[0].Uuid)
// }
// 
// func (s *ListPartsSuite) TestListParts_FilterByName() {
// 	resp, err := s.service.ListParts(nil, model.ListPartsRequest{
// 		Filter: model.PartsFilter{Names: []string{"PartC"}},
// 	})
// 	s.NoError(err)
// 	s.Len(resp.Parts, 1)
// 	s.Equal("PartC", resp.Parts[0].Name)
// }
// 
// func (s *ListPartsSuite) TestListParts_FilterByCategory() {
// 	resp, err := s.service.ListParts(nil, model.ListPartsRequest{
// 		Filter: model.PartsFilter{Categories: []model.Category{1}},
// 	})
// 	s.NoError(err)
// 	s.Len(resp.Parts, 2)
// 	for _, p := range resp.Parts {
// 		s.Equal(model.Category(1), p.Category)
// 	}
// }
// 
// func (s *ListPartsSuite) TestListParts_FilterByManufacturerCountry() {
// 	resp, err := s.service.ListParts(nil, model.ListPartsRequest{
// 		Filter: model.PartsFilter{ManufacturerCountries: []string{"RU"}},
// 	})
// 	s.NoError(err)
// 	s.Len(resp.Parts, 2)
// 	for _, p := range resp.Parts {
// 		s.Equal("RU", p.Manufacturer.Country)
// 	}
// }
// 
// func (s *ListPartsSuite) TestListParts_FilterByTags() {
// 	resp, err := s.service.ListParts(nil, model.ListPartsRequest{
// 		Filter: model.PartsFilter{Tags: []string{"tag3"}},
// 	})
// 	s.NoError(err)
// 	s.Len(resp.Parts, 1)
// 	s.Contains(resp.Parts[0].Tags, "tag3")
// }
// 
// func (s *ListPartsSuite) TestListParts_MultipleFilters() {
// 	resp, err := s.service.ListParts(nil, model.ListPartsRequest{
// 		Filter: model.PartsFilter{
// 			Categories:            []model.Category{1},
// 			ManufacturerCountries: []string{"RU"},
// 			Tags:                  []string{"tag1"},
// 		},
// 	})
// 	s.NoError(err)
// 	s.Len(resp.Parts, 2)
// 	for _, p := range resp.Parts {
// 		s.Equal(model.Category(1), p.Category)
// 		s.Equal("RU", p.Manufacturer.Country)
// 		s.Contains(p.Tags, "tag1")
// 	}
// }
// 
// func TestListPartsSuite(t *testing.T) {
// 	suite.Run(t, new(ListPartsSuite))
// }
