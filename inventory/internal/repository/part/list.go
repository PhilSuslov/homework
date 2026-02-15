package part

import (
	"context"

	"github.com/PhilSuslov/homework/inventory/internal/model"
	repoConverter "github.com/PhilSuslov/homework/inventory/internal/repository/converter"
	repoModel "github.com/PhilSuslov/homework/inventory/internal/repository/model"
)

func (r *Repository) ListParts(ctx context.Context, req model.ListPartsRequest) (model.ListPartsResponse, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var result []repoModel.Part
	rp := repoConverter.ListPartsRequestToRepoModel(req)

	for _, part := range r.Parts {
		if matchFilterList(*part, &rp) {
			result = append(result, *part)
		}
	}
	return model.ListPartsResponse{
		Parts: repoConverter.ListPartsResponseToModel(result),
	}, nil
}

func matchFilterList(part repoModel.Part, f *repoModel.ListPartsRequest) bool {
	if f == nil {
		return true // нет фильтра — все подходят
	}

	// Проверка UUID (логическое ИЛИ внутри)
	if len(f.Filter.Uuids) > 0 {
		match := false
		for _, id := range f.Filter.Uuids {
			if id == part.Uuid {
				match = true
				break
			}
		}
		if !match {
			return false
		}
	}

	// Проверка имен
	if len(f.Filter.Names) > 0 {
		match := false
		for _, name := range f.Filter.Names {
			if name == part.Name {
				match = true
				break
			}
		}
		if !match {
			return false
		}
	}

	// Проверка категорий
	if len(f.Filter.Categories) > 0 {
		match := false
		for _, cat := range f.Filter.Categories {
			if cat == part.Category {
				match = true
				break
			}
		}
		if !match {
			return false
		}
	}

	// Проверка стран производителей
	if len(f.Filter.ManufacturerCountries) > 0 {
		match := false
		for _, country := range f.Filter.ManufacturerCountries {
			if country == part.Manufacturer.Country {
				match = true
				break
			}
		}
		if !match {
			return false
		}
	}

	// Проверка тегов (любое совпадение с тегами части)
	if len(f.Filter.Tags) > 0 {
		match := false
		for _, tag := range f.Filter.Tags {
			for _, partTag := range part.Tags {
				if tag == partTag {
					match = true
					break
				}
			}
			if match {
				break
			}
		}
		if !match {
			return false
		}
	}

	return true
}

//
// func containsUUID(arr []string, v uuid.UUID) bool {
// 	for _, x := range arr {
// 		u, err := uuid.Parse(x)
// 		if err != nil {
// 			continue
// 		}
// 		if u == v {
// 			return true
// 		}
// 	}
// 	return false
// }
//
// func containsString(arr []string, v string) bool {
// 	for _, x := range arr {
// 		if x == v {
// 			return true
// 		}
// 	}
// 	return false
// }
//
// func containsCategory(arr []model.Category, v model.Category) bool {
// 	for _, x := range arr {
// 		if x == v {
// 			return true
// 		}
// 	}
// 	return false
// }
//
// func containsAny(filter []string, tags []string) bool {
// 	for _, f := range filter {
// 		for _, t := range tags {
// 			if f == t {
// 				return true
// 			}
// 		}
// 	}
// 	return false
// }
