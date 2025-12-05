package category

// Category カテゴリエンティティ
type Category struct {
	CategoryID   string `json:"category_id"`
	Name         string `json:"name"`
	DisplayOrder int    `json:"display_order"`
	ImageURL     string `json:"image_url"`
}
