package table

import (
	"time"

	"github.com/google/uuid"
)

type TableID string
type TableStatus string

const (
	StatusAvailable TableStatus = "available"
	StatusOccupied  TableStatus = "occupied"
	StatusBilling   TableStatus = "billing"
)

// Table テーブルエンティティ
type Table struct {
	TableID         TableID     `json:"table_id"`
	Status          TableStatus `json:"status"`
	CurrentOrdersID *uuid.UUID  `json:"current_orders_id,omitempty"`
	LastUpdated     time.Time   `json:"last_updated"`
}
