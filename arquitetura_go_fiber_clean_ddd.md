# Arquitetura Go/Fiber: Clean Architecture + DDD

## 1. Introdução

Este documento descreve uma arquitetura de backend robusta e escalável para APIs REST em Go utilizando o framework Fiber, baseada nos princípios de **Clean Architecture** e **Domain-Driven Design (DDD)**.

### 1.1 Visão Geral

A arquitetura é organizada em camadas concêntricas onde cada camada tem responsabilidades bem definidas e as dependências fluem de fora para dentro, sempre em direção ao domínio. Isso garante que a lógica de negócio permaneça isolada de frameworks, bancos de dados e detalhes de infraestrutura.

```
┌─────────────────────────────────────────┐
│   Entrypoint (HTTP - Fiber)             │  ← Camada mais externa
│   - Controllers                          │
│   - Middleware                           │
│   - Routes                               │
│   - Payloads (DTOs)                      │
├─────────────────────────────────────────┤
│   Application (Use Cases)               │
│   - Orquestração da lógica de negócio   │
│   - Coordenação entre gateways          │
├─────────────────────────────────────────┤
│   Core Domain (Coração da Aplicação)   │  ← Camada mais interna
│   - Entities (Objetos ricos)            │
│   - Gateways (Interfaces)               │
│   - Services de Domínio                 │
│   - Value Objects                        │
│   - Errors & Constants                  │
├─────────────────────────────────────────┤
│   Data Provider (Persistência)          │
│   - Models (GORM)                        │
│   - Gateway Implementations             │
│   - Mappers                              │
└─────────────────────────────────────────┘
```

### 1.2 Princípios Fundamentais

**Inversão de Dependência**
- Camadas externas dependem de camadas internas
- O domínio define interfaces (gateways)
- A infraestrutura implementa essas interfaces
- Nunca o contrário

**Separação de Responsabilidades**
- Cada camada tem um papel claro e único
- Não há mistura de lógica de negócio com infraestrutura
- HTTP não conhece banco de dados diretamente
- Use cases não conhecem frameworks

**Testabilidade Independente**
- Cada camada pode ser testada isoladamente
- Mocks facilitados pelas interfaces
- Lógica de negócio testável sem banco de dados
- Controllers testáveis sem servidor HTTP

### 1.3 Vantagens

- **Manutenibilidade**: Código organizado e fácil de navegar
- **Escalabilidade**: Fácil adicionar novas features sem quebrar existentes
- **Testabilidade**: Alta cobertura de testes unitários e integração
- **Independência de Frameworks**: Trocar Fiber por outro framework não afeta o domínio
- **Independência de Banco**: Trocar PostgreSQL por MongoDB é viável
- **Reusabilidade**: Use cases podem ser reutilizados (HTTP, CLI, gRPC)
- **Clareza**: Desenvolvedor sabe exatamente onde colocar código novo

---

## 2. Estrutura de Diretórios

```
backend/
├── cmd/api/                    # Entry point da aplicação
│   └── main.go                # Inicialização, migrations, seeders
│
├── config/                     # Configuração & Dependency Injection
│   ├── config.go              # Estrutura de configuração (env vars)
│   ├── dependency_injection.go # Container principal
│   └── dependency_injection_*.go # Containers por feature
│
├── core/                       # CAMADA DE DOMÍNIO (Clean Architecture)
│   ├── domain/
│   │   ├── entity/            # Objetos ricos de negócio (User, Role, etc.)
│   │   ├── gateway/           # Interfaces de repositórios
│   │   ├── service/           # Serviços de domínio (HasherService, JWTService)
│   │   ├── errors/            # Erros semânticos do negócio
│   │   ├── constants/         # Constantes do sistema (permissions, roles)
│   │   ├── valueobject/       # Value objects imutáveis (Pagination, SortOrder)
│   │   └── repository/        # Criteria para queries dinâmicas
│   └── validation/            # Validação customizada (wrappers, validators)
│
├── application/               # CAMADA DE APLICAÇÃO (Use Cases)
│   └── usecase/              # Casos de uso organizados por feature
│       ├── auth/             # Login, RefreshToken, etc.
│       ├── users/            # CreateUser, ListUsers, UpdateUser, etc.
│       ├── roles/            # CreateRole, AssignPermissions, etc.
│       └── ...
│
├── dataprovider/             # CAMADA DE DADOS (Persistência)
│   └── database/
│       ├── model/            # Models GORM (representação de tabelas)
│       ├── gateway/          # Implementações dos gateways (GORM-based)
│       └── mapper/           # Conversão Model ↔ Entity
│
├── entrypoint/http/          # CAMADA DE ENTRADA (HTTP/Fiber)
│   ├── controller/           # Handlers HTTP (parse request, call use case)
│   ├── middleware/           # Middleware (JWT, CORS, RateLimit, etc.)
│   ├── routes/               # Definição de rotas por feature
│   └── payload/              # DTOs de Request/Response
│       ├── request/
│       └── response/
│
├── infrastructure/           # Implementações de serviços técnicos
│   ├── auth/                # Implementação de JWTService
│   └── hasher/              # Implementação de HasherService (bcrypt)
│
├── migrations/              # Migrations do banco de dados
│   ├── up/                  # Scripts de criação
│   └── down/                # Scripts de rollback
│
├── seeders/                 # Dados iniciais (roles, permissions, test users)
│
└── docs/                    # Documentação da API (Swagger, guias)
```

### 2.1 Fluxo de Dependências

```
HTTP Request
    ↓
Controller (entrypoint/http/controller)
    ↓
Use Case (application/usecase)
    ↓
Gateway Interface (core/domain/gateway) ← Define o contrato
    ↓
Gateway Implementation (dataprovider/database/gateway) ← Implementa o contrato
    ↓
Database (PostgreSQL + GORM)
```

**Importante**: As dependências SEMPRE fluem de fora para dentro. O domínio nunca importa nada de `entrypoint`, `dataprovider` ou `infrastructure`.

---

## 3. Camadas da Arquitetura

### 3.1 Camada de Domínio (Core)

Esta é a camada mais importante e mais interna da aplicação. Contém as regras de negócio puras e não depende de nada externo.

#### 3.1.1 Entities (`core/domain/entity/`)

**Conceito**: Entities são objetos ricos que representam conceitos centrais do negócio. Não são apenas estruturas de dados (DTOs), mas sim objetos com comportamento e lógica.

**Características:**
- **Campos privados** com getters públicos (encapsulamento)
- **Construtores** para criação (`NewUser`, `NewRole`)
- **Métodos de reconstrução** para carregar do banco (`ReconstructUser`, `ReconstructRole`)
- **Lógica de negócio** em métodos (validação, regras, comportamentos)
- **Imutabilidade** onde faz sentido
- **Identidade única** (geralmente UUID)

**Exemplo Conceitual:**
```go
// Entity User
type User struct {
    id        uuid.UUID      // Privado - só acessível via getter
    email     string
    password  string
    roleID    uuid.UUID
    role      *Role         // Lazy loading
    createdAt time.Time
    updatedAt time.Time
    deletedAt *time.Time    // Soft delete
}

// Constructor: cria nova entidade
func NewUser(email, password string, roleID uuid.UUID) (*User, error) {
    user := &User{
        id:        uuid.New(),
        email:     email,
        password:  password,
        roleID:    roleID,
        createdAt: time.Now(),
        updatedAt: time.Now(),
    }
    if err := user.Validate(); err != nil {
        return nil, err
    }
    return user, nil
}

// Reconstruction: reconstrói do banco de dados
func ReconstructUser(id uuid.UUID, email, password string, ...) (*User, error) {
    return &User{
        id:        id,
        email:     email,
        password:  password,
        createdAt: createdAt,
        updatedAt: updatedAt,
        deletedAt: deletedAt,
    }, nil
}

// Getters: encapsulamento
func (u *User) ID() uuid.UUID { return u.id }
func (u *User) Email() string { return u.email }

// Lógica de negócio
func (u *User) Validate() error {
    if u.email == "" {
        return errors.New("email is required")
    }
    if !isValidEmail(u.email) {
        return errors.New("invalid email format")
    }
    return nil
}

func (u *User) SoftDelete() {
    now := time.Now()
    u.deletedAt = &now
}

func (u *User) IsActive() bool {
    return u.deletedAt == nil
}

func (u *User) HasPermission(permSlug string) bool {
    if u.role == nil {
        return false
    }
    return u.role.HasPermission(permSlug)
}
```

