---
applyTo:
  - "**/*_test.go"
  - "**/*.test.ts"
  - "**/*.test.tsx"
  - "**/*.spec.ts"
  - "**/*.spec.tsx"
---

# Testing Instructions for Copilot

## Overview

Testing is essential for ensuring code quality and preventing regressions. This project uses different testing approaches for backend (Go) and frontend (TypeScript/React).

## Backend Testing (Go)

### Test Structure

```
backend/
├── app/
│   ├── handlers/
│   │   └── user_handler_test.go
│   ├── usecases/
│   │   └── user_usecase_test.go
│   └── infrastructure/
│       └── repositories/
│           └── user_repository_test.go
```

### Testing Framework

- **Testing Package**: Standard `testing` package
- **Assertions**: Use `testify/assert` or standard comparisons
- **Mocking**: Interface-based mocking

### Unit Test Pattern

```go
package usecases

import (
    "testing"
    "errors"
    "github.com/yourusername/pinu/backend/app/domain/entities"
    "github.com/yourusername/pinu/backend/app/domain/repositories"
)

// Mock repository
type MockUserRepository struct {
    FindByIDFunc func(id int) (*entities.User, error)
}

func (m *MockUserRepository) FindByID(id int) (*entities.User, error) {
    if m.FindByIDFunc != nil {
        return m.FindByIDFunc(id)
    }
    return nil, errors.New("not implemented")
}

func (m *MockUserRepository) FindByEmail(email string) (*entities.User, error) {
    return nil, errors.New("not implemented")
}

func (m *MockUserRepository) Create(user *entities.User) error {
    return errors.New("not implemented")
}

func (m *MockUserRepository) Update(user *entities.User) error {
    return errors.New("not implemented")
}

func (m *MockUserRepository) Delete(id int) error {
    return errors.New("not implemented")
}

// Table-driven tests
func TestUserUsecase_GetUserByID(t *testing.T) {
    tests := []struct {
        name    string
        id      int
        mockFn  func(id int) (*entities.User, error)
        want    *entities.User
        wantErr bool
    }{
        {
            name: "success - user found",
            id:   1,
            mockFn: func(id int) (*entities.User, error) {
                return &entities.User{
                    ID:    1,
                    Email: "test@example.com",
                    Name:  "Test User",
                }, nil
            },
            want: &entities.User{
                ID:    1,
                Email: "test@example.com",
                Name:  "Test User",
            },
            wantErr: false,
        },
        {
            name: "error - invalid id",
            id:   -1,
            mockFn: func(id int) (*entities.User, error) {
                return nil, nil
            },
            want:    nil,
            wantErr: true,
        },
        {
            name: "error - user not found",
            id:   999,
            mockFn: func(id int) (*entities.User, error) {
                return nil, errors.New("user not found")
            },
            want:    nil,
            wantErr: true,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // Setup
            mockRepo := &MockUserRepository{
                FindByIDFunc: tt.mockFn,
            }
            usecase := NewUserUsecase(mockRepo)

            // Execute
            got, err := usecase.GetUserByID(tt.id)

            // Assert
            if (err != nil) != tt.wantErr {
                t.Errorf("GetUserByID() error = %v, wantErr %v", err, tt.wantErr)
                return
            }

            if tt.want != nil && got != nil {
                if got.ID != tt.want.ID || got.Email != tt.want.Email {
                    t.Errorf("GetUserByID() = %v, want %v", got, tt.want)
                }
            }
        })
    }
}
```

### Integration Tests (Repository Layer)

```go
func TestUserRepository_FindByID_Integration(t *testing.T) {
    // Skip if not in integration test mode
    if testing.Short() {
        t.Skip("Skipping integration test")
    }

    // Setup test database
    db := setupTestDB(t)
    defer db.Close()

    repo := NewUserRepository(db)

    // Insert test data
    testUser := &entities.User{
        Email: "test@example.com",
        Name:  "Test User",
    }
    err := repo.Create(testUser)
    if err != nil {
        t.Fatalf("Failed to create test user: %v", err)
    }

    // Test FindByID
    found, err := repo.FindByID(testUser.ID)
    if err != nil {
        t.Errorf("FindByID() error = %v", err)
    }
    if found.Email != testUser.Email {
        t.Errorf("FindByID() email = %v, want %v", found.Email, testUser.Email)
    }
}

func setupTestDB(t *testing.T) *sql.DB {
    // Setup test database connection
    // Use Docker container or in-memory database
    db, err := sql.Open("postgres", "postgresql://test:test@localhost:5432/test_db")
    if err != nil {
        t.Fatalf("Failed to connect to test database: %v", err)
    }

    // Run migrations or setup schema
    // Clean up existing data
    _, err = db.Exec("TRUNCATE TABLE users RESTART IDENTITY CASCADE")
    if err != nil {
        t.Fatalf("Failed to clean test database: %v", err)
    }

    return db
}
```

