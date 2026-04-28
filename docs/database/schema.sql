-- SQL dump generated using DBML (dbml.dbdiagram.io)
-- Database: PostgreSQL
-- Generated at: 2026-04-28T15:38:53.815Z

CREATE TYPE "table_status" AS ENUM (
  'available',
  'occupied',
  'billing'
);

CREATE TYPE "order_status" AS ENUM (
  'pending',
  'preparing',
  'served',
  'cancelled'
);

CREATE TYPE "order_group_status" AS ENUM (
  'open',
  'closed',
  'cancelled'
);

CREATE TABLE "settings" (
  "key" varchar(255) PRIMARY KEY,
  "value" text NOT NULL,
  "created_at" timestamptz DEFAULT (CURRENT_TIMESTAMP),
  "updated_at" timestamptz DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE "users" (
  "user_id" uuid PRIMARY KEY,
  "login_id" varchar(255) UNIQUE NOT NULL,
  "password_hash" varchar(255) NOT NULL,
  "password_salt" varchar(255) NOT NULL,
  "name" varchar(255) NOT NULL,
  "is_admin" boolean NOT NULL DEFAULT false,
  "created_at" timestamptz DEFAULT (CURRENT_TIMESTAMP),
  "updated_at" timestamptz DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE "categories" (
  "category_id" varchar(255) PRIMARY KEY,
  "name" varchar(255) NOT NULL,
  "display_order" integer NOT NULL DEFAULT 0,
  "image_url" varchar(255),
  "created_at" timestamptz DEFAULT (CURRENT_TIMESTAMP),
  "updated_at" timestamptz DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE "menus" (
  "menu_id" varchar(255) PRIMARY KEY,
  "name" varchar(255) NOT NULL,
  "description" text,
  "price" decimal(10,2) NOT NULL,
  "image_url" varchar(255),
  "is_sold_out" boolean NOT NULL DEFAULT false,
  "category_id" varchar(255) NOT NULL,
  "created_at" timestamptz DEFAULT (CURRENT_TIMESTAMP),
  "updated_at" timestamptz DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE "menu_options" (
  "menu_option_id" varchar(255) PRIMARY KEY,
  "name" varchar(255) NOT NULL,
  "price" decimal(10,2) NOT NULL,
  "created_at" timestamptz DEFAULT (CURRENT_TIMESTAMP),
  "updated_at" timestamptz DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE "menu_option_assignments" (
  "menu_id" varchar(255) NOT NULL,
  "menu_option_id" varchar(255) NOT NULL,
  PRIMARY KEY ("menu_id", "menu_option_id")
);

CREATE TABLE "tables" (
  "table_id" varchar(255) PRIMARY KEY,
  "qr_token" uuid NOT NULL,
  "status" table_status NOT NULL DEFAULT 'available',
  "current_table_session_id" uuid,
  "last_updated" timestamptz DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE "table_sessions" (
  "table_session_id" uuid PRIMARY KEY,
  "table_id" varchar(255) NOT NULL,
  "is_revoked" boolean NOT NULL DEFAULT false,
  "created_at" timestamptz DEFAULT (CURRENT_TIMESTAMP),
  "last_used" timestamptz DEFAULT (CURRENT_TIMESTAMP),
  "expires_at" timestamptz NOT NULL
);

CREATE TABLE "order_groups" (
  "orders_id" uuid PRIMARY KEY,
  "table_session_id" uuid NOT NULL,
  "status" order_group_status NOT NULL DEFAULT 'open',
  "created_at" timestamptz DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE "order_items" (
  "order_item_id" varchar(255) PRIMARY KEY,
  "orders_id" uuid NOT NULL,
  "menu_id" varchar(255) NOT NULL,
  "quantity" integer NOT NULL,
  "price_at_order" decimal(10,2) NOT NULL,
  "status" order_status NOT NULL DEFAULT 'pending',
  "created_at" timestamptz DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE "order_item_options" (
  "order_item_id" varchar(255) NOT NULL,
  "menu_option_id" varchar(255) NOT NULL,
  PRIMARY KEY ("order_item_id", "menu_option_id")
);

CREATE TABLE "sessions" (
  "session_id" varchar(255) PRIMARY KEY,
  "user_id" uuid NOT NULL,
  "refresh_token" varchar(512) NOT NULL,
  "is_revoked" boolean NOT NULL DEFAULT false,
  "created_at" timestamptz NOT NULL DEFAULT (NOW()),
  "expires_at" timestamptz NOT NULL
);

CREATE INDEX ON "categories" ("display_order");

CREATE INDEX "idx_menus_category_id" ON "menus" ("category_id");

CREATE INDEX "idx_menus_is_sold_out" ON "menus" ("is_sold_out");

CREATE UNIQUE INDEX "ux_tables_qr_token" ON "tables" ("qr_token");

CREATE INDEX "idx_tables_current_table_session_id" ON "tables" ("current_table_session_id");

CREATE INDEX "idx_table_sessions_table_id" ON "table_sessions" ("table_id");

CREATE INDEX "idx_order_groups_table_session_id" ON "order_groups" ("table_session_id");

CREATE INDEX "idx_order_groups_created_at" ON "order_groups" ("created_at");

CREATE UNIQUE INDEX "ux_order_groups_session_open" ON "order_groups" ("table_session_id");

CREATE INDEX "idx_order_items_orders_id" ON "order_items" ("orders_id");

CREATE INDEX "idx_order_items_menu_id" ON "order_items" ("menu_id");

CREATE INDEX "idx_order_items_created_at" ON "order_items" ("created_at");

CREATE INDEX "idx_order_item_options_order_item_id" ON "order_item_options" ("order_item_id");

CREATE INDEX "idx_order_item_options_menu_option_id" ON "order_item_options" ("menu_option_id");

COMMENT ON TABLE "settings" IS 'アプリケーション全体の設定を保存（例：ロゴのURL、店舗のテーブル総数）';

COMMENT ON COLUMN "settings"."key" IS '設定キー';

COMMENT ON COLUMN "settings"."value" IS '設定値';

COMMENT ON TABLE "users" IS '従業員・管理者などのユーザーを表す';

COMMENT ON COLUMN "users"."user_id" IS 'ユーザーID';

COMMENT ON COLUMN "users"."login_id" IS 'ログインID';

COMMENT ON COLUMN "users"."password_hash" IS 'パスワードハッシュ';

COMMENT ON COLUMN "users"."password_salt" IS 'パスワードソルト';

COMMENT ON COLUMN "users"."name" IS '表示名';

COMMENT ON COLUMN "users"."is_admin" IS '管理者フラグ';

COMMENT ON COLUMN "users"."created_at" IS '作成時間';

COMMENT ON TABLE "categories" IS 'メニューの分類を管理';

COMMENT ON COLUMN "categories"."category_id" IS 'カテゴリID';

COMMENT ON COLUMN "categories"."name" IS 'カテゴリ名';

COMMENT ON COLUMN "categories"."display_order" IS '表示順';

COMMENT ON COLUMN "categories"."image_url" IS '画像URL';

COMMENT ON TABLE "menus" IS '提供される各メニューの詳細情報を保存';

COMMENT ON COLUMN "menus"."menu_id" IS 'メニューID';

COMMENT ON COLUMN "menus"."name" IS 'メニュー名';

COMMENT ON COLUMN "menus"."description" IS '説明';

COMMENT ON COLUMN "menus"."price" IS '価格';

COMMENT ON COLUMN "menus"."image_url" IS '画像URL';

COMMENT ON COLUMN "menus"."is_sold_out" IS '売り切れフラグ';

COMMENT ON COLUMN "menus"."category_id" IS 'カテゴリID';

COMMENT ON TABLE "menu_options" IS 'メニューに追加できるオプション（トッピング、サイズ変更など）を管理';

COMMENT ON COLUMN "menu_options"."menu_option_id" IS 'オプションID';

COMMENT ON COLUMN "menu_options"."name" IS 'オプション名';

COMMENT ON COLUMN "menu_options"."price" IS '追加価格';

COMMENT ON TABLE "menu_option_assignments" IS 'どのメニューにどのオプションが利用可能かを示す中間テーブル';

COMMENT ON TABLE "tables" IS '店舗内の物理テーブルを表す。固定 QR 識別子、テーブル状態、現在のテーブルセッションIDを管理';

COMMENT ON COLUMN "tables"."table_id" IS 'テーブルID';

COMMENT ON COLUMN "tables"."qr_token" IS '固定 QR 識別子。QR 読み取り時に table_id 解決へ利用する';

COMMENT ON COLUMN "tables"."status" IS '状態（available: 空席 / occupied: 使用中 / billing: 会計待ち）';

COMMENT ON COLUMN "tables"."current_table_session_id" IS '現在のテーブルセッションID';

COMMENT ON COLUMN "tables"."last_updated" IS '最終更新';

COMMENT ON TABLE "table_sessions" IS '各テーブルのセッション（QR入店〜会計・有効期限まで）を表す';

COMMENT ON COLUMN "table_sessions"."table_session_id" IS 'テーブルセッションID';

COMMENT ON COLUMN "table_sessions"."table_id" IS 'テーブルID';

COMMENT ON COLUMN "table_sessions"."is_revoked" IS '無効化フラグ';

COMMENT ON COLUMN "table_sessions"."created_at" IS '作成時間';

COMMENT ON COLUMN "table_sessions"."last_used" IS '最終利用時刻';

COMMENT ON COLUMN "table_sessions"."expires_at" IS '有効期限';

COMMENT ON TABLE "order_groups" IS 'セッション配下の注文グループ。statusでopen/closed/cancelledを管理';

COMMENT ON COLUMN "order_groups"."orders_id" IS '注文グループID';

COMMENT ON COLUMN "order_groups"."table_session_id" IS 'テーブルセッションID';

COMMENT ON COLUMN "order_groups"."status" IS 'グループ状態（open/closed/cancelled）';

COMMENT ON COLUMN "order_groups"."created_at" IS '作成時間';

COMMENT ON TABLE "order_items" IS '各注文グループに含まれるメニュー項目を管理';

COMMENT ON COLUMN "order_items"."order_item_id" IS '注文明細ID';

COMMENT ON COLUMN "order_items"."orders_id" IS '注文グループID';

COMMENT ON COLUMN "order_items"."menu_id" IS 'メニューID';

COMMENT ON COLUMN "order_items"."quantity" IS '数量';

COMMENT ON COLUMN "order_items"."price_at_order" IS '注文時の価格（後の価格変更に影響されない）';

COMMENT ON COLUMN "order_items"."status" IS '注文ステータス（pending, preparing, served, cancelled）';

COMMENT ON COLUMN "order_items"."created_at" IS '作成時間';

COMMENT ON TABLE "order_item_options" IS '注文された各メニュー項目にどのオプションが選択されたかを記録する中間テーブル';

COMMENT ON TABLE "sessions" IS 'ユーザーのログインセッション情報を管理';

COMMENT ON COLUMN "sessions"."session_id" IS 'セッションID';

COMMENT ON COLUMN "sessions"."user_id" IS 'ユーザーID';

COMMENT ON COLUMN "sessions"."refresh_token" IS 'リフレッシュトークン';

COMMENT ON COLUMN "sessions"."is_revoked" IS '無効化フラグ';

COMMENT ON COLUMN "sessions"."created_at" IS '作成時間';

COMMENT ON COLUMN "sessions"."expires_at" IS '有効期限';

ALTER TABLE "menus" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("category_id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "menu_option_assignments" ADD FOREIGN KEY ("menu_id") REFERENCES "menus" ("menu_id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "menu_option_assignments" ADD FOREIGN KEY ("menu_option_id") REFERENCES "menu_options" ("menu_option_id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "tables" ADD FOREIGN KEY ("current_table_session_id") REFERENCES "table_sessions" ("table_session_id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "table_sessions" ADD FOREIGN KEY ("table_id") REFERENCES "tables" ("table_id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "order_groups" ADD FOREIGN KEY ("table_session_id") REFERENCES "table_sessions" ("table_session_id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "order_items" ADD FOREIGN KEY ("orders_id") REFERENCES "order_groups" ("orders_id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "order_items" ADD FOREIGN KEY ("menu_id") REFERENCES "menus" ("menu_id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "order_item_options" ADD FOREIGN KEY ("order_item_id") REFERENCES "order_items" ("order_item_id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "order_item_options" ADD FOREIGN KEY ("menu_option_id") REFERENCES "menu_options" ("menu_option_id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "sessions" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id") DEFERRABLE INITIALLY IMMEDIATE;
