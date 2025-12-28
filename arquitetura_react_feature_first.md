# Arquitetura React: Feature-First + Layered Architecture

## 1. Introdução

Este documento descreve uma arquitetura frontend robusta e escalável para aplicações React/TypeScript, baseada nos princípios de **Feature-First Organization** e **Layered Architecture**.

### 1.1 Visão Geral

A arquitetura é organizada em camadas concêntricas onde cada camada tem responsabilidades bem definidas e as dependências fluem de cima para baixo. Features são organizadas de forma auto-contida, facilitando escalabilidade e manutenção.

```
┌─────────────────────────────────────────┐
│   App Layer (Configuração Global)       │  ← Camada de aplicação
│   - Routes (React Router)                │
│   - Layouts (AppLayout, PublicLayout)   │
│   - Entry Point (main.tsx, App.tsx)     │
├─────────────────────────────────────────┤
│   Features Layer (Feature-First)        │  ← Camada de features
│   auth/  users/  jobs/  dashboard/     │
│   ├── components (UI presentation)      │
│   ├── hooks (React Query)               │
│   ├── pages (Route components)          │
│   ├── schemas (Zod validation)          │
│   └── types (Feature-specific)          │
├─────────────────────────────────────────┤
│   Shared Layer (Código Compartilhado)  │  ← Camada compartilhada
│   ├── components (UI: Button, Form)    │
│   ├── hooks (usePermissions, etc.)     │
│   ├── stores (Zustand: auth)           │
│   ├── types (Shared types)             │
│   ├── constants (Permissions)          │
│   └── utils (Helper functions)         │
├─────────────────────────────────────────┤
│   Infrastructure Layer                  │  ← Camada de infraestrutura
│   ├── api/                              │
│   │   ├── axios-client.ts               │
│   │   └── services/ (API layer)        │
│   └── config/                           │
│       ├── env.ts (Env variables)        │
│       └── constants.ts                  │
└─────────────────────────────────────────┘
```

### 1.2 Princípios Fundamentais

**Feature-First Organization**
- Código organizado por features, não por tipo de arquivo
- Cada feature contém tudo que precisa: components, hooks, pages, schemas
- Features são auto-contidas e podem ser movidas/removidas facilmente
- Reduz acoplamento entre diferentes partes da aplicação

**Layered Architecture**
- Camadas superiores dependem de camadas inferiores
- App depende de Features
- Features dependem de Shared
- Shared depende de Infrastructure
- Nunca o contrário (evita dependências circulares)

**Separation of Concerns**
- UI components apenas apresentam dados
- Hooks encapsulam lógica de negócio e data fetching
- Services abstraem chamadas de API
- State management separado por tipo (global, server, local)

**Type Safety First**
- TypeScript estrito em todo o projeto
- Type inference de schemas Zod
- Contratos claros entre camadas via interfaces
- Path aliases para imports limpos

### 1.3 Vantagens

- **Escalabilidade**: Fácil adicionar novas features sem afetar existentes
- **Manutenibilidade**: Código organizado e fácil de navegar
- **Reutilização**: Shared layer com componentes e hooks reutilizáveis
- **Testabilidade**: Cada camada pode ser testada isoladamente
- **Developer Experience**: Type safety, hot reload, imports limpos
- **Performance**: Code splitting automático, React Query caching
- **Onboarding**: Desenvolvedores novos entendem a estrutura rapidamente

---

## 2. Estrutura de Diretórios

```
frontend/
├── src/
│   ├── app/                     # APP LAYER
│   │   ├── components/          # Layouts, ProtectedRoute
│   │   ├── hooks/              # App-level hooks
│   │   ├── providers/          # Global providers
│   │   └── routes.tsx          # Route configuration
│   │
│   ├── features/               # FEATURES LAYER (Feature-First)
│   │   ├── auth/              # Authentication feature
│   │   │   ├── components/    # LoginForm, RegisterForm
│   │   │   ├── hooks/         # useLogin, useRegister
│   │   │   ├── pages/         # LoginPage, RegisterPage
│   │   │   ├── schemas/       # login.schema.ts, register.schema.ts
│   │   │   └── types/         # auth-specific types
│   │   │
│   │   ├── users/             # Users CRUD feature
│   │   │   ├── components/    # UserForm, UsersTable, UserDetailCard
│   │   │   ├── hooks/         # useUsersList, useUserCreate, useUserUpdate
│   │   │   ├── pages/         # UsersListPage, UserCreatePage, UserEditPage
│   │   │   ├── schemas/       # user-form.schema.ts
│   │   │   └── index.ts       # Barrel export
│   │   │
│   │   ├── jobs/              # Jobs feature
│   │   ├── dashboard/         # Dashboard feature
│   │   ├── profile/           # User profile feature
│   │   └── {feature}/         # Other features...
│   │
│   ├── shared/                # SHARED LAYER
│   │   ├── components/        # Reusable UI components
│   │   │   ├── ui/           # Radix UI components (Button, Input, etc.)
│   │   │   ├── can.tsx       # Permission gate component
│   │   │   ├── pagination.tsx
│   │   │   └── ...
│   │   │
│   │   ├── hooks/            # Shared custom hooks
│   │   │   ├── use-permissions.ts
│   │   │   ├── use-backend-errors.ts
│   │   │   └── ...
│   │   │
│   │   ├── stores/           # Global state (Zustand)
│   │   │   └── auth.store.ts
│   │   │
│   │   ├── types/            # Shared TypeScript types
│   │   │   ├── auth.types.ts
│   │   │   ├── user.types.ts
│   │   │   ├── api.types.ts
│   │   │   └── error.types.ts
│   │   │
│   │   ├── constants/        # App constants
│   │   │   └── permissions.ts
│   │   │
│   │   ├── services/         # Shared services (usually empty)
│   │   ├── lib/              # Utility functions
│   │   │   └── utils.ts      # cn() utility
│   │   └── utils/            # Helper functions
│   │       └── form-errors.ts
│   │
│   ├── infrastructure/       # INFRASTRUCTURE LAYER
│   │   ├── api/
│   │   │   ├── axios-client.ts      # Axios instance with interceptors
│   │   │   └── services/            # API service layer
│   │   │       ├── auth.service.ts
│   │   │       ├── user.service.ts
│   │   │       ├── job.service.ts
│   │   │       └── ...
│   │   └── config/
│   │       ├── env.ts              # Environment variables
│   │       └── constants.ts        # Infrastructure constants
│   │
│   ├── index.css             # Tailwind + Design System (CSS variables)
│   ├── main.tsx              # Entry point (providers setup)
│   └── App.tsx               # Root component
│
├── public/                   # Static assets
├── package.json              # Dependencies
├── vite.config.ts            # Vite configuration
├── tsconfig.json             # TypeScript config (path aliases)
├── tailwind.config.js        # Tailwind CSS config
├── .eslintrc.cjs             # ESLint config
└── postcss.config.js         # PostCSS config
```

### 2.1 Fluxo de Dependências

```
User Interaction (UI)
    ↓
Page Component (features/users/pages/users-list-page.tsx)
    ↓
Custom Hook (features/users/hooks/use-users-list.ts) → React Query
    ↓
Service (infrastructure/api/services/user.service.ts)
    ↓
Axios Client (infrastructure/api/axios-client.ts)
    ↓
Backend API
```

**Importante**: As dependências SEMPRE fluem de cima para baixo:
- App → Features → Shared → Infrastructure
- Features NUNCA importam de outros Features
- Shared NUNCA importa de Features ou App
- Infrastructure NUNCA importa de camadas superiores

---

## 3. Camadas da Arquitetura

### 3.1 Infrastructure Layer (Camada Mais Externa)

Esta camada lida com detalhes técnicos de infraestrutura: comunicação HTTP, configuração de ambiente, e integrações externas.

#### 3.1.1 API Client (`infrastructure/api/axios-client.ts`)

**Conceito**: Singleton do Axios configurado com interceptors para autenticação, refresh de token, e tratamento de erros.

**Características:**
- **Base URL** configurável via env
- **Timeout** de 30 segundos
- **Request Interceptor**: Adiciona token JWT automaticamente
- **Response Interceptor**: Unwraps `data.data`, refresh automático de token em 401
- **Error Handling**: Padroniza erros para toda aplicação

**Exemplo Conceitual:**
```typescript
import axios from 'axios'
import { useAuthStore } from '@shared/stores/auth.store'

const apiClient = axios.create({
  baseURL: env.apiUrl, // http://localhost:8300
  timeout: 30000,
  headers: {
    'Content-Type': 'application/json',
  },
})

// Request interceptor: adiciona token
apiClient.interceptors.request.use(
  (config) => {
    const { accessToken } = useAuthStore.getState()
    if (accessToken) {
      config.headers.Authorization = `Bearer ${accessToken}`
    }
    return config
  },
  (error) => Promise.reject(error)
)

// Response interceptor: unwrap data, handle errors
apiClient.interceptors.response.use(
  (response) => {
    // Unwrap backend response: { success: true, data: {...} }
    if (response.data?.data !== undefined) {
      return { ...response, data: response.data.data }
    }
    return response
  },
  async (error) => {
    const originalRequest = error.config

    // Token refresh on 401
    if (error.response?.status === 401 && !originalRequest._retry) {
      originalRequest._retry = true

      try {
        const { refreshToken } = useAuthStore.getState()
        const response = await authService.refreshToken(refreshToken)
        const { access_token } = response

        useAuthStore.getState().setAuth({ accessToken: access_token, ... })
        originalRequest.headers.Authorization = `Bearer ${access_token}`

        return apiClient(originalRequest) // Retry original request
      } catch (refreshError) {
        useAuthStore.getState().clearAuth()
        window.location.href = '/login'
        return Promise.reject(refreshError)
      }
    }

    return Promise.reject(error)
  }
)

export { apiClient }
```

**Responsabilidades:**
- Configurar cliente HTTP
- Adicionar token de autenticação
- Refresh automático de token
- Unwrap respostas do backend
- Tratamento global de erros

**NÃO faz:**
- Lógica de negócio
- Manipulação de state
- Transformação de dados (isso é do service)

#### 3.1.2 Services (`infrastructure/api/services/`)

