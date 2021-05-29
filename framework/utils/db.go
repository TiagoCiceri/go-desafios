package utils

import (
	"log"
	"os"

	"github.com/TiagoCiceri/go-desafios/domain"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/lib/pg"
)

func ConnectDB() *gorm.DB {

	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dsn := os.Getenv("dns")

	db, err := gorm.Open("postgres", dsn)

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
		panic(err)
	}

	//defer db.Close()

	db.AutoMigrate(&domain.User{})

	return db
}
