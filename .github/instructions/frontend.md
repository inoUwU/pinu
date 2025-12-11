---
applyTo:
  - "frontend/**"
  - "**/package.json"
  - "**/tsconfig.json"
  - "**/*.tsx"
  - "**/*.ts"
---

# Frontend (Next.js + TypeScript) Instructions for Copilot

## Technology Stack

- **Framework**: Next.js (App Router)
- **Language**: TypeScript (strict mode)
- **UI Components**: Shadcn UI
- **Styling**: Tailwind CSS
- **Form Handling**: Conform
- **Validation**: Zod
- **Data Fetching**: SWR
- **Package Manager**: pnpm
- **Monorepo**: Turbo Repo

## Project Structure

```
frontend/
├── apps/
│   ├── client/          # Customer-facing application
│   │   ├── app/         # Next.js App Router
│   │   ├── components/  # React components
│   │   ├── lib/         # Utilities and helpers
│   │   └── public/      # Static assets
│   └── admin/           # Admin/staff application
│       ├── app/
│       ├── components/
│       ├── lib/
│       └── public/
├── packages/            # Shared packages
│   ├── ui/             # Shared UI components
│   ├── types/          # Shared TypeScript types
│   └── utils/          # Shared utilities
├── package.json
├── turbo.json          # Turbo Repo configuration
└── pnpm-workspace.yaml
```

## Next.js App Router Structure

```
app/
├── (routes)/           # Route groups
│   ├── menu/          # Menu page
│   ├── order/         # Order page
│   └── checkout/      # Checkout page
├── api/               # API routes
│   └── orders/        # Order endpoints
├── layout.tsx         # Root layout
├── page.tsx           # Home page
├── loading.tsx        # Loading state
├── error.tsx          # Error boundary
└── not-found.tsx      # 404 page
```

## TypeScript Guidelines

### Strict Mode
TypeScript strict mode is enabled. Always:
- Define explicit types
- Avoid `any` type (use `unknown` if necessary)
- Handle null/undefined explicitly
- Use type guards for type narrowing

Example:
```tsx
// Good
interface User {
  id: number;
  name: string;
  email: string;
}

function getUser(id: number): User | null {
  // Implementation
}

// Bad - avoid any
function getUser(id: any): any {
  // Implementation
}
```

### Type Definitions
Create type definitions in separate files:

```typescript
// types/user.ts
export interface User {
  id: number;
  name: string;
  email: string;
  role: 'admin' | 'staff' | 'customer';
}

export type CreateUserInput = Omit<User, 'id'>;
export type UpdateUserInput = Partial<Omit<User, 'id'>>;
```

## Component Patterns

### Server Components (Default)
Use Server Components by default for better performance:

```tsx
// app/menu/page.tsx
import { getMenuItems } from '@/lib/api/menu';

export default async function MenuPage() {
  const menuItems = await getMenuItems();
  
  return (
    <div>
      <h1>メニュー</h1>
      <MenuList items={menuItems} />
    </div>
  );
}
```

### Client Components
Use Client Components only when needed (interactivity, hooks, browser APIs):

```tsx
'use client';

import { useState } from 'react';
import { Button } from '@/components/ui/button';

export function Counter() {
  const [count, setCount] = useState(0);
  
  return (
    <div>
      <p>カウント: {count}</p>
      <Button onClick={() => setCount(count + 1)}>
        増やす
      </Button>
    </div>
  );
}
```

### Component Structure
```tsx
// components/menu/menu-item.tsx
import { cn } from '@/lib/utils';
import { MenuItem as MenuItemType } from '@/types/menu';

interface MenuItemProps {
  item: MenuItemType;
  onSelect?: (item: MenuItemType) => void;
  className?: string;
}

export function MenuItem({ item, onSelect, className }: MenuItemProps) {
  return (
    <div className={cn("p-4 border rounded-lg", className)}>
      <h3 className="text-lg font-bold">{item.name}</h3>
      <p className="text-gray-600">{item.description}</p>
      <p className="text-xl font-bold">¥{item.price.toLocaleString()}</p>
      {onSelect && (
        <button onClick={() => onSelect(item)}>
          選択
        </button>
      )}
    </div>
  );
}
```

## Form Handling with Conform + Zod

### Define Schema
```typescript
// lib/schemas/order.ts
import { z } from 'zod';

export const orderSchema = z.object({
  tableId: z.number().min(1, 'テーブル番号を選択してください'),
  items: z.array(z.object({
    menuId: z.number(),
    quantity: z.number().min(1, '数量は1以上を指定してください'),
  })).min(1, '少なくとも1つの商品を選択してください'),
  notes: z.string().optional(),
});

export type OrderInput = z.infer<typeof orderSchema>;
```

