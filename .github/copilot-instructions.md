# Pinu - GitHub Copilot Instructions

## Project Overview

Pinu is a QR Order System designed for single-operator small restaurant businesses. The system follows a Port & Adapter (Hexagonal) architecture with a clear separation of concerns.

**Tech Stack:**
- Backend: Go 1.22+ with Fiber framework
- Frontend: Next.js (App Router) with TypeScript, Shadcn UI, and Tailwind CSS
- Database: PostgreSQL 17.5
- Infrastructure: Docker + Task runner
- Monorepo: Turbo Repo for frontend

## Language Requirements

- **Code**: Write all code in English (variables, functions, comments in code)
- **Documentation/Comments**: Write documentation and commit messages in Japanese
- **User Interface**: Japanese text for UI elements

## Development Environment

### Setup Commands
```bash
task setup         # Initial setup (environment + dependencies)
task start         # Start dev environment (DB + Backend + Frontend)
task dev           # Start with adminer included
task stop          # Stop environment
task fresh         # Clean + setup + start
```

### Service URLs
- Backend API: http://localhost:8000
- Frontend: http://localhost:3000
- Adminer (DB UI): http://localhost:8080

## Architecture Principles

### Backend (Port & Adapter / Hexagonal)

**Layer Structure:**
```
Controllers (app/handlers/) 
  → UseCases (app/usecases/)
    → Repository Interfaces (app/domain/repositories/)
      → Repository Implementations (app/infrastructure/repositories/)
        → Database (PostgreSQL)
```

**Key Rules:**
- Use Dependency Inversion: Inner layers never depend on outer layers
- No service layer: Implement business logic directly in UseCases
- Pointer types: Services and repositories must use pointer types
- Interface types: Use `Interface` not `*Interface` (e.g., `repositories.IUserRepository`)
- Dependency Injection: Use DI container at `app/middleware/injection.go`

**Naming Conventions:**
- Public structs: Start with uppercase (e.g., `UserUsecaseImpl`)
- Private structs: Start with lowercase
- Interfaces: Prefix with `I` (e.g., `IUserRepository`)

### Frontend (Next.js)

**Structure:**
- `apps/client/`: Customer-facing application
- `apps/admin/`: Admin/staff application  
- `packages/`: Shared packages across apps

**Key Technologies:**
- TypeScript strict mode enabled
- Validation: Zod schemas
- Forms: Conform library
- Data fetching: SWR
- UI Components: Shadcn UI + Tailwind CSS

## Code Quality Standards

### General Principles
- Use meaningful variable and function names
- Write concise and specific comments
- Avoid magic numbers; define as named constants
- Follow DRY (Don't Repeat Yourself) principle
- Apply SOLID and KISS (Keep It Simple, Stupid) principles
- Emphasize type safety with type definitions and annotations
- Handle errors carefully, considering all failure cases
- Structure code for testability
- Follow Single Responsibility Principle
- Consider performance and scalability
- Be aware of security risks (SQL injection, XSS, etc.)
- Write clear documentation and type comments
- Minimize dependencies and aim for loose coupling
- Remove unused code, variables, and functions

### Security Requirements
- Never commit secrets or credentials
- Use `.env` files for configuration (see `.env.example`)
- Hash passwords with bcrypt
- Prevent SQL injection using ORM or prepared statements
- Validate and sanitize all user input
- Implement proper error handling without exposing sensitive data

## Testing & Quality

### Running Tests
```bash
task test          # Run all tests
task lint          # Lint all code
task build         # Build all components
task deploy:prep   # Build + test (pre-deployment)
```

### Backend Testing
```bash
cd backend && go test ./... -v
cd backend && golangci-lint run
```

### Frontend Testing
```bash
cd frontend && pnpm test
cd frontend && pnpm lint
cd frontend && pnpm format  # Biome formatter
```

## Database

**Schema:** See `.docs/database/er_diagram.md`

**Main Tables:**
- `tables`: Physical restaurant table management
- `table_sessions`: Customer sessions (QR scan to checkout)
- `order_groups`: Order groups within sessions
- `order_items`: Individual order items
- `menus`: Menu items
- `categories`: Menu categories
- `users`: Staff and admin users

**Operations:**
```bash
task db:start      # Start PostgreSQL
task db:stop       # Stop PostgreSQL  
task db:reset      # Reset database (delete all data)
task db:logs       # View PostgreSQL logs
```

## Commit Guidelines

Follow Conventional Commits format:

```
<type>(<scope>): <description>

[optional body]

[optional footer(s)]
```

**Types:**
- `feat`: New feature (✨)
- `fix`: Bug fix (🐛)
- `docs`: Documentation (📝)
- `style`: Code style (🎨)
- `refactor`: Refactoring (♻️)
- `perf`: Performance (⚡️)
- `test`: Tests (✅)
- `chore`: Build/tools (🔧)

**Breaking Changes:** Add `!` after type or use footer `BREAKING CHANGE:`

## Documentation References

- Architecture: `.docs/current-architecture.md`
- Architecture Decisions: `.docs/architecture-decisions.md`
- Technology Selection: `.docs/technology_selection.md`
- Commit Guidelines: `.docs/commit-guidelines.md`
- Requirements: `.docs/spec/requirements_definition.md`
- Features: `.docs/spec/feature_draft.md`
- ER Diagram: `.docs/database/er_diagram.md`

## Common Issues & Solutions

### Database Connection Error
```bash
docker ps
task db:logs
task db:stop && task db:start
```

### Port Conflicts
Default ports: 5432 (PostgreSQL), 8000 (Backend), 3000 (Frontend)
Override in `.env` file if needed.

### Dependency Errors
```bash
# Backend
cd backend && go mod tidy && go mod download

# Frontend  
cd frontend && rm -rf node_modules pnpm-lock.yaml && pnpm install
```

## Target Users

This system is designed for **elderly, single-operator small restaurant businesses**. Keep the UI/UX extremely simple and intuitive.
