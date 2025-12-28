package service

// HasherService defines the interface for password hashing operations
type HasherService interface {
	// Hash generates a bcrypt hash for the provided password
	Hash(password string) (string, error)

	// Compare verifies if the provided password matches the hashed password
	Compare(hashedPassword, password string) error
}