**Responsabilidades:**
- Validar seus próprios dados
- Executar regras de negócio sobre si mesmo
- Manter consistência interna
- Expor comportamento através de métodos

**NÃO são responsáveis por:**
- Persistência (isso é do gateway)
- Conversão de/para formatos externos (isso é do mapper)
- Lógica de aplicação (isso é do use case)

#### 3.1.2 Gateways (`core/domain/gateway/`)

**Conceito**: Gateways são interfaces que definem como o domínio acessa o mundo externo (banco de dados, APIs, etc.). Representam a **inversão de dependência**.

**Características:**
- São apenas **interfaces**, não implementações
- Definidas no domínio
- Implementadas na camada de dados
- Permitem trocar implementações sem afetar o domínio

**Exemplo Conceitual:**
```go
// Interface definida no domínio
type UserGateway interface {
    // CRUD
    Save(ctx context.Context, user *entity.User) error
    FindByID(ctx context.Context, id uuid.UUID) (*entity.User, error)
    FindByEmail(ctx context.Context, email string) (*entity.User, error)
    Update(ctx context.Context, user *entity.User) error
    Delete(ctx context.Context, id uuid.UUID) error

    // Queries
    FindByCriteria(ctx context.Context, criteria criteria.Criteria,
                   pagination valueobject.Pagination,
                   sortOrder valueobject.SortOrder) ([]*entity.User, error)
    CountByCriteria(ctx context.Context, criteria criteria.Criteria) (int64, error)

    // Utilities
    ExistsByID(ctx context.Context, id uuid.UUID) (bool, error)
    ExistsByEmail(ctx context.Context, email string) (bool, error)
}
```

**Responsabilidades:**
- Definir o contrato de acesso a dados
- Fornecer operações necessárias ao domínio
- Abstrair a tecnologia de persistência

**Por que usar?**
- Use cases não sabem se é PostgreSQL, MongoDB ou in-memory
- Facilita testes (mock do gateway)
- Permite trocar banco de dados sem mudar use cases

#### 3.1.3 Services de Domínio (`core/domain/service/`)

**Conceito**: Serviços de domínio encapsulam lógica de negócio que não pertence a nenhuma entidade específica ou que envolve múltiplas entidades.

**Exemplos Comuns:**
- `HasherService`: Hash e verificação de senhas (bcrypt, argon2)
- `JWTService`: Geração e validação de tokens
- `EmailService`: Envio de emails (interface definida aqui)

**Características:**
- São **stateless** (sem estado interno)
- Geralmente são **interfaces** (implementadas na infra)
- Lógica pura de domínio, não infraestrutura

**Exemplo:**
```go
// Interface no domínio
type HasherService interface {
    Hash(password string) (string, error)
    Compare(hashedPassword, password string) error
}

// Interface no domínio
type JWTService interface {
    GenerateAccessToken(userID, roleID uuid.UUID, email string) (string, error)
    GenerateRefreshToken(userID uuid.UUID) (string, error)
    ValidateToken(token string, tokenType TokenType) (*TokenClaims, error)
}
```

**Implementação**: Fica em `infrastructure/auth/` ou `infrastructure/hasher/`.

#### 3.1.4 Value Objects (`core/domain/valueobject/`)

**Conceito**: Value Objects são objetos imutáveis que representam valores sem identidade própria. Dois value objects com os mesmos atributos são considerados iguais.

**Características:**
- **Imutáveis**: Não mudam após criação
- **Sem identidade**: Não têm ID, comparados por valor
- **Validação**: Sempre em estado válido

**Exemplos:**
```go
// Pagination: representa paginação de forma reutilizável
type Pagination struct {
    Page     int
    PageSize int
}

func NewPagination(page, pageSize int) Pagination {
    if page < 1 {
        page = 1
    }
    if pageSize < 1 || pageSize > 100 {
        pageSize = 10
    }
    return Pagination{Page: page, PageSize: pageSize}
}

func (p Pagination) Offset() int {
    return (p.Page - 1) * p.PageSize
}

// SortOrder: representa ordenação
type SortOrder struct {
    Field     string
    Direction string // "asc" ou "desc"
}
```

#### 3.1.5 Errors (`core/domain/errors/`)

**Conceito**: Erros semânticos específicos do domínio, não erros técnicos.

**Exemplo:**
```go
var (
    ErrUserNotFound           = errors.New("user not found")
    ErrUserEmailExists        = errors.New("user email already exists")
    ErrUserInvalidCredentials = errors.New("invalid credentials")
    ErrUserInactive           = errors.New("user is inactive")
)
```

**Por que usar?**
- Comunicação clara de problemas de negócio
- Controllers podem mapear para HTTP status codes
- Facilita tratamento específico

#### 3.1.6 Constants (`core/domain/constants/`)

**Conceito**: Valores fixos do sistema que fazem parte das regras de negócio.

**Exemplos:**
```go
// Permissions
const (
    CreateUsers = "users.create"
    ReadUsers   = "users.read"
    UpdateUsers = "users.update"
    DeleteUsers = "users.delete"
    ManageUsers = "users.*"  // Wildcard
)

// Role Slugs
const (
    AdminRole     = "admin"
    PartnerRole   = "partner"
    RecruiterRole = "recruiter"
)
```

#### 3.1.7 Validation (`core/validation/`)

**Conceito**: Wrapper sobre bibliotecas de validação com regras customizadas de domínio.

**Características:**
- Usa `go-playground/validator/v10` internamente
- Adiciona validações customizadas (permission_slug, password strength, etc.)
- Centraliza configuração de validação

---

### 3.2 Camada de Aplicação (Use Cases)

**Conceito**: Use cases orquestram a lógica de negócio, coordenando entities, gateways e services para executar uma operação específica do sistema.

#### 3.2.1 Estrutura de um Use Case

Cada use case é um arquivo separado com uma struct e método `Execute()`.

**Exemplo Conceitual:**
```go
// Arquivo: application/usecase/users/create_user.go

// Dados de entrada
type CreateUserData struct {
    Email    string
    Password string
    RoleID   uuid.UUID
}

// Use Case struct
type CreateUser struct {
    userGateway   gateway.UserGateway      // Dependência de gateway
    roleGateway   gateway.RoleGateway
    hasherService service.HasherService    // Dependência de service
}

// Constructor com injeção de dependências
func NewCreateUser(
    userGateway gateway.UserGateway,
    roleGateway gateway.RoleGateway,
    hasherService service.HasherService,
) *CreateUser {
    return &CreateUser{
        userGateway:   userGateway,
        roleGateway:   roleGateway,
        hasherService: hasherService,
    }
}

// Método de execução
func (uc *CreateUser) Execute(ctx context.Context, data CreateUserData) (*entity.User, error) {
    // 1. Validar se email já existe
    exists, err := uc.userGateway.ExistsByEmail(ctx, data.Email)
    if err != nil {
        return nil, err
    }
    if exists {
        return nil, domainErrors.ErrUserEmailExists
    }

    // 2. Validar se role existe
    _, err = uc.roleGateway.FindByID(ctx, data.RoleID)
    if err != nil {
        return nil, errors.New("role not found")
    }

    // 3. Hash da senha usando domain service
    hashedPassword, err := uc.hasherService.Hash(data.Password)
    if err != nil {
        return nil, err
    }

    // 4. Criar entidade (validação interna)
    user, err := entity.NewUser(data.Email, hashedPassword, data.RoleID)
    if err != nil {
        return nil, err
    }

    // 5. Persistir via gateway
    if err := uc.userGateway.Save(ctx, user); err != nil {
        return nil, err
    }

    return user, nil
}
```

#### 3.2.2 Responsabilidades dos Use Cases

**Fazem:**
- Orquestrar fluxo de negócio
- Validar regras de aplicação (email único, role existe)
- Chamar gateways e services
- Compor entities
- Tratar erros e retornar resultados

**NÃO fazem:**
- Validação de formato (isso é do DTO/payload)
- Lógica de persistência (isso é do gateway)
- Parsing de HTTP request (isso é do controller)
- Regras internas de entidade (isso é da entity)

#### 3.2.3 Organização

Use cases são organizados por feature em diretórios separados:

```
application/usecase/
├── users/
│   ├── create_user.go
│   ├── list_users.go
│   ├── get_user.go
│   ├── update_user.go
│   └── delete_user.go
├── auth/
│   ├── login.go
│   ├── refresh_token.go
│   └── logout.go
└── roles/
    ├── create_role.go
    ├── assign_permissions.go
    └── list_roles.go
```