### Use in Component
```tsx
'use client';

import { useForm } from '@conform-to/react';
import { parseWithZod } from '@conform-to/zod';
import { orderSchema } from '@/lib/schemas/order';

export function OrderForm() {
  const [form, fields] = useForm({
    onValidate({ formData }) {
      return parseWithZod(formData, { schema: orderSchema });
    },
    onSubmit(event, context) {
      event.preventDefault();
      // Handle submission
    },
  });
  
  return (
    <form {...form.props}>
      {/* Form fields */}
    </form>
  );
}
```

## Data Fetching with SWR

### Basic Usage
```tsx
'use client';

import useSWR from 'swr';
import { Menu } from '@/types/menu';

const fetcher = (url: string) => fetch(url).then(r => r.json());

export function MenuList() {
  const { data, error, isLoading } = useSWR<Menu[]>(
    '/api/menus',
    fetcher
  );
  
  if (error) return <div>エラーが発生しました</div>;
  if (isLoading) return <div>読み込み中...</div>;
  if (!data) return null;
  
  return (
    <div className="grid grid-cols-2 gap-4">
      {data.map(item => (
        <MenuItem key={item.id} item={item} />
      ))}
    </div>
  );
}
```

### Mutations
```tsx
'use client';

import useSWR, { mutate } from 'swr';

export function useOrder() {
  const { data, error } = useSWR('/api/orders/current');
  
  const addItem = async (menuId: number, quantity: number) => {
    await fetch('/api/orders/items', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ menuId, quantity }),
    });
    
    // Revalidate
    mutate('/api/orders/current');
  };
  
  return { order: data, error, addItem };
}
```

## Styling with Tailwind CSS

### Best Practices
- Use Tailwind utility classes
- Use `cn()` helper for conditional classes
- Create reusable components with Shadcn UI
- Follow mobile-first approach
- Use consistent spacing scale

Example:
```tsx
import { cn } from '@/lib/utils';

interface CardProps {
  children: React.ReactNode;
  variant?: 'default' | 'outlined';
  className?: string;
}

export function Card({ children, variant = 'default', className }: CardProps) {
  return (
    <div
      className={cn(
        "rounded-lg p-4",
        variant === 'default' && "bg-white shadow-md",
        variant === 'outlined' && "border border-gray-300",
        className
      )}
    >
      {children}
    </div>
  );
}
```

## Shadcn UI Components

### Installation
```bash
npx shadcn-ui@latest add button
npx shadcn-ui@latest add card
npx shadcn-ui@latest add dialog
```

### Usage
```tsx
import { Button } from '@/components/ui/button';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';

export function OrderSummary() {
  return (
    <Card>
      <CardHeader>
        <CardTitle>注文内容</CardTitle>
      </CardHeader>
      <CardContent>
        {/* Content */}
        <Button className="w-full">注文を確定</Button>
      </CardContent>
    </Card>
  );
}
```

## API Routes (App Router)

```typescript
// app/api/orders/route.ts
import { NextRequest, NextResponse } from 'next/server';

export async function GET(request: NextRequest) {
  try {
    const response = await fetch('http://localhost:8000/api/v1/orders');
    const data = await response.json();
    
    return NextResponse.json(data);
  } catch (error) {
    return NextResponse.json(
      { error: 'Failed to fetch orders' },
      { status: 500 }
    );
  }
}

export async function POST(request: NextRequest) {
  try {
    const body = await request.json();
    
    // Validate with Zod
    const validatedData = orderSchema.parse(body);
    
    const response = await fetch('http://localhost:8000/api/v1/orders', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(validatedData),
    });
    
    const data = await response.json();
    return NextResponse.json(data);
  } catch (error) {
    return NextResponse.json(
      { error: 'Failed to create order' },
      { status: 500 }
    );
  }
}
```

## Error Handling

### Error Boundaries
```tsx
// app/error.tsx
'use client';

export default function Error({
  error,
  reset,
}: {
  error: Error & { digest?: string };
  reset: () => void;
}) {
  return (
    <div className="flex flex-col items-center justify-center min-h-screen">
      <h2 className="text-2xl font-bold mb-4">エラーが発生しました</h2>
      <p className="text-gray-600 mb-4">{error.message}</p>
      <button
        onClick={reset}
        className="px-4 py-2 bg-blue-500 text-white rounded"
      >
        再試行
      </button>
    </div>
  );
}
```