**Conceito**: Services abstraem chamadas de API, fornecendo uma interface limpa para as camadas superiores. Um service por recurso.

**Características:**
- Métodos CRUD: `getAll`, `getById`, `create`, `update`, `delete`
- Type-safe: usa interfaces TypeScript para request/response
- Usa `apiClient` singleton
- Retorna Promises com tipos específicos

**Exemplo Conceitual - User Service:**
```typescript
import { apiClient } from '../axios-client'
import type { User, UserRequest, PaginatedResponse } from '@shared/types'

export const userService = {
  async getAll(params?: { page?: number; limit?: number }): Promise<PaginatedResponse<User>> {
    const response = await apiClient.get<PaginatedResponse<User>>('/users', { params })
    return response.data
  },

  async getById(id: string): Promise<User> {
    const response = await apiClient.get<User>(`/users/${id}`)
    return response.data
  },

  async create(userData: UserRequest): Promise<User> {
    const response = await apiClient.post<User>('/users', userData)
    return response.data
  },

  async update(id: string, userData: Partial<UserRequest>): Promise<User> {
    const response = await apiClient.put<User>(`/users/${id}`, userData)
    return response.data
  },

  async delete(id: string): Promise<void> {
    await apiClient.delete(`/users/${id}`)
  },
}
```

**Responsabilidades:**
- Encapsular chamadas HTTP
- Definir contratos de API (tipos de request/response)
- Passar parâmetros e payloads
- Retornar dados tipados

**NÃO faz:**
- Gerenciar cache (isso é do React Query)
- Lógica de negócio
- Manipulação de state

#### 3.1.3 Config (`infrastructure/config/`)

**Conceito**: Configuração de ambiente e constantes de infraestrutura.

**env.ts:**
```typescript
const env = {
  apiUrl: import.meta.env.VITE_API_URL || 'http://localhost:8300',
  environment: import.meta.env.MODE, // 'development' | 'production'
}

export { env }
```

**Por que usar?**
- Centraliza configuração
- Type-safe environment variables
- Fácil trocar entre ambientes (dev, staging, prod)

---

### 3.2 Shared Layer (Camada Compartilhada)

Código verdadeiramente compartilhado entre múltiplas features. Apenas coloque aqui o que é realmente reutilizado.

#### 3.2.1 Components (`shared/components/`)

**Conceito**: Componentes UI reutilizáveis, genéricos e sem lógica de negócio específica.

**UI Components (`shared/components/ui/`):**
- Base components do Radix UI estilizados com Tailwind
- Button, Input, Select, Checkbox, Dialog, Card, etc.
- Acessíveis por padrão (Radix UI)
- Consistentes com design system

**Exemplo - Button Component:**
```typescript
import { forwardRef } from 'react'
import { cn } from '@shared/lib/utils'

interface ButtonProps extends React.ButtonHTMLAttributes<HTMLButtonElement> {
  variant?: 'primary' | 'secondary' | 'outline' | 'ghost'
  size?: 'sm' | 'md' | 'lg'
}

const Button = forwardRef<HTMLButtonElement, ButtonProps>(
  ({ className, variant = 'primary', size = 'md', ...props }, ref) => {
    return (
      <button
        ref={ref}
        className={cn(
          // Base styles
          'inline-flex items-center justify-center rounded-md font-medium transition-colors',
          'focus-visible:outline-none focus-visible:ring-2 disabled:opacity-50',
          // Variants
          {
            'bg-primary text-white hover:bg-primary/90': variant === 'primary',
            'bg-secondary text-white hover:bg-secondary/90': variant === 'secondary',
            'border border-input hover:bg-accent': variant === 'outline',
            'hover:bg-accent': variant === 'ghost',
          },
          // Sizes
          {
            'h-9 px-3 text-sm': size === 'sm',
            'h-10 px-4': size === 'md',
            'h-11 px-6 text-lg': size === 'lg',
          },
          className
        )}
        {...props}
      />
    )
  }
)

export { Button }
```

**Common Components:**
- `can.tsx`: Permission-based conditional rendering
- `pagination.tsx`: Pagination component
- `confirm-dialog.tsx`: Confirmation modal
- `spinner.tsx`: Loading spinner

**Responsabilidades:**
- Apresentação visual
- Acessibilidade
- Variantes e tamanhos
- Composição via props

**NÃO fazem:**
- API calls
- State management de features
- Lógica de negócio

#### 3.2.2 Hooks (`shared/hooks/`)

**Conceito**: Custom hooks reutilizáveis que encapsulam lógica comum.

**usePermissions Hook:**
```typescript
import { useAuthStore } from '@shared/stores/auth.store'

export const usePermissions = () => {
  const user = useAuthStore((state) => state.user)
  const role = user?.role

  const hasPermission = (permissionSlug: string): boolean => {
    if (!role?.permissions) return false

    return role.permissions.some((perm) => {
      // Exact match
      if (perm === permissionSlug) return true

      // Wildcard match: "users.*" matches "users.create"
      if (perm.endsWith('.*')) {
        const prefix = perm.slice(0, -2)
        return permissionSlug.startsWith(prefix + '.')
      }

      // Super admin: "*" matches everything
      return perm === '*'
    })
  }

  const hasAnyPermission = (permissions: string[]): boolean => {
    return permissions.some((perm) => hasPermission(perm))
  }

  const hasAllPermissions = (permissions: string[]): boolean => {
    return permissions.every((perm) => hasPermission(perm))
  }

  const hasRole = (roleSlug: string): boolean => {
    return role?.slug === roleSlug
  }

  const isAdmin = (): boolean => {
    return hasRole('admin') || hasPermission('*')
  }

  return {
    hasPermission,
    hasAnyPermission,
    hasAllPermissions,
    hasRole,
    isAdmin,
    permissions: role?.permissions || [],
    role: role?.slug,
  }
}
```

**useBackendErrors Hook:**
```typescript
import { UseFormSetError } from 'react-hook-form'
import { toast } from 'sonner'

interface BackendErrorResponse {
  success: false
  message: string
  errors?: Record<string, string[]>
}

interface UseBackendErrorsOptions {
  setError: UseFormSetError<any>
  onGlobalError?: (message: string) => void
  fieldMapping?: Record<string, string> // backend field → form field
}

export const useBackendErrors = (options: UseBackendErrorsOptions) => {
  const { setError, onGlobalError, fieldMapping = {} } = options

  const handleBackendErrors = (errorResponse: BackendErrorResponse) => {
    // Global error
    if (!errorResponse.errors) {
      const message = errorResponse.message || 'An error occurred'
      onGlobalError?.(message) || toast.error(message)
      return
    }

    // Field errors
    Object.entries(errorResponse.errors).forEach(([field, messages]) => {
      const formField = fieldMapping[field] || field
      setError(formField, {
        type: 'manual',
        message: messages[0], // First error message
      })
    })
  }

  return { handleBackendErrors }
}
```

**Responsabilidades:**
- Encapsular lógica reutilizável
- Abstrair complexidade
- Type-safe
- Sem side effects não declarados

#### 3.2.3 Stores (`shared/stores/`)

**Conceito**: State management global com Zustand. Apenas para estado verdadeiramente global (auth, user, theme).

**Auth Store:**
```typescript
import { create } from 'zustand'
import { persist } from 'zustand/middleware'

interface UserProfile {
  id: string
  email: string
  first_name: string
  last_name: string
  role: {
    slug: string
    permissions: string[]
  }
}

interface AuthState {
  user: UserProfile | null
  accessToken: string | null
  refreshToken: string | null
  isAuthenticated: boolean
  _hasHydrated: boolean

  // Actions
  setAuth: (data: {
    user: UserProfile
    accessToken: string
    refreshToken: string
  }) => void
  clearAuth: () => void
  updateUser: (user: UserProfile) => void
  setHasHydrated: (state: boolean) => void
}

export const useAuthStore = create<AuthState>()(
  persist(
    (set) => ({
      user: null,
      accessToken: null,
      refreshToken: null,
      isAuthenticated: false,
      _hasHydrated: false,

      setAuth: (data) =>
        set({
          user: data.user,
          accessToken: data.accessToken,
          refreshToken: data.refreshToken,
          isAuthenticated: true,
        }),

      clearAuth: () =>
        set({
          user: null,
          accessToken: null,
          refreshToken: null,
          isAuthenticated: false,
        }),

      updateUser: (user) => set({ user }),

      setHasHydrated: (state) => set({ _hasHydrated: state }),
    }),
    {
      name: 'auth-storage', // localStorage key
      onRehydrateStorage: () => (state) => {
        state?.setHasHydrated(true)
      },
    }
  )
)
```

**Responsabilidades:**
- Gerenciar estado global (auth, user)
- Persistir em localStorage (via middleware)
- Actions para mutação de state
- Selectors para leitura de state

**Quando usar:**
- Auth/user information
- Theme (light/dark mode)
- App-level preferences

**Quando NÃO usar:**
- Server data (use React Query)
- Form state (use React Hook Form)
- UI state local (use useState)

#### 3.2.4 Types (`shared/types/`)

**Conceito**: Type definitions compartilhadas entre features.

**auth.types.ts:**
```typescript
export interface LoginRequest {
  email: string
  password: string
}

export interface LoginResponse {
  access_token: string
  refresh_token: string
  token_type: string
  expires_in: number
  user: UserProfile
}

export interface RefreshTokenResponse {
  access_token: string
  token_type: string
  expires_in: number
}
```

**user.types.ts:**
```typescript
export interface User {
  uuid: string
  email: string
  first_name: string
  last_name: string
  phone?: string
  role_id: string
  user_type: string
  created_at: string
  updated_at: string
}

export interface UserRequest {
  email: string
  first_name: string
  last_name: string
  phone?: string
  role_id: string
  user_type: string
}
```

**api.types.ts:**
```typescript
export interface PaginatedResponse<T> {
  data: T[]
  total: number
  limit: number
  offset: number
}

export interface ApiSuccessResponse<T> {
  success: true
  data: T
  message?: string
}

export interface ApiErrorResponse {
  success: false
  message: string
  code?: string
  errors?: Record<string, string[]>
}
```

#### 3.2.5 Constants (`shared/constants/`)

