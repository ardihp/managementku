package seeds

import (
	"fmt"
	"kanban-board/backend-api/database"
	"kanban-board/backend-api/models"

	"gorm.io/gorm"
)

func WorkspaceTypeSeed() {
	workspace_types := []models.WorkspaceType{
		{Name: "Social Media Specialist", Color: "yellow"},
		{Name: "KOL Management", Color: "blue"},
	}

	database.DB.Transaction(func(trx *gorm.DB) error {
		for _, workspace_type := range workspace_types {
			if err := trx.Save(&workspace_type).Error; err != nil {
				return err
			}
		}

		return nil
	})

	fmt.Println("Workspace types seeded successfully")
}
