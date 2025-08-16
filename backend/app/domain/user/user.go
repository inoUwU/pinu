package user

import "time"

type UserID string
type LoginID string

type User struct {
	USER_ID       UserID
	LOGIN_ID      LoginID
	PASSWORD_HASH string
	PASSWORD_SALT string
	NAME          string
	IS_ADMIN      bool
	CREATED_AT    time.Time
}