---

### 3.3 Camada de Dados (Data Provider)

**Conceito**: Camada responsável por persistência, implementando as interfaces de gateway definidas no domínio.

#### 3.3.1 Models (`dataprovider/database/model/`)

**Conceito**: Representam tabelas do banco de dados usando GORM. São estruturas anêmicas (sem lógica), apenas mapeamento ORM.

**Exemplo:**
```go
type User struct {
    UUID      uuid.UUID  `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
    Email     string     `gorm:"uniqueIndex;not null"`
    Password  string     `gorm:"not null"`
    RoleID    uuid.UUID  `gorm:"type:uuid;not null"`
    Role      Role       `gorm:"foreignKey:RoleID"`
    CreatedAt time.Time  `gorm:"not null"`
    UpdatedAt time.Time  `gorm:"not null"`
    DeletedAt *time.Time `gorm:"index"`
}

func (User) TableName() string {
    return "users"
}
```

**Características:**
- Tags GORM para mapeamento (`gorm:"..."`)
- Relacionamentos definidos (ForeignKey, Many2Many)
- Soft delete via `DeletedAt`
- Sem lógica de negócio

**Padrão BaseModel:**
```go
type BaseModel struct {
    UUID      uuid.UUID  `gorm:"primaryKey;type:uuid"`
    CreatedAt time.Time  `gorm:"not null"`
    UpdatedAt time.Time  `gorm:"not null"`
    DeletedAt *time.Time `gorm:"index"`
}
```

#### 3.3.2 Gateway Implementations (`dataprovider/database/gateway/`)

**Conceito**: Implementação concreta das interfaces de gateway usando GORM.

**Exemplo:**
```go
type UserGatewayGORM struct {
    db     *gorm.DB
    mapper *mapper.UserMapper
}

func NewUserGatewayGORM(db *gorm.DB, mapper *mapper.UserMapper) *UserGatewayGORM {
    return &UserGatewayGORM{db: db, mapper: mapper}
}

func (g *UserGatewayGORM) Save(ctx context.Context, user *entity.User) error {
    userModel := g.mapper.ToModel(user)
    return g.db.WithContext(ctx).Create(userModel).Error
}

func (g *UserGatewayGORM) FindByID(ctx context.Context, id uuid.UUID) (*entity.User, error) {
    var userModel model.User
    err := g.db.WithContext(ctx).
        Where("uuid = ? AND deleted_at IS NULL", id).
        First(&userModel).Error

    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, domainErrors.ErrUserNotFound
        }
        return nil, err
    }

    return g.mapper.ToDomain(&userModel)
}
```

**Responsabilidades:**
- Implementar operações de persistência
- Usar mapper para converter Model ↔ Entity
- Tratar erros GORM e mapear para erros de domínio
- Adicionar filtros (soft delete)

#### 3.3.3 Mappers (`dataprovider/database/mapper/`)

**Conceito**: Converte entre Model (banco) e Entity (domínio).

**Exemplo:**
```go
type UserMapper struct{}

// Model → Entity
func (m *UserMapper) ToDomain(userModel *model.User) (*entity.User, error) {
    return entity.ReconstructUser(
        userModel.UUID,
        userModel.Email,
        userModel.Password,
        userModel.RoleID,
        userModel.CreatedAt,
        userModel.UpdatedAt,
        userModel.DeletedAt,
    )
}

// Entity → Model
func (m *UserMapper) ToModel(user *entity.User) *model.User {
    return &model.User{
        UUID:      user.ID(),
        Email:     user.Email(),
        Password:  user.Password(),
        RoleID:    user.RoleID(),
        CreatedAt: user.CreatedAt(),
        UpdatedAt: user.UpdatedAt(),
        DeletedAt: user.DeletedAt(),
    }
}
```

**Por que usar?**
- Mantém separação entre camadas
- Entity não conhece GORM
- Model não tem lógica de negócio
- Facilita mudanças em uma camada sem afetar outra

---

### 3.4 Camada de Entrada (Entrypoint/HTTP)

**Conceito**: Camada que expõe a aplicação via HTTP usando Fiber. Não contém lógica de negócio.

#### 3.4.1 Controllers (`entrypoint/http/controller/`)

**Conceito**: Recebem requisições HTTP, delegam para use cases, retornam respostas HTTP.

**Exemplo:**
```go
type UserController struct {
    createUserUseCase *users.CreateUser
    listUsersUseCase  *users.ListUsers
    validator         *validation.Validator
}

func (c *UserController) Create(ctx *fiber.Ctx) error {
    // 1. Parse do request
    var req payload.CreateUserRequest
    if err := ctx.BodyParser(&req); err != nil {
        return ctx.Status(400).JSON(response.ErrorResponse{
            Success: false,
            Message: "Invalid request body",
        })
    }

    // 2. Validar payload
    if err := c.validator.Validate(req); err != nil {
        return ctx.Status(400).JSON(response.ValidationErrorResponse{
            Success: false,
            Errors:  err,
        })
    }

    // 3. Chamar use case
    user, err := c.createUserUseCase.Execute(ctx.Context(), usecase.CreateUserData{
        Email:    req.Email,
        Password: req.Password,
        RoleID:   uuid.MustParse(req.RoleID),
    })

    // 4. Tratar erros
    if err != nil {
        if errors.Is(err, domainErrors.ErrUserEmailExists) {
            return ctx.Status(409).JSON(response.ErrorResponse{
                Success: false,
                Message: "Email already exists",
            })
        }
        return ctx.Status(500).JSON(response.ErrorResponse{
            Success: false,
            Message: "Internal server error",
        })
    }

    // 5. Formatar resposta
    return ctx.Status(201).JSON(response.SuccessResponse{
        Success: true,
        Data:    mapper.ToUserResponse(user),
    })
}
```

**Responsabilidades:**
- Parse de request body/params/query
- Validação de formato (via validator)
- Chamar use case apropriado
- Mapear erros de domínio para HTTP status codes
- Formatar resposta JSON

**NÃO fazem:**
- Lógica de negócio
- Acesso direto ao banco
- Validação de regras de negócio

#### 3.4.2 Middleware (`entrypoint/http/middleware/`)

**Conceito**: Interceptam requisições antes de chegarem aos controllers.

**Exemplos Comuns:**
- `JWTMiddleware`: Autenticação e autorização
- `CORSMiddleware`: Cross-Origin Resource Sharing
- `RateLimitMiddleware`: Limite de requisições
- `LoggingMiddleware`: Logs de requisições

**Exemplo de JWT Middleware:**
```go
type JWTMiddleware struct {
    jwtService  service.JWTService
    roleGateway gateway.RoleGateway
}

func (m *JWTMiddleware) RequireAuth() fiber.Handler {
    return func(c *fiber.Ctx) error {
        // Extrair token do header
        authHeader := c.Get("Authorization")
        token := strings.TrimPrefix(authHeader, "Bearer ")

        // Validar token
        claims, err := m.jwtService.ValidateToken(token, service.AccessToken)
        if err != nil {
            return c.Status(401).JSON(...)
        }

        // Injetar no contexto
        c.Locals("userID", claims.UserID)
        c.Locals("roleID", claims.RoleID)

        return c.Next()
    }
}