**Conceito**: Constantes da aplicação, principalmente permissões.

**permissions.ts:**
```typescript
export const PERMISSIONS = {
  // Users
  USERS_CREATE: 'users.create',
  USERS_READ: 'users.read',
  USERS_UPDATE: 'users.update',
  USERS_DELETE: 'users.delete',
  USERS_MANAGE: 'users.*',

  // Jobs
  JOBS_CREATE: 'jobs.create',
  JOBS_READ: 'jobs.read',
  JOBS_UPDATE: 'jobs.update',
  JOBS_DELETE: 'jobs.delete',
  JOBS_MANAGE: 'jobs.*',

  // Admin
  ADMIN_ACCESS: 'admin.*',
  SUPER_ADMIN: '*',
} as const

export type Permission = typeof PERMISSIONS[keyof typeof PERMISSIONS]
```

#### 3.2.6 Lib/Utils (`shared/lib/utils.ts`)

**Conceito**: Funções utilitárias puras.

**cn() - Class Name Utility:**
```typescript
import { type ClassValue, clsx } from 'clsx'
import { twMerge } from 'tailwind-merge'

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs))
}

// Uso:
<Button className={cn('bg-primary', isLoading && 'opacity-50', className)} />
```

**Por que usar?**
- Merge classes Tailwind corretamente
- Evita conflitos de classes (ex: `bg-red-500` sobrescreve `bg-blue-500`)
- Suporta conditional classes

---

### 3.3 Features Layer (Feature-First)

Cada feature é auto-contida com seus próprios components, hooks, pages, schemas e types.

#### 3.3.1 Estrutura de uma Feature

```
features/users/
├── components/          # UI components específicos da feature
│   ├── user-form.tsx
│   ├── users-table.tsx
│   ├── user-detail-card.tsx
│   └── index.ts        # Barrel export
├── hooks/              # React Query hooks
│   ├── use-users-list.ts
│   ├── use-user-detail.ts
│   ├── use-user-create.ts
│   ├── use-user-update.ts
│   ├── use-user-delete.ts
│   └── index.ts
├── pages/              # Route components
│   ├── users-list-page.tsx
│   ├── user-create-page.tsx
│   ├── user-edit-page.tsx
│   ├── user-detail-page.tsx
│   └── index.ts
├── schemas/            # Zod validation schemas
│   └── user-form.schema.ts
└── index.ts            # Barrel export (optional)
```

#### 3.3.2 Components (`features/{feature}/components/`)

**Conceito**: Componentes de apresentação específicos da feature. Podem usar lógica, mas não fazem API calls diretamente.

**Exemplo - UserForm:**
```typescript
import { useForm } from 'react-hook-form'
import { zodResolver } from '@hookform/resolvers/zod'
import { Button, Form, FormField, Input } from '@shared/components/ui'
import { createUserSchema, type CreateUserFormData } from '../schemas/user-form.schema'

interface UserFormProps {
  onSubmit: (data: CreateUserFormData) => void
  defaultValues?: Partial<CreateUserFormData>
  isLoading?: boolean
}

export const UserForm = ({ onSubmit, defaultValues, isLoading }: UserFormProps) => {
  const form = useForm<CreateUserFormData>({
    resolver: zodResolver(createUserSchema),
    defaultValues: defaultValues || {
      email: '',
      first_name: '',
      last_name: '',
      role_id: '',
      user_type: 'ADMIN',
    },
  })

  return (
    <Form {...form}>
      <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-4">
        <FormField
          control={form.control}
          name="email"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Email</FormLabel>
              <FormControl>
                <Input type="email" {...field} />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />

        <FormField
          control={form.control}
          name="first_name"
          render={({ field }) => (
            <FormItem>
              <FormLabel>First Name</FormLabel>
              <FormControl>
                <Input {...field} />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />

        {/* More fields... */}

        <Button type="submit" disabled={isLoading}>
          {isLoading ? 'Saving...' : 'Save User'}
        </Button>
      </form>
    </Form>
  )
}
```

**Responsabilidades:**
- Render UI
- Gerenciar form state (via React Hook Form)
- Validar input (via Zod)
- Emitir eventos (onSubmit, onClick, etc.)

**NÃO fazem:**
- API calls (isso é do hook)
- State management global (isso é do store)

#### 3.3.3 Hooks (`features/{feature}/hooks/`)

**Conceito**: Custom hooks que encapsulam lógica de data fetching usando React Query.

**Query Hook - useUsersList:**
```typescript
import { useQuery } from '@tanstack/react-query'
import { userService } from '@infrastructure/api/services/user.service'

interface UseUsersListParams {
  page?: number
  limit?: number
}

export const useUsersList = (params: UseUsersListParams = {}) => {
  return useQuery({
    queryKey: ['users', params],
    queryFn: () => userService.getAll(params),
    staleTime: 5 * 60 * 1000, // 5 minutes
  })
}

// Uso em componente:
const { data, isLoading, error } = useUsersList({ page: 1, limit: 10 })
```

**Mutation Hook - useUserCreate:**
```typescript
import { useMutation, useQueryClient } from '@tanstack/react-query'
import { useNavigate } from 'react-router-dom'
import { toast } from 'sonner'
import { userService } from '@infrastructure/api/services/user.service'
import { useBackendErrors } from '@shared/hooks/use-backend-errors'
import type { UserRequest } from '@shared/types'

export const useUserCreate = (setError: UseFormSetError<any>) => {
  const queryClient = useQueryClient()
  const navigate = useNavigate()
  const { handleBackendErrors } = useBackendErrors({
    setError,
    onGlobalError: (msg) => toast.error(msg),
  })

  return useMutation({
    mutationFn: (userData: UserRequest) => userService.create(userData),
    onSuccess: (newUser) => {
      // Invalidate queries to refetch
      queryClient.invalidateQueries({ queryKey: ['users'] })

      toast.success('User created successfully!')
      navigate(`/users/${newUser.uuid}`)
    },
    onError: (error: any) => {
      if (error.response?.data) {
        handleBackendErrors(error.response.data)
      } else {
        toast.error('Failed to create user')
      }
    },
  })
}

// Uso em componente:
const createMutation = useUserCreate(form.setError)
const handleSubmit = (data: UserRequest) => {
  createMutation.mutate(data)
}
```

**Responsabilidades:**
- Encapsular React Query logic
- Chamar services
- Invalidar queries após mutations
- Tratar erros
- Navigate após sucesso

**Vantagens:**
- Caching automático
- Refetching em background
- Loading/error states
- Invalidation após mutations

#### 3.3.4 Pages (`features/{feature}/pages/`)

**Conceito**: Componentes de rota que compõem outros componentes e hooks.

**Exemplo - UsersListPage:**
```typescript
import { useState } from 'react'
import { Link } from 'react-router-dom'
import { Button, Card } from '@shared/components/ui'
import { Can } from '@shared/components/can'
import { PERMISSIONS } from '@shared/constants/permissions'
import { UsersTable } from '../components/users-table'
import { useUsersList } from '../hooks/use-users-list'

export const UsersListPage = () => {
  const [page, setPage] = useState(1)
  const { data, isLoading, error } = useUsersList({ page, limit: 10 })

  if (isLoading) return <div>Loading...</div>
  if (error) return <div>Error loading users</div>

  return (
    <div className="space-y-4">
      <div className="flex justify-between items-center">
        <h1 className="text-2xl font-bold">Users</h1>

        <Can permission={PERMISSIONS.USERS_CREATE}>
          <Link to="/users/create">
            <Button>Create User</Button>
          </Link>
        </Can>
      </div>

      <Card>
        <UsersTable
          users={data?.data || []}
          total={data?.total || 0}
          page={page}
          onPageChange={setPage}
        />
      </Card>
    </div>
  )
}
```

**Responsabilidades:**
- Compor components
- Usar hooks para data fetching
- Gerenciar UI state (pagination, filters)
- Navigation
- Permission checks

**NÃO fazem:**
- Lógica de negócio complexa
- API calls diretas
- Manipulação direta de DOM

#### 3.3.5 Schemas (`features/{feature}/schemas/`)

**Conceito**: Schemas de validação Zod para forms.

**Exemplo - user-form.schema.ts:**
```typescript
import { z } from 'zod'

export const createUserSchema = z.object({
  email: z.string().min(1, 'Email is required').email('Invalid email'),
  first_name: z.string().min(1, 'First name is required').max(100),
  last_name: z.string().min(1, 'Last name is required').max(100),
  phone: z.string().max(20).optional().or(z.literal('')),
  role_id: z.string().uuid('Invalid role').min(1, 'Role is required'),
  user_type: z.enum(['ADMIN', 'PARTNER', 'RECRUITER', 'CLIENT', 'TALENT']),
})

export const updateUserSchema = z.object({
  email: z.string().email('Invalid email').optional(),
  first_name: z.string().min(1).max(100).optional(),
  last_name: z.string().min(1).max(100).optional(),
  phone: z.string().max(20).optional().or(z.literal('')),
  role_id: z.string().uuid('Invalid role').optional(),
})

// Type inference
export type CreateUserFormData = z.infer<typeof createUserSchema>
export type UpdateUserFormData = z.infer<typeof updateUserSchema>
```

**Por que usar?**
- Type-safe validation
- Type inference automática
- Mensagens de erro customizáveis
- Integração perfeita com React Hook Form
- Schemas reutilizáveis

---

### 3.4 App Layer (Camada de Aplicação)

Camada de configuração global: routing, layouts, providers.

#### 3.4.1 Routes (`app/routes.tsx`)

**Conceito**: Configuração centralizada de todas as rotas da aplicação.

