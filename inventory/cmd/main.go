package main

import (
	"context"
	"log"
	"net"

	pb "github.com/PhilSuslov/homework/shared/pkg/proto/inventory/v1"
	"github.com/google/uuid"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type InventoryService struct {
	pb.UnimplementedInventoryServiceServer
	parts map[string]*pb.Part
}

func NewInventoryService() *InventoryService {
	service := &InventoryService{
		parts: make(map[string]*pb.Part),
	}

	id := uuid.New()
	service.parts[id.String()] = &pb.Part{
		Uuid: id.String(),
		Name: "Main Engine",
		Category: pb.Category_CATEGORY_ENGINE,
		Manufacturer: &pb.Manufacturer{Country: "German"},
		Tags: []string{"main", "engine"},
	}
	
	return service
}

func (s *InventoryService) GetPart(ctx context.Context, req *pb.GetPartRequest) (*pb.GetPartResponse, error) {
	part, ok := s.parts[req.Uuid]
	if !ok {
		return nil, status.Error(codes.NotFound, "part not found")
	}

	return &pb.GetPartResponse{
		Part: part,
	}, nil

}

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

// func matchFilterGet(part *pb.Part, f *pb.ListPartsRequest) bool {
// 	if f == nil {
// 		return true
// 	}
// 
// 	if len(f.Filter.Uuids) > 0 && !containsUUID(f.Filter.Uuids, part.Uuid) {
// 		return false
// 	}
// 	if len(f.Filter.Names) > 0 && !containsString(f.Filter.Names, part.Name) {
// 		return false
// 	}
// 	if len(f.Filter.Categories) > 0 && !containsCategory(f.Filter.Categories, part.Category) {
// 		return false
// 	}
// 	if len(f.Filter.ManufacturerCountries) > 0 && !containsString(f.Filter.ManufacturerCountries, part.Manufacturer.Country) {
// 		return false
// 	}
// 	if len(f.Filter.Tags) > 0 && !containsAny(f.Filter.Tags, part.Tags) {
// 		return false
// 	}
// 
// 	return true
// }

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

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()

	reflection.Register(grpcServer)
	

	inventoryService := NewInventoryService()
	pb.RegisterInventoryServiceServer(grpcServer, inventoryService)

	log.Println("📦 Inventory service started on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
