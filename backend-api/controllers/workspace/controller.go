package workspace

import (
	"gorm.io/gorm"
)

type WorkspaceController struct {
	DB *gorm.DB
}

func NewWorkspaceController(db *gorm.DB) *WorkspaceController {
	return &WorkspaceController{DB: db}
}
