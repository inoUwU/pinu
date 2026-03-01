package table

import (
	"errors"
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

var ErrInvalidTableStatusTransition = errors.New("invalid table status transition")

// Table テーブルエンティティ
type Table struct {
	TableID               TableID     `json:"table_id"`
	Status                TableStatus `json:"status"`
	CurrentTableSessionID *uuid.UUID  `json:"current_table_session_id,omitempty"`
	LastUpdated           time.Time   `json:"last_updated"`
}

func (t *Table) CanTransitionTo(nextStatus TableStatus) bool {
	if t == nil {
		return false
	}

	if t.Status == nextStatus {
		return true
	}

	switch t.Status {
	case StatusAvailable:
		return nextStatus == StatusOccupied
	case StatusOccupied:
		return nextStatus == StatusBilling || nextStatus == StatusAvailable
	case StatusBilling:
		return nextStatus == StatusAvailable
	default:
		return false
	}
}

func (t *Table) TransitionTo(nextStatus TableStatus, now time.Time) error {
	if !t.CanTransitionTo(nextStatus) {
		return ErrInvalidTableStatusTransition
	}

	t.Status = nextStatus
	t.LastUpdated = now

	if nextStatus == StatusAvailable {
		t.CurrentTableSessionID = nil
	}

	return nil
}

func (t *Table) HasActiveSession() bool {
	return t != nil && t.CurrentTableSessionID != nil
}
