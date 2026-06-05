package main

import (
	"fmt"
	"log"

	"noise-complaint-backend/src/config"
	"noise-complaint-backend/src/database"
	"noise-complaint-backend/src/models"
	"noise-complaint-backend/src/seed"

	"gorm.io/gorm"
)

func main() {
	cfg := config.Load()

	err := database.Connect(&cfg.Database)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	db := database.GetDB()

	fmt.Println("Running database migrations...")

	err = db.AutoMigrate(
		&models.Complaint{},
		&models.SamplingPoint{},
		&models.NoiseReading{},
		&models.TimePeriod{},
		&models.EvidenceAttachment{},
		&models.RectificationMeasure{},
		&models.RetestRecord{},
		&models.StatusTransition{},
		&models.RuleConfig{},
		&models.AnomalyEvent{},
		&models.AuditLog{},
		&models.ArchiveSnapshot{},
	)
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	fmt.Println("Database migrations completed!")

	fmt.Println("Seeding initial data...")

	var count int64
	db.Model(&models.Complaint{}).Count(&count)
	if count == 0 {
		seed.SeedAll()
		fmt.Println("Seed data inserted successfully!")
	} else {
		fmt.Println("Database already contains data, skipping seed.")
	}

	fmt.Println("Initialization completed!")
}
