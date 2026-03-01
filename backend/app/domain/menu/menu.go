package menu

const MinMenuPrice = 0.0

// Menu メニューエンティティ
type Menu struct {
	MenuID      string  `json:"menu_id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImageURL    string  `json:"image_url"`
	IsSoldOut   bool    `json:"is_sold_out"`
	CategoryID  string  `json:"category_id"`
}

func (m *Menu) IsOrderable() bool {
	if m == nil {
		return false
	}

	if m.IsSoldOut {
		return false
	}

	if m.Name == "" {
		return false
	}

	return m.Price >= MinMenuPrice
}