func (m *JWTMiddleware) RequirePermission(permSlug string) fiber.Handler {
    return func(c *fiber.Ctx) error {
        // Pegar roleID do contexto
        roleID := c.Locals("roleID").(uuid.UUID)

        // Buscar role com permissões
        role, err := m.roleGateway.FindByIDWithPermissions(c.Context(), roleID)
        if err != nil {
            return c.Status(403).JSON(...)
        }

        // Verificar permissão (com wildcard)
        if !role.HasPermission(permSlug) {
            return c.Status(403).JSON(...)
        }

        return c.Next()
    }
}
```

#### 3.4.3 Routes (`entrypoint/http/routes/`)

**Conceito**: Definem endpoints e aplicam middlewares.

**Exemplo:**
```go
func SetupUserRoutes(
    app *fiber.App,
    userController *controller.UserController,
    jwtMiddleware *middleware.JWTMiddleware,
) {
    userGroup := app.Group("/api/v1/users")

    // Rotas protegidas
    protected := userGroup.Group("/")
    protected.Use(jwtMiddleware.RequireAuth())

    protected.Get("/",
        jwtMiddleware.RequirePermission(constants.ReadUsers),
        userController.List)

    protected.Get("/:id",
        jwtMiddleware.RequirePermission(constants.ReadUsers),
        userController.GetByID)

    protected.Post("/",
        jwtMiddleware.RequirePermission(constants.CreateUsers),
        userController.Create)

    protected.Put("/:id",
        jwtMiddleware.RequirePermission(constants.UpdateUsers),
        userController.Update)

    protected.Delete("/:id",
        jwtMiddleware.RequirePermission(constants.DeleteUsers),
        userController.Delete)
}
```

#### 3.4.4 Payloads (`entrypoint/http/payload/`)

**Conceito**: DTOs (Data Transfer Objects) para request e response.

**Request:**
```go
type CreateUserRequest struct {
    Email    string `json:"email" validate:"required,email"`
    Password string `json:"password" validate:"required,min=8"`
    RoleID   string `json:"role_id" validate:"required,uuid4"`
}
```

**Response:**
```go
type UserResponse struct {
    ID        uuid.UUID `json:"id"`
    Email     string    `json:"email"`
    RoleID    uuid.UUID `json:"role_id"`
    CreatedAt time.Time `json:"created_at"`
}

type SuccessResponse struct {
    Success bool         `json:"success"`
    Data    UserResponse `json:"data"`
}

type ErrorResponse struct {
    Success bool   `json:"success"`
    Message string `json:"message"`
    Code    string `json:"code,omitempty"`
}
```

---

## 4. Padrões Arquiteturais Aplicados

### 4.1 Clean Architecture

**Princípio**: As dependências de código fonte devem apontar apenas para dentro, em direção às políticas de alto nível (domínio).

**Regra de Dependência:**
```
Entrypoint (HTTP) → Application (Use Cases) → Domain (Core) ← Data Provider
```

O domínio não conhece nada sobre HTTP, banco de dados ou frameworks. Isso permite:
- Testar lógica de negócio sem servidor HTTP
- Trocar Fiber por Gin ou Echo sem mudar o domínio
- Trocar PostgreSQL por MongoDB sem mudar use cases

### 4.2 Dependency Injection

**Conceito**: Objetos recebem suas dependências via constructor, não criam internamente.

**Container por Feature:**
```go
type UserContainer struct {
    UserGateway       gateway.UserGateway
    UserMapper        *mapper.UserMapper
    CreateUserUseCase *users.CreateUser
    ListUsersUseCase  *users.ListUsers
    UserController    *controller.UserController
}

func InitializeUserContainer(db *gorm.DB, hasher service.HasherService, ...) *UserContainer {
    // 1. Criar mapper
    userMapper := &mapper.UserMapper{}

    // 2. Criar gateway
    userGateway := gateway.NewUserGatewayGORM(db, userMapper)

    // 3. Criar use cases
    createUserUseCase := users.NewCreateUser(userGateway, roleGateway, hasher)
    listUsersUseCase := users.NewListUsers(userGateway)

    // 4. Criar controller
    userController := controller.NewUserController(createUserUseCase, listUsersUseCase, ...)

    return &UserContainer{
        UserGateway:       userGateway,
        UserMapper:        userMapper,
        CreateUserUseCase: createUserUseCase,
        ListUsersUseCase:  listUsersUseCase,
        UserController:    userController,
    }
}
```

**Vantagens:**
- Facilita testes (mock de dependências)
- Ciclo de vida controlado
- Mudanças localizadas

### 4.3 Gateway/Repository Pattern

**Conceito**: Abstrai acesso a dados através de interface definida no domínio.

**Por que usar:**
- Domínio não conhece tecnologia de persistência
- Facilita trocar de banco (SQL → NoSQL)
- Mocks triviais para testes

**Fluxo:**
```
Use Case → Gateway Interface (domínio) → Gateway Implementation (infra) → GORM → PostgreSQL
```

### 4.4 Factory Pattern

**Conceito**: Centraliza criação de objetos complexos.

**Exemplo**: Criar diferentes tipos de user profile (Talent, Recruiter, Admin).

```go
type UserFactory struct{}

func (f *UserFactory) CreateUser(userType string, data map[string]interface{}) (*entity.User, error) {
    switch userType {
    case "talent":
        return f.createTalentUser(data)
    case "recruiter":
        return f.createRecruiterUser(data)
    default:
        return nil, errors.New("unknown user type")
    }
}
```

### 4.5 Strategy Pattern

**Conceito**: Define família de algoritmos, encapsula cada um e torna intercambiáveis.

**Exemplo**: Different user profiles with different behaviors.

```go
// Interface
type UserProfile interface {
    Type() string
    Validate() error
    CanApplyToJobs() bool
    CanPostJobs() bool
}

// Implementações
type TalentProfile struct { ... }
func (t *TalentProfile) CanApplyToJobs() bool { return true }
func (t *TalentProfile) CanPostJobs() bool { return false }

type RecruiterProfile struct { ... }
func (r *RecruiterProfile) CanApplyToJobs() bool { return false }
func (r *RecruiterProfile) CanPostJobs() bool { return true }

// User delega para profile
type User struct {
    profile UserProfile
}

func (u *User) CanApplyToJobs() bool {
    return u.profile.CanApplyToJobs()
}
```

### 4.6 Criteria/Query Builder Pattern

**Conceito**: Constrói queries dinamicamente de forma type-safe.

**Exemplo:**
```go
type UserCriteria struct {
    email  *string
    roleID *uuid.UUID
    active *bool
}

func NewUserCriteria() *UserCriteria {
    return &UserCriteria{}
}

func (c *UserCriteria) WithEmail(email string) *UserCriteria {
    c.email = &email
    return c
}

func (c *UserCriteria) WithActive(active bool) *UserCriteria {
    c.active = &active
    return c
}

func (c *UserCriteria) Apply(db *gorm.DB) *gorm.DB {
    if c.email != nil {
        db = db.Where("email = ?", *c.email)
    }
    if c.active != nil && *c.active {
        db = db.Where("deleted_at IS NULL")
    }
    return db
}

// Uso:
criteria := NewUserCriteria().WithEmail("test@example.com").WithActive(true)
users, err := userGateway.FindByCriteria(ctx, criteria, pagination, sortOrder)
```

**Vantagens:**
- Type-safe
- Reutilizável
- Composable
- Sem SQL strings

### 4.7 Soft Delete Pattern

**Conceito**: Deleção lógica ao invés de física. Registros são marcados como deletados mas não removidos.

**Implementação em todas as camadas:**

**Model:**
```go
type BaseModel struct {
    DeletedAt *time.Time `gorm:"index"`
}
```

**Entity:**
```go
func (u *User) SoftDelete() {
    now := time.Now()
    u.deletedAt = &now
}

func (u *User) IsActive() bool {
    return u.deletedAt == nil
}
```

**Gateway:**
```go
func (g *UserGatewayGORM) FindAll(ctx context.Context) ([]*entity.User, error) {
    var models []model.User
    // Filtra automaticamente deletados
    err := g.db.WithContext(ctx).Where("deleted_at IS NULL").Find(&models).Error
    ...
}
```

**Criteria:**
```go
type ActiveCriteria struct{}

func (c *ActiveCriteria) Apply(db *gorm.DB) *gorm.DB {
    return db.Where("deleted_at IS NULL")
}
```

**Vantagens:**
- Auditoria completa
- Recuperação possível
- Integridade referencial mantida

---

## 5. Implementação de RBAC (Role-Based Access Control)

### 5.1 Modelo de Dados

**Estrutura:**
```
Roles ←→ RolePermissions ←→ Permissions
  ↓
