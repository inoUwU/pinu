package presenters

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// UserPresenter ユーザーのプレゼンター
type UserPresenter struct {
	userService *services.UserService
}

// TODO: プレゼンターは何に依存するのか調べる.
