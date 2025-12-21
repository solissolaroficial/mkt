# Etapa 1: Build stage
FROM golang:1.21-alpine AS builder

# Definir o diretório de trabalho
WORKDIR /app

# Copiar arquivos go mod
COPY go.mod go.sum ./

# Baixar dependências
RUN go mod download

# Copiar código fonte
COPY . .

# Construir a aplicação
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Etapa 2: Runtime stage
FROM alpine:latest

# Instalar ca-certificates para requisições HTTPS
RUN apk --no-cache add ca-certificates

# Criar diretório de trabalho
WORKDIR /root/

# Copiar o binário da etapa de build
COPY --from=builder /app/main .

# Copiar arquivo .env se existir
COPY --from=builder /app/.env .env

# Expor a porta 8080
EXPOSE 8080

# Comando para executar a aplicação
CMD ["./main"]