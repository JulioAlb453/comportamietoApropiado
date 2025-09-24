package domain

type DirectorUser struct {
	ID        int    `json:"id"`
	Name      string `json:"nombre"`
	apellidos string `json:"apellidos"`
	email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
