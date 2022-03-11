package libs

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

func NewDatabase(env Env, logger Logger) Database {
	username := env.DBUsername
	password := env.DBPassword
	host := env.DBHost
	port := env.DBPort
	dbname := env.DBName

	dns := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s", username, password, host, port, dbname)

	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{
		Logger: logger.GetGormLogger(),
	})

	if err != nil {
		logger.Info("Url: ", dns)
		logger.Panic(err)
	}

	logger.Info("Database connection established")

	return Database{
		DB: db,
	}
}
