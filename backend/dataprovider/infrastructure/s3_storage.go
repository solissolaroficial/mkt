package infrastructure

import (
	"context"
	"fmt"
	"io"
	"log"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	awscfg "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

// S3Storage implementa StorageGateway usando AWS S3
type S3Storage struct {
	client    *s3.Client
	uploader  *manager.Uploader
	bucket    string
	region    string
	publicURL string
}

// NewS3Storage cria novo storage S3
func NewS3Storage(
	accessKey string,
	secretKey string,
	region string,
	bucket string,
	publicURL string,
) (gateway.StorageGateway, error) {
	// Se não houver credenciais, retorna nil (storage desabilitado)
	if accessKey == "" || secretKey == "" {
		log.Println("[S3 STORAGE] Not configured - storage disabled")
		return nil, nil
	}

	// Configurar AWS SDK
	cfg, err := awscfg.LoadDefaultConfig(context.TODO(),
		awscfg.WithRegion(region),
		awscfg.WithCredentialsProvider(
			aws.CredentialsProviderFunc(func(ctx context.Context) (aws.Credentials, error) {
				return aws.Credentials{
					AccessKeyID:     accessKey,
					SecretAccessKey: secretKey,
				}, nil
			}),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to load AWS config: %w", err)
	}

	// Criar cliente S3
	client := s3.NewFromConfig(cfg)

	// Criar uploader com configurações otimizadas
	// PartSize de 5MB permite uploads eficientes de arquivos grandes
	uploader := manager.NewUploader(client, func(u *manager.Uploader) {
		u.PartSize = 5 * 1024 * 1024 // 5MB por parte
	})

	log.Printf("[S3 STORAGE] Initialized - Bucket: %s, Region: %s", bucket, region)

	return &S3Storage{
		client:    client,
		uploader:  uploader,
		bucket:    bucket,
		region:    region,
		publicURL: publicURL,
	}, nil
}

// ValidateBucket verifica se o bucket S3 existe e está acessível
// Útil para validar configurações na inicialização da aplicação
func (s *S3Storage) ValidateBucket(ctx context.Context) error {
	log.Printf("[S3 STORAGE] Validating bucket: %s", s.bucket)

	_, err := s.client.HeadBucket(ctx, &s3.HeadBucketInput{
		Bucket: aws.String(s.bucket),
	})
	if err != nil {
		log.Printf("[S3 STORAGE] Bucket validation failed: %v", err)
		return fmt.Errorf("failed to validate S3 bucket: %w", err)
	}

	log.Printf("[S3 STORAGE] Bucket validated successfully: %s", s.bucket)
	return nil
}

// UploadFile faz upload de um arquivo para o S3
func (s *S3Storage) UploadFile(ctx context.Context, reader io.Reader, key string, contentType string) (string, error) {
	log.Printf("[S3 STORAGE] Uploading file: %s", key)

	// Upload com multipart upload automático
	result, err := s.uploader.Upload(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(s.bucket),
		Key:         aws.String(key),
		Body:        reader,
		ContentType: aws.String(contentType),
		// Nota: ACL foi removido pois não é mais suportado em versões recentes do AWS SDK v2
		// Configure as permissões do bucket via Bucket Policy ou IAM
	})
	if err != nil {
		log.Printf("[S3 STORAGE] Error uploading file: %v", err)
		return "", fmt.Errorf("failed to upload file: %w", err)
	}

	log.Printf("[S3 STORAGE] File uploaded successfully: %s (Location: %s)", key, result.Location)

	// Retornar URL pública
	return s.GetFileURL(key), nil
}

// DeleteFile deleta um arquivo do S3
func (s *S3Storage) DeleteFile(ctx context.Context, key string) error {
	log.Printf("[S3 STORAGE] Deleting file: %s", key)

	_, err := s.client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		log.Printf("[S3 STORAGE] Error deleting file: %v", err)
		return fmt.Errorf("failed to delete file: %w", err)
	}

	log.Printf("[S3 STORAGE] File deleted successfully: %s", key)
	return nil
}

// GetFileURL retorna a URL pública de um arquivo
func (s *S3Storage) GetFileURL(key string) string {
	if s.publicURL != "" {
		// Usar URL customizada (ex: CloudFront)
		return strings.TrimSuffix(s.publicURL, "/") + "/" + key
	}
	// Usar URL padrão do S3
	return fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", s.bucket, s.region, key)
}

// GetPresignedURL retorna uma URL assinada temporária para acesso ao arquivo
// Útil para arquivos privados que precisam de acesso temporário
func (s *S3Storage) GetPresignedURL(ctx context.Context, key string, expiration time.Duration) (string, error) {
	log.Printf("[S3 STORAGE] Generating presigned URL for key: %s with expiration: %v", key, expiration)

	presignClient := s3.NewPresignClient(s.client)

	req, err := presignClient.PresignGetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key),
	}, s3.WithPresignExpires(expiration))
	if err != nil {
		log.Printf("[S3 STORAGE] Error generating presigned URL: %v", err)
		return "", fmt.Errorf("failed to generate presigned URL: %w", err)
	}

	log.Printf("[S3 STORAGE] Presigned URL generated successfully: %s", req.URL)

	return req.URL, nil
}
