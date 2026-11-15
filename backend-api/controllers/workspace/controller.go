package workspace

import (
	"kanban-board/backend-api/database"

	"gorm.io/gorm"
)

type WorkspaceController struct {
	DB *gorm.DB
}

func NewWorkspaceController() *WorkspaceController {
	return &WorkspaceController{
		DB: database.DB,
	}
}