**Exemplo:**
```typescript
import { lazy, Suspense } from 'react'
import { createBrowserRouter, RouterProvider } from 'react-router-dom'
import { ProtectedRoute } from './components/protected-route'
import { PublicLayout } from './components/public-layout'
import { AppLayout } from './components/app-layout'

// Lazy load pages
const HomePage = lazy(() => import('@features/home/pages/home-page'))
const LoginPage = lazy(() => import('@features/auth/pages/login-page'))
const UsersListPage = lazy(() => import('@features/users/pages/users-list-page'))
const UserCreatePage = lazy(() => import('@features/users/pages/user-create-page'))

const router = createBrowserRouter([
  // Public routes
  {
    element: <PublicLayout />,
    children: [
      { path: '/', element: <HomePage /> },
      { path: '/login', element: <LoginPage /> },
      { path: '/register', element: <RegisterPage /> },
    ],
  },

  // Protected routes
  {
    element: (
      <ProtectedRoute>
        <AppLayout />
      </ProtectedRoute>
    ),
    children: [
      { path: '/dashboard', element: <DashboardPage /> },
      { path: '/profile', element: <ProfilePage /> },

      // Users CRUD
      { path: '/users', element: <UsersListPage /> },
      { path: '/users/create', element: <UserCreatePage /> },
      { path: '/users/:id', element: <UserDetailPage /> },
      { path: '/users/:id/edit', element: <UserEditPage /> },

      // Jobs CRUD
      { path: '/jobs', element: <JobsListPage /> },
      // ...
    ],
  },

  // 404
  { path: '*', element: <NotFoundPage /> },
])

export const AppRoutes = () => {
  return (
    <Suspense fallback={<div>Loading...</div>}>
      <RouterProvider router={router} />
    </Suspense>
  )
}
```

**Responsabilidades:**
- Definir todas as rotas
- Lazy loading de páginas
- Layouts por grupo de rotas
- Protected routes
- 404 handling

#### 3.4.2 Components (`app/components/`)

**ProtectedRoute:**
```typescript
import { Navigate } from 'react-router-dom'
import { useAuthStore } from '@shared/stores/auth.store'

interface ProtectedRouteProps {
  children: React.ReactNode
}

export const ProtectedRoute = ({ children }: ProtectedRouteProps) => {
  const isAuthenticated = useAuthStore((state) => state.isAuthenticated)

  if (!isAuthenticated) {
    return <Navigate to="/login" replace />
  }

  return <>{children}</>
}
```

**AppLayout:**
```typescript
import { Outlet } from 'react-router-dom'
import { Header } from './header'
import { Sidebar } from './sidebar'

export const AppLayout = () => {
  return (
    <div className="flex h-screen">
      <Sidebar />
      <div className="flex-1 flex flex-col">
        <Header />
        <main className="flex-1 overflow-auto p-6">
          <Outlet />
        </main>
      </div>
    </div>
  )
}
```

#### 3.4.3 Entry Point (`main.tsx`)

**Conceito**: Setup de providers e inicialização da aplicação.

```typescript
import React from 'react'
import ReactDOM from 'react-dom/client'
import { QueryClient, QueryClientProvider } from '@tanstack/react-query'
import { Toaster } from 'sonner'
import App from './App'
import './index.css'

const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      staleTime: 5 * 60 * 1000, // 5 minutes
      retry: 1,
    },
  },
})

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <QueryClientProvider client={queryClient}>
      <App />
      <Toaster position="top-right" />
    </QueryClientProvider>
  </React.StrictMode>
)
```

**App.tsx:**
```typescript
import { HydrationBoundary } from '@shared/components/hydration-boundary'
import { AppRoutes } from './app/routes'

function App() {
  return (
    <HydrationBoundary>
      <AppRoutes />
    </HydrationBoundary>
  )
}

export default App
```

**HydrationBoundary:**
```typescript
import { useEffect, useState } from 'react'
import { useAuthStore } from '@shared/stores/auth.store'

export const HydrationBoundary = ({ children }: { children: React.ReactNode }) => {
  const [isHydrated, setIsHydrated] = useState(false)

  useEffect(() => {
    // Wait for Zustand to rehydrate from localStorage
    const unsubscribe = useAuthStore.persist.onFinishHydration(() => {
      setIsHydrated(true)
    })

    return () => unsubscribe()
  }, [])

  if (!isHydrated) {
    return <div>Loading...</div>
  }

  return <>{children}</>
}
```

---

## 4. Padrões Arquiteturais Aplicados

### 4.1 Feature-First Architecture

**Princípio**: Organize código por features, não por tipo de arquivo.

**Tradicional (tipo de arquivo):**
```
src/
├── components/
│   ├── UserForm.tsx
│   ├── UserTable.tsx
│   ├── JobForm.tsx
│   └── JobTable.tsx
├── hooks/
│   ├── useUsers.ts
│   └── useJobs.ts
├── pages/
│   ├── UsersPage.tsx
│   └── JobsPage.tsx
```

**Feature-First:**
```
src/features/
├── users/
│   ├── components/
│   │   ├── UserForm.tsx
│   │   └── UserTable.tsx
│   ├── hooks/
│   │   └── useUsers.ts
│   └── pages/
│       └── UsersPage.tsx
├── jobs/
│   ├── components/
│   │   ├── JobForm.tsx
│   │   └── JobTable.tsx
│   ├── hooks/
│   │   └── useJobs.ts
│   └── pages/
│       └── JobsPage.tsx
```

**Vantagens:**
- Features auto-contidas
- Fácil adicionar/remover features
- Reduz acoplamento
- Equipes podem trabalhar em features independentes
- Código relacionado fica junto

### 4.2 Layered Architecture

**Princípio**: Camadas superiores dependem de camadas inferiores, nunca o contrário.

```
App Layer
  ↓ (pode importar)
Features Layer
  ↓ (pode importar)
Shared Layer
  ↓ (pode importar)
Infrastructure Layer
```

**Regras:**
- **App** pode importar de Features, Shared, Infrastructure
- **Features** podem importar de Shared, Infrastructure (NÃO de outras Features)
- **Shared** pode importar de Infrastructure (NÃO de Features ou App)
- **Infrastructure** não importa de nada acima

**Por que?**
- Evita dependências circulares
- Código reutilizável sempre em camadas inferiores
- Fácil entender fluxo de dados
- Refatoração facilitada

### 4.3 State Management Pattern

**Princípio**: Use a ferramenta certa para cada tipo de state.

**Global State (Zustand):**
- Auth/user information
- Theme (dark mode)
- App preferences

```typescript
const user = useAuthStore((state) => state.user)
```

**Server State (React Query):**
- Dados de API
- Caching automático
- Refetching em background

```typescript
const { data: users } = useUsersList()
```

**Local State (useState):**
- UI state (modais, dropdowns)
- Form state (via React Hook Form)

```typescript
const [isOpen, setIsOpen] = useState(false)
```

**Por que separar?**
- Cada ferramenta é otimizada para seu propósito
- Evita duplicação (não coloque server data no Zustand)
- Performance (React Query cache, Zustand selective subscription)

### 4.4 Data Fetching Pattern

**Fluxo:**
```
Component
  → Custom Hook (React Query)
    → Service
      → Axios Client
        → Backend API
```

**Exemplo Completo:**

**1. Component usa hook:**
```typescript
const { data, isLoading } = useUsersList({ page: 1 })
```

**2. Hook usa React Query:**
```typescript
export const useUsersList = (params) => {
  return useQuery({
    queryKey: ['users', params],
    queryFn: () => userService.getAll(params),
  })
}
```

**3. Service faz request:**
```typescript
export const userService = {
  getAll: (params) => apiClient.get('/users', { params }),
}
```

**4. Axios Client adiciona auth e trata erros:**
```typescript
apiClient.interceptors.request.use((config) => {
  config.headers.Authorization = `Bearer ${token}`
  return config
})
```

**Vantagens:**
- Separation of concerns
- Caching automático (React Query)
- Error handling centralizado
- Type safety
- Fácil testar cada camada

### 4.5 Form Handling Pattern

**Stack**: React Hook Form + Zod + Backend Error Mapping

**Fluxo:**
```
User input
  → Form validation (Zod)
  → Submit handler
  → Mutation hook (React Query)
  → Service (API call)
  → Backend validation
  → Error mapping (backend fields → form fields)
  → Display errors
```

**Exemplo:**

**1. Schema Zod:**
```typescript
const loginSchema = z.object({
  email: z.string().email('Invalid email'),
  password: z.string().min(8, 'Min 8 characters'),
})

type LoginFormData = z.infer<typeof loginSchema>
```

**2. Form setup:**
```typescript
const form = useForm<LoginFormData>({
  resolver: zodResolver(loginSchema),
  defaultValues: { email: '', password: '' },
})
```

**3. Mutation com error handling:**
```typescript
const loginMutation = useLogin(form.setError)

const onSubmit = (data: LoginFormData) => {
  loginMutation.mutate(data)
}
```

**4. Backend error mapping:**
```typescript
useMutation({
  mutationFn: authService.login,
  onError: (error) => {
    handleBackendErrors(error.response.data) // Maps to form fields
  }
})
```

**Vantagens:**
- Type-safe validation
- Field-level errors do backend
- Ótima UX (erros inline)
- Reutilizável

### 4.6 Component Composition Pattern

**Princípio**: Componha pequenos componentes reutilizáveis.

**Hierarquia:**
```
Page Component (UsersListPage)
  ↓ compõe
Feature Components (UsersTable, UserFilters)
  ↓ usam
UI Components (Button, Input, Card)
```

**Exemplo:**
```typescript
// Page Component
<UsersListPage>
  <UsersFilters onFilter={...} />
  <UsersTable users={data} />
  <Pagination page={page} onChange={...} />
</UsersListPage>

// Feature Component
<UsersTable>
  <Card>
    <Table>
      <TableHeader>...</TableHeader>
      <TableBody>
        {users.map(user => (
          <TableRow>
            <TableCell>{user.name}</TableCell>
            <TableCell>
              <Button>Edit</Button>
              <Button>Delete</Button>
            </TableCell>
          </TableRow>
        ))}
      </TableBody>
    </Table>
  </Card>
</UsersTable>
```

**Vantagens:**
- Reusabilidade
- Testabilidade
- Manutenibilidade
- Single Responsibility

### 4.7 Permission-Based Rendering

**Conceito**: Conditional rendering baseado em permissões do usuário.

**Hook usePermissions:**
```typescript
const { hasPermission, isAdmin } = usePermissions()

if (hasPermission('users.create')) {
  // Show create button
}
```

