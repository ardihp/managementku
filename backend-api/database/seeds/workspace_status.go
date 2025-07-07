package seeds

import (
	"fmt"
	"kanban-board/backend-api/database"
	"kanban-board/backend-api/models"

	"gorm.io/gorm"
)

func WorkspaceStatusSeed() {
	statuses := []models.WorkspaceStatus{
		{Name: "Backlog", Color: "gray"},
		{Name: "In Progress", Color: "yellow"},
		{Name: "Completed", Color: "green"},
	}

	database.DB.Transaction(func(trx *gorm.DB) error {
		for index, status := range statuses {
			status.Order = index

			if err := trx.Save(&status).Error; err != nil {
				return err
			}
		}

		return nil
	})

	fmt.Println("Workspace statuses seeded successfully")
}
