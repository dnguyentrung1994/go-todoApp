package entities

import (
	"database/sql/driver"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	customUtils "go-todoApp/utils"
)

type Role string

const (
	GUEST       Role = "GUEST"
	TEAM_MEMBER Role = "TEAM_MEMBER"
	TEAM_LEADER Role = "TEAM_LEADER"
	MODERATOR   Role = "MODERATOR"
	ADMIN       Role = "ADMIN"
)

func (e *Role) Scan(value interface{}) error {
	*e = Role(value.([]byte))
	return nil
}

func (e Role) Value() (driver.Value, error) {
	return string(e), nil
}

type User struct {
	ID        uuid.UUID `json:"-"`
	Username  string    `json:"username" gorm:"size:50;not null;uniqueIndex"`
	Password  string    `json:"-" gorm:"size:200;not null"`
	Email     string    `json:"email" gorm:"size:200;not null"`
	Addresses []Address `json:"addresses"`
	Role      Role      `json:"role" gorm:"type:role;default:'GUEST';not null"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (user *User) BeforeCreate(scope *gorm.DB) (err error) {
	user.ID = uuid.New()
	user.Password = customUtils.GenerateHMAC(user.Password, user.ID.String())
	return
}
