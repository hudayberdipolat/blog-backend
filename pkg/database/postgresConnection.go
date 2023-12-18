package database

import (
	"fmt"
	"log"

	"github.com/hudayberdipolat/blog-backend/pkg/config"
	"github.com/ilyakaznacheev/cleanenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseConnection() error {
	var dbConfig config.DatabaseConfig
	if err := cleanenv.ReadConfig("../../.env", &dbConfig); err != nil {
		return err
	}
	postgresConnection := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		dbConfig.DbHost,
		dbConfig.DbUser,
		dbConfig.DbPassword,
		dbConfig.DbName,
		dbConfig.DbPort,
		dbConfig.DbSslMode,
	)
	db, err := gorm.Open(postgres.Open(postgresConnection), &gorm.Config{})
	if err != nil {
		log.Println("Database Connection error -->>", err.Error())
		return err
	}
	DB = db
	return nil
}
