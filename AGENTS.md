# AGENTS.md

## Project

- **Name**: Pinu
- **Type**: QR Order System for small restaurants
- **Target**: Single-operator small restaurant businesses
- **Backend**: Go 1.22+ + Fiber
- **Frontend**: Next.js + TypeScript + Shadcn UI + Tailwind CSS
- **Database**: PostgreSQL 17.5
- **Infrastructure**: Docker + Task runner
- **Architecture**: Port & Adapter (Hexagonal)

## Setup

### Requirements

- Docker
- Task (`sh -c "$(curl --location https://taskfile.dev/install.sh)" -- -d`)
- Go 1.22+
- Node.js 18+ + pnpm

### Commands

```bash
task install:task  # Install Task runner
task setup         # Setup environment and dependencies
task start         # Start dev environment (DB + Backend + Frontend)
task dev           # Start with adminer
task stop          # Stop environment
task db:reset      # Reset database
task clean         # Full cleanup
task fresh         # Clean + setup + start
```

### URLs

- Backend: http://localhost:8000
- Frontend: http://localhost:3000
- Adminer: http://localhost:8080

## Build/Test/Lint

### All

```bash
task build       # Build all
task test        # Test all
task lint        # Lint all
task deploy:prep # Build + test
```

### Backend (Go)

```bash
cd backend && go mod tidy && go mod download
cd backend && go run cmd/server/main.go
cd backend && go build -o ../bin/backend cmd/server/main.go
cd backend && go test ./... -v
cd backend && golangci-lint run
```

### Frontend (Next.js + TypeScript)

```bash
cd frontend && pnpm install
cd frontend && pnpm dev
cd frontend && pnpm build
cd frontend && pnpm test
cd frontend && pnpm lint
cd frontend && pnpm format  # Biome
```

## Architecture

### Backend: Port & Adapter (Hexagonal)

```
Controllers (Primary Adapter)       → app/handlers/
    ↓
UseCases (Application Layer)        → app/usecases/
    ↓
Repository Interface (Domain Port)  → app/domain/repositories/
    ↓
Repository Implementation           → app/infrastructure/repositories/
    ↓
Database (PostgreSQL)
```

### Layer Responsibilities

**Controllers** (`app/handlers/`)
- HTTP request/response handling
- UseCase invocation
- Presentation layer

**UseCases** (`app/usecases/`)
- Business logic
- Validation
- Application logic
- Input/Output DTO usage

**Domain** (`app/domain/`)
- Entity definitions (`entities/`)
- Repository interfaces (`repositories/`)

**Infrastructure** (`app/infrastructure/`)
- Repository implementations (`repositories/`)
- Database operations

### Principles

- Dependency Inversion: Inner layers do not depend on outer layers
- No service layer: Direct UseCase implementation
- Pointer types: Services and repositories use pointer types
- Interface types: Use `Interface` not `*Interface`

### Frontend

- Next.js App Router
- Turbo Repo monorepo
- `apps/`: client, admin applications
- `packages/`: shared packages

## Code Conventions

Reference: `.github/instructions/general.instructions.md`

### General

- Meaningful variable/function names
- Concise, specific comments
- Constants for magic numbers
- DRY principle
- SOLID and KISS principles
- Type safety: use type definitions and annotations
- Error handling: handle exceptions and failure cases
- Single Responsibility Principle
- Security: prevent injection, XSS

### Go

- Struct names: Public structs start with uppercase (e.g., `UserUsecaseImpl`)
- Interface types: No pointers (e.g., `repositories.IUserRepository`)
- Dependency Injection: Use DI container (`app/middleware/injection.go`)

### TypeScript/Next.js

- TypeScript strict mode
- Validation: Zod
- Forms: Conform
- API: SWR
- UI: Shadcn UI + Tailwind CSS

## Database

### Connection (Default)

- Database: `pinu`
- User: `pinu_user`
- Password: `pinu_pass`
- Port: `5432`
- Config: `.env` (see `.env.example`)

