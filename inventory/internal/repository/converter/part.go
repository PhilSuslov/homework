package converter

import (
	"github.com/PhilSuslov/homework/inventory/internal/model"
	repoModel "github.com/PhilSuslov/homework/inventory/internal/repository/model"

)

func GetPartResponseToRepoModel(info model.Part) repoModel.Part{
	return repoModel.Part{
		Uuid:          info.Uuid,
		Name:          info.Name,
		Description:   info.Description,
		Price:         info.Price,
		StockQuantity: info.StockQuantity,
		Category:      repoModel.Category(info.Category),
		Dimensions:    repoModel.Dimensions(info.Dimensions),
		Manufacturer:  repoModel.Manufacturer(info.Manufacturer),
		Tags:          info.Tags,
		CreatedAt:     info.CreatedAt,
		UpdatedAt:     info.UpdatedAt,
	}
}

func GetPartResponseToModel(repo repoModel.Part) model.Part{
	return model.Part{
		Uuid:          repo.Uuid,
		Name:          repo.Name,
		Description:   repo.Description,
		Price:         repo.Price,
		StockQuantity: repo.StockQuantity,
		Category:      model.Category(repo.Category),
		Dimensions:    model.Dimensions(repo.Dimensions),
		Manufacturer:  model.Manufacturer(repo.Manufacturer),
		Tags:          repo.Tags,
		CreatedAt:     repo.CreatedAt,
		UpdatedAt:     repo.UpdatedAt,
	}
}

func PartsFilterToModel(repo repoModel.PartsFilter) model.PartsFilter{
	return model.PartsFilter{
		Uuids: repo.Uuids,
		Names: repo.Names,
		Categories: RepoCategoriesToModel(repo.Categories),
		ManufacturerCountries: repo.ManufacturerCountries,
		Tags: repo.Tags,
	}
}

func PartsFilterToRepoModel(info model.PartsFilter) repoModel.PartsFilter{
	return repoModel.PartsFilter{
		Uuids: info.Uuids,
		Names: info.Names,
		Categories: RepoCategoriesToRepoModel(info.Categories),
		ManufacturerCountries: info.ManufacturerCountries,
		Tags: info.Tags,
		
	}
}

func ListPartsRequestToRepoModel(info model.ListPartsRequest) repoModel.ListPartsRequest{
	return repoModel.ListPartsRequest{
		Filter: PartsFilterToRepoModel(info.Filter),
	}
}


func ListPartsRequestToModel(repo repoModel.ListPartsRequest) model.ListPartsRequest{
	return model.ListPartsRequest{
		Filter: PartsFilterToModel(repo.Filter),
	}
}

func RepoCategoriesToModel(repo []repoModel.Category) []model.Category{
	res := make([]model.Category, 0, len(repo))
	for _, c := range repo{
		res = append(res, model.Category(c))
	}
	return res
}

func RepoCategoriesToRepoModel(info []model.Category) []repoModel.Category{
	res := make([]repoModel.Category, 0, len(info))
	for _, c := range info{
		res = append(res, repoModel.Category(c))
	}
	return res
}

func ListPartsResponseToRepoModel(info []model.Part) []repoModel.Part{
	res := make([]repoModel.Part, 0, len(info))
	for _, c := range info{
		res = append(res, GetPartResponseToRepoModel(c))
	}
	return res
}

func ListPartsResponseToModel(repo []repoModel.Part) []model.Part{
	res := make([]model.Part, 0, len(repo))
	for _, c := range repo{
		res = append(res, GetPartResponseToModel(c))
	}
	return res
}