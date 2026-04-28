# Project Guidelines

## Code Style

- まず `.github/instructions/general.instructions.md` を最優先で適用し、重複ルールはこのファイルに再記載しない。
- コードは英語（識別子・コード内コメント）、エージェント向け説明文は日本語で記述する。
- Frontend は Biome で lint/format を実行する（root と各 app の `package.json` スクリプトに準拠）。
- Go の domain port は `*Store` / `UnitOfWork` 命名が実装実態。`I*` 接頭辞前提で新規命名しない。

## Architecture

- Backend の依存解決は DI コンテナ（`do`）で一元管理し、ルート組み立て時に Handler を注入する。
- HTTP は `routes -> handlers -> usecases -> repositories` の流れで接続する。
- UseCase は `Store`、`UnitOfWork`、`Logger`（必要なら token maker）を受け取り、業務ロジックを担当する。
- トランザクション境界は `UnitOfWork.Run` を利用し、複数更新をまたぐ処理をまとめる。

## Build and Test

- 標準実行は `task` を入口にする（OS 別 Taskfile に委譲される）。
- 初期化/起動: `task setup`、`task start`、`task dev`、停止は `task stop`。
- Backend: `cd backend && go test ./... -v`、`go build`、`go run app/cmd/main.go`。
- Frontend: `cd frontend && pnpm dev`、`pnpm build`、`pnpm lint`、`pnpm format`。
- 注意: 現状 `frontend` root に `pnpm test` スクリプトは未定義。`task test` 実行時は失敗可能性を前提に確認する。
- 注意: Windows Taskfile では `backend:lint` が未有効。`task lint` 実行時は事前に Taskfile 定義を確認する。

## Project Conventions

- モノレポ運用では「編集対象に最も近い指示ファイル」を優先する。
- 仕様・API 契約・業務フロー・DB 設計に関わる作業では、まず `docs/` 配下の正本を確認して仮説を立て、その後に実装を照合する。
- docs と実装が不一致でも、推測でどちらかに合わせず差分を明示する。変更時は「実コードを正」として docs も更新する。
- 実装・ドキュメントの不整合がある場合、推測で合わせず実コードを正として更新する。
- backend エントリポイントは `app/cmd/main.go` 前提で扱い、`cmd/server/main.go` 前提の記述を増やさない。
- ドキュメント参照は `docs/` を正とし、古い `.docs/` 参照は新規に追加しない。

## Documentation First

- 仕様全体の起点は [docs/spec/reverse_engineered_system_spec.md](docs/spec/reverse_engineered_system_spec.md) を優先する。
- 要件粒度の確認は [docs/spec/requirements_definition.md](docs/spec/requirements_definition.md) を参照する。
- QR / table session / checkout の流れは [docs/spec/session_management_flow.md](docs/spec/session_management_flow.md) を参照する。
- DB 設計は [docs/database/README.md](docs/database/README.md) と [docs/database/schema.dbml](docs/database/schema.dbml) を原本として扱う。
- DB スキーマ変更時は `schema.dbml` と `database/init/init.sql` を同期し、生成物の `schema.sql` は差分確認用として扱う。
- コミット規約が必要な場合は [docs/commit-guidelines.md](docs/commit-guidelines.md) を参照する。

## Integration Points

- API は `/api` 配下で提供され、auth/user/menu/order/analytics 系エンドポイントを Fiber ルートで公開する。
- SSE は API グループに統合されるため、関連変更時はルーティングとイベント配信の両方を確認する。
- Frontend は `frontend/` 直下の単一 Next.js とし、client は `/`、admin は `/admin`、共有要素は workspace package で再利用する。
- DB は PostgreSQL + Bun を利用。Repository 実装でクエリ責務を持たせる。

## Security

- 認証情報は `.env` 経由で管理し、シークレットをコミットしない。
- パスワードは `bcrypt + salt + PEPPER` の既存方式に合わせる。
- JWT は HS256 を使用し、署名方式検証を維持する。
- SQL は Bun のプレースホルダ/クエリビルダを使い、文字列連結での動的クエリ生成を避ける。
- 環境変数名（例: `SECRET_KEY` と `JWT_SECRET`）に差異があるため、認証周辺変更時は両 `.env.example` と実装参照名を必ず突合する。
