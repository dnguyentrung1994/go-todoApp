package repository

import (
	"go-todoApp/libs"

	"gorm.io/gorm"
)

type UserAddressRepository struct {
	libs.Database
	logger libs.Logger
}

// NewUserRepository creates a new user repository
func NewUserAddressRepository(db libs.Database, logger libs.Logger) UserAddressRepository {
	return UserAddressRepository{
		Database: db,
		logger:   logger,
	}
}

// WithTrx enables repository with transaction
func (r UserAddressRepository) WithTrx(trxHandle *gorm.DB) UserAddressRepository {
	if trxHandle == nil {
		r.logger.Error("Transaction Database not found in gin context. ")
		return r
	}
	r.Database.DB = trxHandle
	return r
}
