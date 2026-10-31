package seeds

import (
	"fmt"
	"kanban-board/backend-api/database"
	"kanban-board/backend-api/helpers"
	"kanban-board/backend-api/models"
	"time"

	"gorm.io/gorm"
)

func UserSeeds() {
	users := []models.User{
		{Name: "Joko Dodoh", Username: "Joko", Email: "joko@gmail.com", Role: "manager", Password: helpers.HashPassword("test12345"), IsDeleted: false, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Puhdidi", Username: "Didih", Email: "puhdidi@gmail.com", Role: "employee", Password: helpers.HashPassword("test12345"), IsDeleted: false, CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}

	database.DB.Transaction(func(trx *gorm.DB) error {
		for _, user := range users {
			if err := trx.Save(&user).Error; err != nil {
				return err
			}
		}

		return nil
	})

	fmt.Println("Users seeded successfully")
}
