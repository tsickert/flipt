package model

type User struct {
	Id        string `json:"id,omitempty"`
	Email     string `json:"email,omitempty"`
	Name      string `json:"name,omitempty"`
	Enabled   bool   `json:"enabled,omitempty"`
	CreatedAt int    `json:"createdAt,omitempty"`
	UpdatedAt int    `json:"updatedAt,omitempty"`
}