### Handler Tests

```go
func TestUserHandler_GetUser(t *testing.T) {
    // Setup Fiber app
    app := fiber.New()

    // Create mock usecase
    mockUsecase := &MockUserUsecase{
        GetUserByIDFunc: func(id int) (*entities.User, error) {
            if id == 1 {
                return &entities.User{
                    ID:    1,
                    Email: "test@example.com",
                    Name:  "Test User",
                }, nil
            }
            return nil, errors.New("user not found")
        },
    }

    handler := NewUserHandler(mockUsecase)
    app.Get("/users/:id", handler.GetUser)

    // Test success case
    req := httptest.NewRequest("GET", "/users/1", nil)
    resp, err := app.Test(req)
    if err != nil {
        t.Fatalf("Request failed: %v", err)
    }

    if resp.StatusCode != fiber.StatusOK {
        t.Errorf("Expected status 200, got %d", resp.StatusCode)
    }

    // Test not found case
    req = httptest.NewRequest("GET", "/users/999", nil)
    resp, err = app.Test(req)
    if err != nil {
        t.Fatalf("Request failed: %v", err)
    }

    if resp.StatusCode != fiber.StatusNotFound {
        t.Errorf("Expected status 404, got %d", resp.StatusCode)
    }
}
```

### Running Backend Tests

```bash
# Run all tests
cd backend && go test ./... -v

# Run tests with coverage
cd backend && go test ./... -coverprofile=coverage.out
cd backend && go tool cover -html=coverage.out

# Run specific package tests
cd backend && go test ./app/usecases -v

# Run specific test
cd backend && go test ./app/usecases -v -run TestUserUsecase_GetUserByID

# Run integration tests
cd backend && go test ./... -v -tags=integration

# Skip integration tests (run only unit tests)
cd backend && go test ./... -v -short
```

## Frontend Testing (TypeScript/React)

### Test Structure

```
frontend/
├── apps/
│   ├── client/
│   │   ├── components/
│   │   │   └── menu-item.test.tsx
│   │   ├── lib/
│   │   │   └── utils.test.ts
│   │   └── app/
│   │       └── api/
│   │           └── orders/
│   │               └── route.test.ts
```

### Testing Libraries

- **Test Runner**: Jest or Vitest
- **React Testing**: @testing-library/react
- **User Interactions**: @testing-library/user-event
- **Assertions**: expect (built-in)

### Component Testing

```typescript
// components/menu/menu-item.test.tsx
import { render, screen, fireEvent } from '@testing-library/react';
import { MenuItem } from './menu-item';
import { Menu } from '@/types/menu';

describe('MenuItem', () => {
  const mockItem: Menu = {
    id: 1,
    name: 'カレーライス',
    description: '当店自慢のカレー',
    price: 800,
    category_id: 1,
    is_available: true,
  };

  it('renders menu item with correct information', () => {
    render(<MenuItem item={mockItem} />);

    expect(screen.getByText('カレーライス')).toBeInTheDocument();
    expect(screen.getByText('当店自慢のカレー')).toBeInTheDocument();
    expect(screen.getByText('¥800')).toBeInTheDocument();
  });

  it('calls onSelect when clicked', () => {
    const mockOnSelect = jest.fn();
    render(<MenuItem item={mockItem} onSelect={mockOnSelect} />);

    const button = screen.getByRole('button', { name: /選択/i });
    fireEvent.click(button);

    expect(mockOnSelect).toHaveBeenCalledWith(mockItem);
    expect(mockOnSelect).toHaveBeenCalledTimes(1);
  });

  it('does not show select button when onSelect is not provided', () => {
    render(<MenuItem item={mockItem} />);

    const button = screen.queryByRole('button', { name: /選択/i });
    expect(button).not.toBeInTheDocument();
  });

  it('applies custom className', () => {
    const { container } = render(
      <MenuItem item={mockItem} className="custom-class" />
    );

    expect(container.firstChild).toHaveClass('custom-class');
  });
});
```

