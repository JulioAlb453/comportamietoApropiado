package domain

type DirectorUser struct {
	ID        string    `json:"id"`
	Name      string `json:"nombre"`
	apellidos string `json:"apellidos"`
	email     string `json:"email"`
	password  string `json:"password"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
