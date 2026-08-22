package dto

type CreateUserInput struct {
	Name string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Description *string `json:"description"`
	Icon_rul string `json:"icon_url"`
}
