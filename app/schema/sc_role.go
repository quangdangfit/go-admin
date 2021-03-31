package schema

type Role struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type RoleBodyParam struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
}
