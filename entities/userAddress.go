package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Address struct {
	ID         uuid.UUID `json:"-"`
	UserId     uuid.UUID `json:"-" gorm:"not null"`
	Prefecture string    `json:"prefecture" gorm:"not null"`
	City       string    `json:"city" gorm:"not null"`
	District   string    `json:"district" gorm:"not null"`
	Street     string    `json:"street" gorm:"not null"`
	Additional string    `json:"additional"`
}

func (address *Address) BeforeCreate(scope *gorm.DB) (err error) {
	address.ID = uuid.New()

	return
}
