package structs

type Member struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

type WorkspaceStatus struct {
	Name  string `json:"name"`
	Color string `json:"color"`
}

type WorkspaceType struct {
	Id    uint   `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color"`
}

type WorkspaceRoleResponse struct {
	Id        uint   `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type WorkspaceTypeResponse struct {
	Id        uint   `json:"id"`
	Name      string `json:"name"`
	Color     string `json:"color"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type WorkspaceCardResponse struct {
	Id          uint     `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Members     []Member `json:"members"`
}

type WorkspaceResponse struct {
	Id          uint          `json:"id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Type        WorkspaceType `json:"type"`
	Members     []Member      `json:"members"`
	Tags        []string      `json:"tags"`
	IsArchived  bool          `json:"is_archived"`
	CreatedAt   string        `json:"created_at"`
	UpdatedAt   string        `json:"updated_at"`
}

type WorkspaceCreateRequest struct {
	Name            string   `json:"name" binding:"required"`
	Description     string   `json:"description" binding:"required"`
	WorkspaceTypeId int      `json:"type_id" binding:"required"`
	Tags            []string `json:"tags" binding:"required"`
}

type WorkspaceUpdateRequest struct {
	Name            string   `json:"name" binding:"required"`
	Description     string   `json:"description" binding:"required"`
	WorkspaceTypeId int      `json:"type_id" binding:"required"`
	Tags            []string `json:"tags" binding:"required"`
}

type WorkspaceAssignMember struct {
	WorkspaceId int   `json:"workspace_id" binding:"required"`
	RoleId      int   `json:"role_id" binding:"required"`
	UserId      []int `json:"user_id" binding:"required"`
}

type WorkspaceUpdateStatusRequest struct {
	WorkspaceId int               `json:"workspace_id" binding:"required"`
	Status      []WorkspaceStatus `json:"status"`
}

type WorkspaceCreateRoleRequest struct {
	Name string `json:"name" binding:"required"`
}

type WorkspaceCreateTypeRequest struct {
	Name  string `json:"name" binding:"required"`
	Color string `json:"color" binding:"required"`
}

type WorkspaceStatusCardResponse struct {
	Name       string                  `json:"name"`
	Color      string                  `json:"color"`
	IsArchived bool                    `json:"is_archived"`
	Cards      []WorkspaceCardResponse `json:"cards"`
}

type WorkspaceCreateCardRequest struct {
	WorkspaceId int    `json:"workspace_id" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	Status      string `json:"status" binding:"required"`
}