Users
```

**Junction Table:**
```go
type RolePermission struct {
    RoleID       uuid.UUID `gorm:"primaryKey"`
    PermissionID uuid.UUID `gorm:"primaryKey"`
    Role         Role      `gorm:"constraint:OnDelete:CASCADE"`
    Permission   Permission `gorm:"constraint:OnDelete:CASCADE"`
}
```

**Role:**
```go
type Role struct {
    UUID        uuid.UUID
    Name        string
    Slug        string `gorm:"uniqueIndex"`
    Permissions []Permission `gorm:"many2many:role_permissions"`
}
```

**Permission:**
```go
type Permission struct {
    UUID uuid.UUID
    Name string
    Slug string `gorm:"uniqueIndex"`
}
```

### 5.2 Sistema de Permissões

**Organização por Recurso:**
```go
// core/domain/constants/permissions.go
const (
    // Users
    CreateUsers = "users.create"
    ReadUsers   = "users.read"
    UpdateUsers = "users.update"
    DeleteUsers = "users.delete"
    ManageUsers = "users.*"     // Wildcard

    // Roles
    CreateRoles = "roles.create"
    ReadRoles   = "roles.read"
    ManageRoles = "roles.*"

    // Jobs
    CreateJobs = "jobs.create"
    ReadJobs   = "jobs.read"
    ManageJobs = "jobs.*"
)
```

**Wildcard Matching:**
O wildcard `*` permite que uma permissão cubra múltiplas operações:
- `users.*` cobre `users.create`, `users.read`, `users.update`, `users.delete`
- `*` cobre todas as permissões do sistema

**Implementação na Entity:**
```go
func (r *Role) HasPermission(slug string) bool {
    for _, perm := range r.permissions {
        permSlug := perm.Slug()

        // Match exato
        if permSlug == slug {
            return true
        }

        // Wildcard match: "users.*" matches "users.create"
        if strings.HasSuffix(permSlug, ".*") {
            prefix := strings.TrimSuffix(permSlug, ".*")
            if strings.HasPrefix(slug, prefix+".") {
                return true
            }
        }

        // Super admin: "*" matches everything
        if permSlug == "*" {
            return true
        }
    }
    return false
}
```

### 5.3 Fluxo de Autorização

**1. Request chega com JWT:**
```
Authorization: Bearer eyJhbGciOiJIUzI1NiIs...
```

**2. JWTMiddleware extrai e valida:**
```go
func (m *JWTMiddleware) RequireAuth() fiber.Handler {
    return func(c *fiber.Ctx) error {
        token := extractToken(c)
        claims, err := m.jwtService.ValidateToken(token, AccessToken)
        if err != nil {
            return c.Status(401).JSON(...)
        }

        // Injeta no contexto
        c.Locals("userID", claims.UserID)
        c.Locals("roleID", claims.RoleID)
        c.Locals("email", claims.Email)

        return c.Next()
    }
}
```

**3. RequirePermission verifica permissão:**
```go
func (m *JWTMiddleware) RequirePermission(permSlug string) fiber.Handler {
    return func(c *fiber.Ctx) error {
        roleID := c.Locals("roleID").(uuid.UUID)

        // Busca role com permissões via gateway
        role, err := m.roleGateway.FindByIDWithPermissions(c.Context(), roleID)
        if err != nil {
            return c.Status(403).JSON(...)
        }

        // Verifica permissão (incluindo wildcard)
        if !role.HasPermission(permSlug) {
            return c.Status(403).JSON(response.ErrorResponse{
                Success: false,
                Message: "Insufficient permissions",
                Code:    "FORBIDDEN",
            })
        }

        return c.Next()
    }
}
```

**4. Controller executa normalmente:**
```go
protected.Get("/",
    jwtMiddleware.RequirePermission(constants.ReadUsers),
    userController.List)
```

### 5.4 Seeders RBAC

**Ordem de Execução:**
1. **PermissionSeeder**: Cria todas as permissões do sistema
2. **RoleSeeder**: Cria roles base (Admin, Partner, Recruiter, etc.)
3. **RolePermissionSeeder**: Mapeia permissões a cada role

**Exemplo - PermissionSeeder:**
```go
func (s *PermissionSeeder) Seed(ctx context.Context) error {
    permissions := []struct {
        Name string
        Slug string
    }{
        {"Create Users", constants.CreateUsers},
        {"Read Users", constants.ReadUsers},
        {"Update Users", constants.UpdateUsers},
        {"Delete Users", constants.DeleteUsers},
        {"Manage Users", constants.ManageUsers},
        // ... outras permissões
    }

    for _, p := range permissions {
        exists, _ := s.permissionGateway.ExistsBySlug(ctx, p.Slug)
        if !exists {
            perm, _ := entity.NewPermission(p.Name, p.Slug)
            s.permissionGateway.Save(ctx, perm)
        }
    }
    return nil
}
```

**Exemplo - RolePermissionSeeder:**
```go
func (s *RolePermissionSeeder) Seed(ctx context.Context) error {
    // Admin tem todas as permissões (wildcard)
    adminRole, _ := s.roleGateway.FindBySlug(ctx, constants.AdminRole)
    allPerm, _ := s.permissionGateway.FindBySlug(ctx, "*")
    s.roleGateway.AssignPermissions(ctx, adminRole.ID(), []uuid.UUID{allPerm.ID()})

    // Recruiter pode gerenciar jobs e ler users
    recruiterRole, _ := s.roleGateway.FindBySlug(ctx, constants.RecruiterRole)
    manageJobs, _ := s.permissionGateway.FindBySlug(ctx, constants.ManageJobs)
    readUsers, _ := s.permissionGateway.FindBySlug(ctx, constants.ReadUsers)
    s.roleGateway.AssignPermissions(ctx, recruiterRole.ID(), []uuid.UUID{
        manageJobs.ID(),
        readUsers.ID(),
    })

    return nil
}
```

---

## 6. Configuração e Ambiente

### 6.1 Estrutura de Config

```go
type Config struct {
    Server   ServerConfig
    Database DatabaseConfig
    JWT      JWTConfig
    Redis    RedisConfig
    CORS     CORSConfig
}

type ServerConfig struct {
    Port        string
    Environment string // "development", "production"
}

type DatabaseConfig struct {
    Host            string
    Port            string
    User            string
    Password        string
    DBName          string
    MaxOpenConns    int
    MaxIdleConns    int
    ConnMaxLifetime int // minutos
    ConnMaxIdleTime int // minutos
}

type JWTConfig struct {
    Secret              string
    AccessTokenExpiry   int // horas
    RefreshTokenExpiry  int // dias
}

type CORSConfig struct {
    AllowedOrigins   []string
    AllowedMethods   []string
    AllowedHeaders   []string
    AllowCredentials bool
}
```

**Carregamento via Environment Variables:**
```go
func LoadConfig() *Config {
    return &Config{
        Server: ServerConfig{
            Port:        getEnv("PORT", "8080"),
            Environment: getEnv("ENVIRONMENT", "development"),
        },
        Database: DatabaseConfig{
            Host:            getEnv("DB_HOST", "localhost"),
            Port:            getEnv("DB_PORT", "5432"),
            User:            getEnv("DB_USER", "postgres"),
            Password:        getEnv("DB_PASSWORD", ""),
            DBName:          getEnv("DB_NAME", "myapp"),
            MaxOpenConns:    getEnvAsInt("DB_MAX_OPEN_CONNS", 25),
            MaxIdleConns:    getEnvAsInt("DB_MAX_IDLE_CONNS", 10),
            ConnMaxLifetime: getEnvAsInt("DB_CONN_MAX_LIFETIME", 5),
            ConnMaxIdleTime: getEnvAsInt("DB_CONN_MAX_IDLE_TIME", 1),
        },
        JWT: JWTConfig{
            Secret:             getEnv("JWT_SECRET", ""),
            AccessTokenExpiry:  getEnvAsInt("JWT_ACCESS_EXPIRY", 24),
            RefreshTokenExpiry: getEnvAsInt("JWT_REFRESH_EXPIRY", 7),
        },
    }
}

func getEnv(key, defaultValue string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return defaultValue
}
```

**Validação:**
```go
func (c *Config) Validate() error {
    if c.JWT.Secret == "" {
        return errors.New("JWT_SECRET is required")
    }
    if c.Database.Password == "" && c.Server.Environment == "production" {
        return errors.New("DB_PASSWORD is required in production")
    }
    return nil
}
```

### 6.2 Connection Pool

```go
func setupDatabase(cfg *config.Config) (*gorm.DB, error) {
    dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        cfg.Database.Host,
        cfg.Database.Port,
        cfg.Database.User,
        cfg.Database.Password,
        cfg.Database.DBName,
    )

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, err
    }

    sqlDB, err := db.DB()
    if err != nil {
        return nil, err
    }

    // Connection pool settings
    sqlDB.SetMaxOpenConns(cfg.Database.MaxOpenConns)           // 25
    sqlDB.SetMaxIdleConns(cfg.Database.MaxIdleConns)           // 10
    sqlDB.SetConnMaxLifetime(time.Duration(cfg.Database.ConnMaxLifetime) * time.Minute)
    sqlDB.SetConnMaxIdleTime(time.Duration(cfg.Database.ConnMaxIdleTime) * time.Minute)

    return db, nil
}
```

---

## 7. Migrations e Seeders

### 7.1 Estratégia de Migrations

**Estrutura de Diretórios:**
```
migrations/
├── up/
│   ├── 000_create_utility_functions.up.sql
│   ├── 001_create_permissions_table.up.sql
│   └── 002_create_roles_table.up.sql
└── down/
    ├── 000_create_utility_functions.down.sql
    ├── 001_create_permissions_table.down.sql
    └── 002_create_roles_table.down.sql
