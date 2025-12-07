---
applyTo:
  - "backend/**"
  - "**/go.mod"
  - "**/go.sum"
---

# Backend (Go + Fiber) Instructions for Copilot

## Technology Stack

- **Language**: Go 1.22+
- **Web Framework**: Fiber (Express-inspired framework)
- **Database**: PostgreSQL 17.5
- **Architecture**: Port & Adapter (Hexagonal Architecture)

## Architecture Layers

```
Controllers (app/handlers/)
    ↓
UseCases (app/usecases/)
    ↓
Repository Interfaces (app/domain/repositories/)
    ↓
Repository Implementations (app/infrastructure/repositories/)
    ↓
Database (PostgreSQL)
```

### Layer Responsibilities

**1. Controllers (`app/handlers/`)**
- HTTP request/response handling
- Request validation and parsing
- Response formatting (JSON)
- HTTP status code management
- UseCase invocation
- Error handling and formatting

Example:
```go
func (h *UserHandler) GetUser(c *fiber.Ctx) error {
    id, err := c.ParamsInt("id")
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid user ID",
        })
    }
    
    user, err := h.userUsecase.GetUserByID(id)
    if err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": err.Error(),
        })
    }
    
    return c.JSON(user)
}
```

**2. UseCases (`app/usecases/`)**
- Business logic implementation
- Input validation
- Data transformation
- Application logic coordination
- Error handling

Example:
```go
type UserUsecaseImpl struct {
    userRepo repositories.IUserRepository
}

func NewUserUsecase(userRepo repositories.IUserRepository) *UserUsecaseImpl {
    return &UserUsecaseImpl{
        userRepo: userRepo,
    }
}

func (u *UserUsecaseImpl) GetUserByID(id int) (*entities.User, error) {
    if id <= 0 {
        return nil, fmt.Errorf("invalid user ID")
    }
    
    user, err := u.userRepo.FindByID(id)
    if err != nil {
        return nil, fmt.Errorf("failed to get user: %w", err)
    }
    
    return user, nil
}
```

**3. Domain (`app/domain/`)**
- Entity definitions (`entities/`)
- Repository interfaces (`repositories/`)
- Business rules and constraints
- No external dependencies

Example Entity:
```go
// app/domain/entities/user.go
package entities

import "time"

type User struct {
    ID        int       `json:"id"`
    Email     string    `json:"email"`
    Name      string    `json:"name"`
    Password  string    `json:"-"` // Never expose in JSON
    Role      string    `json:"role"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
```

Example Repository Interface:
```go
// app/domain/repositories/user_repository.go
package repositories

import "github.com/yourusername/pinu/backend/app/domain/entities"

type IUserRepository interface {
    FindByID(id int) (*entities.User, error)
    FindByEmail(email string) (*entities.User, error)
    Create(user *entities.User) error
    Update(user *entities.User) error
    Delete(id int) error
}
```

**4. Infrastructure (`app/infrastructure/`)**
- Repository implementations (`repositories/`)
- Database operations
- External service integrations
- File system operations

Example:
```go
// app/infrastructure/repositories/user_repository.go
package repositories

import (
    "database/sql"
    "github.com/yourusername/pinu/backend/app/domain/entities"
    "github.com/yourusername/pinu/backend/app/domain/repositories"
)

type UserRepository struct {
    db *sql.DB
}

func NewUserRepository(db *sql.DB) repositories.IUserRepository {
    return &UserRepository{db: db}
}

func (r *UserRepository) FindByID(id int) (*entities.User, error) {
    var user entities.User
    err := r.db.QueryRow(
        "SELECT id, email, name, role, created_at, updated_at FROM users WHERE id = $1",
        id,
    ).Scan(&user.ID, &user.Email, &user.Name, &user.Role, &user.CreatedAt, &user.UpdatedAt)
    
    if err != nil {
        return nil, err
    }
    
    return &user, nil
}
```

## Naming Conventions

### Structs and Types
- **Public structs**: Start with uppercase (e.g., `UserUsecaseImpl`, `UserHandler`)
- **Private structs**: Start with lowercase (e.g., `userRequest`, `errorResponse`)
- **Implementation structs**: Suffix with `Impl` (e.g., `UserUsecaseImpl`)
- **Interfaces**: Prefix with `I` (e.g., `IUserRepository`, `IUserUsecase`)

### Functions and Methods
- Use camelCase for unexported functions
- Use PascalCase for exported functions
- Be descriptive: `GetUserByID`, not `GetUser`

### Variables
- Short names for short scopes: `i`, `err`, `ok`
- Descriptive names for larger scopes: `userRepository`, `orderUsecase`
- Constants: ALL_CAPS with underscores: `MAX_RETRY_COUNT`

## Dependency Injection

Use the DI container located at `app/middleware/injection.go`:

```go
// app/middleware/injection.go
package middleware

import (
    "database/sql"
    "github.com/yourusername/pinu/backend/app/handlers"
    "github.com/yourusername/pinu/backend/app/usecases"
    "github.com/yourusername/pinu/backend/app/infrastructure/repositories"
)

type Container struct {
    UserHandler *handlers.UserHandler
    // Add other handlers
}