**Componente Can:**
```typescript
// Simple permission
<Can permission={PERMISSIONS.USERS_CREATE}>
  <Button>Create User</Button>
</Can>

// Multiple permissions (any)
<Can permission={[PERMISSIONS.USERS_UPDATE, PERMISSIONS.USERS_DELETE]} any>
  <Button>Manage</Button>
</Can>

// Admin check
<Can admin>
  <AdminPanel />
</Can>

// Inverse
<Can permission={PERMISSIONS.USERS_DELETE} not>
  <span>You cannot delete users</span>
</Can>
```

**Implementação do Can:**
```typescript
interface CanProps {
  permission?: string | string[]
  any?: boolean
  admin?: boolean
  not?: boolean
  children: React.ReactNode
}

export const Can = ({ permission, any, admin, not, children }: CanProps) => {
  const { hasPermission, hasAnyPermission, isAdmin } = usePermissions()

  let hasAccess = false

  if (admin) {
    hasAccess = isAdmin()
  } else if (Array.isArray(permission)) {
    hasAccess = any
      ? hasAnyPermission(permission)
      : hasAllPermissions(permission)
  } else if (permission) {
    hasAccess = hasPermission(permission)
  }

  if (not) {
    hasAccess = !hasAccess
  }

  return hasAccess ? <>{children}</> : null
}
```

---

## 5. Implementação de RBAC (Frontend)

### 5.1 Permission System

**Conceito**: Sincronizar permissões do frontend com backend, permitindo controle granular de acesso.

**Constantes de Permissão:**
```typescript
// shared/constants/permissions.ts
export const PERMISSIONS = {
  // Users
  USERS_CREATE: 'users.create',
  USERS_READ: 'users.read',
  USERS_UPDATE: 'users.update',
  USERS_DELETE: 'users.delete',
  USERS_MANAGE: 'users.*', // Wildcard

  // Jobs
  JOBS_CREATE: 'jobs.create',
  JOBS_READ: 'jobs.read',
  JOBS_PUBLISH: 'jobs.publish',
  JOBS_ARCHIVE: 'jobs.archive',
  JOBS_MANAGE: 'jobs.*',

  // Applications
  APPLICATIONS_CREATE: 'applications.create',
  APPLICATIONS_READ: 'applications.read',
  APPLICATIONS_APPROVE: 'applications.approve',
  APPLICATIONS_REJECT: 'applications.reject',

  // Admin
  ADMIN_ACCESS: 'admin.*',
  SUPER_ADMIN: '*', // All permissions
} as const
```

**User Type com Permissões:**
```typescript
interface UserRoleWithPermissions {
  slug: string
  permissions: string[]
}

interface UserProfile {
  id: string
  email: string
  role: UserRoleWithPermissions
}
```

### 5.2 usePermissions Hook

**Conceito**: Hook para verificação de permissões em qualquer componente.

**Implementação:**
```typescript
import { useAuthStore } from '@shared/stores/auth.store'

export const usePermissions = () => {
  const user = useAuthStore((state) => state.user)
  const role = user?.role

  const hasPermission = (permissionSlug: string): boolean => {
    if (!role?.permissions) return false

    return role.permissions.some((perm) => {
      // Exact match: "users.create" === "users.create"
      if (perm === permissionSlug) return true

      // Wildcard match: "users.*" matches "users.create"
      if (perm.endsWith('.*')) {
        const prefix = perm.slice(0, -2) // Remove ".*"
        return permissionSlug.startsWith(prefix + '.')
      }

      // Super admin: "*" matches everything
      return perm === '*'
    })
  }

  const hasAnyPermission = (permissions: string[]): boolean => {
    return permissions.some((perm) => hasPermission(perm))
  }

  const hasAllPermissions = (permissions: string[]): boolean => {
    return permissions.every((perm) => hasPermission(perm))
  }

  const hasRole = (roleSlug: string): boolean => {
    return role?.slug === roleSlug
  }

  const isAdmin = (): boolean => {
    return hasRole('admin') || hasPermission('*')
  }

  return {
    hasPermission,
    hasAnyPermission,
    hasAllPermissions,
    hasRole,
    isAdmin,
    permissions: role?.permissions || [],
    role: role?.slug,
  }
}
```

**Uso em Componentes:**
```typescript
const { hasPermission, isAdmin } = usePermissions()

// Conditional logic
const handleDelete = () => {
  if (!hasPermission('users.delete')) {
    toast.error('No permission to delete')
    return
  }
  // Delete logic
}

// Conditional rendering
{isAdmin() && <AdminPanel />}
```

### 5.3 Can Component

**Conceito**: Componente declarativo para conditional rendering baseado em permissões.

**Props:**
```typescript
interface CanProps {
  permission?: string | string[]  // Single or multiple permissions
  any?: boolean                   // If true, need ANY permission (not all)
  admin?: boolean                 // Admin-only check
  not?: boolean                   // Inverse check
  children: React.ReactNode
}
```

**Implementação:**
```typescript
import { usePermissions } from '@shared/hooks/use-permissions'

export const Can = ({
  permission,
  any = false,
  admin = false,
  not = false,
  children
}: CanProps) => {
  const { hasPermission, hasAnyPermission, hasAllPermissions, isAdmin } = usePermissions()

  let hasAccess = false

  // Admin check
  if (admin) {
    hasAccess = isAdmin()
  }
  // Multiple permissions
  else if (Array.isArray(permission)) {
    hasAccess = any
      ? hasAnyPermission(permission)
      : hasAllPermissions(permission)
  }
  // Single permission
  else if (permission) {
    hasAccess = hasPermission(permission)
  }
  // No permission specified = show to all
  else {
    hasAccess = true
  }

  // Inverse check
  if (not) {
    hasAccess = !hasAccess
  }

  return hasAccess ? <>{children}</> : null
}
```

**Exemplos de Uso:**

```typescript
// 1. Single permission
<Can permission={PERMISSIONS.USERS_CREATE}>
  <Button onClick={handleCreate}>Create User</Button>
</Can>

// 2. Multiple permissions - need ANY
<Can permission={[PERMISSIONS.USERS_UPDATE, PERMISSIONS.USERS_DELETE]} any>
  <DropdownMenu>
    <DropdownMenuItem>Edit</DropdownMenuItem>
    <DropdownMenuItem>Delete</DropdownMenuItem>
  </DropdownMenu>
</Can>

// 3. Multiple permissions - need ALL
<Can permission={[PERMISSIONS.USERS_READ, PERMISSIONS.JOBS_READ]}>
  <Link to="/admin/reports">Reports</Link>
</Can>

// 4. Admin only
<Can admin>
  <Link to="/admin">Admin Panel</Link>
</Can>

// 5. Inverse check (show if DON'T have permission)
<Can permission={PERMISSIONS.USERS_DELETE} not>
  <Alert>You cannot delete users</Alert>
</Can>

// 6. Nested conditions
<Can permission={PERMISSIONS.USERS_READ}>
  <UsersTable users={users}>
    <Can permission={PERMISSIONS.USERS_UPDATE}>
      <Button>Edit</Button>
    </Can>
    <Can permission={PERMISSIONS.USERS_DELETE}>
      <Button variant="destructive">Delete</Button>
    </Can>
  </UsersTable>
</Can>
```

### 5.4 Protected Routes

**Conceito**: Proteger rotas inteiras baseado em autenticação (e opcionalmente permissões).

**ProtectedRoute Básico (apenas autenticação):**
```typescript
import { Navigate } from 'react-router-dom'
import { useAuthStore } from '@shared/stores/auth.store'

interface ProtectedRouteProps {
  children: React.ReactNode
}

export const ProtectedRoute = ({ children }: ProtectedRouteProps) => {
  const isAuthenticated = useAuthStore((state) => state.isAuthenticated)

  if (!isAuthenticated) {
    return <Navigate to="/login" replace />
  }

  return <>{children}</>
}
```

**ProtectedRoute Avançado (com permissões):**
```typescript
import { Navigate } from 'react-router-dom'
import { useAuthStore } from '@shared/stores/auth.store'
import { usePermissions } from '@shared/hooks/use-permissions'

interface ProtectedRouteProps {
  children: React.ReactNode
  permission?: string
  admin?: boolean
}

export const ProtectedRoute = ({
  children,
  permission,
  admin = false
}: ProtectedRouteProps) => {
  const isAuthenticated = useAuthStore((state) => state.isAuthenticated)
  const { hasPermission, isAdmin } = usePermissions()

  // Not authenticated
  if (!isAuthenticated) {
    return <Navigate to="/login" replace />
  }

  // Admin required
  if (admin && !isAdmin()) {
    return <Navigate to="/unauthorized" replace />
  }

  // Permission required
  if (permission && !hasPermission(permission)) {
    return <Navigate to="/unauthorized" replace />
  }

  return <>{children}</>
}
```

**Uso em Routes:**
```typescript
// routes.tsx
{
  path: '/users',
  element: (
    <ProtectedRoute permission={PERMISSIONS.USERS_READ}>
      <UsersListPage />
    </ProtectedRoute>
  ),
},
{
  path: '/admin',
  element: (
    <ProtectedRoute admin>
      <AdminDashboard />
    </ProtectedRoute>
  ),
}
```

**Fluxo de Autorização:**
```
1. User tenta acessar /users
2. ProtectedRoute verifica isAuthenticated
   - Se false → redirect /login
3. ProtectedRoute verifica permission (users.read)
   - Se false → redirect /unauthorized
4. Se ambos ok → renderiza UsersListPage
```

---

## 6. State Management

### 6.1 Zustand (Global State)

**Quando usar:**
- Auth/user information
- Theme (dark/light mode)
- App-level preferences
- Qualquer state que precisa persistir em localStorage

**Características:**
- Lightweight (1KB)
- Middleware de persist automático
- Selective subscription (performance)
- DevTools support