```

**Utility Functions (Trigger para updated_at):**
```sql
-- 000_create_utility_functions.up.sql
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;
```

**GORM AutoMigrate:**
```go
func runMigrations(db *gorm.DB) error {
    // Ordem importa: dependências primeiro
    return db.AutoMigrate(
        &model.Permission{},
        &model.Role{},
        &model.User{},
        &model.Job{},
        &model.RolePermission{},
    )
}
```

**Ordem de Execução:**
1. Utility functions
2. Enums
3. Tabelas base (sem FK)
4. Tabelas com FK
5. Junction tables
6. Indexes e triggers

### 7.2 Seeders

**Padrão de Seeder:**
```go
type PermissionSeeder struct {
    db                *gorm.DB
    permissionGateway gateway.PermissionGateway
}

func (s *PermissionSeeder) Seed(ctx context.Context) error {
    // 1. Verificar se já foi seeded
    count, _ := s.permissionGateway.Count(ctx)
    if count > 0 {
        log.Println("Permissions already seeded, skipping...")
        return nil
    }

    // 2. Criar e salvar
    permissions := []struct{ Name, Slug string }{
        {"Create Users", constants.CreateUsers},
        {"Read Users", constants.ReadUsers},
        // ...
    }

    for _, p := range permissions {
        perm, _ := entity.NewPermission(p.Name, p.Slug)
        s.permissionGateway.Save(ctx, perm)
    }

    return nil
}
```

**Ordem de Execução:**
```go
func runSeeds(container *config.Container) error {
    ctx := context.Background()

    // 1. Permissions primeiro
    permSeeder := seeders.NewPermissionSeeder(container.DB, container.PermissionGateway)
    permSeeder.Seed(ctx)

    // 2. Roles
    roleSeeder := seeders.NewRoleSeeder(container.DB, container.RoleGateway)
    roleSeeder.Seed(ctx)

    // 3. RolePermissions (mapeia permissões a roles)
    rpSeeder := seeders.NewRolePermissionSeeder(container.DB, container.RoleGateway, container.PermissionGateway)
    rpSeeder.Seed(ctx)

    // 4. Users de teste
    userSeeder := seeders.NewUserSeeder(container.DB, container.UserGateway, container.RoleGateway)
    userSeeder.Seed(ctx)

    return nil
}
```

**Idempotência:**
Seeders devem ser idempotentes (podem rodar múltiplas vezes sem duplicar dados):
```go
exists, _ := s.permissionGateway.ExistsBySlug(ctx, slug)
if !exists {
    // Criar
}
```

---

## 8. Tratamento de Erros

### 8.1 Erros de Domínio

**Definição:**
```go
// core/domain/errors/user_errors.go
var (
    ErrUserNotFound           = errors.New("user not found")
    ErrUserEmailExists        = errors.New("user email already exists")
    ErrUserInvalidCredentials = errors.New("invalid credentials")
    ErrUserInactive           = errors.New("user is inactive")
)
```

**Propagação:**
```
Gateway → Use Case → Controller
```

Gateway mapeia erros GORM para erros de domínio:
```go
if errors.Is(err, gorm.ErrRecordNotFound) {
    return nil, domainErrors.ErrUserNotFound
}
```

Use case retorna erro de domínio:
```go
user, err := uc.userGateway.FindByID(ctx, id)
if err != nil {
    return nil, err // ErrUserNotFound
}
```

Controller mapeia para HTTP status:
```go
if errors.Is(err, domainErrors.ErrUserNotFound) {
    return c.Status(404).JSON(...)
}
if errors.Is(err, domainErrors.ErrUserEmailExists) {
    return c.Status(409).JSON(...)
}
```

### 8.2 Erros de Validação

**Validação de Formato (Payload):**
```go
type CreateUserRequest struct {
    Email    string `json:"email" validate:"required,email"`
    Password string `json:"password" validate:"required,min=8,user_password"`
}

