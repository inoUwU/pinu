package main

import (
	"database/sql"
	"fmt"
	"inoUwU/pinu/app/api"
	mylogger "inoUwU/pinu/app/infrastructure/logger"
	"inoUwU/pinu/app/middleware"
	"log"
	"log/slog"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/extra/bundebug"

	// PostgreSQLドライバーを匿名インポート
	_ "github.com/lib/pq"
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

	app.Use(middleware.Recover())

	// TODO: production環境では、CORS設定を適切に行う必要があります
	// TODO: Read .env file for origin settings

	// corsミドルウェアを設定
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000,http://localhost:5500",
		AllowHeaders:     "Origin, Content-Type, Accept, Cache-Control",
		AllowCredentials: true,
		AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS",
		ExposeHeaders:    "Content-Length, Content-Type, Connection, Cache-Control",
	}))

	// ロガーの設定
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})

	slogger := slog.New(handler)
	appLogger := mylogger.NewSlogLogger(slogger)

	// 依存性注入コンテナの設定
	injector := middleware.Injection(db, appLogger)

	// APIルートを設定
	api.SetupRoutes(app, injector)

	log.Printf("サーバーをポート %s で起動します", port)
	log.Fatal(app.Listen(":" + port))
}

func initDatabase() (*bun.DB, error) {

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")

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

	// クエリを標準出力する設定
	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
	))

	db.NewSelect()

	return db, nil
}