**Exemplo - Auth Store:**
```typescript
import { create } from 'zustand'
import { persist } from 'zustand/middleware'

interface AuthState {
  user: UserProfile | null
  accessToken: string | null
  refreshToken: string | null
  isAuthenticated: boolean

  setAuth: (data: { user, accessToken, refreshToken }) => void
  clearAuth: () => void
  updateUser: (user: UserProfile) => void
}

export const useAuthStore = create<AuthState>()(
  persist(
    (set) => ({
      user: null,
      accessToken: null,
      refreshToken: null,
      isAuthenticated: false,

      setAuth: (data) => set({
        user: data.user,
        accessToken: data.accessToken,
        refreshToken: data.refreshToken,
        isAuthenticated: true,
      }),

      clearAuth: () => set({
        user: null,
        accessToken: null,
        refreshToken: null,
        isAuthenticated: false,
      }),

      updateUser: (user) => set({ user }),
    }),
    {
      name: 'auth-storage', // localStorage key
    }
  )
)
```

**Uso:**
```typescript
// Selective subscription (apenas re-render quando user mudar)
const user = useAuthStore((state) => state.user)

// Actions
const { setAuth, clearAuth } = useAuthStore()
```

### 6.2 React Query (Server State)

**Quando usar:**
- Dados de API (users, jobs, etc.)
- Qualquer dado que vem do servidor
- Listas, detalhes, etc.

**Características:**
- Caching automático
- Refetching em background
- Invalidation após mutations
- Loading/error states
- Retry logic
- Pagination support

**Query Example:**
```typescript
export const useUsersList = (params: { page: number; limit: number }) => {
  return useQuery({
    queryKey: ['users', params], // Cache key
    queryFn: () => userService.getAll(params),
    staleTime: 5 * 60 * 1000, // 5 minutes
    retry: 1,
  })
}

// Uso
const { data, isLoading, error, refetch } = useUsersList({ page: 1, limit: 10 })
```

**Mutation Example:**
```typescript
export const useUserCreate = () => {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: (userData: UserRequest) => userService.create(userData),
    onSuccess: () => {
      // Invalidate e refetch
      queryClient.invalidateQueries({ queryKey: ['users'] })
      toast.success('User created!')
    },
  })
}

// Uso
const createMutation = useUserCreate()
createMutation.mutate(userData)
```

**Query Invalidation:**
```typescript
// Invalidar todas queries de users
queryClient.invalidateQueries({ queryKey: ['users'] })

// Invalidar query específica
queryClient.invalidateQueries({ queryKey: ['users', { id: '123' }] })

// Refetch imediato
queryClient.refetchQueries({ queryKey: ['users'] })
```

### 6.3 Local State (useState)

**Quando usar:**
- UI state (modais, dropdowns, tabs)
- State que não precisa persistir
- State local a um componente

**Exemplos:**
```typescript
// Modal state
const [isOpen, setIsOpen] = useState(false)

// Form state (gerenciado por React Hook Form)
const form = useForm()

// Filter state
const [searchTerm, setSearchTerm] = useState('')
const [selectedRole, setSelectedRole] = useState<string | null>(null)
```

**Quando NÃO usar:**
- Não duplique server state (use React Query)
- Não use para auth (use Zustand)
- Evite prop drilling excessivo (considere Context ou Zustand)

---

## 7. API Integration

### 7.1 Axios Client

**Configuração Base:**
```typescript
const apiClient = axios.create({
  baseURL: env.apiUrl, // 'http://localhost:8300'
  timeout: 30000,
  headers: {
    'Content-Type': 'application/json',
  },
})
```

**Request Interceptor (adiciona token):**
```typescript
apiClient.interceptors.request.use(
  (config) => {
    const { accessToken } = useAuthStore.getState()
    if (accessToken) {
      config.headers.Authorization = `Bearer ${accessToken}`
    }
    return config
  },
  (error) => Promise.reject(error)
)
```

**Response Interceptor (unwrap, token refresh):**
```typescript
apiClient.interceptors.response.use(
  (response) => {
    // Unwrap: { success: true, data: {...} } → {...}
    if (response.data?.data !== undefined) {
      return { ...response, data: response.data.data }
    }
    return response
  },
  async (error) => {
    // Token refresh on 401
    if (error.response?.status === 401 && !error.config._retry) {
      error.config._retry = true

      try {
        const newToken = await refreshToken()
        error.config.headers.Authorization = `Bearer ${newToken}`
        return apiClient(error.config) // Retry
      } catch {
        clearAuth()
        window.location.href = '/login'
      }
    }

    return Promise.reject(error)
  }
)
```

### 7.2 Services Layer

**Padrão de Service:**
```typescript
export const userService = {
  getAll: (params) => apiClient.get('/users', { params }),
  getById: (id) => apiClient.get(`/users/${id}`),
  create: (data) => apiClient.post('/users', data),
  update: (id, data) => apiClient.put(`/users/${id}`, data),
  delete: (id) => apiClient.delete(`/users/${id}`),
}
```

**Type-Safe Service:**
```typescript
import type { User, UserRequest, PaginatedResponse } from '@shared/types'

export const userService = {
  async getAll(params?: GetUsersParams): Promise<PaginatedResponse<User>> {
    const response = await apiClient.get<PaginatedResponse<User>>('/users', { params })
    return response.data
  },

  async getById(id: string): Promise<User> {
    const response = await apiClient.get<User>(`/users/${id}`)
    return response.data
  },

  async create(data: UserRequest): Promise<User> {
    const response = await apiClient.post<User>('/users', data)
    return response.data
  },
}
```

### 7.3 Error Handling

**Backend Error Response:**
```typescript
interface BackendErrorResponse {
  success: false
  message: string
  code?: string
  errors?: Record<string, string[]> // Field errors
}
```

**Handling em Mutation:**
```typescript
useMutation({
  mutationFn: userService.create,
  onError: (error: AxiosError<BackendErrorResponse>) => {
    const errorData = error.response?.data

    if (errorData?.errors) {
      // Field errors → map to form
      handleBackendErrors(errorData)
    } else {
      // Global error → toast
      toast.error(errorData?.message || 'An error occurred')
    }
  },
})
```

**useBackendErrors Hook:**
```typescript
export const useBackendErrors = (options: {
  setError: UseFormSetError<any>
  fieldMapping?: Record<string, string>
}) => {
  const handleBackendErrors = (errorResponse: BackendErrorResponse) => {
    if (!errorResponse.errors) return

    Object.entries(errorResponse.errors).forEach(([field, messages]) => {
      const formField = options.fieldMapping?.[field] || field
      options.setError(formField, {
        type: 'manual',
        message: messages[0],
      })
    })
  }

  return { handleBackendErrors }
}
```

---

## 8. Routing

### 8.1 React Router Setup

**Provider em main.tsx:**
```typescript
import { BrowserRouter } from 'react-router-dom'

<BrowserRouter>
  <App />
</BrowserRouter>
```

**Route Configuration:**
```typescript
import { Routes, Route } from 'react-router-dom'

<Routes>
  {/* Public */}
  <Route path="/" element={<HomePage />} />
  <Route path="/login" element={<LoginPage />} />

  {/* Protected */}
  <Route
    path="/dashboard"
    element={
      <ProtectedRoute>
        <DashboardPage />
      </ProtectedRoute>
    }
  />
</Routes>
```

**Lazy Loading:**
```typescript
import { lazy, Suspense } from 'react'

const UsersListPage = lazy(() => import('@features/users/pages/users-list-page'))

<Suspense fallback={<div>Loading...</div>}>
  <UsersListPage />
</Suspense>
```

### 8.2 Route Organization

**Por Feature:**
```
/users          → UsersListPage
/users/create   → UserCreatePage
/users/:id      → UserDetailPage
/users/:id/edit → UserEditPage

/jobs          → JobsListPage
/jobs/create   → JobCreatePage
/jobs/:id      → JobDetailPage
```

**Nested Routes com Layout:**
```typescript
<Route element={<AppLayout />}>
  <Route path="/dashboard" element={<DashboardPage />} />
  <Route path="/users" element={<UsersListPage />} />
  <Route path="/jobs" element={<JobsListPage />} />
</Route>
```

### 8.3 Navigation Patterns

**Programmatic Navigation:**
```typescript
import { useNavigate } from 'react-router-dom'

const navigate = useNavigate()

// Navigate após success
onSuccess: (user) => {
  toast.success('User created!')
  navigate(`/users/${user.id}`)
}

// Navigate back
navigate(-1)

// Navigate com replace (não adiciona ao history)
navigate('/login', { replace: true })
```

**Link Component:**
```typescript
import { Link } from 'react-router-dom'

<Link to="/users/create">
  <Button>Create User</Button>
</Link>
```

**NavLink (com active state):**
```typescript
import { NavLink } from 'react-router-dom'

<NavLink
  to="/users"
  className={({ isActive }) => isActive ? 'active' : ''}
>
  Users
</NavLink>
```

---

## 9. Form Handling

### 9.1 React Hook Form

**Setup:**
```typescript
import { useForm } from 'react-hook-form'
import { zodResolver } from '@hookform/resolvers/zod'

const form = useForm<FormData>({
  resolver: zodResolver(schema),
  defaultValues: { ... },
})
```

**FormField Pattern:**
```typescript
<FormField
  control={form.control}
  name="email"
  render={({ field }) => (
    <FormItem>
      <FormLabel>Email</FormLabel>
      <FormControl>
        <Input type="email" {...field} />
      </FormControl>
      <FormMessage /> {/* Shows validation errors */}
    </FormItem>
  )}
/>
```

**Submit Handler:**
```typescript
const onSubmit = (data: FormData) => {
  createMutation.mutate(data)
}

<form onSubmit={form.handleSubmit(onSubmit)}>
```

### 9.2 Zod Validation

**Schema Definition:**
```typescript
import { z } from 'zod'

export const userSchema = z.object({
  email: z.string().min(1, 'Required').email('Invalid email'),
  first_name: z.string().min(1, 'Required').max(100, 'Too long'),
  phone: z.string().max(20).optional().or(z.literal('')),
  password: z.string()
    .min(8, 'Min 8 characters')
    .regex(/[A-Z]/, 'Need uppercase')
    .regex(/[0-9]/, 'Need number'),
})
```

**Type Inference:**
```typescript
type UserFormData = z.infer<typeof userSchema>

// Agora UserFormData é:
// {
//   email: string
//   first_name: string
//   phone?: string
//   password: string
// }
```

**Custom Refinements:**
```typescript
z.object({
  password: z.string().min(8),
  confirm_password: z.string(),
}).refine((data) => data.password === data.confirm_password, {
  message: 'Passwords do not match',
  path: ['confirm_password'], // Error será exibido neste campo
})
```

