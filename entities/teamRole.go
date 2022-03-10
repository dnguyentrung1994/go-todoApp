package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TeamRole struct {
	ID    uuid.UUID `json:"-"`
	Users []*User   `gorm:"many2many:user_team_role;"`
}

func (m *TeamRole) BeforeCreate(scope *gorm.DB) (err error) {
	m.ID = uuid.New()
	return
}
