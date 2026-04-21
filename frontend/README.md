# Frontend

`frontend/` は client (`/`) と admin (`/admin`) を内包する単一の Next.js アプリです。

## 開発

```bash
pnpm install
cp .env.sample .env.local
pnpm dev
```

- client: <http://localhost:3000>
- admin: <http://localhost:3000/admin>

## よく使うコマンド

```bash
pnpm dev
pnpm build
pnpm lint
pnpm format
```

## 構成

- `app/(client)` - client 向けルート
- `app/admin` - admin 向けルート
- `components`, `hooks`, `lib`, `store` - フロントエンド共通コード
- `packages/ui` - shared UI package
