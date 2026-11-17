package database

import (
	"fmt"
	"kanban-board/backend-api/config"
	"kanban-board/backend-api/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dbUser := config.GetEnv("DB_USER", "root")
	dbPass := config.GetEnv("DB_PASS", "")
	dbHost := config.GetEnv("DB_HOST", "localhost")
	dbPort := config.GetEnv("DB_PORT", "3306")
	dbName := config.GetEnv("DB_NAME", "")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPass, dbHost, dbPort, dbName)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	fmt.Println("Database connected successfully!")
}

func MigrateDB() {
	DB.Transaction(func(tx *gorm.DB) error {
		fmt.Println("Starting database migration")

		err := tx.AutoMigrate(&models.User{}, &models.Workspace{}, &models.WorkspaceMember{},
			&models.WorkspaceCardStatus{}, &models.WorkspaceCard{}, &models.WorkspaceCardMember{},
			&models.HistoryLog{}, &models.UserHistoryLog{}, &models.WorkspaceHistoryLog{},
			&models.WorkspaceCardActivity{})

		if err != nil {
			return fmt.Errorf("Database migration error, %w", err)
		}

		fmt.Println("Database migrated successfully!")
		return nil
	})
}