### Try-Catch in Components
```tsx
'use client';

import { useState } from 'react';

export function OrderButton() {
  const [error, setError] = useState<string | null>(null);
  
  const handleOrder = async () => {
    try {
      setError(null);
      const response = await fetch('/api/orders', {
        method: 'POST',
        body: JSON.stringify(orderData),
      });
      
      if (!response.ok) {
        throw new Error('注文に失敗しました');
      }
      
      // Success handling
    } catch (err) {
      setError(err instanceof Error ? err.message : '不明なエラー');
    }
  };
  
  return (
    <>
      {error && <div className="text-red-500">{error}</div>}
      <button onClick={handleOrder}>注文する</button>
    </>
  );
}
```

## Accessibility

- Use semantic HTML elements
- Add ARIA labels where needed
- Ensure keyboard navigation works
- Maintain sufficient color contrast
- Test with screen readers

```tsx
<button
  aria-label="メニューを追加"
  onClick={handleAdd}
  className="p-2 rounded-full hover:bg-gray-100"
>
  <PlusIcon className="w-6 h-6" />
</button>
```

## Performance Optimization

### Image Optimization
```tsx
import Image from 'next/image';

<Image
  src="/images/menu/curry.jpg"
  alt="カレー"
  width={300}
  height={200}
  priority={false}
  placeholder="blur"
/>
```

### Dynamic Imports
```tsx
import dynamic from 'next/dynamic';

const HeavyComponent = dynamic(() => import('@/components/heavy-component'), {
  loading: () => <p>読み込み中...</p>,
  ssr: false,
});
```

### Memoization
```tsx
import { memo, useMemo } from 'react';

export const MenuItem = memo(function MenuItem({ item }) {
  const formattedPrice = useMemo(
    () => item.price.toLocaleString('ja-JP'),
    [item.price]
  );
  
  return <div>{formattedPrice}円</div>;
});
```

## Testing

### Component Tests
```typescript
import { render, screen } from '@testing-library/react';
import { MenuItem } from './menu-item';

describe('MenuItem', () => {
  it('renders menu item correctly', () => {
    const item = {
      id: 1,
      name: 'カレー',
      price: 800,
      description: '当店自慢のカレー',
    };
    
    render(<MenuItem item={item} />);
    
    expect(screen.getByText('カレー')).toBeInTheDocument();
    expect(screen.getByText('¥800')).toBeInTheDocument();
  });
});
```

## UI/UX Considerations for Elderly Users

This system targets elderly single-operator restaurant owners:

1. **Large Touch Targets**: Minimum 44x44px for buttons
2. **High Contrast**: Use clear, high-contrast colors
3. **Simple Navigation**: Minimize steps, clear hierarchy
4. **Large Text**: Use readable font sizes (16px minimum)
5. **Clear Labels**: Use descriptive, simple Japanese text
6. **Error Messages**: Clear, actionable error messages
7. **Confirmation Dialogs**: Confirm important actions
8. **Loading States**: Show clear loading indicators

Example:
```tsx
<Button
  size="lg"
  className="text-lg min-h-[44px] min-w-[44px]"
>
  注文を確定する
</Button>
```

## Development Commands

```bash
cd frontend

# Install dependencies
pnpm install

# Run dev server
pnpm dev

# Build for production
pnpm build

# Run tests
pnpm test

# Lint
pnpm lint

# Format with Biome
pnpm format
```

## Common Patterns

### Loading States
```tsx
export function MenuPage() {
  return (
    <Suspense fallback={<MenuSkeleton />}>
      <MenuList />
    </Suspense>
  );
}
```

### Error States
```tsx
if (error) {
  return <ErrorMessage message="データの読み込みに失敗しました" />;
}
```

### Empty States
```tsx
if (items.length === 0) {
  return (
    <div className="text-center py-12">
      <p className="text-gray-500">まだ注文がありません</p>
    </div>
  );
}
```

## Key Principles

1. **Server Components First**: Use Server Components by default
2. **Client Components Only When Needed**: For interactivity
3. **Type Safety**: Use TypeScript strictly
4. **Validation**: Validate all user input with Zod
5. **Accessibility**: Build accessible components
6. **Performance**: Optimize images, lazy load when appropriate
7. **User-Friendly**: Design for elderly users
8. **Error Handling**: Handle errors gracefully
9. **Loading States**: Always show loading indicators
10. **Mobile-First**: Design for mobile, enhance for desktop
