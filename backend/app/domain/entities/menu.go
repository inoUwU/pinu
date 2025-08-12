package entities

type Menu struct {
	CategoryID   string `json:"category_id,omitempty"`
	Name         string `json:"name,omitempty"`
	DisplayOrder int    `json:"display_order,omitempty"`
	ImageURL     string `json:"image_url,omitempty"`
}
