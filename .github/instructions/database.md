---
applyTo: 
  - "**/database/**"
  - "**/migrations/**"
  - "**/repositories/**"
  - "**/*repository*.go"
  - "**/*_test.go"
---

# Database Instructions for Copilot

## Database System

- **DBMS**: PostgreSQL 17.5
- **ORM**: Not specified (use database/sql or equivalent)
- **Connection**: Managed via `.env` configuration
- **Schema**: See `.docs/database/er_diagram.md`

## Connection Configuration

Default settings:
```
DB_HOST=localhost
DB_PORT=5432
DB_NAME=pinu
DB_USER=pinu_user
DB_PASSWORD=pinu_pass
```

## Schema Design Principles

### Table Structure

**Core Tables:**
1. `tables` - Physical restaurant tables
   - Manage table numbers, seating capacity, status
   
2. `table_sessions` - Customer sessions
   - Track from QR code scan to checkout
   - One session per table at a time
   
3. `order_groups` - Order batches within a session
   - Multiple order groups per session
   - Used for tracking when orders were placed
   
4. `order_items` - Individual ordered items
   - Links to menus and order_groups
   - Tracks quantity, price at time of order
   
5. `menus` - Menu items
   - Product information, pricing, availability
   - Links to categories
   
6. `categories` - Menu categories
   - Organizational structure for menus
   
7. `users` - Staff and admin accounts
   - Authentication and authorization

### Naming Conventions

- **Tables**: Plural snake_case (e.g., `table_sessions`, `order_items`)
- **Columns**: snake_case (e.g., `created_at`, `user_id`)
- **Foreign Keys**: `{table}_id` format (e.g., `table_id`, `menu_id`)
- **Timestamps**: Always include `created_at`, `updated_at` where appropriate
- **Soft Deletes**: Use `deleted_at` column (nullable timestamp)

## Repository Pattern

When working with database code in Go:

### Interface Definition (Domain Layer)

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

### Implementation (Infrastructure Layer)

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
    // Implementation
}
```

### Key Rules

1. **Interfaces in Domain**: Define repository interfaces in `app/domain/repositories/`
2. **Implementations in Infrastructure**: Implement in `app/infrastructure/repositories/`
3. **Pointer Returns**: Repository methods return pointer types (e.g., `*entities.User`)
4. **Error Handling**: Always return errors, never panic
5. **Transactions**: Support transaction context when needed
6. **Prepared Statements**: Use prepared statements to prevent SQL injection

## Database Operations Best Practices

### Security

- **SQL Injection Prevention**: Always use parameterized queries
  ```go
  // Good
  row := db.QueryRow("SELECT * FROM users WHERE email = $1", email)
  
  // Bad - SQL injection vulnerability!
  query := fmt.Sprintf("SELECT * FROM users WHERE email = '%s'", email)
  ```

- **Password Storage**: Use bcrypt for hashing
  ```go
  hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
  ```

### Query Optimization

- Use indexes for frequently queried columns
- Avoid SELECT * when possible, specify needed columns
- Use JOINs efficiently
- Consider pagination for large result sets
- Use database connection pooling

### Error Handling

```go
if err != nil {
    if errors.Is(err, sql.ErrNoRows) {
        // Handle not found case
        return nil, fmt.Errorf("user not found")
    }
    // Handle other errors
    return nil, fmt.Errorf("database error: %w", err)
}
```

## Testing Database Code

### Unit Tests

- Mock repository interfaces for testing UseCases
- Don't test database in UseCase tests
- Test repository implementations separately

### Integration Tests

- Use test database or Docker container
- Clean up test data after each test
- Test actual database interactions

```go
func TestUserRepository_FindByID(t *testing.T) {
    // Setup test database
    db := setupTestDB(t)
    defer db.Close()
    
    repo := NewUserRepository(db)
    
    // Test implementation
    user, err := repo.FindByID(1)
    assert.NoError(t, err)
    assert.NotNil(t, user)
}
```

## Migrations

When creating migrations:

1. **File Naming**: Use timestamp prefix (e.g., `20240101120000_create_users_table.sql`)
2. **Up/Down**: Provide both up and down migrations
3. **Idempotent**: Migrations should be safe to run multiple times
4. **Data Preservation**: Never lose user data
5. **Backward Compatible**: Consider backwards compatibility

## Performance Considerations

- **Indexes**: Add indexes for foreign keys and frequently queried columns
- **Connection Pooling**: Configure appropriate pool size
- **Query Optimization**: Use EXPLAIN ANALYZE to understand query performance
- **Caching**: Consider caching for frequently accessed, rarely changed data
- **Batch Operations**: Use batch inserts/updates when processing multiple records

## Task Commands

```bash
task db:start      # Start PostgreSQL container
task db:stop       # Stop PostgreSQL container
task db:reset      # Reset database (WARNING: deletes all data)
task db:logs       # View PostgreSQL logs
```

## Common Issues

### Connection Refused
- Check if Docker container is running: `docker ps`
- Verify port 5432 is not in use
- Check `.env` configuration

### Permission Denied
- Verify database user has correct permissions
- Check database and schema ownership

### Migration Errors
- Review migration SQL syntax
- Check for constraint violations
- Verify foreign key relationships exist
