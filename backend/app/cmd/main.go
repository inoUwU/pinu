package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/extra/bundebug"
	"inoUwU/pinu/app/api"
)

func main() {
	// ルートの.envファイルを読み込む
	err := godotenv.Load("../.env")
	if err != nil {
		log.Printf(".envファイルの読み込みに失敗しました: %v", err)
	}

	// データベース接続
	db, err := initDatabase()
	if err != nil {
		log.Fatalf("データベース接続に失敗しました: %v", err)
	}
	defer db.Close()

	port := os.Getenv("BACKEND_PORT")
	if port == "" {
		port = "8000" // デフォルトポート
	}

	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			return c.Status(code).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
	})

	// ミドルウェア設定
	app.Use(logger.New())
	app.Use(cors.New())

	// ヘルスチェック
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Pinu API is running!")
	})

	// 依存性注入コンテナの設定
	injector := Injection(db)

	// APIルートを設定
	api.SetupRoutes(app, injector)

	log.Printf("サーバーをポート %s で起動します", port)
	log.Fatal(app.Listen(":" + port))
}

func initDatabase() (*sql.DB, error) {

	user := os.Getenv("DATABASE_USER")
	password := os.Getenv("DATABASE_PASSWORD")
	host := os.Getenv("DATABASE_HOST")
	port := os.Getenv("DATABASE_PORT")
	dbname := os.Getenv("DATABASE_NAME")

	if user == "" || password == "" || host == "" || port == "" || dbname == "" {
		return nil, fmt.Errorf("データベース接続情報が不足しています")
	}

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, dbname)
	pool, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("データベース接続エラー: %w", err)
	}

	// 接続テスト
	if err := pool.Ping(); err != nil {
		return nil, fmt.Errorf("データベース接続テストエラー: %w", err)
	}

	// Postgre SQL用の ダイアレクトを設定
	db := bun.NewDB(pool, pgdialect.New())
	defer db.Close()

	// クエリを標準出力する設定
	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
	))

	return db.DB, nil
}
