package seeds

import (
	"go-boilerplate/internal/core/domain/models"
	"go-boilerplate/utils/conv"
	"log"

	"gorm.io/gorm"
)

func SeedAdmin(db *gorm.DB) {
	// Implementation for seeding admin user
	bytes, err := conv.HashPassword("admin123")
	if err != nil {
		log.Printf("%s: %v", err.Error(), err)
	}

	admin := models.User{
		Name:       "Super Admin",
		Email:      "superadmin@admin.com",
		Password:   bytes,
		IsVerified: true,
		Latitude:   "",
		Longitude:  "",
	}

	// FirstOrCreate, parameter pertama non_existing, parameter kedua kondisi pencarian
	if err := db.FirstOrCreate(&admin, models.User{Email: admin.Email}).Error; err != nil {
		log.Printf("Failed to seed admin user: %v", err)
	} else {
		log.Printf("Seeded admin user: %s,%s", admin.Email, admin.Name)
	}
}
