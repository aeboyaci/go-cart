package models

type Product struct {
	BaseModel
	Name              string  `json:"name" validate:"required"`
	Description       string  `json:"description" validate:"required"`
	Price             float64 `json:"price" validate:"required,gt=0"`
	QuantityAvailable int     `json:"quantityAvailable" validate:"required,gt=0"`
}
