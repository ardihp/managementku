package structs

type Member struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

type WorkspaceStatus struct {
	Id         uint   `json:"id"`
	Order      int    `json:"order"`
	Name       string `json:"name"`
	Color      string `json:"color"`
	IsArchived bool   `json:"is_archived"`
}

type WorkspaceType struct {
	Id    uint   `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color"`
}

type WorkspaceCardStatus struct {
	Id         uint   `json:"id"`
	Name       string `json:"name"`
	Color      string `json:"color"`
	IsArchived bool   `json:"is_archived"`
	Order      int    `json:"order"`
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

type WorkspaceCardActivityResponse struct {
	Id          int    `json:"id"`
	UserCreated string `json:"user_created"`
	Type        string `json:"type"`
	Message     string `json:"message"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type WorkspaceCardResponse struct {
	Id          uint                `json:"id"`
	Title       string              `json:"title"`
	Description string              `json:"description"`
	Status      WorkspaceCardStatus `json:"status"`
	UserCreated string              `json:"user_created"`
	Members     []Member            `json:"members"`
	IsArchived  bool                `json:"is_archived"`
	IsCompleted bool                `json:"is_completed"`
	CreatedAt   string              `json:"created_at"`
	UpdatedAt   string              `json:"updated_at"`
}

type WorkspaceResponse struct {
	Id          uint            `json:"id"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Type        WorkspaceType   `json:"type"`
	Status      WorkspaceStatus `json:"status,omitempty"`
	Members     []Member        `json:"members"`
	Tags        []string        `json:"tags"`
	IsArchived  bool            `json:"is_archived"`
	CreatedAt   string          `json:"created_at"`
	UpdatedAt   string          `json:"updated_at"`
}

type WorkspaceByStatusResponse struct {
	Id         uint                `json:"id"`
	Name       string              `json:"name"`
	Color      string              `json:"color"`
	Workspaces []WorkspaceResponse `json:"workspaces"`
	IsArchived bool                `json:"is_archived"`
}

type WorkspaceCreateRequest struct {
	Name        string   `json:"name" binding:"required"`
	Description string   `json:"description" binding:"required"`
	TypeId      int      `json:"type_id" binding:"required"`
	StatusId    int      `json:"status_id" binding:"required"`
	Tags        []string `json:"tags,omitempty"`
}

type WorkspaceUpdateRequest struct {
	Name        string   `json:"name" binding:"required"`
	Description string   `json:"description" binding:"required"`
	TypeId      int      `json:"type_id" binding:"required"`
	StatusId    int      `json:"status_id" binding:"required"`
	Tags        []string `json:"tags,omitempty"`
}

type WorkspaceAssignMember struct {
	WorkspaceId int   `json:"workspace_id" binding:"required"`
	UserId      []int `json:"user_id" binding:"required"`
}

type WorkspaceCardAssignMember struct {
	CardId int   `json:"card_id" binding:"required"`
	UserId []int `json:"user_id" binding:"required"`
}

type WorkspaceWatchCardRequest struct {
	UserId int `json:"user_id" binding:"required"`
}

type WorkspaceCreateRoleRequest struct {
	Name string `json:"name" binding:"required"`
}

type WorkspaceCreateTypeRequest struct {
	Name  string `json:"name" binding:"required"`
	Color string `json:"color" binding:"required"`
}

type WorkspaceUpdateTypeRequest struct {
	Id    int    `json:"id" binding:"required"`
	Name  string `json:"name" binding:"required"`
	Color string `json:"color" binding:"required"`
}

type WorkspaceUpdateTypeOrderRequest struct {
	OrderStatus []int `json:"order_status" binding:"required"`
}

type WorkspaceCardByStatusResponse struct {
	Id         int                     `json:"id"`
	Name       string                  `json:"name"`
	Color      string                  `json:"color"`
	IsArchived bool                    `json:"is_archived"`
	Cards      []WorkspaceCardResponse `json:"cards"`
}

type WorkspaceCreateCardRequest struct {
	StatusId    int    `json:"status_id" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
}

type WorkspaceUpdateCardRequest struct {
	StatusId    int    `json:"status_id" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
}

type WorkspaceCreateStatusRequest struct {
	Name  string `json:"name" binding:"required"`
	Color string `json:"color" binding:"required"`
}

type WorkspaceCreateCardStatusRequest struct {
	Name  string `json:"name" binding:"required"`
	Color string `json:"color" binding:"required"`
}

type WorkspaceUpdateStatusRequest struct {
	Id    int    `json:"id" binding:"required"`
	Name  string `json:"name" binding:"required"`
	Color string `json:"color" binding:"required"`
}

type WorkspaceUpdateCardStatusRequest struct {
	Name  string `json:"name" binding:"required"`
	Color string `json:"color" binding:"required"`
}

type WorkspaceOrder struct {
	StatusId    int   `json:"status_id" binding:"required"`
	WorkspaceId []int `json:"workspace_id" binding:"required"`
}

type WorkspaceUpdateOrderRequest struct {
	WorkspaceId int              `json:"workspace_id" binding:"required"`
	OrderStatus []WorkspaceOrder `json:"order_status" binding:"required"`
}

type WorkspaceCardUpdateOrderStatusRequest struct {
	OrderStatus []int `json:"order_status" binding:"required"`
}

type WorkspaceCardOrder struct {
	StatusId int   `json:"status_id" binding:"required"`
	CardId   []int `json:"card_id" binding:"required"`
}

type WorkspaceCardUpdateOrderRequest struct {
	CardId      int                  `json:"card_id" binding:"required"`
	OrderStatus []WorkspaceCardOrder `json:"order_status" binding:"required"`
}

type WorkspaceCardCreateActivityRequst struct {
	Message string `json:"message"`
}
