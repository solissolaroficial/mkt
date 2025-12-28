package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// DatabaseConfig holds database configuration
type DatabaseConfig struct {
	Host            string
	Port            int
	User            string
	Password        string
	DBName          string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime int // in seconds
	ConnMaxIdleTime int // in seconds
}

// JWTConfig holds JWT configuration
type JWTConfig struct {
	Secret                  string
	AccessTokenExpiryHours  int
	RefreshTokenExpiryHours int
}

// ServerConfig holds server configuration
type ServerConfig struct {
	Port        int
	Environment string
}

// Config holds all application configuration
type Config struct {
	Database DatabaseConfig
	JWT      JWTConfig
	Server   ServerConfig
}

// getEnv returns environment variable value or default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// getEnvAsInt returns environment variable value as int or default value
func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

// LoadConfig loads configuration from .env file and environment variables
func LoadConfig() (*Config, error) {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		// Continue even if .env file is not found
		// This allows running with environment variables only
		fmt.Println("Warning: .env file not found, using environment variables")
	}

	config := &Config{
		Database: DatabaseConfig{
			Host:            getEnv("DB_HOST", "localhost"),
			Port:            getEnvAsInt("DB_PORT", 5432),
			User:            getEnv("DB_USER", "postgres"),
			Password:        getEnv("DB_PASSWORD", ""),
			DBName:          getEnv("DB_NAME", "solis_mkt"),
			MaxOpenConns:    getEnvAsInt("DB_MAX_OPEN_CONNS", 25),
			MaxIdleConns:    getEnvAsInt("DB_MAX_IDLE_CONNS", 5),
			ConnMaxLifetime: getEnvAsInt("DB_CONN_MAX_LIFETIME", 300),
			ConnMaxIdleTime: getEnvAsInt("DB_CONN_MAX_IDLE_TIME", 60),
		},
		JWT: JWTConfig{
			Secret:                  getEnv("JWT_SECRET", ""),
			AccessTokenExpiryHours:  getEnvAsInt("JWT_ACCESS_TOKEN_EXPIRY_HOURS", 1),
			RefreshTokenExpiryHours: getEnvAsInt("JWT_REFRESH_TOKEN_EXPIRY_HOURS", 24),
		},
		Server: ServerConfig{
			Port:        getEnvAsInt("SERVER_PORT", 8500),
			Environment: getEnv("SERVER_ENVIRONMENT", "development"),
		},
	}

	return config, nil
}

// Validate validates the configuration
func (c *Config) Validate() error {
	// Validate JWT secret is not empty
	if c.JWT.Secret == "" {
		return fmt.Errorf("JWT_SECRET cannot be empty")
	}

	// Validate DB password is present in production
	if c.Server.Environment == "production" && c.Database.Password == "" {
		return fmt.Errorf("DB_PASSWORD cannot be empty in production environment")
	}

	// Validate database name
	if c.Database.DBName == "" {
		return fmt.Errorf("DB_NAME cannot be empty")
	}

	// Validate database user
	if c.Database.User == "" {
		return fmt.Errorf("DB_USER cannot be empty")
	}

	// Validate port ranges
	if c.Database.Port <= 0 || c.Database.Port > 65535 {
		return fmt.Errorf("DB_PORT must be between 1 and 65535")
	}

	if c.Server.Port <= 0 || c.Server.Port > 65535 {
		return fmt.Errorf("SERVER_PORT must be between 1 and 65535")
	}

	// Validate JWT expiry hours
	if c.JWT.AccessTokenExpiryHours <= 0 {
		return fmt.Errorf("JWT_ACCESS_TOKEN_EXPIRY_HOURS must be greater than 0")
	}

	if c.JWT.RefreshTokenExpiryHours <= 0 {
		return fmt.Errorf("JWT_REFRESH_TOKEN_EXPIRY_HOURS must be greater than 0")
	}

	return nil
}
