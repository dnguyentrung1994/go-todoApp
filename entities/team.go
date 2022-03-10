package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Team struct {
	ID          uuid.UUID `json:"-"`
	TeamName    string    `json:"teamName" gorm:"size:50;not null;index:idx_teamname,unique"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt
	Roles       []TeamRole
}

func (m *Team) BeforeCreate(scope *gorm.DB) (err error) {
	m.ID = uuid.New()
	return
}
