-- Drop existing types and tables
DROP TYPE IF EXISTS table_status CASCADE;
DROP TYPE IF EXISTS order_status CASCADE;
-- 追加: order_group_status を先に掃除
DROP TYPE IF EXISTS order_group_status CASCADE;
DROP TABLE IF EXISTS
    settings, users, categories, menus, menu_options, menu_option_assignments,
    tables, table_sessions, order_groups, order_items, order_item_options 
CASCADE;

-- ENUM types
CREATE TYPE table_status AS ENUM ('available', 'occupied', 'billing');
CREATE TYPE order_status AS ENUM ('pending', 'preparing', 'served', 'cancelled');
-- 追加: 注文グループの状態
CREATE TYPE order_group_status AS ENUM ('open', 'closed', 'cancelled');

-- Settings
CREATE TABLE settings (
    key VARCHAR(255) PRIMARY KEY,
    value TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Users
CREATE TABLE users (
    user_id UUID PRIMARY KEY,
    login_id VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    password_salt VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    is_admin BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Categories
CREATE TABLE categories (
    category_id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    display_order INTEGER NOT NULL DEFAULT 0,
    image_url VARCHAR(255),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Menus
CREATE TABLE menus (
    menu_id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    price DECIMAL(10,2) NOT NULL,
    image_url VARCHAR(255),
    is_sold_out BOOLEAN NOT NULL DEFAULT FALSE,
    category_id VARCHAR(255) NOT NULL REFERENCES categories(category_id),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Menu Options
CREATE TABLE menu_options (
    menu_option_id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    price DECIMAL(10,2) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
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
    -- 変更: current_orders_id -> current_table_session_id
    current_table_session_id UUID, -- 外部キー後付け
    last_updated TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Table Sessions（外部キーは後で）
CREATE TABLE table_sessions (
    table_session_id UUID PRIMARY KEY,
    table_id VARCHAR(255) NOT NULL, -- 外部キー後付け
    -- 変更: orders_id を削除
    is_revoked BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    -- 追加: 最終利用時刻
    last_used TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    expires_at TIMESTAMP WITH TIME ZONE NOT NULL
);

-- Order Groups（外部キーは後で）
CREATE TABLE order_groups (
    orders_id UUID PRIMARY KEY,
    table_session_id UUID NOT NULL,
    -- 追加: グループ状態（開いている/締めた/キャンセル）
    status order_group_status NOT NULL DEFAULT 'open',
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

-- Sessions（ユーザーのログインセッション）
CREATE TABLE sessions (
    -- 変更: PostgreSQLにstring型はないためVARCHARに修正
    session_id VARCHAR(255) PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(user_id),
    refresh_token VARCHAR(512) NOT NULL,
    is_revoked BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    expires_at TIMESTAMP WITH TIME ZONE NOT NULL
);

-- 外部キー制約の追加（循環回避）
ALTER TABLE table_sessions
    ADD CONSTRAINT fk_table_sessions_table FOREIGN KEY (table_id) REFERENCES tables(table_id);

ALTER TABLE order_groups
    ADD CONSTRAINT fk_order_groups_table_session FOREIGN KEY (table_session_id) REFERENCES table_sessions(table_session_id);

-- 変更: tables -> order_groups 参照を廃止し、現在のセッション参照に切替
ALTER TABLE tables
    ADD CONSTRAINT fk_tables_current_table_session FOREIGN KEY (current_table_session_id) REFERENCES table_sessions(table_session_id);

-- Indexes
-- 変更: table_sessions.orders_id のインデックスは不要
-- CREATE INDEX idx_table_sessions_orders_id ON table_sessions(orders_id);
CREATE INDEX idx_table_sessions_table_id ON table_sessions(table_id);
CREATE INDEX idx_tables_current_table_session_id ON tables(current_table_session_id);
CREATE INDEX idx_order_items_orders_id ON order_items(orders_id);
CREATE INDEX idx_order_item_options_order_item_id ON order_item_options(order_item_id);
CREATE INDEX idx_order_item_options_menu_option_id ON order_item_options(menu_option_id);
-- 同一セッション内で開いている注文グループは高々1つ
CREATE UNIQUE INDEX ux_order_groups_session_open
  ON order_groups(table_session_id)
  WHERE status = 'open';

-- Comments
COMMENT ON TABLE settings IS 'アプリ全体の設定（ロゴURLなど）';
COMMENT ON TABLE users IS 'ユーザーアカウント';
COMMENT ON TABLE categories IS 'メニュー分類';
COMMENT ON TABLE menus IS '提供メニュー';
COMMENT ON TABLE menu_options IS 'オプション設定';
COMMENT ON TABLE menu_option_assignments IS 'メニューとオプションの関係';
COMMENT ON TABLE tables IS '店舗の物理テーブル（current_table_session_idで現在のセッションを参照）';
COMMENT ON TABLE table_sessions IS 'テーブルごとのセッション';
COMMENT ON TABLE order_groups IS '注文グループ（セッション配下）';
COMMENT ON TABLE order_items IS '注文明細';
COMMENT ON TABLE order_item_options IS '明細に付属したオプション';

COMMENT ON COLUMN order_items.price_at_order IS '注文時の価格（後の価格変更に影響されない）';
COMMENT ON COLUMN tables.status IS 'available: 空席 / occupied: 使用中 / billing: 会計待ち';
COMMENT ON COLUMN order_items.status IS '注文ステータス（pending, preparing, served, cancelled）';
COMMENT ON COLUMN order_groups.status IS '注文グループの状態（open, closed, cancelled）';

-- KPI: 注文・収益指標（メニュー指向：注文数・総収益・平均注文額・販売アイテム数）
CREATE OR REPLACE VIEW v_kpi_orders_menu AS
SELECT
  (SELECT COUNT(*) FROM order_groups) AS total_orders,
  (SELECT COUNT(*) FROM order_groups WHERE created_at >= date_trunc('day', now())) AS orders_today,
  COALESCE((SELECT SUM(oi.price_at_order * oi.quantity) FROM order_items oi WHERE oi.status <> 'cancelled'), 0) AS total_revenue,
  CASE
    WHEN (SELECT COUNT(*) FROM order_groups) = 0 THEN 0
    ELSE COALESCE((SELECT SUM(oi.price_at_order * oi.quantity) FROM order_items oi WHERE oi.status <> 'cancelled'),0)::numeric
         / NULLIF((SELECT COUNT(*) FROM order_groups), 0)
  END AS average_order_value,
  COALESCE((SELECT SUM(oi.quantity) FROM order_items oi WHERE oi.status <> 'cancelled'), 0) AS total_items_sold;

-- トップメニュー（過去30日、数量順）
CREATE OR REPLACE VIEW v_top_menus_30d AS
SELECT
  m.menu_id,
  m.name,
  SUM(oi.quantity) AS quantity_sold,
  SUM(oi.price_at_order * oi.quantity) AS revenue
FROM order_items oi
JOIN menus m ON oi.menu_id = m.menu_id
JOIN order_groups og ON oi.orders_id = og.orders_id
WHERE oi.status <> 'cancelled'
  AND og.created_at >= now() - INTERVAL '30 days'
GROUP BY m.menu_id, m.name
ORDER BY quantity_sold DESC;

-- カテゴリ別売上（過去30日）
CREATE OR REPLACE VIEW v_sales_by_category_30d AS
SELECT
  c.category_id,
  c.name AS category_name,
  SUM(oi.quantity) AS total_quantity,
  SUM(oi.price_at_order * oi.quantity) AS total_revenue,
  CASE WHEN SUM(oi.quantity) = 0 THEN 0 ELSE SUM(oi.price_at_order * oi.quantity)::numeric / NULLIF(SUM(oi.quantity),0) END AS avg_price_per_item
FROM order_items oi
JOIN menus m ON oi.menu_id = m.menu_id
JOIN categories c ON m.category_id = c.category_id
JOIN order_groups og ON oi.orders_id = og.orders_id
WHERE oi.status <> 'cancelled'
  AND og.created_at >= now() - INTERVAL '30 days'
GROUP BY c.category_id, c.name
ORDER BY total_revenue DESC;

-- メニュー別パフォーマンス（累計）
CREATE OR REPLACE VIEW v_menu_performance AS
SELECT
  m.menu_id,
  m.name,
  COALESCE(SUM(oi.quantity),0) AS times_ordered,
  COALESCE(SUM(oi.price_at_order * oi.quantity),0) AS revenue,
  CASE WHEN SUM(oi.quantity) = 0 THEN 0 ELSE SUM(oi.price_at_order * oi.quantity)::numeric / NULLIF(SUM(oi.quantity),0) END AS avg_price_at_order,
  CASE WHEN COUNT(DISTINCT oi.orders_id) = 0 THEN 0 ELSE SUM(oi.quantity)::numeric / NULLIF(COUNT(DISTINCT oi.orders_id),0) END AS avg_quantity_per_order
FROM menus m
LEFT JOIN order_items oi ON m.menu_id = oi.menu_id AND oi.status <> 'cancelled'
GROUP BY m.menu_id, m.name
ORDER BY revenue DESC;

-- メニュー日次トレンド（過去30日：メニュー別日次売上）
CREATE OR REPLACE VIEW v_menu_daily_trend_30d AS
SELECT
  (og.created_at::date) AS day,
  m.menu_id,
  m.name,
  SUM(oi.quantity) AS quantity_sold,
  SUM(oi.price_at_order * oi.quantity) AS revenue
FROM order_groups og
JOIN order_items oi ON og.orders_id = oi.orders_id
JOIN menus m ON oi.menu_id = m.menu_id
WHERE og.created_at >= date_trunc('day', now() - INTERVAL '29 days')
  AND oi.status <> 'cancelled'
GROUP BY day, m.menu_id, m.name
ORDER BY day, quantity_sold DESC;

-- 日次トレンド（既存：過去30日：日ごとの注文数・収益）
CREATE OR REPLACE VIEW v_orders_daily_30d AS
SELECT
  (og.created_at::date) AS day,
  COUNT(DISTINCT og.orders_id) AS orders_count,
  SUM(oi.price_at_order * oi.quantity) FILTER (WHERE oi.status <> 'cancelled') AS revenue
FROM order_groups og
LEFT JOIN order_items oi ON og.orders_id = oi.orders_id
WHERE og.created_at >= date_trunc('day', now() - INTERVAL '29 days')
GROUP BY day
ORDER BY day;

-- 追加インデックス（メニュー／カテゴリ／注文集計向け）
CREATE INDEX IF NOT EXISTS idx_order_items_menu_id ON order_items(menu_id);
CREATE INDEX IF NOT EXISTS idx_order_items_created_at ON order_items(created_at);
CREATE INDEX IF NOT EXISTS idx_menus_category_id ON menus(category_id);
CREATE INDEX IF NOT EXISTS idx_menus_is_sold_out ON menus(is_sold_out);
CREATE INDEX IF NOT EXISTS idx_order_groups_created_at ON order_groups(created_at);
