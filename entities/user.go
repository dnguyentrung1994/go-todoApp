package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	customUtils "go-todoApp/utils"
)

type User struct {
	ID        uuid.UUID     `json:"-"`
	Username  string        `json:"username" gorm:"size:50;not null;index:idx_username,unique"`
	Password  string        `json:"-" gorm:"size:200;not null"`
	Email     string        `json:"email" gorm:"size:200;not null"`
	Addresses []UserAddress `json:"addresses"`
	CreatedAt time.Time     `json:"createdAt"`
	UpdatedAt time.Time     `json:"updatedAt"`
	DeletedAt gorm.DeletedAt
	Roles     []*TeamRole `json:"roles" gorm:"many2many:user_team_role;"`
}

func (user *User) BeforeCreate(scope *gorm.DB) (err error) {
	user.ID = uuid.New()
	user.Password = customUtils.GenerateHMAC(user.Password, user.ID.String())
	return
}
