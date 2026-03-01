package models

import (
	"time"

	"github.com/uptrace/bun"
)

// TableModel テーブルのインフラストラクチャーモデル
type TableModel struct {
	bun.BaseModel `bun:"table:tables"`

	TableID               string    `bun:"table_id,pk"`
	Status                string    `bun:"status,notnull"`
	CurrentTableSessionID *string   `bun:"current_table_session_id"`
	LastUpdated           time.Time `bun:"last_updated,notnull,default:current_timestamp"`
}
