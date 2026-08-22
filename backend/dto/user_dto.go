package dto

type CreateUserInput struct {
	Name        string  `json:"name" binding:"required"`
	Email       string  `json:"email" binding:"required"`
	Password    string  `json:"password" binding:"required"`
	Description *string `json:"description"`
	Icon_rul    string  `json:"icon_url"`
}

type UpdateUserInput struct {
	Name        *string `json:"name"`
	Email       *string `json:"email"`
	Password    *string `json:"password"`
	Description *string `json:"description"`
	Icon_url    *string `json:"icon_url"`
}
