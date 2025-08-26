package output

import (
	"time"

	"inoUwU/pinu/app/domain/menu"
)

// GetMenusOutput メニュー一覧取得の出力
type GetMenusOutput struct {
	Menus []*menu.Menu `json:"menus"`
	Count int          `json:"count"`
}

// GetMenuByIDOutput メニュー単一取得の出力
type GetMenuByIDOutput struct {
	Menu *menu.Menu `json:"menu"`
}

// GetMenusByCategoryOutput カテゴリ別メニュー取得の出力
type GetMenusByCategoryOutput struct {
	Menus      []*menu.Menu `json:"menus"`
	Count      int          `json:"count"`
	CategoryID string       `json:"category_id"`
}

// CreateMenuOutput メニュー作成の出力
type CreateMenuOutput struct {
	Menu        *menu.Menu `json:"menu"`
	Message     string     `json:"message"`
	CreatedAt   time.Time  `json:"created_at"`
	MenuID      string     `json:"menu_id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Price       float64    `json:"price"`
	ImageURL    string     `json:"image_url"`
	IsSoldOut   bool       `json:"is_sold_out"`
	CategoryID  string     `json:"category_id"`
}

// UpdateMenuOutput メニュー更新の出力
type UpdateMenuOutput struct {
	Menu      *menu.Menu `json:"menu"`
	Message   string     `json:"message"`
	UpdatedAt time.Time  `json:"updated_at"`
}

// DeleteMenuOutput メニュー削除の出力
type DeleteMenuOutput struct {
	MenuID    string    `json:"menu_id"`
	Message   string    `json:"message"`
	DeletedAt time.Time `json:"deleted_at"`
}
