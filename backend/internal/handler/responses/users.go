package responses

// UserResponse represents the response payload for user-related operations.
type UserResponse struct {
	Username        string `json:"username"`
	Role            string `json:"role"`
	RoleDescription string `json:"role_description"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
	DeletedAt       string `json:"deleted_at"`
}