func NewContainer(db *sql.DB) *Container {
    // Repository layer
    userRepo := repositories.NewUserRepository(db)
    
    // UseCase layer
    userUsecase := usecases.NewUserUsecase(userRepo)
    
    // Handler layer
    userHandler := handlers.NewUserHandler(userUsecase)
    
    return &Container{
        UserHandler: userHandler,
    }
}
```

## Error Handling

### Best Practices
- Always check errors immediately
- Wrap errors with context using `fmt.Errorf` with `%w`
- Return errors, don't log and ignore
- Use custom error types for business logic errors
- Never panic in production code

Example:
```go
func (u *UserUsecaseImpl) CreateUser(input CreateUserInput) (*entities.User, error) {
    // Validation
    if input.Email == "" {
        return nil, fmt.Errorf("email is required")
    }
    
    // Check if user exists
    existingUser, err := u.userRepo.FindByEmail(input.Email)
    if err != nil && !errors.Is(err, sql.ErrNoRows) {
        return nil, fmt.Errorf("failed to check existing user: %w", err)
    }
    if existingUser != nil {
        return nil, fmt.Errorf("user with email already exists")
    }
    
    // Hash password
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
    if err != nil {
        return nil, fmt.Errorf("failed to hash password: %w", err)
    }
    
    // Create user
    user := &entities.User{
        Email:    input.Email,
        Name:     input.Name,
        Password: string(hashedPassword),
        Role:     "user",
    }
    
    if err := u.userRepo.Create(user); err != nil {
        return nil, fmt.Errorf("failed to create user: %w", err)
    }
    
    return user, nil
}
```

## Fiber Framework Specifics

### Routing
```go
func SetupRoutes(app *fiber.App, container *middleware.Container) {
    api := app.Group("/api")
    v1 := api.Group("/v1")
    
    // User routes
    users := v1.Group("/users")
    users.Get("/", container.UserHandler.GetUsers)
    users.Get("/:id", container.UserHandler.GetUser)
    users.Post("/", container.UserHandler.CreateUser)
    users.Put("/:id", container.UserHandler.UpdateUser)
    users.Delete("/:id", container.UserHandler.DeleteUser)
}
```

### Middleware
```go
// CORS
app.Use(cors.New(cors.Config{
    AllowOrigins: "http://localhost:3000",
    AllowMethods: "GET,POST,PUT,DELETE",
    AllowHeaders: "Origin, Content-Type, Accept, Authorization",
}))

// Logger
app.Use(logger.New())

// Recovery from panics
app.Use(recover.New())
```

## Testing

### Unit Tests
- Test each layer independently
- Mock dependencies using interfaces
- Use table-driven tests for multiple cases
- Test error cases

Example:
```go
func TestUserUsecase_GetUserByID(t *testing.T) {
    tests := []struct {
        name    string
        id      int
        want    *entities.User
        wantErr bool
    }{
        {
            name: "valid user",
            id:   1,
            want: &entities.User{ID: 1, Email: "test@example.com"},
            wantErr: false,
        },
        {
            name: "invalid id",
            id:   -1,
            want: nil,
            wantErr: true,
        },
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // Setup mock repository
            mockRepo := &MockUserRepository{}
            usecase := usecases.NewUserUsecase(mockRepo)
            
            // Test
            got, err := usecase.GetUserByID(tt.id)
            if (err != nil) != tt.wantErr {
                t.Errorf("GetUserByID() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if !reflect.DeepEqual(got, tt.want) {
                t.Errorf("GetUserByID() = %v, want %v", got, tt.want)
            }
        })
    }
}
```

## Security Best Practices

1. **Password Hashing**: Always use bcrypt
2. **SQL Injection**: Use parameterized queries (e.g., `$1`, `$2`)
3. **Input Validation**: Validate all user input in UseCases
4. **Authentication**: Implement JWT or session-based auth
5. **Authorization**: Check user permissions before operations
6. **CORS**: Configure allowed origins properly
7. **Rate Limiting**: Implement rate limiting for APIs
8. **Sensitive Data**: Never log passwords or tokens
9. **Error Messages**: Don't expose internal details to clients

## Performance Considerations

- Use connection pooling for database
- Implement pagination for list endpoints
- Cache frequently accessed data
- Use indexes on database queries
- Profile code with pprof when needed
- Consider using goroutines for concurrent operations (carefully!)

## Development Commands

```bash
cd backend

# Install dependencies
go mod tidy && go mod download

# Run server (with hot reload via Air)
go run cmd/server/main.go

# Build binary
go build -o ../bin/backend cmd/server/main.go

# Run tests
go test ./... -v

# Run specific test
go test ./app/usecases -v -run TestUserUsecase

# Lint code
golangci-lint run

# Format code
gofmt -w .
go mod tidy
```

## Common Patterns

### DTO (Data Transfer Object)
Use separate structs for input/output:

```go
type CreateUserInput struct {
    Email    string `json:"email" validate:"required,email"`
    Name     string `json:"name" validate:"required"`
    Password string `json:"password" validate:"required,min=8"`
}

type UserResponse struct {
    ID        int       `json:"id"`
    Email     string    `json:"email"`
    Name      string    `json:"name"`
    Role      string    `json:"role"`
    CreatedAt time.Time `json:"created_at"`
}
```

### Validation
Use a validation library or implement custom validation:

```go
func (i *CreateUserInput) Validate() error {
    if i.Email == "" {
        return fmt.Errorf("email is required")
    }
    if !isValidEmail(i.Email) {
        return fmt.Errorf("invalid email format")
    }
    if len(i.Password) < 8 {
        return fmt.Errorf("password must be at least 8 characters")
    }
    return nil
}
```

## Key Principles

1. **Dependency Inversion**: Inner layers don't depend on outer layers
2. **No Service Layer**: Business logic goes directly in UseCases
3. **Pointer Types**: Services and repositories use pointer types
4. **Interface Types**: Use `IRepository` not `*IRepository`
5. **Single Responsibility**: Each function/struct has one clear purpose
6. **Fail Fast**: Validate input early, return errors immediately
7. **Testability**: Write code that's easy to test with mocks
