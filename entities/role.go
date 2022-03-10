package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Role struct {
	ID          uuid.UUID `json:"-"`
	RoleName    string    `json:"roleName" gorm:"size:50;not null;uniqueIndex"`
	Description string    `json:"description"`
	Teams       []TeamRole
}

func (m *Role) BeforeCreate(scope *gorm.DB) (err error) {
	m.ID = uuid.New()
	return
}
