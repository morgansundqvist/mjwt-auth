package mjwtauth

// AuthUser represents the minimal user information needed for authentication.
type AuthUser interface {
	GetID() string
	GetUsername() string
	GetPasswordHash() string
}
