package dtos

// CategoryCreateDTO para criação de categoria
type CategoryCreateDTO struct {
	Name        string `json:"category_name" validate:"required,min=2,max=100"`
	Description string `json:"category_description" validate:"max=500"`
	Active      bool   `json:"category_active"`
	Icon        string `json:"category_icon" validate:"max=100"`
	Main        bool   `json:"category_main"`
}

// CategoryUpdateDTO para atualização de categoria
type CategoryUpdateDTO struct {
	Name        *string `json:"category_name,omitempty" validate:"omitempty,min=2,max=100"`
	Description *string `json:"category_description,omitempty" validate:"omitempty,max=500"`
	Active      *bool   `json:"category_active,omitempty"`
	Icon        *string `json:"category_icon,omitempty" validate:"omitempty,max=100"`
	Main        *bool   `json:"category_main,omitempty"`
}

// CategoryResponseDTO para resposta de categoria
type CategoryResponseDTO struct {
	ID          uint   `json:"category_id"`
	Name        string `json:"category_name"`
	Description string `json:"category_description"`
	Active      bool   `json:"category_active"`
	Icon        string `json:"category_icon"`
	Main        bool   `json:"category_main"`
}
