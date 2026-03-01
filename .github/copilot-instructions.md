# Pinu - GitHub Copilot Instructions

このワークスペースでは、まず `AGENTS.md` と `.github/instructions/general.instructions.md` を優先して従ってください。

## 最優先ルール

- コード（識別子・コード内コメント）は英語で記述する。
- エージェント向け説明・ドキュメントは日本語で記述する。
- 実装は既存アーキテクチャ（DI + UseCase + Repository）に合わせ、推測で層構造を変更しない。

## 実行コマンド

- 標準入口は `task`（`setup/start/dev/stop`）。
- Backend: `cd backend && go test ./... -v`、`go run app/cmd/main.go`。
- Frontend: `cd frontend && pnpm dev`、`pnpm build`、`pnpm lint`、`pnpm format`。

## 重要な注意点

- Backend エントリポイントは `backend/app/cmd/main.go` を前提にする。
- Frontend root の `pnpm test` は未定義のため、`task test` は失敗可能性を考慮する。
- Windows の `task lint` は `backend:lint` 未有効の影響を受ける可能性がある。
- 環境変数名の差異（`SECRET_KEY` と `JWT_SECRET`）に注意し、認証変更時は `.env.example` と実装参照名を突合する。

## 参照

- 詳細な運用規約: `AGENTS.md`
- 全般コーディング規約: `.github/instructions/general.instructions.md`
