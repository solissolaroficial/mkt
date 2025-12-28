FROM golang:1.25-alpine

# Instalar dependências necessárias
RUN apk add --no-cache git ca-certificates

WORKDIR /app

# Instalar air para hot reload
RUN go install github.com/air-verse/air@v1.52.3

# Copiar arquivos go mod
COPY go.mod go.sum ./
RUN go mod download

# Copiar arquivo .air.toml
COPY .air.toml .

# Copiar script de entrada
COPY docker-entrypoint.sh /usr/local/bin/
RUN chmod +x /usr/local/bin/docker-entrypoint.sh

# Criar diretório tmp para o air
RUN mkdir -p tmp

EXPOSE 8080

CMD ["/usr/local/bin/docker-entrypoint.sh"]