### Hook Testing

```typescript
// lib/hooks/use-order.test.ts
import { renderHook, act, waitFor } from '@testing-library/react';
import { useOrder } from './use-order';

// Mock SWR
jest.mock('swr', () => ({
  __esModule: true,
  default: jest.fn(),
}));

import useSWR from 'swr';

describe('useOrder', () => {
  beforeEach(() => {
    jest.clearAllMocks();
  });

  it('returns order data when available', () => {
    const mockOrder = {
      id: 1,
      table_id: 5,
      items: [],
      total: 0,
    };

    (useSWR as jest.Mock).mockReturnValue({
      data: mockOrder,
      error: null,
      isLoading: false,
    });

    const { result } = renderHook(() => useOrder());

    expect(result.current.order).toEqual(mockOrder);
    expect(result.current.error).toBeNull();
  });

  it('handles error state', () => {
    const mockError = new Error('Failed to fetch');

    (useSWR as jest.Mock).mockReturnValue({
      data: null,
      error: mockError,
      isLoading: false,
    });

    const { result } = renderHook(() => useOrder());

    expect(result.current.order).toBeNull();
    expect(result.current.error).toEqual(mockError);
  });

  it('adds item to order', async () => {
    global.fetch = jest.fn(() =>
      Promise.resolve({
        ok: true,
        json: () => Promise.resolve({ success: true }),
      })
    ) as jest.Mock;

    (useSWR as jest.Mock).mockReturnValue({
      data: { id: 1, items: [] },
      error: null,
      isLoading: false,
      mutate: jest.fn(),
    });

    const { result } = renderHook(() => useOrder());

    await act(async () => {
      await result.current.addItem(1, 2);
    });

    expect(global.fetch).toHaveBeenCalledWith('/api/orders/items', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ menuId: 1, quantity: 2 }),
    });
  });
});
```

### Utility Function Testing

```typescript
// lib/utils.test.ts
import { formatPrice, calculateTotal } from './utils';

describe('formatPrice', () => {
  it('formats price with comma separator', () => {
    expect(formatPrice(1000)).toBe('¥1,000');
    expect(formatPrice(1234567)).toBe('¥1,234,567');
  });

  it('handles zero', () => {
    expect(formatPrice(0)).toBe('¥0');
  });

  it('handles negative numbers', () => {
    expect(formatPrice(-500)).toBe('¥-500');
  });
});

describe('calculateTotal', () => {
  it('calculates total correctly', () => {
    const items = [
      { price: 800, quantity: 2 },
      { price: 500, quantity: 1 },
    ];
    expect(calculateTotal(items)).toBe(2100);
  });

  it('returns 0 for empty array', () => {
    expect(calculateTotal([])).toBe(0);
  });
});
```

### Form Testing with Conform + Zod

```typescript
// components/order-form.test.tsx
import { render, screen, fireEvent, waitFor } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { OrderForm } from './order-form';

describe('OrderForm', () => {
  it('validates required fields', async () => {
    render(<OrderForm />);

    const submitButton = screen.getByRole('button', { name: /注文を確定/i });
    fireEvent.click(submitButton);

    await waitFor(() => {
      expect(screen.getByText(/テーブル番号を選択してください/i)).toBeInTheDocument();
    });
  });

  it('submits form with valid data', async () => {
    const mockOnSubmit = jest.fn();
    render(<OrderForm onSubmit={mockOnSubmit} />);

    // Fill in form
    const tableSelect = screen.getByLabelText(/テーブル番号/i);
    await userEvent.selectOptions(tableSelect, '5');

    const submitButton = screen.getByRole('button', { name: /注文を確定/i });
    fireEvent.click(submitButton);

    await waitFor(() => {
      expect(mockOnSubmit).toHaveBeenCalled();
    });
  });
});
```

### Running Frontend Tests

