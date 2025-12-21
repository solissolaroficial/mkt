# Etapa 1: Build stage
FROM node:18-alpine AS builder

# Definir o diretório de trabalho
WORKDIR /app

# Copiar arquivos de dependências
COPY package*.json ./

# Instalar dependências
RUN npm ci

# Copiar código fonte
COPY . .

# Construir a aplicação
RUN npm run build

# Etapa 2: Runtime stage com servidor de desenvolvimento
FROM node:18-alpine

# Instalar serve para servir os arquivos estáticos
RUN npm install -g serve

# Criar diretório de trabalho
WORKDIR /app

# Copiar o build da etapa anterior
COPY --from=builder /app/dist ./dist

# Expor a porta 80
EXPOSE 80

# Comando para servir os arquivos estáticos
CMD ["serve", "-s", "dist", "-l", "80"]