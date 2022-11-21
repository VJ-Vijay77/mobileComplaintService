package initializers

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	dsn := "postgresql://vijay:12345@postgres:5432/services?sslmode=disable"
	// dsn := "host=postgres port=5432 user=vijay password=12345 dbname=services sslmode=disable timezone=UTC"
	var err error

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Couldnt connect to database\nerr = %v", err)
	}
}
