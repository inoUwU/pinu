package menu_option

type MenuOptionID string

// MenuOption メニューオプションエンティティ
type MenuOption struct {
	MenuOptionID MenuOptionID `json:"menu_option_id" bun:",pk"`
	Name         string       `json:"name"`
	Price        float64      `json:"price"`
}
