package repository

import (
	"go-todoApp/libs"

	"gorm.io/gorm"
)

type TeamRepository struct {
	libs.Database
	logger libs.Logger
}

// NewUserRepository creates a new user repository
func NewTeamRepository(db libs.Database, logger libs.Logger) UserRepository {
	return UserRepository{
		Database: db,
		logger:   logger,
	}
}

// WithTrx enables repository with transaction
func (r TeamRepository) WithTrx(trxHandle *gorm.DB) TeamRepository {
	if trxHandle == nil {
		r.logger.Error("Transaction Database not found in gin context. ")
		return r
	}
	r.Database.DB = trxHandle
	return r
}
