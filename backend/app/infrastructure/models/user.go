package models

import (
	"github.com/uptrace/bun"
	"time"
)

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`

	USER_ID       string    `bun:",pk"`
	LOGIN_ID      string    `bun:""`
	NAME          string    `bun:""`
	PASSWORD_HASH string    `bun:""`
	PASSWORD_SALT string    `bun:""`
	IS_ADMIN      bool      `bun:""`
	CREATED_AT    time.Time `bun:""`
}
