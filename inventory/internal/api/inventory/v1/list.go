package v1

import (
	"context"

	pb "github.com/PhilSuslov/homework/shared/pkg/proto/inventory/v1"
	"github.com/google/uuid"


)

func (s *InventoryService) ListParts(ctx context.Context, req *pb.ListPartsRequest) (*pb.ListPartsResponse, error) {
    var result []*pb.Part

    for _, part := range s.parts {
        if matchFilterList(part, req) {
            result = append(result, part)
        }
    }

    return &pb.ListPartsResponse{
        Parts: result,
    }, nil
}


func matchFilterList(part *pb.Part, f *pb.ListPartsRequest) bool {
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

func containsUUID(arr []string, v uuid.UUID) bool {
	for _, x := range arr {
		u, err := uuid.Parse(x)
		if err != nil{
			continue 
		}
		if u == v {
			return true
		}
	}
	return false
}

func containsString(arr []string, v string) bool {
	for _, x := range arr {
		if x == v {
			return true
		}
	}
	return false
}

func containsCategory(arr []pb.Category, v pb.Category) bool {
	for _, x := range arr {
		if x == v {
			return true
		}
	}
	return false
}

func containsAny(filter []string, tags []string) bool {
	for _, f := range filter {
		for _, t := range tags {
			if f == t {
				return true
			}
		}
	}
	return false
}
