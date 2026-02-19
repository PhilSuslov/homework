package converter

import (
	"time"

	"github.com/PhilSuslov/homework/inventory/internal/model"
	inventory_v1 "github.com/PhilSuslov/homework/shared/pkg/proto/inventory/v1"
	"github.com/samber/lo"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func InventoryGetToModel(info *inventory_v1.GetPartRequest) model.GetPartRequest {
	return model.GetPartRequest{Uuid: info.Uuid}
}

func InventoryGetPartResponseToModel(repo *inventory_v1.Part) model.Part {
	var createdAt *time.Time
	if repo.CreatedAt != nil {
		createdAt = lo.ToPtr(repo.CreatedAt.AsTime())
	}

	var updatedAt *time.Time
	if repo.UpdatedAt != nil {
		updatedAt = lo.ToPtr(repo.UpdatedAt.AsTime())
	}

	return model.Part{
		Uuid:          repo.Uuid,
		Name:          repo.Name,
		Description:   repo.Description,
		Price:         repo.Price,
		StockQuantity: repo.StockQuantity,
		Category:      model.Category(repo.Category),
		Dimensions:    InventoryDimensionsToModel(repo.Dimensions),
		Manufacturer:  InventoryManufacturerToModel(repo.Manufacturer),
		Tags:          repo.Tags,
		CreatedAt:     *createdAt,
		UpdatedAt:     *updatedAt,
	}
}

func InventoryGetPartResponseNoteToModel(repo *inventory_v1.Part) model.Note {
	var createdAt *time.Time
	if repo.CreatedAt != nil {
		createdAt = lo.ToPtr(repo.CreatedAt.AsTime())
	}

	var updatedAt *time.Time
	if repo.UpdatedAt != nil {
		updatedAt = lo.ToPtr(repo.UpdatedAt.AsTime())
	}

	return model.Note{
		Body: model.Part{
			Uuid:          repo.Uuid,
			Name:          repo.Name,
			Description:   repo.Description,
			Price:         repo.Price,
			StockQuantity: repo.StockQuantity,
			Category:      model.Category(repo.Category),
			Dimensions:    InventoryDimensionsToModel(repo.Dimensions),
			Manufacturer:  InventoryManufacturerToModel(repo.Manufacturer),
			Tags:          repo.Tags,
			CreatedAt:     *createdAt,
			UpdatedAt:     *updatedAt,
		},
	}
}

func InventoryGetPartResponseToNote(note model.Note) *inventory_v1.Part {
	var createdAt *timestamppb.Timestamp
	if &note.Body.CreatedAt != nil {
		createdAt = timestamppb.New(note.Body.CreatedAt)
	}

	var updatedAt *timestamppb.Timestamp
	if &note.Body.UpdatedAt != nil {
		updatedAt = timestamppb.New(note.Body.UpdatedAt)
	}

	return &inventory_v1.Part{
		Uuid:          note.Body.Uuid,
		Name:          note.Body.Name,
		Description:   note.Body.Description,
		Price:         note.Body.Price,
		StockQuantity: note.Body.StockQuantity,
		Category:      inventory_v1.Category(note.Body.Category),
		Dimensions:    InventoryDimensionsToProto(note.Body.Dimensions),
		Manufacturer:  InventoryManufacturerToProto(note.Body.Manufacturer),
		Tags:          note.Body.Tags,
		CreatedAt:     createdAt,
		UpdatedAt:     updatedAt,
	}
}


func InventoryGetPartResponseToProto(proto model.Part) *inventory_v1.Part {
	var createdAt *timestamppb.Timestamp
	if &proto.CreatedAt != nil {
		createdAt = timestamppb.New(proto.CreatedAt)
	}

	var updatedAt *timestamppb.Timestamp
	if &proto.UpdatedAt != nil {
		updatedAt = timestamppb.New(proto.UpdatedAt)
	}

	return &inventory_v1.Part{
		Uuid:          proto.Uuid,
		Name:          proto.Name,
		Description:   proto.Description,
		Price:         proto.Price,
		StockQuantity: proto.StockQuantity,
		Category:      inventory_v1.Category(proto.Category),
		Dimensions:    InventoryDimensionsToProto(proto.Dimensions),
		Manufacturer:  InventoryManufacturerToProto(proto.Manufacturer),
		Tags:          proto.Tags,
		CreatedAt:     createdAt,
		UpdatedAt:     updatedAt,
	}
}

func InventoryPartsFilterToModel(repo *inventory_v1.PartsFilter) model.PartsFilter {
	return model.PartsFilter{
		Uuids:                 repo.Uuids,
		Names:                 repo.Names,
		Categories:            InventoryCategoriesToModel(&repo.Categories),
		ManufacturerCountries: repo.ManufacturerCountries,
		Tags:                  repo.Tags,
	}
}

func InventoryPartsFilterToProto(proto model.PartsFilter) *inventory_v1.PartsFilter {
	return &inventory_v1.PartsFilter{
		Uuids:                 proto.Uuids,
		Names:                 proto.Names,
		Categories:            *InventoryCategoriesToProto(proto.Categories),
		ManufacturerCountries: proto.ManufacturerCountries,
		Tags:                  proto.Tags,
	}
}

func InventoryCategoriesToModel(info *[]inventory_v1.Category) []model.Category {
	res := make([]model.Category, 0, len(*info))
	for _, c := range *info {
		res = append(res, model.Category(c))
	}
	return res
}

func InventoryCategoriesToProto(proto []model.Category) *[]inventory_v1.Category {
	res := make([]inventory_v1.Category, 0, len(proto))
	for _, c := range proto {
		res = append(res, inventory_v1.Category(c))
	}
	return &res
}

func InventoryListPartsRequestToModel(info *inventory_v1.ListPartsRequest) model.ListPartsRequest {
	return model.ListPartsRequest{
		Filter: InventoryPartsFilterToModel(info.Filter),
	}
}

func InventoryListPartsResponseToProto(info *inventory_v1.ListPartsResponse) model.ListPartsResponse {
	parts := make([]model.Note, 0, len(info.Parts))

	for _, p := range info.Parts {
		parts = append(parts, InventoryGetPartResponseNoteToModel(p))
	}

	return model.ListPartsResponse{
		Parts: parts,
	}
}

func InventoryListPartsResponseToModel(info model.ListPartsResponse) *inventory_v1.ListPartsResponse {
	parts := make([]*inventory_v1.Part, 0, len(info.Parts))

	for _, p := range info.Parts {
		parts = append(parts, InventoryGetPartResponseToNote(p))
	}

	return &inventory_v1.ListPartsResponse{
		Parts: parts,
	}
}

func InventoryDimensionsToModel(info *inventory_v1.Dimensions) model.Dimensions {
	return model.Dimensions{
		Length: info.Length,
		Width:  info.Width,
		Height: info.Height,
		Weight: info.Weight,
	}
}

func InventoryDimensionsToProto(info model.Dimensions) *inventory_v1.Dimensions {
	return &inventory_v1.Dimensions{
		Length: info.Length,
		Width:  info.Width,
		Height: info.Height,
		Weight: info.Weight,
	}
}

func InventoryManufacturerToModel(info *inventory_v1.Manufacturer) model.Manufacturer {
	return model.Manufacturer{
		Name:    info.Name,
		Country: info.Country,
		Website: info.Website,
	}
}

func InventoryManufacturerToProto(info model.Manufacturer) *inventory_v1.Manufacturer {
	return &inventory_v1.Manufacturer{
		Name:    info.Name,
		Country: info.Country,
		Website: info.Website,
	}
}

