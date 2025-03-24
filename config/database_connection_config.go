package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/joho/godotenv"
	"godp.abdanhafidz.com/models"
)

var DB *gorm.DB
var err error
var Salt string

func init() {
	godotenv.Load()
	if err != nil {
		fmt.Println("Gagal membaca file .env")
		return
	}
	os.Setenv("TZ", "Asia/Jakarta")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	Salt := os.Getenv("SALT")
	dsn := "host=" + dbHost + " user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName + " port=" + dbPort + " sslmode=disable TimeZone=Asia/Jakarta"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{TranslateError: true})
	if err != nil {
		panic(err)
	}
	if Salt == "" {
		Salt = "D3f4u|t"
	}

	// Call AutoMigrateAll to perform auto-migration
	AutoMigrateAll(DB)
}

func AutoMigrateAll(db *gorm.DB) {
	// Enable logger to see SQL logs
	db.Logger.LogMode(logger.Info)

	// Auto-migrate all models
	err := db.AutoMigrate(
		&models.Account{},
		&models.AccountDetails{},
		&models.EmailVerification{},
		&models.ExternalAuth{},
		&models.FCM{},
		&models.ForgotPassword{},
	)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Migration completed successfully.")
}
