package controllers

import (
	"kanban-board/backend-api/controllers/workspace"

	"gorm.io/gorm"
)

type AppControllers struct {
	Workspace *workspace.WorkspaceController
}

func InitControllers(db *gorm.DB) *AppControllers {
	return &AppControllers{
		Workspace: workspace.NewWorkspaceController(db),
	}
}
