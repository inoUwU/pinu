package main

import (
	"database/sql"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/uptrace/bun/driver/sqliteshim"
	"log"
	"os"
)

func main() {
	// ルートの.envファイルを読み込む
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf(".envファイルの読み込みに失敗しました: %v", err)
	}

	port := os.Getenv("BACKEND_PORT")
	if port == "" {
		port = "8000" // デフォルトポート
	}

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":" + port)

}
