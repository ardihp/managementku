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

type WorkspaceStatusResponse struct {
	WorkspaceId uint              `json:"workspace_id"`
	Status      []WorkspaceStatus `json:"status"`
}

type WorkspaceRoleReponse struct {
	Id        uint   `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type WorkspaceResponse struct {
	Id          uint     `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Category    string   `json:"category"`
	Members     []Member `json:"members"`
	CreatedAt   string   `json:"created_at"`
	UpdatedAt   string   `json:"updated_at"`
}

type WorkspaceCreateRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Category    string `json:"category" binding:"required"`
}

type WorkspaceUpdateRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Category    string `json:"category" binding:"required"`
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
