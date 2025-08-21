package settings

// Settings 設定エンティティ
type Settings struct {
	Key   string `json:"key" bun:",pk"`
	Value string `json:"value"`
}
