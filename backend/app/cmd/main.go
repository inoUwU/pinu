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
	"github.com/uptrace/bun/driver/sqliteshim"
	"inoUwU/pinu/app/api"
)

func main() {
	// ルートの.envファイルを読み込む
	err := godotenv.Load(".env")
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

	// APIルートを設定
	api.SetupRoutes(app, db)

	log.Printf("サーバーをポート %s で起動します", port)
	log.Fatal(app.Listen(":" + port))
}

// initDatabase データベースを初期化する
func initDatabase() (*sql.DB, error) {
	// 開発環境ではSQLiteを使用
	dbPath := os.Getenv("DATABASE_PATH")
	if dbPath == "" {
		dbPath = "./database.db"
	}

	db, err := sql.Open(sqliteshim.ShimName, dbPath)
	if err != nil {
		return nil, fmt.Errorf("データベースオープンエラー: %w", err)
	}

	// 接続テスト
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("データベース接続テストエラー: %w", err)
	}

	// テーブル作成（本番環境では別途マイグレーションツールを使用）
	err = createTables(db)
	if err != nil {
		return nil, fmt.Errorf("テーブル作成エラー: %w", err)
	}

	return db, nil
}

// createTables テーブルを作成する（サンプル）
func createTables(db *sql.DB) error {
	createUserTable := `
	CREATE TABLE IF NOT EXISTS users (
		id TEXT PRIMARY KEY,
		name TEXT NOT NULL,
		email TEXT UNIQUE NOT NULL
	);`

	_, err := db.Exec(createUserTable)
	if err != nil {
		return err
	}

	// サンプルデータの挿入
	insertSampleData := `
	INSERT OR IGNORE INTO users (id, name, email) VALUES 
	('1', 'John Doe', 'john@example.com'),
	('2', 'Jane Smith', 'jane@example.com'),
	('3', 'Bob Johnson', 'bob@example.com');`

	_, err = db.Exec(insertSampleData)
	return err
}
