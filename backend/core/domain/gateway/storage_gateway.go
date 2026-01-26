package gateway

import (
	"context"
	"io"
	"time"
)

// StorageGateway define interface para armazenamento de arquivos
// Esta interface segue o princípio de Inversão de Dependência
type StorageGateway interface {
	// UploadFile faz upload de um arquivo para o storage
	// Retorna a URL pública do arquivo
	UploadFile(ctx context.Context, reader io.Reader, key string, contentType string) (string, error)

	// DeleteFile deleta um arquivo do storage
	DeleteFile(ctx context.Context, key string) error

	// GetFileURL retorna a URL pública de um arquivo
	GetFileURL(key string) string

	// GetPresignedURL retorna uma URL assinada temporária para acesso ao arquivo
	// Útil para arquivos privados que precisam de acesso temporário
	GetPresignedURL(ctx context.Context, key string, expiration time.Duration) (string, error)

	// ValidateBucket verifica se o bucket S3 existe e está acessível
	// Útil para validar configurações na inicialização da aplicação
	ValidateBucket(ctx context.Context) error
}