// No controller:
if err := c.validator.Validate(req); err != nil {
    return c.Status(400).JSON(response.ValidationErrorResponse{
        Success: false,
        Errors:  formatValidationErrors(err),
    })
}
```

**Validação de Negócio (Entity):**
```go
func (u *User) Validate() error {
    if u.email == "" {
        return errors.New("email is required")
    }
    if !u.IsValidEmail() {
        return errors.New("invalid email format")
    }
    return nil
}
```

### 8.3 Resposta HTTP Padronizada

**Sucesso:**
```json
{
  "success": true,
  "data": { ... },
  "message": "User created successfully"
}
```

**Erro de Validação:**
```json
{
  "success": false,
  "message": "Validation failed",
  "errors": {
    "email": "Email is required",
    "password": "Password must be at least 8 characters"
  }
}
```

**Erro de Negócio:**
```json
{
  "success": false,
  "message": "Email already exists",
  "code": "USER_EMAIL_EXISTS"
}
```

---

## 9. Request/Response Pattern

### 9.1 DTOs de Request

```go
// entrypoint/http/payload/request/user/create_user_request.go
type CreateUserRequest struct {
    Email    string `json:"email" validate:"required,email"`
    Password string `json:"password" validate:"required,min=8"`
    RoleID   string `json:"role_id" validate:"required,uuid4"`
}
```

**Características:**
- Tags `json` para unmarshaling
- Tags `validate` para validação de formato
- Organizados por feature em subdiretórios

### 9.2 DTOs de Response

```go
// entrypoint/http/payload/response/user/user_response.go
type UserResponse struct {
    ID        uuid.UUID `json:"id"`
    Email     string    `json:"email"`
    RoleID    uuid.UUID `json:"role_id"`
    RoleName  string    `json:"role_name,omitempty"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

type UserSuccessResponse struct {
    Success bool         `json:"success"`
    Message string       `json:"message"`
    Data    UserResponse `json:"data"`
}

type UserListResponse struct {
    Success bool           `json:"success"`
    Data    []UserResponse `json:"data"`
    Meta    MetaResponse   `json:"meta"`
}

type MetaResponse struct {
    Page      int   `json:"page"`
    PageSize  int   `json:"page_size"`
    TotalRows int64 `json:"total_rows"`
    TotalPages int  `json:"total_pages"`
}
```

### 9.3 Payload Mapper

```go
type UserPayloadMapper struct{}

func (m *UserPayloadMapper) ToResponse(user *entity.User) UserResponse {
    return UserResponse{
        ID:        user.ID(),
        Email:     user.Email(),
        RoleID:    user.RoleID(),
        CreatedAt: user.CreatedAt(),
        UpdatedAt: user.UpdatedAt(),
    }
}

func (m *UserPayloadMapper) ToResponseList(users []*entity.User) []UserResponse {
    responses := make([]UserResponse, len(users))
    for i, user := range users {
        responses[i] = m.ToResponse(user)
    }
    return responses
}
```

---

## 10. Entry Point e Inicialização

### 10.1 Ordem de Inicialização (main.go)

```go
func main() {
    // 1. Setup de logger
    setupLogger()

    // 2. Carregar configuração
    cfg := config.LoadConfig()
    if err := cfg.Validate(); err != nil {
        log.Fatal("Invalid config:", err)
    }

    // 3. Criar container de DI
    container := config.NewContainer(cfg)

    // 4. Conectar ao banco de dados
    db, err := setupDatabase(cfg)
    if err != nil {
        log.Fatal("Database connection failed:", err)
    }
    container.DB = db

    // 5. Inicializar componentes (gateways, use cases, controllers)
    if err := container.Initialize(); err != nil {
        log.Fatal("Container initialization failed:", err)
    }

    // 6. Setup do Fiber app
    app := setupFiber(cfg)

    // 7. Registrar todas as rotas
    setupRoutes(app, container)

    // 8. Executar migrations
    if err := runMigrations(db); err != nil {
        log.Fatal("Migrations failed:", err)
    }

    // 9. Executar seeders (apenas em dev)
    if cfg.Server.Environment == "development" {
        runSeeds(container)
    }

    // 10. Iniciar servidor com graceful shutdown
    startServer(app, cfg)
}
```

### 10.2 Graceful Shutdown

```go
func startServer(app *fiber.App, cfg *config.Config) {
    // Canal para sinais do OS
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

    addr := fmt.Sprintf(":%s", cfg.Server.Port)

    // Goroutine para iniciar servidor
    go func() {
        log.Printf("Server starting on %s", addr)
        if err := app.Listen(addr); err != nil {
            log.Printf("Server error: %v", err)
        }
    }()

    // Aguardar sinal de shutdown
    <-quit
    log.Println("Shutting down server...")

    // Timeout de 30 segundos para finalizar requisições em andamento
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()

    if err := app.ShutdownWithContext(ctx); err != nil {
        log.Printf("Server forced to shutdown: %v", err)
    }

    log.Println("Server exited gracefully")
}
```

---

## 11. Checklist: Criando uma Nova Feature

Use este template ao adicionar uma nova feature (exemplo: "Teams").

### Passo 1: Camada de Domínio

- [ ] **Criar Entity** (`core/domain/entity/team.go`)
  - Campos privados com getters
  - Constructor `NewTeam(name, description string) (*Team, error)`
  - Reconstruction `ReconstructTeam(id uuid.UUID, ...) (*Team, error)`
  - Métodos de lógica de negócio (`Validate()`, `AddMember()`, `RemoveMember()`)
  - Soft delete support (`SoftDelete()`, `IsActive()`)

- [ ] **Criar Gateway Interface** (`core/domain/gateway/team_gateway.go`)
  ```go
  type TeamGateway interface {
      Save(ctx context.Context, team *entity.Team) error
      FindByID(ctx context.Context, id uuid.UUID) (*entity.Team, error)
      Update(ctx context.Context, team *entity.Team) error
      Delete(ctx context.Context, id uuid.UUID) error
      FindByCriteria(...) ([]*entity.Team, error)
      CountByCriteria(...) (int64, error)
      ExistsByID(ctx context.Context, id uuid.UUID) (bool, error)
  }
  ```

- [ ] **Criar Erros de Domínio** (`core/domain/errors/team_errors.go`)
  ```go
  var (
      ErrTeamNotFound    = errors.New("team not found")
      ErrTeamNameExists  = errors.New("team name already exists")
  )
  ```

- [ ] **(Opcional) Criar Value Objects** se necessário

- [ ] **(Opcional) Criar Criteria** (`core/domain/repository/criteria/team_criteria.go`)
  ```go
  type TeamCriteria struct {
      name   *string
      active *bool
  }
  ```

### Passo 2: Camada de Aplicação

- [ ] **Criar Use Cases** (`application/usecase/teams/`)
  - `create_team.go`: CreateTeam use case
  - `list_teams.go`: ListTeams use case
  - `get_team.go`: GetTeam use case
  - `update_team.go`: UpdateTeam use case
  - `delete_team.go`: DeleteTeam use case

Cada use case deve ter:
```go
type CreateTeam struct {
    teamGateway gateway.TeamGateway
    // outros gateways/services necessários
}

func NewCreateTeam(teamGateway gateway.TeamGateway) *CreateTeam {
    return &CreateTeam{teamGateway: teamGateway}
}

func (uc *CreateTeam) Execute(ctx context.Context, data CreateTeamData) (*entity.Team, error) {
    // 1. Validações de aplicação
    // 2. Criar entity
    // 3. Salvar via gateway
    // 4. Retornar resultado
}
```

### Passo 3: Camada de Dados

- [ ] **Criar Model GORM** (`dataprovider/database/model/team.go`)
  ```go
  type Team struct {
      UUID        uuid.UUID `gorm:"primaryKey;type:uuid"`
      Name        string    `gorm:"uniqueIndex;not null"`
      Description *string
      CreatedAt   time.Time  `gorm:"not null"`
      UpdatedAt   time.Time  `gorm:"not null"`
      DeletedAt   *time.Time `gorm:"index"`
  }
  ```

- [ ] **Criar Mapper** (`dataprovider/database/mapper/team_mapper.go`)
  ```go
  type TeamMapper struct{}

  func (m *TeamMapper) ToDomain(model *model.Team) (*entity.Team, error) {
      return entity.ReconstructTeam(...)
  }

  func (m *TeamMapper) ToModel(team *entity.Team) *model.Team {
      return &model.Team{...}
  }
  ```

- [ ] **Criar Gateway Implementation** (`dataprovider/database/gateway/team_gateway.go`)
  ```go
  type TeamGatewayGORM struct {
      db     *gorm.DB
      mapper *mapper.TeamMapper
  }

  func (g *TeamGatewayGORM) Save(ctx context.Context, team *entity.Team) error {
      teamModel := g.mapper.ToModel(team)
      return g.db.WithContext(ctx).Create(teamModel).Error
  }
  // ... outras implementações
  ```

### Passo 4: Camada HTTP

- [ ] **Criar Payloads** (`entrypoint/http/payload/`)
  - `request/team/create_team_request.go`
  ```go
  type CreateTeamRequest struct {
      Name        string `json:"name" validate:"required,min=3"`
      Description string `json:"description,omitempty"`
  }
  ```
  - `response/team/team_response.go`
  ```go
  type TeamResponse struct {
      ID          uuid.UUID `json:"id"`
      Name        string    `json:"name"`
      Description string    `json:"description,omitempty"`
      CreatedAt   time.Time `json:"created_at"`
  }
  ```
  - Mapper de payload

- [ ] **Criar Controller** (`entrypoint/http/controller/team_controller.go`)
  ```go
  type TeamController struct {
      createTeamUseCase *teams.CreateTeam
      listTeamsUseCase  *teams.ListTeams
      // ...
  }

  func (c *TeamController) Create(ctx *fiber.Ctx) error {
      // Parse, validate, call use case, format response
  }
  ```

- [ ] **Criar Routes** (`entrypoint/http/routes/team_routes.go`)
  ```go
  func SetupTeamRoutes(app *fiber.App, controller *controller.TeamController, jwt *middleware.JWTMiddleware) {
      teamGroup := app.Group("/api/v1/teams")
      protected := teamGroup.Group("/")
      protected.Use(jwt.RequireAuth())

      protected.Get("/", jwt.RequirePermission(constants.ReadTeams), controller.List)
      protected.Post("/", jwt.RequirePermission(constants.CreateTeams), controller.Create)
      // ...
  }
  ```

### Passo 5: RBAC (se aplicável)

- [ ] **Adicionar Constantes de Permissão** (`core/domain/constants/permissions.go`)
  ```go
  const (
      CreateTeams = "teams.create"
      ReadTeams   = "teams.read"
      UpdateTeams = "teams.update"
      DeleteTeams = "teams.delete"
      ManageTeams = "teams.*"
  )
  ```

- [ ] **Atualizar PermissionSeeder** para incluir novas permissões

- [ ] **Atualizar RolePermissionSeeder** para mapear permissões a roles apropriadas

### Passo 6: Dependency Injection

- [ ] **Criar Container** (`config/dependency_injection_teams.go`)
  ```go
  type TeamContainer struct {
      TeamGateway       gateway.TeamGateway
      TeamMapper        *mapper.TeamMapper
      CreateTeamUseCase *teams.CreateTeam
      ListTeamsUseCase  *teams.ListTeams
      TeamController    *controller.TeamController
  }

  func InitializeTeamContainer(db *gorm.DB, validator *validation.Validator) *TeamContainer {
      teamMapper := &mapper.TeamMapper{}
      teamGateway := gateway.NewTeamGatewayGORM(db, teamMapper)

      createTeamUseCase := teams.NewCreateTeam(teamGateway)
      listTeamsUseCase := teams.NewListTeams(teamGateway)

      teamController := controller.NewTeamController(createTeamUseCase, listTeamsUseCase, validator)

      return &TeamContainer{
          TeamGateway:       teamGateway,
          TeamMapper:        teamMapper,
          CreateTeamUseCase: createTeamUseCase,
          ListTeamsUseCase:  listTeamsUseCase,
          TeamController:    teamController,
      }
  }
  ```

- [ ] **Registrar no Container Principal** (`config/dependency_injection.go`)
  ```go
  type Container struct {
      // ... outros containers
      TeamContainer *TeamContainer
  }

  func (c *Container) Initialize() error {
      // ... outras inicializações
      c.TeamContainer = InitializeTeamContainer(c.DB, c.Validator)
      return nil
  }
  ```

### Passo 7: Migrations

- [ ] **Adicionar model ao AutoMigrate** (`cmd/api/main.go`)
  ```go
  db.AutoMigrate(
      // ... outros models
      &model.Team{},
  )
  ```

- [ ] **(Opcional) Criar migration SQL manual** se necessário (`migrations/up/`, `migrations/down/`)

### Passo 8: Seeders (opcional)

- [ ] **Criar Seeder** (`seeders/team_seeder.go`) se precisar dados iniciais
  ```go
  type TeamSeeder struct {
      db          *gorm.DB
      teamGateway gateway.TeamGateway
  }

  func (s *TeamSeeder) Seed(ctx context.Context) error {
      // Criar teams de exemplo
  }
  ```

- [ ] **Adicionar ao runSeeds** em `main.go`

### Passo 9: Testes

- [ ] **Unit tests para entity** (validação, lógica de negócio)
- [ ] **Unit tests para use cases** (mock do gateway)
- [ ] **Integration tests para gateway** (banco de dados real/test)
- [ ] **E2E tests para endpoints** (servidor HTTP)

---

## 12. Stack Tecnológico Recomendado

| Componente | Tecnologia | Versão Mínima |
|------------|------------|---------------|
| **Linguagem** | Go | 1.22+ |
| **Framework HTTP** | Fiber | v2 |
| **Database** | PostgreSQL | 14+ |
| **ORM** | GORM | v2 |
| **Authentication** | golang-jwt/jwt | v5 |
| **Password Hashing** | bcrypt (golang.org/x/crypto) | latest |
| **Validation** | go-playground/validator | v10 |
| **UUID** | google/uuid | latest |
| **API Documentation** | swaggo/swag | latest |
| **Environment** | godotenv | latest (dev) |

**Dependências Opcionais:**
- **Redis**: Cache, sessions, rate limiting
- **Logrus/Zap**: Logging estruturado
- **Viper**: Configuração avançada
- **Testify**: Assertions e mocks para testes

---

## 13. Boas Práticas e Recomendações

### 13.1 Separação de Responsabilidades

**Regra**: Cada camada tem uma responsabilidade única e bem definida.

- **Domínio**: Regras de negócio puras
- **Aplicação**: Orquestração de casos de uso
- **Dados**: Persistência
- **HTTP**: Entrada/saída via web

**NÃO faça:**
- Controller chamando gateway diretamente (pule use case)
- Use case conhecendo detalhes HTTP (Request, Response)
- Entity conhecendo banco de dados (GORM)
- Gateway com lógica de negócio

### 13.2 Encapsulamento

**Regra**: Campos de entities devem ser privados, acessados via getters.

**Por quê?**
- Protege invariantes
- Permite validação
- Facilita mudanças futuras

```go
// BOM
type User struct {
    id    uuid.UUID // privado
    email string    // privado
}
func (u *User) ID() uuid.UUID { return u.id }
func (u *User) Email() string { return u.email }

// RUIM
type User struct {
    ID    uuid.UUID // público - qualquer um pode mudar
    Email string    // público - sem controle
}
```

### 13.3 Immutabilidade

**Regra**: Value objects devem ser imutáveis.

```go
// BOM
type Pagination struct {
    Page     int
    PageSize int
}
// Sem setters, criar novo objeto para mudar

// RUIM
func (p *Pagination) SetPage(page int) {
    p.Page = page // mutação
}
```

### 13.4 Inversão de Dependência

**Regra**: Módulos de alto nível não dependem de módulos de baixo nível. Ambos dependem de abstrações.

```go
// BOM
// Use case depende de interface (abstração)
type CreateUser struct {
    userGateway gateway.UserGateway // interface
}

// RUIM
// Use case depende de implementação concreta
type CreateUser struct {
    userGateway *gateway.UserGatewayGORM // concreto
}
```

### 13.5 Soft Delete por Padrão

**Regra**: Use soft delete para preservar dados e auditoria.

**Quando NÃO usar:**
- Dados sensíveis (GDPR, LGPD)
- Dados temporários (sessions, tokens)
- Tabelas de log

**Implementação:**
- `DeletedAt *time.Time` em todos os models
- Filtro automático em queries
- Método `IsActive()` nas entities

### 13.6 Validation em Múltiplas Camadas

**Camadas de Validação:**

1. **DTO/Payload** (formato): `validate:"required,email,min=3"`
2. **Entity** (negócio): `Validate()` method
3. **Banco** (constraints): `UNIQUE`, `NOT NULL`, `CHECK`

```go
// Payload: formato
type CreateUserRequest struct {
    Email string `validate:"required,email"`
}

// Entity: regras de negócio
func (u *User) Validate() error {
    if !u.IsValidBusinessEmail() {
        return errors.New("business email required")
    }
    return nil
}

// Banco: constraints
CREATE TABLE users (
    email VARCHAR(255) UNIQUE NOT NULL
);
```

### 13.7 Erros Semânticos

**Regra**: Use erros customizados de domínio, não erros técnicos.

```go
// BOM
var ErrUserNotFound = errors.New("user not found")
return domainErrors.ErrUserNotFound

// RUIM
return gorm.ErrRecordNotFound // expõe detalhe de infra
```

**Mapeamento no Controller:**
```go
if errors.Is(err, domainErrors.ErrUserNotFound) {
    return c.Status(404).JSON(...)
}
```

### 13.8 Context Propagation

**Regra**: Passe `context.Context` em todas as operações assíncronas.

**Por quê?**
- Permite cancelamento
- Timeout de operações
- Valores no contexto (user ID, trace ID)

```go
// BOM
func (uc *CreateUser) Execute(ctx context.Context, data UserData) (*entity.User, error) {
    user, err := uc.userGateway.Save(ctx, user) // propaga context
    ...
}

// RUIM
func (uc *CreateUser) Execute(data UserData) (*entity.User, error) {
    user, err := uc.userGateway.Save(context.Background(), user) // context fixo
    ...
}
```

### 13.9 Dependency Injection

**Regra**: Injete dependências via constructor, não crie internamente.

```go
// BOM
type CreateUser struct {
    userGateway gateway.UserGateway
}
func NewCreateUser(userGateway gateway.UserGateway) *CreateUser {
    return &CreateUser{userGateway: userGateway}
}

// RUIM
type CreateUser struct {}
func (uc *CreateUser) Execute(...) {
    userGateway := gateway.NewUserGatewayGORM(...) // cria internamente
}
```

**Benefícios:**
- Facilita testes (mock de gateway)
- Controle de ciclo de vida
- Configuração centralizada

### 13.10 Migrations Versionadas

**Regra**: Use migrations versionadas com up e down.

**Formato do nome:**
```
{timestamp}_{description}.{up|down}.sql
1704906420_create_users_table.up.sql
1704906420_create_users_table.down.sql
```

**Características:**
- **Versionadas**: Timestamp para ordem
- **Reversíveis**: Sempre crie down migration
- **Idempotentes**: Podem rodar múltiplas vezes
- **Incrementais**: Uma mudança por migration

```sql
-- up
CREATE TABLE IF NOT EXISTS users (...);

-- down
DROP TABLE IF EXISTS users;
```

---

## Conclusão

Esta arquitetura combina os melhores princípios de **Clean Architecture** e **Domain-Driven Design** para criar sistemas backend escaláveis, testáveis e manuteníveis em Go.

**Principais Takeaways:**

1. **Separação Clara de Camadas**: Cada camada tem responsabilidade única
2. **Inversão de Dependência**: Domínio define interfaces, infra implementa
3. **Testabilidade**: Cada camada pode ser testada isoladamente
4. **Independência de Frameworks**: Lógica de negócio não depende de Fiber, GORM, etc.
5. **Escalabilidade**: Fácil adicionar features sem quebrar existentes
6. **Manutenibilidade**: Código organizado e fácil de navegar

**Quando usar esta arquitetura:**
- APIs REST de médio a grande porte
- Sistemas com lógica de negócio complexa
- Projetos que precisam de alta testabilidade
- Aplicações que podem mudar requisitos técnicos (trocar banco, framework)

**Quando NÃO usar:**
- Protótipos rápidos ou MVPs muito simples
- Scripts ou ferramentas CLI básicas
- APIs extremamente simples (CRUD puro sem regras)

Use o **Checklist de Features** (seção 11) como guia prático para adicionar novas funcionalidades seguindo todos os padrões descritos neste documento.
