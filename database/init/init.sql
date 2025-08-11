-- Drop existing types and tables
DROP TYPE IF EXISTS table_status CASCADE;
DROP TYPE IF EXISTS order_status CASCADE;
DROP TABLE IF EXISTS 
    settings, users, categories, menus, menu_options, menu_option_assignments,
    tables, table_sessions, order_groups, order_items, order_item_options 
CASCADE;

-- ENUM types
CREATE TYPE table_status AS ENUM ('available', 'occupied', 'billing');
CREATE TYPE order_status AS ENUM ('pending', 'preparing', 'served', 'cancelled');

-- Settings
CREATE TABLE settings (
    key VARCHAR(255) PRIMARY KEY,
    value TEXT NOT NULL
);

-- Users
CREATE TABLE users (
    user_id UUID PRIMARY KEY,
    login_id VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    password_salt VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    is_admin BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Categories
CREATE TABLE categories (
    category_id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    display_order INTEGER NOT NULL DEFAULT 0,
    image_url VARCHAR(255)
);

-- Menus
CREATE TABLE menus (
    menu_id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    price DECIMAL(10,2) NOT NULL,
    image_url VARCHAR(255),
    is_sold_out BOOLEAN NOT NULL DEFAULT FALSE,
    category_id VARCHAR(255) NOT NULL REFERENCES categories(category_id)
);

-- Menu Options
CREATE TABLE menu_options (
    menu_option_id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    price DECIMAL(10,2) NOT NULL
);

-- Menu Option Assignments
CREATE TABLE menu_option_assignments (
    menu_id VARCHAR(255) NOT NULL,
    menu_option_id VARCHAR(255) NOT NULL,
    PRIMARY KEY (menu_id, menu_option_id),
    FOREIGN KEY (menu_id) REFERENCES menus(menu_id) ON DELETE CASCADE,
    FOREIGN KEY (menu_option_id) REFERENCES menu_options(menu_option_id) ON DELETE CASCADE
);

-- Tables（外部キーは後で）
CREATE TABLE tables (
    table_id VARCHAR(255) PRIMARY KEY,
    status table_status NOT NULL DEFAULT 'available',
    current_orders_id UUID, -- 外部キー後付け
    last_updated TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Table Sessions（外部キーは後で）
CREATE TABLE table_sessions (
    table_session_id UUID PRIMARY KEY,
    table_id VARCHAR(255) NOT NULL, -- 外部キー後付け
    orders_id UUID NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    last_used TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    expires_at TIMESTAMP WITH TIME ZONE NOT NULL
);

-- Order Groups（外部キーは後で）
CREATE TABLE order_groups (
    orders_id UUID PRIMARY KEY,
    table_session_id UUID NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Order Items
CREATE TABLE order_items (
    order_item_id VARCHAR(255) PRIMARY KEY,
    orders_id UUID NOT NULL REFERENCES order_groups(orders_id),
    menu_id VARCHAR(255) NOT NULL REFERENCES menus(menu_id),
    quantity INTEGER NOT NULL,
    price_at_order DECIMAL(10,2) NOT NULL,
    status order_status NOT NULL DEFAULT 'pending',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Order Item Options
CREATE TABLE order_item_options (
    order_item_id VARCHAR(255) NOT NULL,
    menu_option_id VARCHAR(255) NOT NULL,
    PRIMARY KEY (order_item_id, menu_option_id),
    FOREIGN KEY (order_item_id) REFERENCES order_items(order_item_id) ON DELETE CASCADE,
    FOREIGN KEY (menu_option_id) REFERENCES menu_options(menu_option_id) ON DELETE CASCADE
);

-- Sessions
CREATE TABLE sessions (
    session_id UUID PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(user_id),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    expires_at TIMESTAMP NOT NULL,
    ip_address TEXT,
    user_agent TEXT
);

-- 外部キー制約の追加（循環回避）
ALTER TABLE table_sessions
    ADD CONSTRAINT fk_table_sessions_table FOREIGN KEY (table_id) REFERENCES tables(table_id);

ALTER TABLE order_groups
    ADD CONSTRAINT fk_order_groups_table_session FOREIGN KEY (table_session_id) REFERENCES table_sessions(table_session_id);

ALTER TABLE tables
    ADD CONSTRAINT fk_tables_current_orders FOREIGN KEY (current_orders_id) REFERENCES order_groups(orders_id);

-- Indexes
CREATE INDEX idx_table_sessions_orders_id ON table_sessions(orders_id);
CREATE INDEX idx_order_items_orders_id ON order_items(orders_id);
CREATE INDEX idx_order_item_options_order_item_id ON order_item_options(order_item_id);
CREATE INDEX idx_order_item_options_menu_option_id ON order_item_options(menu_option_id);

-- Comments
COMMENT ON TABLE settings IS 'アプリ全体の設定（ロゴURLなど）';
COMMENT ON TABLE users IS 'ユーザーアカウント';
COMMENT ON TABLE categories IS 'メニュー分類';
COMMENT ON TABLE menus IS '提供メニュー';
COMMENT ON TABLE menu_options IS 'オプション設定';
COMMENT ON TABLE menu_option_assignments IS 'メニューとオプションの関係';
COMMENT ON TABLE tables IS '店舗の物理テーブル';
COMMENT ON TABLE table_sessions IS 'テーブルごとのセッション';
COMMENT ON TABLE order_groups IS '注文グループ';
COMMENT ON TABLE order_items IS '注文明細';
COMMENT ON TABLE order_item_options IS '明細に付属したオプション';

COMMENT ON COLUMN order_items.price_at_order IS '注文時の価格（後の価格変更に影響されない）';
COMMENT ON COLUMN tables.status IS 'available: 空席 / occupied: 使用中 / billing: 会計待ち';
COMMENT ON COLUMN order_items.status IS '注文ステータス（pending, preparing, served, cancelled）';
