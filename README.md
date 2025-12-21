# Solis Hub - Sistema de Gestão de Marketing

## Estrutura do Projeto

```
solis-hub/
├── backend/          # API REST em Go + Fiber
├── frontend/         # Aplicação React + TypeScript
└── docker-compose.yml # Orquestração dos serviços
```

## Tecnologias

- **Backend**: Go, Fiber, GORM, PostgreSQL
- **Frontend**: React, TypeScript, Vite
- **Banco de Dados**: PostgreSQL
- **Containerização**: Docker e Docker Compose

## Como Executar o Projeto

### Pré-requisitos
- Docker e Docker Compose instalados

### Passos para Executar

1. **Clonar o repositório**
   ```bash
   git clone <repository-url>
   cd solis-hub
   ```

2. **Configurar variáveis de ambiente**
   ```bash
   # Copiar arquivo de exemplo
   cp backend/.env.example backend/.env
   
   # Editar as variáveis conforme necessário
   nano backend/.env
   ```

3. **Executar com Docker Compose**
   ```bash
   # Construir e iniciar todos os serviços
   docker-compose up --build
   
   # Ou em modo detached
   docker-compose up --build -d
   ```

4. **Acessar as aplicações**
   - Frontend: http://localhost:3500
   - Backend API: http://localhost:8500
   - Saúde da API: http://localhost:8500/health

### Estrutura de Portas

- **Frontend**: 3500:80 (mapeado para porta 3500 no host)
- **Backend**: 8500:8080 (mapeado para porta 8500 no host)
- **PostgreSQL**: 5432 (não exposto externamente, apenas interno)

## Desenvolvimento

### Backend
```bash
cd backend
go mod tidy
go run main.go
```

### Frontend
```bash
cd frontend
npm install
npm run dev
```

## API Endpoints

### Autenticação
- `POST /api/auth/login` - Login de usuário
- `POST /api/auth/logout` - Logout
- `GET /api/auth/profile` - Perfil do usuário

### KPIs
- `GET /api/kpis` - Lista todos os KPIs
- `GET /api/kpis/:id` - Detalhes de um KPI
- `PUT /api/kpis/:id/data/:month` - Atualiza valor realizado
- `PUT /api/kpis/:id/meta/:month` - Atualiza meta

### Tarefas
- `GET /api/tasks` - Lista tarefas
- `POST /api/tasks` - Cria nova tarefa
- `GET /api/tasks/:id` - Detalhes da tarefa
- `PUT /api/tasks/:id` - Atualiza tarefa
- `DELETE /api/tasks/:id` - Exclui tarefa

### E muitas outras rotas...

## Banco de Dados

O PostgreSQL é inicializado automaticamente pelo Docker Compose. Os dados são persistidos no volume `postgres_data`.

## Variáveis de Ambiente

### Backend
- `DB_HOST` - Host do banco de dados
- `DB_PORT` - Porta do banco
- `DB_USER` - Usuário do banco
- `DB_PASSWORD` - Senha do banco
- `DB_NAME` - Nome do banco
- `PORT` - Porta do servidor backend
- `JWT_SECRET` - Chave secreta para tokens JWT

### Frontend
- `VITE_API_URL` - URL da API backend

## Deploy

Para produção, altere as seguintes configurações:

1. **Segurança**:
   - Altere o `JWT_SECRET` para um valor seguro
   - Configure HTTPS
   - Ajuste as configurações de CORS

2. **Banco de Dados**:
   - Use senhas fortes
   - Configure backups regulares
   - Ajuste as configurações de performance

3. **Monitoramento**:
   - Configure logs centralizados
   - Implemente métricas de monitoramento
   - Configure alertas

## Contribuição

1. Faça um fork do projeto
2. Crie uma branch para sua feature
3. Faça commit das suas mudanças
4. Abra um Pull Request

## Licença

MIT License - veja o arquivo LICENSE para detalhes.