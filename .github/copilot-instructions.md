# Pinu - GitHub Copilot Instructions

このワークスペースでは、まず `AGENTS.md` と `.github/instructions/general.instructions.md` を優先して従ってください。

## 最優先ルール

- コード（識別子・コード内コメント）は英語で記述する。
- エージェント向け説明・ドキュメントは日本語で記述する。
- 実装は既存アーキテクチャ（DI + UseCase + Repository）に合わせ、推測で層構造を変更しない。
- 仕様・API 契約・業務フロー・DB 設計に関わる作業では、まず `docs/` 配下を参照して仮説を立て、その後に実装差分を照合する。
- docs と実装が不一致な場合は差分を明示し、推測で埋めず、実コードを正として docs 更新の要否も判断する。

## 実行コマンド

- 標準入口は `task`（`setup/start/dev/stop`）。
- Backend: `cd backend && go test ./... -v`、`go run app/cmd/main.go`。
- Frontend: `cd frontend && pnpm dev`、`pnpm build`、`pnpm lint`、`pnpm format`。

## 参照順

- 全体仕様: [docs/spec/reverse_engineered_system_spec.md](../docs/spec/reverse_engineered_system_spec.md)
- 要件確認: [docs/spec/requirements_definition.md](../docs/spec/requirements_definition.md)
- QR / table session / checkout フロー: [docs/spec/session_management_flow.md](../docs/spec/session_management_flow.md)
- DB 設計: [docs/database/README.md](../docs/database/README.md)、[docs/database/schema.dbml](../docs/database/schema.dbml)
- コミット規約: [docs/commit-guidelines.md](../docs/commit-guidelines.md)

## 重要な注意点

- Backend エントリポイントは `backend/app/cmd/main.go` を前提にする。
- Frontend root の `pnpm test` は未定義のため、`task test` は失敗可能性を考慮する。
- Windows の `task lint` は `backend:lint` 未有効の影響を受ける可能性がある。
- 環境変数名の差異（`SECRET_KEY` と `JWT_SECRET`）に注意し、認証変更時は `.env.example` と実装参照名を突合する。
- Frontend の lint / format は Biome を前提にする。
- DB スキーマ変更時は `docs/database/schema.dbml` と `database/init/init.sql` の同期を前提にする。
- ドキュメント参照は `docs/` を正とし、古い `.docs/` 参照は新規に追加しない。

## 参照

- 詳細な運用規約: `AGENTS.md`
- 全般コーディング規約: `.github/instructions/general.instructions.md`