### 9.3 Backend Error Mapping

**Problema**: Backend retorna erros com field names diferentes do form.

**Solução**:
```typescript
const { handleBackendErrors } = useBackendErrors({
  setError: form.setError,
  fieldMapping: {
    'backend_field': 'form_field', // Ex: 'user_type' → 'userType'
  },
})

createMutation.mutate(data, {
  onError: (error) => {
    if (error.response?.data?.errors) {
      handleBackendErrors(error.response.data)
    }
  },
})
```

**Fluxo:**
```
Backend error: { errors: { email: ['Email already exists'] } }
  ↓
handleBackendErrors() mapeia para form field
  ↓
form.setError('email', { message: 'Email already exists' })
  ↓
<FormMessage /> mostra erro embaixo do campo
```

---

## 10. Styling & Design System

### 10.1 Tailwind CSS

**Utility-First Approach:**
```tsx
<button className="bg-primary text-white px-4 py-2 rounded-md hover:bg-primary/90">
  Click me
</button>
```

**Responsive Design:**
```tsx
<div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
```

**Dark Mode:**
```tsx
<div className="bg-white dark:bg-gray-900 text-black dark:text-white">
```

### 10.2 Design System (CSS Variables)

**index.css:**
```css
@layer base {
  :root {
    /* Colors (HSL) */
    --background: 0 0% 100%;
    --foreground: 217 19% 27%;
    --primary: 217 33% 17%;
    --accent: 217 83% 56%;
    --success: 158 84% 39%;
    --error: 0 84% 60%;

    /* Shadows */
    --shadow-sm: 0 2px 8px hsl(217 19% 27% / 0.06);
    --shadow-md: 0 4px 16px hsl(217 19% 27% / 0.1);

    /* Gradients */
    --gradient-primary: linear-gradient(135deg, ...);
  }

  .dark {
    --background: 217 33% 10%;
    --foreground: 210 20% 98%;
    /* ... */
  }
}
```

**Uso em Tailwind:**
```tsx
<div className="bg-background text-foreground">
<div className="bg-primary text-white">
```

### 10.3 UI Component Library

**Base: Radix UI + Tailwind Styling**

Componentes disponíveis:
- Button, Input, Select, Checkbox
- Card, Dialog, DropdownMenu
- Form components (Form, FormField, FormItem, FormLabel, FormMessage)
- Avatar, Badge, Separator
- Skeleton, Spinner
- Pagination

**Por que Radix UI?**
- Acessibilidade built-in
- Unstyled (você estiliza com Tailwind)
- Keyboard navigation
- ARIA attributes automáticos

### 10.4 Utility Functions

**cn() - Class Name Merge:**
```typescript
import { cn } from '@shared/lib/utils'

<Button className={cn(
  'bg-primary', // Base
  isLoading && 'opacity-50', // Conditional
  className // User override
)} />
```

**Por que usar?**
- Merge classes corretamente
- Sobrescreve classes conflitantes (`bg-red-500` sobrescreve `bg-blue-500`)
- Suporta conditional classes

---

## 11. TypeScript Configuration

### 11.1 Path Aliases

**tsconfig.json:**
```json
{
  "compilerOptions": {
    "baseUrl": ".",
    "paths": {
      "@/*": ["./src/*"],
      "@app/*": ["./src/app/*"],
      "@features/*": ["./src/features/*"],
      "@shared/*": ["./src/shared/*"],
      "@infrastructure/*": ["./src/infrastructure/*"]
    }
  }
}
```

**Vite Config:**
```typescript
export default defineConfig({
  resolve: {
    alias: {
      '@': path.resolve(__dirname, './src'),
      '@app': path.resolve(__dirname, './src/app'),
      '@features': path.resolve(__dirname, './src/features'),
      '@shared': path.resolve(__dirname, './src/shared'),
      '@infrastructure': path.resolve(__dirname, './src/infrastructure'),
    },
  },
})
```

**Uso:**
```typescript
import { Button } from '@shared/components/ui/button'
import { useUsersList } from '@features/users/hooks'
import { userService } from '@infrastructure/api/services/user.service'
```

### 11.2 Type Safety

**Strict Mode:**
```json
{
  "strict": true,
  "noUnusedLocals": true,
  "noUnusedParameters": true,
  "noFallthroughCasesInSwitch": true,
}
```

**Type Inference de Zod:**
```typescript
const schema = z.object({ name: z.string() })
type FormData = z.infer<typeof schema> // { name: string }
```

**Generic Types:**
```typescript
interface PaginatedResponse<T> {
  data: T[]
  total: number
}

const users: PaginatedResponse<User> = await userService.getAll()
```

### 11.3 Type Organization

**Shared Types:**
- `shared/types/auth.types.ts`
- `shared/types/user.types.ts`
- `shared/types/api.types.ts`

**Feature Types:**
- `features/users/types/` (se houver types específicos)

**Service Types:**
- Co-located com services em `infrastructure/api/services/`

---

## 12. Build & Development

### 12.1 Vite Configuration

**Fast HMR:**
- Hot Module Replacement em < 100ms
- Não precisa rebuild completo

**Build Optimization:**
- Code splitting automático
- Tree shaking
- Minificação

**Environment Variables:**
```typescript
// .env
VITE_API_URL=http://localhost:8300

// Uso
import.meta.env.VITE_API_URL
```

### 12.2 Development Scripts

**package.json:**
```json
{
  "scripts": {
    "dev": "vite",                    // Desenvolvimento local
    "build": "tsc && vite build",     // Build de produção
    "lint": "eslint . --ext ts,tsx",  // ESLint check
    "preview": "vite preview",        // Preview de build
    "test": "vitest",                 // Testes
  }
}
```

**Comandos:**
```bash
npm run dev      # Inicia dev server (localhost:3000)
npm run build    # Build para produção (pasta dist/)
npm run lint     # Verifica erros de linting
npm run preview  # Preview do build
```

---

## 13. Checklist: Criando uma Nova Feature

Use este template para adicionar uma nova feature (exemplo: "Products").

### Passo 1: Infrastructure Layer

- [ ] **Criar Types** (`shared/types/product.types.ts`)
  ```typescript
  export interface Product {
    id: string
    name: string
    description: string
    price: number
    created_at: string
  }

  export interface ProductRequest {
    name: string
    description: string
    price: number
  }

  export interface ProductResponse extends Product {}
  ```

- [ ] **Criar Service** (`infrastructure/api/services/product.service.ts`)
  ```typescript
  import { apiClient } from '../axios-client'
  import type { Product, ProductRequest, PaginatedResponse } from '@shared/types'

  export const productService = {
    getAll: (params) => apiClient.get<PaginatedResponse<Product>>('/products', { params }),
    getById: (id: string) => apiClient.get<Product>(`/products/${id}`),
    create: (data: ProductRequest) => apiClient.post<Product>('/products', data),
    update: (id: string, data: Partial<ProductRequest>) =>
      apiClient.put<Product>(`/products/${id}`, data),
    delete: (id: string) => apiClient.delete(`/products/${id}`),
  }
  ```

### Passo 2: Shared Layer (se aplicável)

- [ ] **Adicionar Permissões** (`shared/constants/permissions.ts`)
  ```typescript
  export const PERMISSIONS = {
    // ... existing
    PRODUCTS_CREATE: 'products.create',
    PRODUCTS_READ: 'products.read',
    PRODUCTS_UPDATE: 'products.update',
    PRODUCTS_DELETE: 'products.delete',
    PRODUCTS_MANAGE: 'products.*',
  }
  ```

- [ ] **(Opcional) Criar hooks compartilhados** se necessário

### Passo 3: Feature Layer

- [ ] **Criar diretório** `features/products/`

- [ ] **Criar Schemas** (`features/products/schemas/product-form.schema.ts`)
  ```typescript
  import { z } from 'zod'

  export const createProductSchema = z.object({
    name: z.string().min(1, 'Name is required').max(100),
    description: z.string().min(1, 'Description is required'),
    price: z.number().min(0, 'Price must be positive'),
  })

  export const updateProductSchema = createProductSchema.partial()

  export type CreateProductFormData = z.infer<typeof createProductSchema>
  export type UpdateProductFormData = z.infer<typeof updateProductSchema>
  ```

- [ ] **Criar Hooks** (`features/products/hooks/`)

  **use-products-list.ts:**
  ```typescript
  import { useQuery } from '@tanstack/react-query'
  import { productService } from '@infrastructure/api/services/product.service'

  export const useProductsList = (params = {}) => {
    return useQuery({
      queryKey: ['products', params],
      queryFn: () => productService.getAll(params),
      staleTime: 5 * 60 * 1000,
    })
  }
  ```

  **use-product-create.ts:**
  ```typescript
  import { useMutation, useQueryClient } from '@tanstack/react-query'
  import { toast } from 'sonner'
  import { productService } from '@infrastructure/api/services/product.service'

  export const useProductCreate = () => {
    const queryClient = useQueryClient()

    return useMutation({
      mutationFn: productService.create,
      onSuccess: () => {
        queryClient.invalidateQueries({ queryKey: ['products'] })
        toast.success('Product created!')
      },
    })
  }
  ```

  **index.ts** (barrel export):
  ```typescript
  export * from './use-products-list'
  export * from './use-product-detail'
  export * from './use-product-create'
  export * from './use-product-update'
  export * from './use-product-delete'
  ```

- [ ] **Criar Components** (`features/products/components/`)

  **product-form.tsx:**
  ```typescript
  import { useForm } from 'react-hook-form'
  import { zodResolver } from '@hookform/resolvers/zod'
  import { Button, Form, FormField, Input } from '@shared/components/ui'
  import { createProductSchema, type CreateProductFormData } from '../schemas'

  interface ProductFormProps {
    onSubmit: (data: CreateProductFormData) => void
    defaultValues?: Partial<CreateProductFormData>
    isLoading?: boolean
  }

  export const ProductForm = ({ onSubmit, defaultValues, isLoading }: ProductFormProps) => {
    const form = useForm<CreateProductFormData>({
      resolver: zodResolver(createProductSchema),
      defaultValues,
    })

    return (
      <Form {...form}>
        <form onSubmit={form.handleSubmit(onSubmit)}>
          <FormField name="name" control={form.control} ... />
          <FormField name="description" control={form.control} ... />
          <FormField name="price" control={form.control} ... />
          <Button type="submit" disabled={isLoading}>Save</Button>
        </form>
      </Form>
    )
  }
  ```

  **products-table.tsx**, **product-detail-card.tsx**, **index.ts**

