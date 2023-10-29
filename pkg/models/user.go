package models

import "go-cart/pkg/common/types"

type User struct {
	BaseModel
	Email     string         `json:"email" validate:"required" gorm:"uniqueIndex"`
	Password  string         `json:"password" validate:"required"`
	FirstName string         `json:"firstName" validate:"required"`
	LastName  string         `json:"lastName" validate:"required"`
	Role      types.UserRole `json:"-"`
}
