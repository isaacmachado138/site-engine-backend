package dtos

// SiteCategoryCreateDTO para criar ligação site-categoria
type SiteCategoryCreateDTO struct {
	SiteID     uint `json:"site_id" binding:"required"`
	CategoryID uint `json:"category_id" binding:"required"`
}

// SiteCategoryResponseDTO para resposta da ligação site-categoria
type SiteCategoryResponseDTO struct {
	SiteID     uint `json:"site_id"`
	CategoryID uint `json:"category_id"`
}
