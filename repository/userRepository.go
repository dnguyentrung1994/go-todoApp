package repository

import (
	"go-todoApp/libs"

	"gorm.io/gorm"
)

type UserRepository struct {
	libs.Database
	logger libs.Logger
}

// NewUserRepository creates a new user repository
func NewUserRepository(db libs.Database, logger libs.Logger) UserRepository {
	return UserRepository{
		Database: db,
		logger:   logger,
	}
}

// WithTrx enables repository with transaction
func (r UserRepository) WithTrx(trxHandle *gorm.DB) UserRepository {
	if trxHandle == nil {
		r.logger.Error("Transaction Database not found in gin context. ")
		return r
	}
	r.Database.DB = trxHandle
	return r
}