### Operations

```bash
task db:start  # Start PostgreSQL
task db:stop   # Stop PostgreSQL
task db:reset  # Reset (delete all data)
task db:logs   # Show logs
```

### Schema

Reference: `.docs/database/er_diagram.md`

Main tables:
- `tables`: Physical table management
- `table_sessions`: Session (QR entry to checkout)
- `order_groups`: Order groups per session
- `order_items`: Order items
- `menus`: Menu items
- `categories`: Menu categories
- `users`: Staff/admins

## Commits/PRs

### Format: Conventional Commits

Reference: `.docs/commit-guidelines.md`

```
<type>(<scope>): <description>

[optional body]

[optional footer(s)]
```

### Types

| Emoji | Type     | Description                |
|-------|----------|----------------------------|
| ✨    | feat     | New feature                |
| 🐛    | fix      | Bug fix                    |
| 📝    | docs     | Documentation              |
| 🎨    | style    | Code style                 |
| ♻️    | refactor | Refactoring                |
| ⚡️    | perf     | Performance                |
| ✅    | test     | Tests                      |
| 🔧    | chore    | Build/tools                |

### Breaking Changes

Add `!` or footer `BREAKING CHANGE:`

```
feat(api)!: drop support for legacy v1 endpoints

BREAKING CHANGE: API clients must now use the v2 endpoints.
```

### PR Rules

- Follow Conventional Commits
- Tests pass
- No lint errors
- Reference issue (Closes: #XX)

## Project Structure

```
.
├── backend/
│   ├── app/
│   │   ├── handlers/        # Controllers (HTTP handlers)
│   │   ├── usecases/        # Application Layer
│   │   ├── domain/          # Domain Layer
│   │   │   ├── entities/
│   │   │   └── repositories/
│   │   ├── infrastructure/  # Infrastructure Layer
│   │   │   └── repositories/
│   │   └── middleware/      # DI, auth
│   ├── cmd/server/          # Entry point
│   ├── go.mod
│   └── go.sum
├── frontend/                # Next.js (Turbo Repo)
│   ├── apps/
│   │   ├── client/          # Client app
│   │   └── admin/           # Admin app
│   ├── packages/            # Shared packages
│   ├── package.json
│   └── turbo.json
├── database/
│   ├── init/                # Init SQL
│   └── postgresql.conf
├── .docs/
│   ├── current-architecture.md
│   ├── architecture-decisions.md
│   ├── technology_selection.md
│   ├── commit-guidelines.md
│   ├── spec/
│   └── database/
├── docker/
├── compose.yml
├── Taskfile.yml
├── .env                     # Excluded from git
└── .env.example
```

## Documentation

- Architecture: `.docs/current-architecture.md`
- Architecture decisions: `.docs/architecture-decisions.md`
- Technology selection: `.docs/technology_selection.md`
- Commit guidelines: `.docs/commit-guidelines.md`
- Requirements: `.docs/spec/requirements_definition.md`
- Features: `.docs/spec/feature_draft.md`
- ER diagram: `.docs/database/er_diagram.md`

## Security

- No secrets in commits
- `.env` excluded via `.gitignore`
- Create `.env` from `.env.example`
- Password hashing: bcrypt
- SQL injection prevention: ORM or prepared statements

## Troubleshooting

### Database connection error

```bash
docker ps
task db:logs
task db:stop
task db:start
```

### Port conflicts

Default ports:
- `5432`: PostgreSQL
- `8000`: Backend
- `3000`: Frontend

Change in `.env`:
```
DB_PORT=5433
BACKEND_PORT=8001
FRONTEND_PORT=3001
```

### Dependency errors

```bash
# Backend
cd backend && go mod tidy && go mod download

# Frontend
cd frontend && rm -rf node_modules pnpm-lock.yaml && pnpm install
```

## Constraints

- License: MIT
- Code language: English
- Documentation/comments: Japanese
- UI/UX: Simple design for elderly operators
