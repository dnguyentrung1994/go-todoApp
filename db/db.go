package db

import (
	"fmt"
	entities "go-todoApp/entities"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

var db *gorm.DB
var err error

func ConnectToDB() {
	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "8080")
	database := getEnv("DB_NAME", "test")
	username := getEnv("DB_USERNAME", "admin")
	password := getEnv("DB_PASSWORD", "password")

	dbInfo := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		username,
		password,
		host,
		port,
		database,
	)
	db, err = gorm.Open(postgres.Open(dbInfo), &gorm.Config{})
	if err != nil {
		log.Println("Failed to connect to database")
		panic(err)
	}

	log.Println("Database connected")

	//create the role enum if not exists before migrating tables
	db.Exec("DO $$ BEGIN CREATE TYPE role AS ENUM ('GUEST', 'TEAM_MEMBER', 'TEAM_LEADER', 'MODERATOR', 'ADMIN'); EXCEPTION WHEN duplicate_object THEN null; END $$")

	db.AutoMigrate(&entities.User{}, &entities.Address{})
}

func GetDB() *gorm.DB {
	return db
}
