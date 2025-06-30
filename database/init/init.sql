-- pinu DB initialize SQL

-- Drop existing types and tables to ensure a clean setup
DROP TYPE IF EXISTS table_status CASCADE;
DROP TYPE IF EXISTS order_status CASCADE;
DROP TABLE IF EXISTS settings, users, tables, table_sessions, orders, categories, menus, menu_options, order_item_options, menu_option_assignments CASCADE;

-- Create custom ENUM types for status management
-- This enhances data integrity by restricting column values.

CREATE TYPE table_status AS ENUM (
    'available',  -- 空席
    'occupied',   -- 使用中
    'billing'     -- 会計済み
);

CREATE TYPE order_status AS ENUM (
    'pending',    -- 受付待ち
    'preparing',  -- 調理中
    'served',     -- 提供済み
    'cancelled'   -- キャンセル
);

-- Create tables

-- Application settings table
CREATE TABLE settings (
    key VARCHAR(255) PRIMARY KEY,
    value TEXT NOT NULL
);

-- User accounts table
CREATE TABLE users (
    user_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    username VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    is_admin BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Categories for menus
CREATE TABLE categories (
    category_id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    display_order INTEGER NOT NULL DEFAULT 0
);

-- Menus table
CREATE TABLE menus (
    menu_id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    price DECIMAL(10, 0) NOT NULL,
    image_url VARCHAR(255),
    is_sold_out BOOLEAN NOT NULL DEFAULT FALSE,
    category_id VARCHAR(255) NOT NULL,
    FOREIGN KEY (category_id) REFERENCES categories(category_id)
);

-- Options for menus (e.g., toppings, size changes)
CREATE TABLE menu_options (
    menu_option_id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    price DECIMAL(10, 0) NOT NULL
);

-- Assignment table for many-to-many relationship between menus and options
CREATE TABLE menu_option_assignments (
    menu_id VARCHAR(255) NOT NULL,
    menu_option_id VARCHAR(255) NOT NULL,
    PRIMARY KEY (menu_id, menu_option_id), -- Composite Key
    FOREIGN KEY (menu_id) REFERENCES menus(menu_id) ON DELETE CASCADE,
    FOREIGN KEY (menu_option_id) REFERENCES menu_options(menu_option_id) ON DELETE CASCADE
);

-- Physical tables in the restaurant
CREATE TABLE tables (
    table_id VARCHAR(255) PRIMARY KEY,
    status table_status NOT NULL DEFAULT 'available',
    current_orders_id UUID, -- Can be nullable
    last_updated TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Customer session management
CREATE TABLE table_sessions (
    table_session_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    table_id VARCHAR(255) NOT NULL,
    orders_id UUID NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    last_used TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    expires_at TIMESTAMP WITH TIME ZONE NOT NULL,
    FOREIGN KEY (table_id) REFERENCES tables(table_id)
);

-- Individual order items
CREATE TABLE orders (
    order_item_id VARCHAR(255) PRIMARY KEY,
    orders_id UUID NOT NULL,
    menu_id VARCHAR(255) NOT NULL,
    quantity INTEGER NOT NULL,
    price_at_order DECIMAL(10, 0) NOT NULL,
    status order_status NOT NULL DEFAULT 'pending',
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (menu_id) REFERENCES menus(menu_id)
    -- Note: A foreign key to table_sessions(orders_id) might be beneficial
);

-- Options selected for each order item
CREATE TABLE order_item_options (
    order_item_id VARCHAR(255) NOT NULL,
    menu_option_id VARCHAR(255) NOT NULL,
    PRIMARY KEY (order_item_id, menu_option_id), -- Composite Key
    FOREIGN KEY (order_item_id) REFERENCES orders(order_item_id) ON DELETE CASCADE,
    FOREIGN KEY (menu_option_id) REFERENCES menu_options(menu_option_id) ON DELETE CASCADE
);

-- Add comments on tables and columns for clarity
COMMENT ON TABLE settings IS 'アプリケーション全体の設定を保存します（例：ロゴのURL、店舗のテーブル総数）。';
COMMENT ON TABLE users IS 'ユーザーアカウント（スタッフ、管理者）を管理します。';
COMMENT ON TABLE categories IS 'メニューの分類を管理します。';
COMMENT ON TABLE menus IS '提供される各メニューの詳細情報を保存します。';
COMMENT ON TABLE menu_options IS 'メニューに追加できるオプション（トッピング、サイズ変更など）を管理します。';
COMMENT ON TABLE menu_option_assignments IS 'どのメニューにどのオプションが利用可能かを示す中間テーブルです。';
COMMENT ON TABLE tables IS '店舗内の物理テーブルを表します。';
COMMENT ON TABLE table_sessions IS '顧客の注文セッションを管理します。QRコード読み取りで生成されます。';
COMMENT ON TABLE orders IS '各注文に含まれる個別のメニュー項目を管理します。';
COMMENT ON TABLE order_item_options IS '注文された各メニュー項目にどのオプションが選択されたかを記録する中間テーブルです。';

COMMENT ON COLUMN tables.status IS 'テーブルの状態（available: 空席, occupied: 使用中, billing: 会計済み）';
COMMENT ON COLUMN orders.status IS '注文の状態（pending: 受付待ち, preparing: 調理中, served: 提供済み, cancelled: キャンセル）';
COMMENT ON COLUMN orders.price_at_order IS 'メニュー価格の変更に影響されないよう、注文時の価格を記録します。';
