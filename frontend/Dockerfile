# Etapa 1: Build stage (para produção)
FROM node:18-alpine AS builder

# Definir o diretório de trabalho
WORKDIR /app

# Copiar arquivos de dependências
COPY package*.json ./

# Instalar dependências
RUN npm install

# Copiar código fonte
COPY . .

# Definir argumento de build para a URL da API
ARG VITE_API_URL
ENV VITE_API_URL=$VITE_API_URL

# Construir a aplicação
RUN npm run build

# Etapa 2: Runtime stage com hot reload (desenvolvimento)
FROM node:18-alpine

# Criar diretório de trabalho
WORKDIR /app

# Copiar package.json
COPY package*.json ./

# Instalar dependências
RUN npm install

# Copiar código fonte para hot reload
COPY . .

# Definir argumento de build para a URL da API
ARG VITE_API_URL
ENV VITE_API_URL=$VITE_API_URL

# Expor a porta 3000 (porta padrão do Vite)
EXPOSE 3000

# Comando para executar o Vite em modo dev com hot reload
CMD ["npm", "run", "dev"]