```bash
# Frontend is a single Next.js app
cd frontend

# Run the configured frontend checks
pnpm lint
pnpm build

# If a test script is added later, run it from the frontend root
pnpm test
```

## Test Best Practices

### General Principles

1. **Test Behavior, Not Implementation**: Test what the code does, not how it does it
2. **Write Tests First (TDD)**: Consider writing tests before implementation
3. **Keep Tests Simple**: One test should test one thing
4. **Use Descriptive Names**: Test names should describe what they test
5. **AAA Pattern**: Arrange, Act, Assert
6. **DRY in Tests**: Extract common setup into helper functions
7. **Independent Tests**: Tests should not depend on each other
8. **Fast Tests**: Keep tests fast to encourage frequent running

### Naming Conventions

```go
// Go
func TestFunctionName_Scenario_ExpectedResult(t *testing.T)
func TestUserUsecase_GetUserByID_ReturnsErrorWhenIDIsInvalid(t *testing.T)
```

```typescript
// TypeScript
describe('Component/Function name', () => {
  it('does something in some scenario', () => {
    // Test
  });
});
```

### What to Test

**Backend:**
- ✅ Business logic in UseCases
- ✅ Data validation
- ✅ Error handling
- ✅ Repository implementations (integration tests)
- ✅ HTTP handlers (status codes, response format)
- ❌ External libraries (trust they work)
- ❌ Simple getters/setters without logic

**Frontend:**
- ✅ Component rendering with different props
- ✅ User interactions (clicks, input, forms)
- ✅ Conditional rendering
- ✅ Form validation
- ✅ Data fetching hooks
- ✅ Utility functions
- ❌ Third-party components (trust they work)
- ❌ Styling/CSS (use visual regression tests separately)

### Mock Data

Create reusable mock data:

```typescript
// test/mocks/menu.ts
export const mockMenuItem: Menu = {
  id: 1,
  name: 'カレーライス',
  description: '当店自慢のカレー',
  price: 800,
  category_id: 1,
  is_available: true,
  created_at: '2024-01-01T00:00:00Z',
  updated_at: '2024-01-01T00:00:00Z',
};

export const mockMenuItems: Menu[] = [
  mockMenuItem,
  {
    id: 2,
    name: 'ラーメン',
    description: '醤油ラーメン',
    price: 700,
    category_id: 1,
    is_available: true,
    created_at: '2024-01-01T00:00:00Z',
    updated_at: '2024-01-01T00:00:00Z',
  },
];
```

### Test Coverage Goals

- **Minimum**: 70% code coverage
- **Target**: 80-90% code coverage
- **Critical Paths**: 100% coverage (authentication, payment, orders)

### Continuous Integration

Tests should run automatically:
- On every push
- On pull requests
- Before deployment

```yaml
# .github/workflows/test.yml
name: Tests

on: [push, pull_request]

jobs:
  backend-tests:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      - uses: actions/setup-go@v4
        with:
          go-version: '1.22'
      - name: Run backend tests
        run: cd backend && go test ./... -v

  frontend-tests:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      - uses: pnpm/action-setup@v2
      - uses: actions/setup-node@v3
        with:
          node-version: '18'
      - name: Install dependencies
        run: cd frontend && pnpm install
      - name: Run frontend tests
        run: cd frontend && pnpm test
```

## Debugging Tests

### Go Tests
```bash
# Run with verbose output
go test -v

# Run specific test with print statements
go test -v -run TestName

# Debug with delve
dlv test -- -test.run TestName
```

### TypeScript Tests
```bash
# Run in watch mode
pnpm test:watch

# Debug in VS Code (add breakpoints and run debug configuration)
# Or use console.log for simple debugging
```

## Common Testing Mistakes to Avoid

1. ❌ Testing implementation details instead of behavior
2. ❌ Not cleaning up after tests (database, files, etc.)
3. ❌ Tests that depend on execution order
4. ❌ Overly complex test setup
5. ❌ Not testing error cases
6. ❌ Hardcoding values without explanation
7. ❌ Not mocking external dependencies
8. ❌ Flaky tests (tests that sometimes pass, sometimes fail)
9. ❌ Skipping tests instead of fixing them
10. ❌ Not updating tests when code changes
