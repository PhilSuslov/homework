package model

import "time"

// GetPart ответ
type GetPartRequest struct {
	Uuid string
}

// GetPart ответ
// type GetPartResponse struct {
// 	Part Part
// }

type GetPartResponse struct {
	Part Note
}

type Note struct{
	// ID primitive.ObjectID `bson:"_id, omitempty"`
	OrderUUID string `bson:"_id"`
	Body Part `bson:"body"`
}

// PartsFilter для фильтрации в ListParts
type PartsFilter struct {
	Uuids                 []string
	Names                 []string
	Categories            []Category
	ManufacturerCountries []string
	Tags                  []string
}

// ListParts запрос
type ListPartsRequest struct {
	Filter PartsFilter
}

// ListParts ответ
type ListPartsResponse struct {
	Parts []Note
}

// Part структура
type Part struct {
	Uuid          string
	Name          string
	Description   string
	Price         float64
	StockQuantity int64
	Category      Category
	Dimensions    Dimensions
	Manufacturer  Manufacturer
	Tags          []string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

// Manufacturer структура
type Manufacturer struct {
	Name    string
	Country string
	Website string
}

// Dimensions структура
type Dimensions struct {
	Length float64
	Width  float64
	Height float64
	Weight float64
}

// Category перечисление
type Category int32

const (
	Category_CATEGORY_UNKNOWN_UNSPECIFIED Category = 0
	Category_CATEGORY_ENGINE              Category = 1
	Category_CATEGORY_FUEL                Category = 2
	Category_CATEGORY_PORTHOLE            Category = 3
	Category_CATEGORY_WING                Category = 4
)

// Enum value maps for Category.
var (
	Category_name = map[int32]string{
		0: "CATEGORY_UNKNOWN_UNSPECIFIED",
		1: "CATEGORY_ENGINE",
		2: "CATEGORY_FUEL",
		3: "CATEGORY_PORTHOLE",
		4: "CATEGORY_WING",
	}
	Category_value = map[string]int32{
		"CATEGORY_UNKNOWN_UNSPECIFIED": 0,
		"CATEGORY_ENGINE":              1,
		"CATEGORY_FUEL":                2,
		"CATEGORY_PORTHOLE":            3,
		"CATEGORY_WING":                4,
	}
)
