package structs

type HistoryLogResponse struct {
	ID          int    `json:"id"`
	User        string `json:"user"`
	Type        string `json:"type"`
	Action      string `json:"action"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type ListHistoryLogRequest struct {
	Type string `form:"type" binding:"oneof=general workspace"`
}
