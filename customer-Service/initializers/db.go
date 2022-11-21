package initializers

import (
	"log"

	"github.com/VJ-Vijay77/customerServiceMiniProject/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


var DB *gorm.DB


func ConnectToDB() {
	dsn := "postgresql://vijay:12345@postgres:5432/services?sslmode=disable"
	// dsn := "host=postgres port=5432 user=vijay password=12345 dbname=services sslmode=disable timezone=UTC"
	var err error

	DB,err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Couldnt connect to database\nerr = %v",err)
	}
}

func Migrate() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Complaint{})
	DB.AutoMigrate(&models.Associate{})
	
	// var asso []models.Associate
	var associates = []models.Associate{
		{
			Firstname:"associate1",
			Lastname:"CORP",
			Password:"54321",
			Phone:8590462737,
		},
		{
			Firstname:"associate2",
			Lastname:"LUMP",
			Password:"54321",
			Phone:8590462738,
		},
		{
			Firstname:"associate2",
			Lastname:"THALES",
			Password:"54321",
			Phone:8590462739,
		},
	}
	DB.FirstOrCreate(&associates)
}