- [ ] **Criar Pages** (`features/products/pages/`)

  **products-list-page.tsx:**
  ```typescript
  import { useState } from 'react'
  import { Link } from 'react-router-dom'
  import { Button, Card } from '@shared/components/ui'
  import { Can } from '@shared/components/can'
  import { PERMISSIONS } from '@shared/constants/permissions'
  import { ProductsTable } from '../components'
  import { useProductsList } from '../hooks'

  export const ProductsListPage = () => {
    const [page, setPage] = useState(1)
    const { data, isLoading } = useProductsList({ page })

    return (
      <div className="space-y-4">
        <div className="flex justify-between">
          <h1 className="text-2xl font-bold">Products</h1>
          <Can permission={PERMISSIONS.PRODUCTS_CREATE}>
            <Link to="/products/create">
              <Button>Create Product</Button>
            </Link>
          </Can>
        </div>
        <Card>
          <ProductsTable products={data?.data || []} />
        </Card>
      </div>
    )
  }
  ```

  **product-create-page.tsx**, **product-edit-page.tsx**, **product-detail-page.tsx**, **index.ts**

### Passo 4: App Layer

- [ ] **Adicionar Rotas** (`app/routes.tsx`)
  ```typescript
  import { lazy } from 'react'

  const ProductsListPage = lazy(() => import('@features/products/pages/products-list-page'))
  const ProductCreatePage = lazy(() => import('@features/products/pages/product-create-page'))
  // ...

  // In routes:
  {
    path: '/products',
    element: <ProtectedRoute><ProductsListPage /></ProtectedRoute>
  },
  {
    path: '/products/create',
    element: <ProtectedRoute><ProductCreatePage /></ProtectedRoute>
  },
  {
    path: '/products/:id',
    element: <ProtectedRoute><ProductDetailPage /></ProtectedRoute>
  },
  {
    path: '/products/:id/edit',
    element: <ProtectedRoute><ProductEditPage /></ProtectedRoute>
  }
  ```

- [ ] **Adicionar navegação** em sidebar/header (se aplicável)

### Passo 5: Integração com RBAC

- [ ] **Proteger rotas** com ProtectedRoute
- [ ] **Usar Can component** para botões/ações
  ```typescript
  <Can permission={PERMISSIONS.PRODUCTS_CREATE}>
    <Button>Create Product</Button>
  </Can>

  <Can permission={PERMISSIONS.PRODUCTS_UPDATE}>
    <Button>Edit</Button>
  </Can>

  <Can permission={PERMISSIONS.PRODUCTS_DELETE}>
    <Button variant="destructive">Delete</Button>
  </Can>
  ```

### Passo 6: Testes (opcional mas recomendado)

- [ ] **Unit tests** para hooks (React Query)
- [ ] **Component tests** para forms e tables
- [ ] **Integration tests** para fluxos completos (create, edit, delete)

---

## 14. Stack Tecnológico Recomendado

| Componente | Tecnologia | Versão | Justificativa |
|------------|------------|--------|---------------|
| **Framework** | React | 18+ | Componentes, Hooks, Concurrent features |
| **Build Tool** | Vite | 5+ | HMR rápido, builds otimizados |
| **Language** | TypeScript | 5+ | Type safety, better DX |
| **Routing** | React Router | 6+ | Navegação declarativa, layouts |
| **State (Global)** | Zustand | 4+ | Lightweight, persist, DevTools |
| **State (Server)** | TanStack React Query | 5+ | Caching, refetching, mutations |
| **Forms** | React Hook Form | 7+ | Performance, validation |
| **Validation** | Zod | 3+ | Type-safe schemas, inference |
| **HTTP Client** | Axios | 1+ | Interceptors, error handling |
| **Styling** | Tailwind CSS | 3+ | Utility-first, responsive |
| **UI Components** | Radix UI | latest | Acessibilidade, unstyled |
| **Testing** | Vitest + Testing Library | latest | Testes rápidos, compatíveis |

---

## 15. Boas Práticas e Recomendações

### 15.1 Feature-First Organization

**Regra**: Organize por features, não por tipo de arquivo.

**Por quê?**
- Features auto-contidas
- Fácil adicionar/remover
- Reduz acoplamento
- Código relacionado fica junto

**Evite:**
```
src/
├── components/
│   ├── UserForm.tsx
│   ├── JobForm.tsx
├── hooks/
│   ├── useUsers.ts
│   └── useJobs.ts
```

**Prefira:**
```
src/features/
├── users/
│   ├── components/UserForm.tsx
│   └── hooks/useUsers.ts
├── jobs/
│   ├── components/JobForm.tsx
│   └── hooks/useJobs.ts
```

### 15.2 Separation of Concerns

**Regra**: Cada camada tem uma responsabilidade única.

- **UI Components**: Apenas apresentação
- **Hooks**: Lógica de negócio e data fetching
- **Services**: Chamadas de API
- **Types**: Contratos entre camadas

**NÃO faça:**
```typescript
// Component fazendo API call direto
const MyComponent = () => {
  const [data, setData] = useState([])

  useEffect(() => {
    axios.get('/users').then(res => setData(res.data))
  }, [])
}
```

**FAÇA:**
```typescript
// Component usa hook
const MyComponent = () => {
  const { data } = useUsersList()
}

// Hook usa service
const useUsersList = () => {
  return useQuery({
    queryKey: ['users'],
    queryFn: userService.getAll,
  })
}
```

### 15.3 Type Safety

**Regra**: Use TypeScript estrito.

- Evite `any` - use `unknown` se necessário
- Infira tipos de schemas Zod quando possível
- Use type guards para narrowing
- Generics para reutilização

**Exemplos:**
```typescript
// BOM
const schema = z.object({ name: z.string() })
type FormData = z.infer<typeof schema>

// RUIM
type FormData = any
```

### 15.4 State Management

**Regra**: Use a ferramenta certa para cada tipo de state.

- **Global**: Zustand (auth, theme)
- **Server**: React Query (dados de API)
- **Local**: useState (UI state)

**NÃO duplique server state no Zustand:**
```typescript
// RUIM
const users = useAuthStore((state) => state.users) // NÃO!

// BOM
const { data: users } = useUsersList() // React Query
```

### 15.5 Error Handling

**Regra**: Centralize error handling.

- Interceptors para errors globais
- Toast para mensagens de erro
- Backend error mapping para forms
- Fallback UI para errors

**Exemplo:**
```typescript
// Interceptor
apiClient.interceptors.response.use(
  (response) => response,
  (error) => {
    if (error.response?.status === 500) {
      toast.error('Server error')
    }
    return Promise.reject(error)
  }
)
```

### 15.6 Performance

**Regras:**
- Lazy load pages
- Memoize componentes pesados (React.memo)
- Use React Query caching
- Debounce search inputs
- Code splitting automático

**Exemplo:**
```typescript
// Lazy loading
const UsersPage = lazy(() => import('@features/users/pages/users-list-page'))

// Memoization
const UserCard = React.memo(({ user }) => {
  // Expensive component
})

// Debounce
const [search, setSearch] = useState('')
const debouncedSearch = useDebounce(search, 500)
```

### 15.7 Accessibility

**Regras:**
- Use Radix UI (acessibilidade built-in)
- ARIA labels onde necessário
- Keyboard navigation
- Semantic HTML
- Contrast ratios adequados

**Exemplo:**
```tsx
<button aria-label="Delete user">
  <TrashIcon />
</button>
```

### 15.8 Code Organization

**Regras:**
- Barrel exports (index.ts) para imports limpos
- Co-locate related files
- Consistent naming conventions (kebab-case para arquivos)
- Path aliases para imports absolutos

**Exemplo:**
```typescript
// Barrel export
// features/users/hooks/index.ts
export * from './use-users-list'
export * from './use-user-create'

// Uso
import { useUsersList, useUserCreate } from '@features/users/hooks'
```

### 15.9 Form Handling

**Regras:**
- React Hook Form para performance
- Zod para validação type-safe
- Backend error mapping para UX
- Controlled components quando necessário

**Exemplo:**
```typescript
const form = useForm<FormData>({
  resolver: zodResolver(schema),
})

const { handleBackendErrors } = useBackendErrors({
  setError: form.setError,
})
```

### 15.10 API Integration

**Regras:**
- Services layer para abstração
- TypeScript interfaces para contracts
- Error handling centralizado
- Token refresh automático

**Exemplo:**
```typescript
// Type-safe service
export const userService = {
  getAll: (): Promise<User[]> => apiClient.get('/users'),
}

// Error handling no interceptor
// Token refresh no interceptor
```

---

## Conclusão

Esta arquitetura combina os melhores princípios de **Feature-First Organization** e **Layered Architecture** para criar aplicações React escaláveis, testáveis e manuteníveis.

**Principais Takeaways:**

1. **Feature-First**: Organize por features, não por tipo de arquivo
2. **Layered**: Camadas com responsabilidades claras (App → Features → Shared → Infrastructure)
3. **State Management**: Ferramenta certa para cada tipo (Zustand, React Query, useState)
4. **Type Safety**: TypeScript estrito + Zod para validação
5. **Separation of Concerns**: UI, logic, API calls separados
6. **RBAC**: Permissões granulares com hook + componente

**Quando usar esta arquitetura:**
- Aplicações React de médio a grande porte
- Sistemas com múltiplas features
- Aplicações que precisam escalar
- Projetos que requerem alta testabilidade
- Apps com controle de acesso granular (RBAC)

**Quando NÃO usar:**
- Protótipos rápidos ou MVPs muito simples
- Landing pages estáticas
- Aplicações com apenas 1-2 telas

Use o **Checklist de Features** (seção 13) como guia prático para adicionar novas funcionalidades seguindo todos os padrões descritos neste documento.
