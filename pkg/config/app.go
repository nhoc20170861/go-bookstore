package config

import (
	"fmt"
	"log"

	"github.com/nhoc20170861/go-bookstore/pkg/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func ConnectDatabase() {
	utils.LoadEnv() // Load environment variables
	var err error

	// Set up the connection string to the MySQL database.
	// Replace these values with your actual database credentials.
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		utils.GetEnv("DB_USER", "YOUR_USERNAME"),
		utils.GetEnv("DB_PASSWORD", "YOUR_PASSWORD"),
		utils.GetEnv("DB_HOST", "127.0.0.1"),
		utils.GetEnv("DB_PORT", "3306"),
		utils.GetEnv("DB_NAME", "YOUR_DATABASE_NAME"),
	)

	// Connect to the MySQL database using GORM
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // Enable detailed logs
	})

	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
		panic(err)
	}

	log.Println("Database connection successfully established")
}

func GetDB() *gorm.DB {
	return db
}
